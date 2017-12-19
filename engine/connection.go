package engine

import "github.com/google/uuid"

type connection struct {
	id ConnectionID
}

func (c *connection) ID() ConnectionID {
	return c.id
}

type Connection interface {
	ID() ConnectionID
}

type ConnectionID interface {
	String() string
}

type connectionID uuid.UUID

var space = uuid.New()

func CreateConnectionID() ConnectionID {

	data := [16]byte(uuid.New())
	sliced := data[:]
	return uuid.NewSHA1(space, sliced)
}

func (id *connectionID) String() string {
	return uuid.UUID(*id).String()
}

type ConnectionManager interface {
	GetConnection(connectionID ConnectionID) Connection
	CreateConnection() Connection
}

type connectionManager struct {
	connections map[ConnectionID]Connection
}

func (m *connectionManager) GetConnection(connectionID ConnectionID) Connection {
	return m.connections[connectionID]
}
func (m *connectionManager) CreateConnection() Connection {
	connection := &connection{}
	id := CreateConnectionID()
	connection.id = id
	m.connections[id] = connection
	return connection
}
func CreateConnectionManager() ConnectionManager {
	m := &connectionManager{}
	m.connections = make(map[ConnectionID]Connection)
	return m
}
