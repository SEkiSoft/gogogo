// Copyright (c) 2016 sekisoft
// See License.txt

package api

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/sekisoft/gogogo/model"
)

const (
	WRITE_WAIT  = 5 * time.Second
	PONG_WAIT   = 10 * time.Second
	PING_PERIOD = (PONG_WAIT * 9) / 10
	MAX_SIZE    = 512
	REDIS_WAIT  = 30 * time.Second
)

type WebConn struct {
	WebSocket *websocket.Conn
	Send      chan model.WebSocketMessage
	TokenId   string
	PlayerId  string
}

func NewWebConn(s *Session, ws *websocket.Conn) *WebConn {
	return &WebConn{
		Send:      make(chan model.WebSocketMessage, 64),
		WebSocket: ws,
		PlayerId:  s.Token.PlayerId,
		TokenId:   s.Token.Id,
	}
}

func (c *WebConn) readPump() {
	defer func() {
		webHub.Unregister(c)
		c.WebSocket.Close()
	}()
	c.WebSocket.SetReadLimit(MAX_SIZE)
	c.WebSocket.SetReadDeadline(time.Now().Add(PONG_WAIT))
	c.WebSocket.SetPongHandler(func(string) error {
		c.WebSocket.SetReadDeadline(time.Now().Add(PONG_WAIT))
		return nil
	})

	for {
		var req model.WebSocketRequest
		if err := c.WebSocket.ReadJSON(&req); err != nil {
			return
		} else {
			BaseRoutes.WebSocket.ServeWebSocket(c, &req)
		}
	}
}

func (c *WebConn) writePump() {
	ticker := time.NewTicker(PING_PERIOD)

	defer func() {
		ticker.Stop()
		c.WebSocket.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				c.WebSocket.SetWriteDeadline(time.Now().Add(WRITE_WAIT))
				c.WebSocket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.WebSocket.SetWriteDeadline(time.Now().Add(WRITE_WAIT))
			if err := c.WebSocket.WriteJSON(msg); err != nil {
				return
			}

		case <-ticker.C:
			c.WebSocket.SetWriteDeadline(time.Now().Add(WRITE_WAIT))
			if err := c.WebSocket.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
