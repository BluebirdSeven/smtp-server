package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("220 localhost Simple SMTP Server\r\n"))

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Connection closed:", err)
			return
		}

		input := strings.TrimSpace(string(buf[:n]))
		fmt.Println("Client:", input)

		switch {
		case strings.HasPrefix(input, "HELO"):
			conn.Write([]byte("250 Hello\r\n"))
		case strings.HasPrefix(input, "MAIL FROM:"):
			conn.Write([]byte("250 OK\r\n"))
		case strings.HasPrefix(input, "RCPT TO:"):
			conn.Write([]byte("250 OK\r\n"))
		case strings.HasPrefix(input, "DATA"):
			conn.Write([]byte("354 End data with <CR><LF>.<CR><LF>\r\n"))
		case input == ".":
			conn.Write([]byte("250 Message received\r\n"))
		case strings.HasPrefix(input, "QUIT"):
			conn.Write([]byte("221 Bye\r\n"))
			return
		default:
			conn.Write([]byte("500 Unrecognized command\r\n"))
		}
	}
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
