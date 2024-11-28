package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprint(conn, "Error reading command: %s\n", err)
		return
	}

	// Splice Line by size 2
	parts := strings.SplitN(strings.TrimSpace(line), " ", 2)
	if len(parts) != 2 {
		fmt.Fprintf(conn, "Invalid command format. Excepted format: COMMAND:RESOURCE\n")
		return
	}

	command := parts[0]
	resource := parts[1]
	log.Printf("Received command: %s %s\n", command, resource)

	switch command {
	case "GET":
		handleGet(conn, resource)
	default:
		fmt.Fprintf(conn, "Unknown command: %s\n", command)
	}
}

func handleGet(conn net.Conn, resource string) {
	fmt.Fprintf(conn, "GET command received for resource: %s\n", resource)
}
