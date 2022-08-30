package main

import (
	"fmt"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}

	defer lis.Close()

	println("Waiting for connections")
	var count uint32
	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}
		if count == 0 {
			fmt.Println("Successfully accepted connection")
		}
		count += 1
		defer conn.Close()

		msg := make([]byte, 100)
		if _, err = conn.Read(msg); err != nil {
			panic(err)
		}
		if string(msg) != "quit" || string(msg) != "Quit" {
			fmt.Println(string(msg))
		}
		//		fmt.Println(("HELLO" + " " + string(msg)))
		//		m := []byte("Hi!" + string(msg))
		var sms string
		fmt.Print("You:")
		fmt.Scan(&sms)
		if _, err = conn.Write([]byte(sms)); err != nil {
			panic(err)
		}
	}
}
