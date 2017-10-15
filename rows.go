package sqlmock

import (
	"database/sql/driver"
	"io"
)

// Rows mock
type Rows struct {
	ExpectedColumns    []string
	ExpectedValuesList [][]driver.Value
	ExpectedCloseErr   error
	ExpectedNextErr    error
	pos                int
}

// Columns returns the names of the columns
func (rows *Rows) Columns() []string {
	return rows.ExpectedColumns
}

// Close do nothing
func (rows *Rows) Close() error {
	return rows.ExpectedCloseErr
}

// Next mock
func (rows *Rows) Next(dest []driver.Value) error {
	if rows.ExpectedNextErr != nil {
		return rows.ExpectedNextErr
	}
	if len(rows.ExpectedValuesList) <= rows.pos {
		return io.EOF
	}
	expectedValues := rows.ExpectedValuesList[rows.pos]
	rows.pos++
	for i, val := range expectedValues {
		dest[i] = val
	}
	return nil
}
