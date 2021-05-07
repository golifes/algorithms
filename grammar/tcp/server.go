package main

import (
	"fmt"
	"io"
	"net"
)

const (
	IpAddrPort = "127.0.0.1:8000"
)

func main() {
	listener, err := net.Listen("tcp", IpAddrPort)

	defer listener.Close()
	if err != nil {
		fmt.Printf("server listen error [%+v]", err)
		return
	}
	for {
		conn, err := listener.Accept()
		//defer conn.Close()
		if err != nil {
			fmt.Printf("server accept error [%+v]", err)
			return
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		read, err := conn.Read(buf)

		if err != nil || err == io.EOF {
			fmt.Println("process error", err)
			return
		}
		fmt.Println(string(buf[:read]))
	}
}
