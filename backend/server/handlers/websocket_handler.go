// package handlers

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/karasuneo/websocket/backend/server/models"

// 	"github.com/gorilla/websocket"
// 	// "golang.org/x/net/websocket"
// )

// type WebsocketHandler struct {
// }

// func NewWebsocketHandler(hub *models.Hub) *WebsocketHandler {
// 	return &WebsocketHandler{}
// }

// func HandleWebSocket(c *gin.Context) {
// 	ug := &websocket.Upgrader{
// 		CheckOrigin: func(r *http.Request) bool {
// 			return true
// 		},
// 	}
// 	ws, err := ug.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = ug.Message.Send(ws, fmt.Sprintf("Server: \"%s\" received!", msg))

// 	// client := models.NewClient(ws)
// 	// go client.ReadLoop(h.hub.BroadcastCh, h.hub.UnRegisterCh)
// 	// go client.WriteLoop()
// 	// h.hub.RegisterCh <- client
// }

// func HandleWebSocket(c *gin.Context) {
// 	log.Println("Serving at localhost:8000...")
// 	websocket.Handler(func(ws *websocket.Conn) {
// 		defer ws.Close()

// 		// 初回のメッセージを送信
// 		err := websocket.Message.Send(ws, "Server: Hello, Next.js!")
// 		if err != nil {
// 			// c.Logger().Error(err)
// 		}

// 		for {
// 			// Client からのメッセージを読み込む
// 			msg := ""
// 			err := websocket.Message.Receive(ws, &msg)
// 			if err != nil {
// 				if err.Error() == "EOF" {
// 					log.Println(fmt.Errorf("read %s", err))
// 					break
// 				}
// 				log.Println(fmt.Errorf("read %s", err))
// 				// c.Logger().Error(err)
// 			}

// 			// Client からのメッセージを元に返すメッセージを作成し送信する
// 			err = websocket.Message.Send(ws, fmt.Sprintf("Server: \"%s\" received!", msg))
// 			if err != nil {
// 				// c.Logger().Error(err)
// 			}
// 		}
// 	}).ServeHTTP(c.Writer, c.Request)
// 	// return nil
// }

package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/karasuneo/websocket/backend/server/models"
)

type WebsocketHandler struct {
	hub *models.Hub
}

func NewWebsocketHandler(hub *models.Hub) *WebsocketHandler {
	return &WebsocketHandler{
		hub: hub,
	}
}

func (h *WebsocketHandler) Handle(c *gin.Context) {
	ug := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := ug.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := models.NewClient(ws)
	go client.ReadLoop(h.hub.BroadcastCh, h.hub.UnRegisterCh)
	go client.WriteLoop()
	h.hub.RegisterCh <- client
}
