package main

import (
	"fmt"
	"net"
	"os"
	// Uncomment this block to pass the first stage
	// "net"
	// "os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	var conn net.Conn
	for {
		conn, err = l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	var (
		_   *string
		err error
	)
	for {
		_, err = ReadString(conn)
		if err == nil {
			err = WriteString(conn, "+PONG\r\n")
		}

		if err != nil {
			fmt.Println("An error occurred when handling conn: ", err.Error())
			break
		}
	}
}

func ReadString(conn net.Conn) (*string, error) {
	var (
		n   int
		p   int
		err error
	)
	const bufSize = 2048
	var buf [bufSize]byte
	for p = 0; p < bufSize && (p < 1 || buf[p-1] != '\n'); p += n {
		n, err = conn.Read(buf[p:])
		if err != nil {
			return nil, err
		}
	}
	str := string(buf[:p])
	return &str, nil
}

func WriteString(conn net.Conn, str string) error {
	var (
		n   int
		p   int
		err error
	)
	buf := []byte(str)
	for p = 0; p < len(buf); p += n {
		n, err = conn.Write(buf[p:])
		if err != nil {
			return err
		}
	}
	return nil
}
