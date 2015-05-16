package main

import (
	"fmt"
	"net"
	"os"
)

func handleConn(coon net.Conn) {
	fmt.Printfln("Reading once from connection")
	var buf [1024]byte
	n, err := coon.Read(buf[:])
	if err != nil {
		fmt.Println("Error on read: ", err)
		os.Exit(-1)
	}

	fmt.Println("Client sent:  ", string(buf[:n]))
	coon.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":15440")
	if err != nil {
		fmt.Println("Error on listen: ", err)
	}

	for {
		fmt.Println("Waiting for a connection via Accept")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error on accept: ", err)
			os.Exit(-1)
		}
		go handleConn(conn)
	}
	fmt.Println("Exiting")
}
