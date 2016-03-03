package gontlet

type Transport interface {
	sendOutgoing([]byte)
	serve()
	unregisterConn(conn connection)
	recieve(buff []byte)
}
