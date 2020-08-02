package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.Int("port", 8000, "port at which this server will listen")

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			// failed to establish connection, e.g. connection aborted
			log.Print(err)
		} else {
			// otherwise, handle the connection
			go handleConn(conn)
		}
	}
}

func handleConn(conn net.Conn) {
	const INTERVAL = 1 * time.Second
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(INTERVAL)
	}
}
