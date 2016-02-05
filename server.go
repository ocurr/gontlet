package gontlet

import (
	"fmt"
	"net"
	"log"
)

type Server struct {
	robotConn net.Conn
	send chan []byte
	recv chan []byte
}

func listen(l net.Listener,s *Server) {
	defer l.Close()
	conn, err := l.Accept()
	if err != nil {
		log.Fatal(err)
	}
	s.registerConn(conn)
}

func NewServer(port string) *Server {
	l, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		fmt.Println("Error starting listener: ", err)
	}
	s := Server{
		robotConn: nil,
		send: make(chan []byte, 25),
		recv: make(chan []byte, 25),
	}
	go listen(l,&s)
	return &s
}

func (s* Server) registerConn(c net.Conn) {
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
