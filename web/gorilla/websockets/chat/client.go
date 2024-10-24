package main

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	// maximum time allowed to write a message to a peer
	writeWait = 10 * time.Second

	// maximum time allowed to read the next pong message from a peer
	pongWait = 60 * time.Second

	// send pings to peers with this period
	// must be less than pongWait - why?
	pingPeriod = (pongWait * 9) / 10

	// maximum message size allowed from peer
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

// A client is the middleman between a websocket connection and the hub
type Client struct {
	hub *Hub 

	// The websocket connection
	conn *websocket.Conn

	// Buffered channel of outbound messages 
	send chan []byte
}

// pump messages from the websocket conn to the hub
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c 
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string)error {c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		// todo:
		// read a message from the connection
		// and  then write it to the hub's broadcast channel
	}
	
}

// pump messages from the hub to the websocket conn
func (c *Client) writePump() {

}
