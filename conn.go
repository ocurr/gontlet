package gontlet

import (
	"fmt"
	"net"
	"time"
)

type connection struct {
	conn net.Conn
	send chan []byte
}

func (c connection) sendData(data []byte) {
	c.send <-data
}

func (c connection) writeData(s *Server) {
	for {
		buff := <-c.send
		msg := make([]byte, 1)
		msg[0] = byte(uint8(len(buff)))
		msg = append(msg, buff...)
		_, err := c.conn.Write(msg)
		if err != nil {
			fmt.Printf("Error writing to robot: ", err)
			s.unregister <- c
			break
		}
		time.Sleep(100*time.Millisecond)
	}
}

func (c connection) readData(s *Server) {
	defer c.conn.Close()
	for {
		buff := make([]byte, 1024)
		_, err := c.conn.Read(buff)
		if err != nil {
			fmt.Printf("Error reading from the robot: ", err)
			s.unregister <- c
			break
		}
		s.recv <- buff
	}
}

func (c connection) handle(s *Server) {
	go c.readData(s)
	go c.writeData(s)
}
