package sqlmock

import (
	"database/sql"
	"database/sql/driver"
	"sync"
)

var pool *Driver

func init() {
	pool = &Driver{}
	sql.Register("sqlmock", pool)
}

// Driver implements a mock of database driver
type Driver struct {
	sync.Mutex
}

// Open returns a new connection mock
func (d *Driver) Open(name string) (driver.Conn, error) {
	d.Lock()
	defer d.Unlock()

	return &Conn{}, nil
}
