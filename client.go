package gontlet

import (
	"errors"
	"fmt"
	"net"
	"time"
)

type Client struct {
	conn       connection
	send       chan []byte
	recv       chan []byte
	unregister chan connection
	address    string
}

func newClient(address string) *Client {
	c := &Client{
		send:       make(chan []byte, 1),
		recv:       make(chan []byte, 1),
		unregister: make(chan connection, 1),
		address:    address,
	}
	c.connect()
	return c
}

func (c *Client) connect() {
	var err error
	addr, err := net.ResolveTCPAddr("tcp", c.address)
	if err != nil {
		fmt.Println("Error on client address:", err)
	}
	var conn net.Conn
	err = errors.New("dummy error")
	for err != nil {
		conn, err = net.DialTCP("tcp", nil, addr)
		fmt.Println("unable to find a connection")
		time.Sleep(4 * time.Second)
	}
	fmt.Println("found connection")
	c.conn = connection{conn: conn, send: make(chan []byte, 1)}
	go c.conn.handle(c)
}

func (c *Client) sendOutgoing(data []byte) {
	c.send <- data
}

func (c *Client) getIncoming() []byte {
	return <-c.recv
}

func (c *Client) recieve(buff []byte) {
	c.recv <- buff
}

func (c *Client) unregisterConn(conn connection) {
	c.unregister <- conn
}

func (c *Client) serve() {
	for {
		select {
		case <-c.unregister:
			c.connect()
		case msg := <-c.send:
			c.conn.sendData(msg)
		}
	}
}
