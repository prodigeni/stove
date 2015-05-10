package main

import (
	"fmt"
	"github.com/HearthSim/hs-proto/go/bnet"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = 52525
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections
	sock, err := net.Listen(CONN_TYPE, fmt.Sprintf("%s:%d", CONN_HOST, CONN_PORT))
	if err != nil {
		fmt.Println("Error creating socket:", err.Error())
		os.Exit(1)
	}
	defer sock.Close()

	fmt.Printf("Listening on %s:%d ...\n", CONN_HOST, CONN_PORT)
	for {
		conn, err := sock.Accept()
		if err != nil {
			fmt.Println("Error on incoming connection:", err.Error())
			os.Exit(2)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)

	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Printf("%d bytes received: %q\n", reqLen, buf[:reqLen])
	resp := ""
	conn.Write([]byte(resp))
	conn.Close()
}
