package internal

import (
	"github.com/spf13/afero"
	"testing"
)

func setupDb(t *testing.T, config *Config) *DB {
	db, err := InitDb(config)
	if err != nil {
		t.Error(err)
	}

	if db == nil {
		t.Error(err)
	}

	return db
}

func TestInitDb(t *testing.T) {
	fs := afero.NewMemMapFs()
	config := &Config{
		Storage: &Storage{
			Directory:  "/home/c16a/tmp/test/hydra",
			Filesystem: fs,
		},
	}
	setupDb(t, config)
}

func TestDB_CreateSchema(t *testing.T) {
	fs := afero.NewMemMapFs()
	config := &Config{
		Storage: &Storage{
			Directory:  "/home/c16a/tmp/test/hydra",
			Filesystem: fs,
		},
	}

	db := setupDb(t, config)
	err := db.CreateSchema("app")
	if err != nil {
		t.Error(err)
		return
	}

	// Schema definition list is present
	_, err = fs.Stat("/home/c16a/tmp/test/hydra/schemas.txt")
	if err != nil {
		t.Error(err)
		return
	}

	// Schema folder is present
	_, err = fs.Stat("/home/c16a/tmp/test/hydra/schemas/app")
	if err != nil {
		t.Error(err)
		return
	}
}
