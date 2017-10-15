package sqlmock

import (
	"context"
	"database/sql/driver"
)

// Conn is a connection object.
type Conn struct {
	ExpectedResult     *Result
	ExpectedRows       *Rows
	ExpectedCloseErr   error
	ExpectedPrepareErr error
	ExpectedBeginErr   error
	ExpectedExecErr    error
	ExpectedQueryErr   error
	CanPing            bool
}

// Prepare returns a prepared statement
func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	if c.ExpectedPrepareErr != nil {
		return nil, c.ExpectedPrepareErr
	}
	return &Stmt{conn: c, query: query}, nil
}

// Close do nothing
func (c *Conn) Close() error { return c.ExpectedCloseErr }

// Begin starts a new transaction
func (c *Conn) Begin() (driver.Tx, error) {
	if c.ExpectedBeginErr != nil {
		return nil, c.ExpectedBeginErr
	}
	return &Tx{}, nil
}

// BeginTx implements driver.ConnBeginTx to Conn
func (c *Conn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	if c.ExpectedBeginErr != nil {
		return nil, c.ExpectedBeginErr
	}
	return &Tx{ctx: ctx}, nil
}

// PrepareContext implements driver.ConnPrepareContext to Conn
func (c *Conn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	if c.ExpectedPrepareErr != nil {
		return nil, c.ExpectedPrepareErr
	}
	return &Stmt{ctx: ctx, conn: c, query: query}, nil
}

// Exec implements driver.Execer to Conn
func (c *Conn) Exec(query string, args []driver.Value) (driver.Result, error) {
	if c.ExpectedExecErr != nil {
		return nil, c.ExpectedExecErr
	}
	return c.ExpectedResult, nil
}

// ExecContext implements driver.ExecerContext to Conn
func (c *Conn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	if c.ExpectedExecErr != nil {
		return nil, c.ExpectedExecErr
	}
	return c.ExpectedResult, nil
}

// Ping implements driver.Pinger to Conn
func (c *Conn) Ping(ctx context.Context) error {
	if c.CanPing {
		return nil
	}
	return driver.ErrBadConn
}

// Query implements driver.Queryer to Conn
func (c *Conn) Query(query string, args []driver.Value) (driver.Rows, error) {
	if c.ExpectedQueryErr != nil {
		return nil, c.ExpectedQueryErr
	}
	return c.ExpectedRows, nil
}

// QueryContext implements driver.QueryerContext to Conn
func (c *Conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	if c.ExpectedQueryErr != nil {
		return nil, c.ExpectedQueryErr
	}
	return c.ExpectedRows, nil
}
