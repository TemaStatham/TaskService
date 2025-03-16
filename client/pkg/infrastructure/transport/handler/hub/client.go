package hub

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

type Client struct {
	RoomID uint // roomid => taskid
	UserID uint // userid => userid
	Conn   *websocket.Conn
	send   chan Message
	hub    *Hub
}

func NewClient(userID uint, taskID uint, conn *websocket.Conn, hub *Hub) *Client {
	return &Client{UserID: userID, RoomID: taskID, Conn: conn, send: make(chan Message, 256), hub: hub}
}

func (c *Client) Read() {
	defer func() {
		c.hub.unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		var msg Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			break
		}
		c.hub.broadcast <- msg
	}
}

func (c *Client) Write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			err := c.Conn.WriteJSON(message)
			if err != nil {
				fmt.Println("Ошибка отправки:", err)
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
