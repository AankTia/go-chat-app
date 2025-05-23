package main

import (
	"time"

	"github.com/gorilla/websocket"
)

// client represent a single chat user
type client struct {
	// socket is the web for this client
	socket *websocket.Conn
	// send is a channel on which message are sent
	send chan *message
	// room is the room this client is chating in
	room *room
	// userData holds information abut the user
	userData map[string]interface{}
}

func (c *client) read() {
	defer c.socket.Close()

	for {
		var msg *message

		if err := c.socket.ReadJSON(&msg); err != nil {
			return
		}

		msg.When = time.Now()
		msg.Name = c.userData["name"].(string)
		msg.AvatarURL, _ = c.room.avatar.GetAvatarURL(c)
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			break
		}
	}
}
