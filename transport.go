package gontlet

type Transport interface {
	sendOutgoing([]byte)
	serve()
}
