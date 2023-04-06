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
	defer conn.Close()
	for {
		time.Sleep(time.Second * 3)
		buf := make([]byte, 1000)
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Println("request")
		requestInfo := string(buf[:n])
		fmt.Println(requestInfo)

		if requestInfo == `"close"` {
			fmt.Println("close connection")
			conn.Close()
			return
		}

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
}
