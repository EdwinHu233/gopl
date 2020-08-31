package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		mustCopy(conn, os.Stdin)
		tcp_conn, ok := conn.(*net.TCPConn)
		if ok {
			tcp_conn.CloseWrite()
		}
		done <- struct{}{}
	}()
	mustCopy(os.Stdout, conn)
	<-done
	conn.Close()
}

func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}
