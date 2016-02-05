package gontlet

import (
	"fmt"
	"net"
)

type connection struct {
	conn net.Conn
	send chan []byte
}

func (c connection) SendData(data []byte) {
	fmt.Println("Connection Sending Message: ", string(data))
	c.send <-data
}

func (c connection) WriteData(s *Server) {
	for {
		buff := <-c.send
		fmt.Println("Goroutine Sending Message: ", string(buff))
		msg := make([]byte, 1)
		msg[0] = byte(uint8(len(buff)))
		msg = append(msg, buff...)
		_, err := c.conn.Write(msg)
		if err != nil {
			fmt.Printf("Error writing to robot: ", err)
			s.unregister <- c
			break
		}
	}
}

func (c connection) ReadData(s *Server) {
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

func (c connection) Handle(s *Server) {
	go c.ReadData(s)
	go c.WriteData(s)
}
