package sql

import (
	"testing"
)

func TestNewConnection(t *testing.T) {
	db := NewConnection()

	t.Log("in use: ", db.pool.Stats().InUse)
	t.Log("idle: ", db.pool.Stats().Idle)
	t.Log("max open conn: ", db.pool.Stats().MaxOpenConnections)
	t.Log("max idle conn: ", db.pool.Stats().MaxIdleClosed)
	t.Log("max lifetime conn: ", db.pool.Stats().MaxLifetimeClosed)

	if err := db.pool.Ping(); err != nil {
		t.Error(err)
	}
}
