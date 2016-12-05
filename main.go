package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"os/exec"
	"github.com/gorilla/websocket"

	b64 "encoding/base64"
)

var upgrader = websocket.Upgrader{}

func main() {

	http.Handle("/", http.FileServer(http.Dir("site/")))

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		fmt.Println("Client connection")
		go conn.WriteMessage(1, []byte(runCom("ls -lah")))
		go func(conn *websocket.Conn) {
			ch := time.Tick(3 * time.Second)
			for range ch {
				conn.WriteMessage(1, []byte(runCom("ls -lah")))
			}
		}(conn)
	})
	fmt.Println("Server started")
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server crashed")
}

func fileToBytes(f string) []byte {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return b
}

func runCom(cmd string) string {
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head,parts...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	return b64.StdEncoding.EncodeToString([]byte(out))
}
