package main

type room struct {
	// forward is a channel that holds incoming messages
	// that should be forwarded to the other client
	forward chan []byte
	// join a channel for clients wishig to join the room.
	join chan *client
	// leave is a channel for clients whising to leave the room.
	leave chan *client
	// clients hold all current clients in this room.
	clients map[*client]bool
}