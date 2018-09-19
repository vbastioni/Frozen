package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const bufsize = 128
const message = "Hello world!"

func run(port string) (err error) {
	var c net.Conn
	c, err = net.Dial("tcp4", port)
	if err != nil {
		return
	}
	buf := make([]byte, bufsize)
	copy(buf, message)
	if _, err = c.Write(buf); err != nil {
		return
	}
	buf = make([]byte, bufsize)
	if _, err = c.Read(buf); err != nil {
		return
	}
	log.Println(string(buf))
	return
}

func main() {
	port := os.Args[1]
	fmt.Println(port)
	if err := run(port); err != nil {
		log.Fatalln(err)
	}
}
