package sqlmock

// Result implements driver.Result
type Result struct {
	ExpectedLastInsertID int64
	ExpectedRowsAffected int64
	ExpectedErr          error
}

// LastInsertId returns expected value
func (result *Result) LastInsertId() (int64, error) {
	if result.ExpectedErr != nil {
		return 0, result.ExpectedErr
	}
	return result.ExpectedLastInsertID, nil
}

// RowsAffected returns expected value
func (result *Result) RowsAffected() (int64, error) {
	if result.ExpectedErr != nil {
		return 0, result.ExpectedErr
	}
	return result.ExpectedRowsAffected, nil
}
