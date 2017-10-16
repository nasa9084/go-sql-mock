package sqlmock

import "database/sql/driver"

// RowsOpts is used for arguments of ExpectedRows
type RowsOpts func(*Driver) error

// Columns define ExpectedRows' Columns
func Columns(cols []string) RowsOpts {
	return RowsOpts(func(d *Driver) error {
		d.conn.ExpectedRows.ExpectedColumns = cols
		return nil
	})
}

// ValuesList define ExpectedRows' ValuesList
func ValuesList(valuesList [][]driver.Value) RowsOpts {
	return RowsOpts(func(d *Driver) error {
		d.conn.ExpectedRows.ExpectedValuesList = valuesList
		return nil
	})
}

// AddValues append new Values to ExpectedRows' ValuesList
func AddValues(values ...[]driver.Value) RowsOpts {
	return RowsOpts(func(d *Driver) error {
		for _, value := range values {
			d.conn.ExpectedRows.ExpectedValuesList = append(d.conn.ExpectedRows.ExpectedValuesList, value)
		}
		return nil
	})
}

// ResultOpts is used for arguments of ExpectedResult
type ResultOpts func(*Driver) error

// LastInsertID define ExpectedResult's LastInsertID
func LastInsertID(liid int64) ResultOpts {
	return ResultOpts(func(d *Driver) error {
		d.conn.ExpectedResult.ExpectedLastInsertID = liid
		return nil
	})
}

// RowsAffected define ExpectedResult's RowsAffected
func RowsAffected(ra int64) ResultOpts {
	return ResultOpts(func(d *Driver) error {
		d.conn.ExpectedResult.ExpectedRowsAffected = ra
		return nil
	})
}
