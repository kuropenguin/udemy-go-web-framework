package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", conn)
	fmt.Println("")
	requestData := "request"
	requestByteData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("before")
	conn.Write(requestByteData)
	responseData := make([]byte, 1000)
	n, err := conn.Read(responseData)
	fmt.Println("after")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(responseData[:n]))
}
