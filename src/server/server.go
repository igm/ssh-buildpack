package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"

	"golang.org/x/net/websocket"
)

func main() {
	go func() {
		cmd := exec.Command("./sshd")
		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
		}
		cmd.Wait()
	}()
	http.Handle("/ssh", websocket.Handler(shovel))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func shovel(ws *websocket.Conn) {
	fmt.Println("new websocket client accepted")
	conn, err := net.Dial("tcp", "localhost:2200")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		io.Copy(ws, conn)
	}()
	io.Copy(conn, ws)
}

func handler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("ok"))
}
