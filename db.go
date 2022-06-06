package fmk

// mockgen -source=db.go -destination=mocks/mock_db.go  -self_package IDb, IModel

// IDb required interface for DB connectivity
type IDb interface {
	Disconnect() error
	GetModel(colName string) IModel
}

// IModel required interface for CRUD Model operations
// for both all data associated to a single resource
type IModel interface {
	Insert(doc interface{}) (string, error)
	UpdateByFilter(filter interface{}, toChange interface{}) error
	FindByFilter(filter interface{}, sort interface{}, pagination MPagination, docs interface{}) error
	DeleteByFilter(filter interface{}) error
}
