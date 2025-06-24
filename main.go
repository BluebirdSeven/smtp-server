package main

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintf(conn, "220 localhost Simple SMTP Server\r\n")
	// For production, you'd need full SMTP command handling (HELO, MAIL FROM, RCPT TO, DATA, QUIT, etc.)
}

func main() {
	ln, err := net.Listen("tcp", ":2525") // non-standard SMTP port for testing
	if err != nil {
		log.Fatalf("Failed to bind to port: %v", err)
	}
	log.Println("SMTP server started on port 2525")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Connection error: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}
