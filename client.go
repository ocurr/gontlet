package gontlet

import (
	"fmt"
	"net"
	"time"
)

type Client struct {
	conn    connection
	send    chan []byte
	recv    chan []byte
	address string
}

func newClient(address string) *Client {
	c := &Client{
		send:    make(chan []byte, 1),
		recv:    make(chan []byte, 1),
		address: address,
	}
	c.connect()
	return c
}

func (c *Client) connect() {
	addr, err := net.ResolveTCPAddr("tcp", c.address)
	if err != nil {
		fmt.Println("Error on client address:", err)
	}
	for conn, err := net.DialTCP("tcp", nil, addr); ; {
		if err != nil {
			time.Sleep(4 * time.Second)
		} else {
			c.conn = connection{conn: conn, send: make(chan []byte, 1)}
		}
	}
}

func (c *Client) sendOutgoing(data []byte) {
	c.send <- data
}

func (c *Client) getIncoming() []byte {
	return <-c.recv
}

func (c *Client) serve() {
	for {
		select {
		case msg := <-c.send:
			c.conn.sendData(msg)
		}
	}
}
