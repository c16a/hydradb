package internal

import (
	"testing"
)

func TestInitDb(t *testing.T) {
	config := &Config{}

	db, err := InitDb(config)
	if err != nil {
		t.Fail()
	}

	if db == nil {
		t.Fail()
	}
}
