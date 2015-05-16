package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	ln, err := net.Listen("tcp", ":15440")
	if err != nil {
		fmt.Println("Error on listen: ", err)
		os.Exit(-1)
	}

	fmt.Println("Waiting for a connection via Accept")

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Error on accept: ", err)
		os.Exit(-1)
	}

	fmt.Println("Reading once from connection")

	var buf [1024]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("Error on read: ", err)
		os.Exit(-1)
	}

	fmt.Println("Client sent: ", string(buf[0:n]))
	conn.Close()
	fmt.Println("Exiting")

}
