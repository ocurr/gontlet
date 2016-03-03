package gontlet

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	connections map[connection]bool
	send        chan []byte
	recv        chan []byte
	register    chan connection
	unregister  chan connection
}

func listen(l net.Listener, s *Server) {
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			break
		}
		s.registerConn(connection{conn: c, send: make(chan []byte, 1)})
	}
}

func newServer(port string) *Server {
	l, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		fmt.Println("Error starting listener: ", err)
	}
	s := Server{
		connections: make(map[connection]bool),
		send:        make(chan []byte, 1),
		recv:        make(chan []byte, 1),
		register:    make(chan connection, 3),
		unregister:  make(chan connection, 3),
	}
	go listen(l, &s)
	return &s
}

func (s *Server) registerConn(c connection) {
	s.register <- c
}

func (s *Server) unregisterConn(conn connection) {
	s.unregister <- conn
}

func (s *Server) sendOutgoing(data []byte) {
	s.send <- data
}

func (s *Server) getIncoming() []byte {
	return <-s.recv
}

func (s *Server) recieve(buff []byte) {
	s.recv <- buff
}

func (s *Server) serve() {
	for {
		select {
		case c := <-s.register:
			s.connections[c] = true
			go c.handle(s)
		case c := <-s.unregister:
			if _, ok := s.connections[c]; ok {
				delete(s.connections, c)
				close(c.send)
			}
		case m := <-s.send:
			for c := range s.connections {
				c.sendData(m)
			}
		}
	}
}
