package internal

import (
	"fmt"
	"github.com/spf13/afero"
	"github.com/yalp/jsonpath"
	"go.uber.org/zap"
	"log"
	"os"
	"path"
	"sync"
)

type DB struct {
	config      *Config
	schemaMutex sync.RWMutex
	logger      *zap.Logger
}

func InitDb(config *Config) (*DB, error) {
	db := &DB{
		config:      config,
		schemaMutex: sync.RWMutex{},
		logger:      setupLogger(),
	}
	err := db.setupDataDir()
	return db, err
}

func setupLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}

func (d *DB) setupDataDir() error {
	fileSystem := d.config.Storage.Filesystem
	_, err := fileSystem.Stat(d.config.Storage.Directory)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		// Directory doesn't exist. Make it.
		err = fileSystem.MkdirAll(d.config.Storage.Directory, os.ModeDir)
		if err != nil {
			return err
		}
	}
	return err
}

func (d *DB) Validate() error {
	return nil
}

func (d *DB) CreateSchema(schema string) error {
	fileSystem := d.config.Storage.Filesystem

	d.schemaMutex.RLock()
	defer d.schemaMutex.RUnlock()

	schemasConfigPath := path.Join(d.config.Storage.Directory, "schemas.txt")
	f, err := fileSystem.OpenFile(schemasConfigPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer func(f afero.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	if err != nil {
		d.logger.Error("could not open schemas file", zap.Error(err))
		log.Fatal(err)
	}

	if _, err := f.Write([]byte(fmt.Sprintf("%s\n", schema))); err != nil {
		d.logger.Error("could not update schemas file", zap.Error(err))
		return err
	}

	err = fileSystem.MkdirAll(path.Join(d.config.Storage.Directory, "schemas", schema), os.ModeDir)
	if err != nil {
		d.logger.Error("could not create new schema directory", zap.Error(err))
		return err
	}

	return nil
}

func (d *DB) SchemaExists() bool {
	return false
}

func (d *DB) Save(schema string, item interface{}) error {
	_, err := jsonpath.Read(item, "$..id")
	return err
}

func (d *DB) GetByID(schema string, id string) error {
	panic("implement me")
}
