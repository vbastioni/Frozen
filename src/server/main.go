package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/vbastioni/lib"
)

var (
	port = lib.KVArg{Key: "--port", Regexp: `^[0-9]{4}$`, Required: true}
)

const bufsize = 128

func handleClient(c net.Conn, chstr chan string, chcmd chan int) {
	// for {
	buf := make([]byte, bufsize)
	nr, err := c.Read(buf)
	if err != nil {
		return
	}
	data := string(buf[:nr])
	fmt.Printf("received %s on server\n", data)
	c.Write(buf)
	// }
}

func run(port string) (err error) {
	var sock net.Listener
	if sock, err = net.Listen("tcp4", ":"+port); err != nil {
		return
	}
	fmt.Printf("Server running on port ':%s'\n", port)
	var c net.Conn
	var chans []chan string
	chcmd := make(chan int)
	for {
		c, err = sock.Accept()
		if err != nil {
			return
		}
		chstr := make(chan string)
		chans = append(chans, chstr)
		go handleClient(c, chstr, chcmd)
	}
	return
}

func main() {
	args := []*lib.KVArg{
		&lib.KVArg{Key: "--port", Regexp: "\\d{4}", Required: true, Extended: true},
	}
	if err := lib.GetArgs(args); err != nil {
		fmt.Println(err)
		fmt.Print("usage: ./frozen_server\n")
		os.Exit(1)
	}
	switch s := args[0].GetValue().(type) {
	case string:
		if err := run(s); err != nil {
			log.Fatalln(err)
		}
	default:
		log.Fatalln("Server: fatal error")
	}
}
