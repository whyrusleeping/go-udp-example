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
	// to do udp, you could do it just like the in class demo,
	// but to have more control over the udp connection, i'm going to
	// show a more udp specific way to do things
	uaddr, err := net.ResolveUDPAddr("udp", ":9000")
	if err != nil {
		panic(err)
	}

	list, err := net.ListenUDP("udp", uaddr)
	if err != nil {
		panic(err)
	}

	// loop and read packets
	for {
		// make a buffer of size 4096 to read data into
		mesbuf := make([]byte, 4096)

		nread, from, err := list.ReadFromUDP(mesbuf)
		if err != nil {
			// normally you dont want to panic,
			// this is just because i'm lazy
			panic(err)
		}

		// truncate buf to size of read message
		mesbuf = mesbuf[:nread]

		fmt.Printf("got message of size %d bytes from %s\n", nread, from.String())

		// unmarshal received data into a message struct
		var message MyMessage
		err = json.Unmarshal(mesbuf, &message)
		if err != nil {
			panic(err)
		}

		fmt.Println("got a message!")
		fmt.Printf("seq: %d, mes: %s\n", message.Sequence, message.Message)
	}
}
