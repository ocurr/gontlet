package gontlet

import (
	"fmt"
	"net"
)

type Server struct {
	robotConn net.Conn
	send chan []byte
	recv chan []byte
}

func NewServer() *Server {
	s := Server{
		robotConn: nil,
		send: make(chan []byte, 25),
		recv: make(chan []byte, 25),
	}
	return &s
}

func (s* Server) RegisterConnection(c net.Conn) {
	s.robotConn = c
}

func (s* Server) SendOutgoing(data []byte) {
	s.send <- data
}

func (s* Server) GetIncoming() []byte {
	return <- s.recv
}

func (s* Server) WriteData() {
	defer s.robotConn.Close()
	for {
		buff := <-s.send
		msg := make([]byte, 1)
		msg[0] = byte(uint8(len(buff)))
		msg = append(msg, buff...)
		_, err := s.robotConn.Write(msg)
		if err != nil {
			fmt.Printf("Error writing to robot: ", err)
		}
	}
}

func (s* Server) ReadData() {
	defer s.robotConn.Close()
	for {
		buff := make([]byte, 1024)
		_, err := s.robotConn.Read(buff)
		if err != nil {
			fmt.Printf("Error reading from the robot: ", err)
		}
		s.recv <- buff
	}
}

func (s* Server) Serve() {
	go s.WriteData()
	go s.ReadData()
}
