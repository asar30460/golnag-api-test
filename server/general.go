package server

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Server struct {
	clients   map[int]*Client
	broadcast chan string
}

type Client struct {
	clientID int // 模擬使用者ID
	conn     *websocket.Conn
	receive  chan string // 每個使用者只需要有接收通道，而發送是發給伺服器，讓伺服器決定這筆訊息要讓哪些使用者接收
}

func NewGeneralServer() *Server {
	return &Server{
		clients:   make(map[int]*Client),
		broadcast: make(chan string, 64),
	}
}

func (s *Server) HandleWS(ctx *gin.Context) {
	client_id++
	log.Printf("A new client %d is trying to connect...", client_id)

	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Fatalln("Error upgrading connection:", err)
		return
	}
	defer ws.Close()

	// 每個使用者都有屬於自己的傳送及接收通道
	client := &Client{
		clientID: client_id,
		conn:     ws,
		receive:  make(chan string, 64),
	}
	s.clients[client_id] = client

	// 啟動 Goroutine 來處理接收訊息
	go func() {
		for {
			msg := <-client.receive
			// 寫入接收到的訊息至客戶端
			if err := ws.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				log.Println("Error writing message:", err)
				return
			}
			// log.Printf("Message sent to client %d: %s", client.clientID, msg)
		}
	}()

	// 處理從 WebSocket 接收到的訊息
	for {
		_, msg, err := ws.ReadMessage() // 讀取使用者發送的訊息
		if err != nil {
			log.Printf("Error reading message from client %d: %v", client.clientID, err)
			break
		}
		// log.Printf("Message received from client %d: %s", client.clientID, msg)

		match, _ := regexp.Compile(`{"owner":.*?}`)

		if !match.MatchString(string(msg)) {
			log.Println("No match found")
			continue
		}

		matchStr := match.FindString(string(msg))
		fixedStr := strings.Replace(matchStr, `"password":"***":false`, `"password":"***"`, -1)
		log.Println("String matched:", fixedStr)

		var result Object
		err = json.Unmarshal([]byte(fixedStr), &result)
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("result:%+v", result)

		// 將訊息發送到廣播通道
		s.broadcast <- string(msg)
	}
}

func (s *Server) BroadcastMessage() {
	for {
		msg := <-s.broadcast
		// log.Printf("Broadcasting message: %s", msg)
		for _, client := range s.clients {
			client.receive <- msg
		}
	}
}
