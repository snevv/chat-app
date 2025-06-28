package ws

import "fmt"

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	defaultRoom := &Room{
		ID:      "abc123",
		Name:    "default",
		Clients: make(map[string]*Client),
	}
	return &Hub{
		Rooms: map[string]*Room{
			"abc123": defaultRoom,
		},
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 100),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case cli := <-hub.Register:
			if _, ok := hub.Rooms[cli.RoomID]; ok {
				hub.Rooms[cli.RoomID].Clients[cli.ID] = cli

				hub.Broadcast <- &Message{
					Content:  fmt.Sprintf("User %s has joined the chat", cli.Username),
					RoomID:   cli.RoomID,
					Username: cli.Username,
				}
			}

		case cli := <-hub.Unregister:
			if room, ok := hub.Rooms[cli.RoomID]; ok {
				if _, ok := room.Clients[cli.ID]; ok {
					if len(room.Clients) > 0 {
						hub.Broadcast <- &Message{
							Content:  fmt.Sprintf("User %s has left the chat", cli.Username),
							RoomID:   cli.RoomID,
							Username: cli.Username,
						}
					}

					delete(hub.Rooms[cli.RoomID].Clients, cli.ID)
					close(cli.MsgChan)
				}
			}

		case msg := <-hub.Broadcast:
			if room, ok := hub.Rooms[msg.RoomID]; ok {
				for _, cli := range room.Clients {
					cli.MsgChan <- msg
				}
			}
		}
	}
}
