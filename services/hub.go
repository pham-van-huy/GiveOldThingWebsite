// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package services

type RoomMessage struct {
	room    string
	message string
}

type RoomClient struct {
	room   string
	client *Client
}

type ClientMessage struct {
	message []byte
	client  *Client
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool
	// Hold all room - clients
	rooms map[string][]*Client

	// Inbound messages from the clients.
	broadcast chan *ClientMessage

	broadcastRoom chan *RoomMessage
	// Register requests from the clients.
	register chan *Client

	registerRoom chan *RoomClient
	// Register leave room
	leaveRoom chan *RoomClient
	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:     make(chan *ClientMessage),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		clients:       make(map[*Client]bool),
		rooms:         make(map[string][]*Client),
		registerRoom:  make(chan *RoomClient),
		leaveRoom:     make(chan *RoomClient),
		broadcastRoom: make(chan *RoomMessage),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case registerRoom := <-h.registerRoom:
			registerRoom.client.rooms = append(registerRoom.client.rooms, registerRoom.room)
			if h.rooms[registerRoom.room] == nil {
				h.rooms[registerRoom.room] = []*Client{registerRoom.client}
			} else {
				h.rooms[registerRoom.room] = append(h.rooms[registerRoom.room], registerRoom.client)
			}
		case broadcastRoom := <-h.broadcastRoom:
			for _, client := range h.rooms[broadcastRoom.room] {
				select {
				case client.send <- []byte(broadcastRoom.message):
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case clientMessage := <-h.broadcast:
			client := clientMessage.client
			rooms := client.rooms
			if len(rooms) != 0 {
				for _, room := range rooms {
					for _, client := range h.rooms[room] {
						select {
						case client.send <- clientMessage.message:
						default:
							close(client.send)
							delete(h.clients, client)
						}
					}
				}
			} else {
				for client := range h.clients {
					select {
					case client.send <- clientMessage.message:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}
