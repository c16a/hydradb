package api

type DB interface {
	Save(interface{}) error
	SaveAll(...interface{}) error
	GetByID(string, interface{}) error
}
