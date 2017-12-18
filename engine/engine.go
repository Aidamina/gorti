package engine

type Engine interface {
	ConnectionManager() ConnectionManager
}

type engine struct {
	connectionManager ConnectionManager
}

func (e *engine) ConnectionManager() ConnectionManager {
	return e.connectionManager
}

func CreateEngine() Engine {
	return &engine{}
}
