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
	conn *Conn
}

// Open returns a new connection mock
func (d *Driver) Open(name string) (driver.Conn, error) {
	d.Lock()
	defer d.Unlock()

	if d.conn == nil {
		d.conn = &Conn{}
	}

	return d.conn, nil
}
