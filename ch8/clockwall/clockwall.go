package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type Server struct {
	Name string
	Conn net.Conn
}

func (s *Server) ReadLine() (string, error) {
	reader := bufio.NewReader(s.Conn)
	line, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s disconnected: %v\n", s.Name, err)
		return "", err
	}
	return fmt.Sprintf("%s: %s", s.Name, string(line)), nil
}

func main() {
	flag.Parse()
	args := flag.Args()
	servers := make([]Server, len(args))
	for i, arg := range args {
		pos_equal := strings.Index(arg, "=")
		name := arg[:pos_equal]
		addr := arg[pos_equal+1:]
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}
		servers[i] = Server{name, conn}
	}
	for {
		for _, s := range servers {
			line, err := s.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s", line)
		}
	}
}
