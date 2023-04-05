package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	time.Sleep(time.Second * 5)
	defer conn.Close()
	buf := make([]byte, 1000)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println("request")
	fmt.Println(string(buf[:n]))

	responseData := "response"
	responseByteData, err := json.Marshal(responseData)
	if err != nil {
		panic(err)
	}
	_, err = conn.Write(responseByteData)
	if err != nil {
		panic(err)
	}
}
