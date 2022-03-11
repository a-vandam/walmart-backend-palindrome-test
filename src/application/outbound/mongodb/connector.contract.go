package mongodb

type ConnectorContract interface {
	GetAll(string query)
}
