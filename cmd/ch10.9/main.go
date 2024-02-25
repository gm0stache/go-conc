package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
)

var validReqHeader, _ = regexp.Compile("GET (.+) HTTP/1.1\r\n")

func handleHTTPRequest(conn net.Conn) {
	defer conn.Close()
	buff := make([]byte, 1024)
	size, _ := conn.Read(buff)
	if !validReqHeader.Match(buff[:size]) {
		conn.Write([]byte("HTTP/1.1 500 Internal Server Error\r\n\r\n"))
		return
	}
	filename := validReqHeader.FindSubmatch(buff[:size])[1]
	filepath := fmt.Sprintf("./cmd/ch10.9/pages/%s", filename)
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n<html>Not Found</html>"))
		return
	}
	header := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %d\r\n\r\n", len(fileContent))
	conn.Write([]byte(header))
	conn.Write(fileContent)
}

// startHTTPWorkers enables explicitly limiting the amount of handler processes spawned.
func startHTTPWorkers(n int, incomingConnections <-chan net.Conn) {
	for i := 0; i < n; i++ {
		go func() {
			for conn := range incomingConnections {
				handleHTTPRequest(conn)
			}
		}()
	}
}

func main() {
	const workerCount int = 5
	incomingConns := make(chan net.Conn)
	startHTTPWorkers(workerCount, incomingConns)
	server, err := net.Listen("tcp", "localhost:8008")
	if err != nil {
		log.Fatalf("failed to start server on port 8008\n%+v", err)
		return
	}
	defer server.Close()
	for {
		conn, _ := server.Accept()
		select {
		case incomingConns <- conn:
		default:
			fmt.Println("server is busy")
			conn.Write([]byte("HTTP/1.1 429 Too Many Requests\r\n\r\n<html>Busy</html>"))
			conn.Close()
		}
	}
}
