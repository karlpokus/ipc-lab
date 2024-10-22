package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

const socketPath = "/tmp/test.sock"

func main() {
	// cleanup
	err := os.Remove(socketPath)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}
	s := http.Server{
		Handler: http.FileServer(http.Dir("/tmp")),
	}
	log.Println("listening on", socketPath)
	l, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	s.Serve(l)

	// for {
	// 	conn, err := l.Accept()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		return
	// 	}
	// 	go handler(conn)
	// }
}

func handler(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("conn read error: %v", err)
		return
	}
	log.Printf("Received: %s", buffer[:n])
	conn.Write([]byte("Hello from server!"))
}
