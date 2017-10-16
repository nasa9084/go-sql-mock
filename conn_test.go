package sqlmock_test

import (
	"context"
	"database/sql/driver"
	"errors"
	"io"
	"reflect"
	"testing"

	sqlmock "github.com/nasa9084/go-sql-mock"
)

func TestClose(t *testing.T) {
	candidates := []struct {
		ExpectedCloseErr error
	}{
		{nil},
		{driver.ErrBadConn},
	}
	for _, c := range candidates {
		conn := sqlmock.Conn{
			ExpectedCloseErr: c.ExpectedCloseErr,
		}
		err := conn.Close()
		if err != c.ExpectedCloseErr {
			t.Errorf(`"%s" != "%s"`, err, c.ExpectedCloseErr)
			return
		}
	}
}

func TestPing(t *testing.T) {
	candidates := []struct {
		CanPing  bool
		Expected error
	}{
		{true, nil},
		{false, driver.ErrBadConn},
	}
	for _, c := range candidates {
		conn := sqlmock.Conn{
			CanPing: c.CanPing,
		}
		if err := conn.Ping(context.Background()); err != c.Expected {
			t.Errorf(`"%s" != "%s"`, err, c.Expected)
			return
		}
	}
}

func TestPrepare(t *testing.T) {
	candidates := []struct {
		Name        string
		Query       string
		Args        []driver.Value
		ExpectedErr error
	}{
		{
			Name:        "normal select",
			Query:       "SELECT * FROM mock",
			ExpectedErr: nil,
		}, {
			Name:        "select with placeholder",
			Query:       "SELECT * FROM mock WHERE id=?",
			Args:        []driver.Value{"something"},
			ExpectedErr: nil,
		}, {
			Name:        "empty query",
			Query:       "",
			ExpectedErr: nil,
		}, {
			Name:        "empty query and expectedErr",
			Query:       "",
			ExpectedErr: errors.New("query is empty"),
		},
	}
	for _, c := range candidates {
		t.Log(c.Name)
		conn := sqlmock.Conn{
			ExpectedPrepareErr: c.ExpectedErr,
		}
		stmt, err := conn.Prepare(c.Query)
		if err != c.ExpectedErr {
			t.Errorf(`"%s" != "%s"`, err, c.ExpectedErr)
			return
		}
		if err != nil && stmt != nil {
			t.Errorf("returned Stmt should be nil when error is returned")
			return
		}
	}
}

func TestBegin(t *testing.T) {
	candidates := []struct {
		ExpectedBeginErr error
	}{
		{nil},
		{driver.ErrBadConn},
	}
	for _, c := range candidates {
		conn := sqlmock.Conn{
			ExpectedBeginErr: c.ExpectedBeginErr,
		}
		if _, err := conn.Begin(); err != c.ExpectedBeginErr {
			t.Errorf(`"%s" != "%s"`, err, c.ExpectedBeginErr)
			return
		}
	}
}

func TestConnExec(t *testing.T) {
	candidates := []struct {
		Query           string
		ExpectedExecErr error
		ExpectedResult  *sqlmock.Result
	}{
		{
			"INSERT INTO mock VALUES('hoge');",
			nil,
			&sqlmock.Result{
				ExpectedLastInsertID: 1,
				ExpectedRowsAffected: 1,
			},
		},
		{
			"",
			errors.New("query is empty"),
			nil,
		},
	}
	for _, c := range candidates {
		conn := sqlmock.Conn{
			ExpectedExecErr: c.ExpectedExecErr,
			ExpectedResult:  c.ExpectedResult,
		}
		r, err := conn.Exec(c.Query, nil)
		if err != c.ExpectedExecErr {
			t.Errorf(`"%s" != "%s"`, err, c.ExpectedExecErr)
			return
		}
		if err != nil {
			continue
		}
		if r == nil {
			t.Errorf("result should not be nil")
			return
		}
		liid, err := r.LastInsertId()
		eliid, eerr := c.ExpectedResult.LastInsertId()
		if liid != eliid {
			t.Errorf("%d != %d", liid, eliid)
			return
		}
		if err != eerr {
			t.Errorf(`"%s" != "%s"`, err, eerr)
			return
		}
	}
}

func TestConnQuery(t *testing.T) {
	candidates := []struct {
		Query            string
		ExpectedQueryErr error
		ExpectedRows     *sqlmock.Rows
	}{
		{
			"",
			errors.New(`empty query`),
			nil,
		},
		{
			"SELECT * FROM mock",
			nil,
			&sqlmock.Rows{
				ExpectedColumns: []string{"id", "name", "age"},
				ExpectedValuesList: [][]driver.Value{
					[]driver.Value{"something", "anything", 10},
				},
			},
		},
	}
	for _, c := range candidates {
		conn := sqlmock.Conn{
			ExpectedQueryErr: c.ExpectedQueryErr,
			ExpectedRows:     c.ExpectedRows,
		}
		rows, err := conn.Query(c.Query, nil)
		if err != nil {
			if err != c.ExpectedQueryErr {
				t.Errorf(`"%s" != "%s"`, err, c.ExpectedQueryErr)
				return
			}
			continue
		}
		if rows == nil {
			t.Errorf("ether rows or nil should not be nil")
			return
		}
		if !reflect.DeepEqual(rows.Columns(), c.ExpectedRows.Columns()) {
			t.Errorf("%s != %s", rows.Columns(), c.ExpectedRows.Columns())
			return
		}
		var dest []driver.Value = make([]driver.Value, len(c.ExpectedRows.ExpectedColumns))
		err = rows.Next(dest)
		if err == io.EOF {
			t.Errorf("one row should be returned")
			return
		}
		for i, v := range dest {
			if v != c.ExpectedRows.ExpectedValuesList[0][i] {
				t.Errorf(`"%s" != "%s"`, v, c.ExpectedRows.ExpectedValuesList[0][i])
				return
			}
		}
		err = rows.Next(dest)
		if err != io.EOF {
			t.Errorf("just one row should be returned")
			return
		}
	}
}
