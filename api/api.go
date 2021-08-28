package api

type DbApi interface {
	CreateSchema(schema string) error
	Save(schema string, item interface{}) error
	GetByID(schema string, id string) error
}
