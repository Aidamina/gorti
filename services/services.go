package services

import "github.com/aidamina/gorti/engine"

type services struct {
	engine engine.Engine
}

//Services handles all RTI services
type Services interface {
	Engine() engine.Engine
}

func (s *services) Engine() engine.Engine {
	return s.engine
}

//CreateServices creates a Services instane
func CreateServices() Services {
	return &services{}
}
