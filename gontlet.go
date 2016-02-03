package gontlet

import (
	"fmt"
    "net"
    "log"
)

func handleConnection(conn net.Conn) {
	for {
		defer conn.Close()
		var message string
		n, err := fmt.Scanf("%s",&message)
		if err != nil && err.Error() != "unexpected newline" {
			fmt.Println("input finished", err)
			break
		}

		if n == -1 {
			fmt.Println("the impossible happened")
		}

		if message == "" {
			fmt.Println("please input a message")
		} else if message == "/quit" {
			conn.Write([]byte(message))
			fmt.Println("closing connection")
			conn.Close()
			log.Fatal()
			return
		} else {
			buff := []byte(message)
			msg := make([]byte, 1)
			msg[0] = byte(uint8(len(buff)))
			fmt.Printf("Sending message out (length %d)\n", int(msg[0]))
			msg = append(msg, buff...)
			_, err = conn.Write(msg);
			if message[0] != 0 {
				if err != nil {
					fmt.Printf("error writing to connection", err)
					conn.Close()
					return
				}
			}
		}
	}
}

func main() {
    l, err := net.Listen("tcp", "localhost:8081")
    if err != nil {
        log.Fatal(err)
    }
    defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("accept failed", err)
		}

		go handleConnection(conn)

	}
}
