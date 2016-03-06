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
	c.send <- data
}

func (c connection) writeData(t Transport) {
	for {
		buff := <-c.send
		msg := make([]byte, 1)
		msg[0] = byte(uint8(len(buff)))
		msg = append(msg, buff...)
		_, err := c.conn.Write(msg)
		if err != nil {
			fmt.Printf("Error writing to robot: ", err)
			t.unregisterConn(c)
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func (c connection) readData(t Transport) {
	defer c.conn.Close()
	for {
		buff := make([]byte, 1024)
		_, err := c.conn.Read(buff)
		if err != nil {
			fmt.Printf("Error reading from the robot: ", err)
			t.unregisterConn(c)
			break
		}
		t.recieve(buff)
	}
}

func (c connection) handle(t Transport) {
	go c.readData(t)
	go c.writeData(t)
}
