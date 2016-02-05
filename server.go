package gontlet

import (
	"fmt"
	"net"
	"log"
	"time"
)

type Server struct {
	robotConn net.Conn
	send chan []byte
	recv chan []byte
	unregister chan bool
}

func listen(l net.Listener,s *Server) {
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		if <-s.unregister {
			s.registerConn(conn)
			s.unregister <- false
		}
	}
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
		unregister: make(chan bool, 1),
	}
	s.unregister <- true
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
	for {
		if (s.robotConn != nil) {
			buff := <-s.send
			msg := make([]byte, 1)
			msg[0] = byte(uint8(len(buff)))
			msg = append(msg, buff...)
			_, err := s.robotConn.Write(msg)
			if err != nil {
				if !<-s.unregister {
					s.robotConn.Close()
					s.unregister <- true
				}
				fmt.Printf("Error writing to robot: ", err)
				for <-s.unregister {
					time.Sleep(time.Duration(4) * time.Second)
				}
			}
		}
	}
}

func (s* Server) ReadData() {
	for {
		if s.robotConn != nil {
			buff := make([]byte, 1024)
			_, err := s.robotConn.Read(buff)
			if err != nil {
				fmt.Printf("Error reading from the robot: ", err)
				if !<-s.unregister {
					s.robotConn.Close()
					s.unregister <- true
				}
				for <-s.unregister {
					time.Sleep(time.Duration(4) * time.Second)
				}
			} else {
				s.recv <- buff
			}
		}
	}
}

func (s* Server) Serve() {
	go s.WriteData()
	go s.ReadData()
}
