package services

type services struct {
}

//Services handles all RTI services
type Services interface {
}

//CreateServices creates a Services instane
func CreateServices() Services {
	return &services{}
}
