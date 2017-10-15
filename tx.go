package sqlmock

import "context"

// Tx mock
type Tx struct {
	ctx                 context.Context
	ExpectedCommitErr   error
	ExpectedRollbackErr error
}

// Commit do nothing
func (tx *Tx) Commit() error { return tx.ExpectedCommitErr }

// Rollback do nothing
func (tx *Tx) Rollback() error { return tx.ExpectedRollbackErr }
