package sqlmock_test

import (
	"database/sql"
	"testing"
)

func TestDriver(t *testing.T) {
	db, err := sql.Open("sqlmock", "")
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	defer db.Close()
}
