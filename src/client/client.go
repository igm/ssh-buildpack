package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"golang.org/x/net/websocket"
)

func shovel(con net.Conn) {
	origin := "https://ssh-sample.cfapps.io:4443/"
	url := "wss://ssh-sample.cfapps.io:4443/ssh"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		io.Copy(ws, con)
	}()
	io.Copy(con, ws)
}

func main() {
	l, err := net.Listen("tcp", ":2222")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		fmt.Println("connection accepted")
		if err != nil {
			fmt.Println(err)
		}
		go shovel(conn)
	}
}
