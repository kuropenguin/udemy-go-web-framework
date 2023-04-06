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
	requestAndResponse(conn, "reques1")
	requestAndResponse(conn, "reques2")
	requestAndResponse(conn, "reques3")
	requestAndResponse(conn, "close")
}

func requestAndResponse(conn net.Conn, requestData string) {
	requestByteData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	conn.Write(requestByteData)
	responseData := make([]byte, 1000)
	n, err := conn.Read(responseData)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(responseData[:n]))
}
