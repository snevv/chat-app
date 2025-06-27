package ws

import "github.com/gorilla/websocket"

type Client struct {
	Conn     *websocket.Conn
	MsgChan  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}
