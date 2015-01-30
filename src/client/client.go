package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"golang.org/x/net/websocket"
)

func shovel(con net.Conn, appUrl string) {
	origin := fmt.Sprintf("http://%s", appUrl)
	url := fmt.Sprintf("wss://%s:4443/ssh", appUrl)
	// url := "wss://ssh-sample.cfapps.io:4443/ssh"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ssh session initiated.")
	go func() {
		io.Copy(ws, con)
	}()
	io.Copy(con, ws)
	fmt.Println("ssh session closed.")
}

func main() {
	flag.Usage = Usage
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		return
	}
	l, err := net.Listen("tcp", ":2222")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		fmt.Println("connection accepted, connecting to: ", args[0])
		if err != nil {
			fmt.Println(err)
		}
		go shovel(conn, args[0])
	}
}

func Usage() {
	fmt.Println(`
usage:
  client APP_URL
`)
}
