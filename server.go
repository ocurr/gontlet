package gontlet

import (
	"fmt"
	"net"
)

type Server struct {
	listeners map[string][]net.Conn
	send chan string
	recv chan string
}

func NewServer() Server {
	s := Server{
		listeners: make(map[string][]net.Conn),
		send: make(chan string, 25)
		recv: make(chan string, 25)
	}
}

func (s* Server) RegisterConnection(tableName string, c net.Conn) {
	s.conns[tableName] = append(s.conns[tableName], c)
}

func (s* Server) UpdateListeners(tableName string, data []byte) error {
	for _,conn in range listeners[tableName] {
		_,err := conn.Write(data)
		if err != nil {
			return err
		}
	}
	return nil
}

