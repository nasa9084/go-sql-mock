package sqlmock

import (
	"context"
	"database/sql/driver"
)

// Stmt is a prepared statement mock
type Stmt struct {
	conn             *Conn
	query            string
	ctx              context.Context
	ExpectedCloseErr error
	ExpectedExecErr  error
	ExpectedQueryErr error
}

// Close the statement
func (stmt *Stmt) Close() error { return stmt.ExpectedCloseErr }

// NumInput returns -1, which means this driver doesn't know the number of placeholders
// this causes passing sanity check Exec or Query argument counts
func (stmt *Stmt) NumInput() int { return -1 }

// Exec mock
func (stmt *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	if stmt.ExpectedExecErr != nil {
		return nil, stmt.ExpectedExecErr
	}
	return stmt.conn.Exec(stmt.query, args)
}

// Query mock
func (stmt *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	if stmt.ExpectedQueryErr != nil {
		return nil, stmt.ExpectedQueryErr
	}
	return stmt.conn.Query(stmt.query, args)
}
