package ws

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[string]*Room),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case cli := <-hub.Register:
			if _, ok := hub.Rooms[cli.RoomID]; ok {
				room := hub.Rooms[cli.RoomID]

				if _, ok := room.Clients[cli.ID]; ok {
					room.Clients[cli.ID] = cli
				}
			}

		// case cli := <-hub.Unregister:
		// 	if _, ok := hub.Rooms[cli.RoomID]; ok {
		// 		if _, ok := 
		// 	}
		// }
	}
}
