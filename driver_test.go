package sqlmock_test

import (
	"database/sql"
	"database/sql/driver"
	"testing"

	sqlmock "github.com/nasa9084/go-sql-mock"
)

func TestDriver(t *testing.T) {
	db, err := sql.Open("sqlmock", "")
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	defer db.Close()
}

func TestExpectedRows(t *testing.T) {
	opts := []sqlmock.RowsOpts{
		sqlmock.Columns([]string{"id", "name", "age"}),
		sqlmock.ValuesList([]driver.Value{"something", "Alice", 10}),
	}
	db, err := sql.Open("sqlmock", "")
	if err != nil {
		t.Fatalf("%s", err)
	}
	sqlmock.ExpectedRows(opts...)
	row := db.QueryRow(`SELECT * FROM mock`)

	var some struct {
		ID   string
		Name string
		Age  int
	}
	if err := row.Scan(&some.ID, &some.Name, &some.Age); err != nil {
		t.Errorf("%s", err)
		return
	}
	if some.ID != "something" {
		t.Errorf(`"%s" != "%s"`, some.ID, "something")
		return
	}
	if some.Name != "Alice" {
		t.Errorf(`"%s" != "%s"`, some.Name, "Alice")
		return
	}
	if some.Age != 10 {
		t.Errorf(`%d != %d`, some.Age, 10)
		return
	}
}
