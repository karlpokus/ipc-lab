package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
)

const socketPath = "/tmp/test.sock"

func main() {
	//b, err := netDial()
	b, err := pushhttp()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Received: %s", b)
}

func pushhttp() ([]byte, error) {
	c := http.Client{
		Transport: &http.Transport{
			// create custom dialer?
			// try: DialTLSContext
			//
			// unecrypted connection
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", socketPath)
			},
		},
	}
	res, err := c.Get("http://unix" + "/foo")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

func netDial() ([]byte, error) {
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	_, err = conn.Write([]byte("Hello from client!"))
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	return buffer[:n], nil
}
