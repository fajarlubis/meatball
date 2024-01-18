package sql

import "testing"

func TestInsertOne(t *testing.T) {
	err := InsertOne()

	if err != nil {
		t.Fatal(err)
	}
}
