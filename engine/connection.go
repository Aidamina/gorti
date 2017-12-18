package engine

import "github.com/google/uuid"

type connection struct {
}

type Connection interface {
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
}
