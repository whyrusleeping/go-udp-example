package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type MyMessage struct {
	Sequence int
	Message  string
}

func main() {
	target, err := net.ResolveUDPAddr("udp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	// pass nil for the 'local address'
	con, err := net.DialUDP("udp", nil, target)
	if err != nil {
		panic(err)
	}

	message := &MyMessage{
		Sequence: 17,
		Message:  "Hello World",
	}

	data, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	n, err := con.Write(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("wrote %d bytes!\n", n)
}
