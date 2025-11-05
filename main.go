package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

// --- 허용 Origin 목록 ---
var allowedOrigins = []string{
	"http://localhost",
	"https://localhost",
	"https://onair.gamecats.co",
	"ws://localhost",
	"wss://localhost",
	"wss://onair.gamecats.co",
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		if origin == "" {
			log.Println("[WARN] No Origin header (allowing by default for OBS)")
			return true
		}
		for _, o := range allowedOrigins {
			if strings.HasPrefix(origin, o) {
				return true
			}
		}
		log.Printf("[BLOCKED] Origin not allowed: %s\n", origin)
		return false
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)

func main() {
	defaultPort := "8762"
	port := flag.String("port", defaultPort, "Port number to listen on")
	flag.Parse()

	if len(os.Args) == 2 && os.Args[1][0] != '-' {
		*port = os.Args[1]
	}

	addr := fmt.Sprintf(":%s", *port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("❌ %s 포트가 이미 사용 중입니다.\n", *port)
		fmt.Print("실행 중인 프로그램을 종료하거나 다른 포트를 명시해서 실행할 수 있습니다.\n\n")
		fmt.Print("> relay.exe --port PORT\n\n")
		fmt.Print("엔터 키를 누르면 종료합니다...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return
	}
	defer listener.Close()

	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	fmt.Printf("✅ Local relay server running at ws://localhost:%s/ws\n", *port)
	fmt.Printf("   Allowed origins: %v\n", allowedOrigins)

	log.Fatal(http.Serve(listener, nil))
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("[ERROR] Upgrade error:", err)
		return
	}
	defer ws.Close()

	clients[ws] = true
	log.Printf("[CONNECT] Client connected (%d active)\n", len(clients))

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			delete(clients, ws)
			log.Printf("[DISCONNECT] Client left (%d active)\n", len(clients))
			break
		}
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		log.Printf("[BROADCAST] Message to %d clients : %s\n", len(clients), string(msg))
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				client.Close()
				delete(clients, client)
				log.Println("[WARN] Failed to send message, client removed")
			}
		}
	}
}
