package connector

type IConnector interface {
	Start() error
	Stop()
}
