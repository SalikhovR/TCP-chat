package main

import (
	"fmt"
	"net"
)

func main() {
	var count uint32
	for {
		conn, err := net.Dial("tcp", "localhost:1234")
		if err != nil {
			panic(err)
		}

		defer conn.Close()
		if count == 0 {
			fmt.Println("Successfully connected")
		}
		count += 1
		var sms string
		fmt.Print("He:")
		fmt.Scan(&sms)
		b := make([]byte, 100)
		if _, err = conn.Write([]byte(sms)); err != nil {
			panic(err)
		}

		if _, err = conn.Read(b); err != nil {
			panic(err)
		}
		if string(b) != "quit" || string(b) != "Quit" {
			fmt.Println(string(b))
		}

	}
}
