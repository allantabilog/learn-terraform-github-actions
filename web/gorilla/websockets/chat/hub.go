package main

type Hub struct {

	// map of registered clients
	clients map[*Client]bool

	// messages from clients
	broadcast chan []byte 

	// register requests from clients
	register chan *Client 

	// un-register requests from clients
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast: make(chan []byte),
		register: make(chan *Client),
		unregister: make(chan *Client),
		clients: make(map[*Client]bool),
	}
}


func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true //no security here huh
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				// why does it have to be in a select block?
				select {
				case client.send <- message:
				default: 
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
