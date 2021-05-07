package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("client dial error [%+v]", err)
		return
	}

	defer conn.Close()

	conn.Write([]byte("hello"))
}
