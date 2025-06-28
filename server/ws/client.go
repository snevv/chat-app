package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	MsgChan  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

func (cli *Client) WritePump() {
	defer cli.Conn.Close()

	for {
		msg, ok := <-cli.MsgChan
		if !ok {
			break
		}

		cli.Conn.WriteJSON(msg)
	}
}

func (cli *Client) ReadPump(hub *Hub) {
	defer func() {
		hub.Unregister <- cli
		cli.Conn.Close()
	}()

	for {
		_, m, err := cli.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &Message{
			Content:  string(m),
			RoomID:   cli.RoomID,
			Username: cli.Username,
		}

		hub.Broadcast <- msg
	}
}
