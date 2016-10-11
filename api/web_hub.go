// Copyright (c) 2016 sekisoft
// See License.txt

package api

import (
	"fmt"

	l4g "github.com/alecthomas/log4go"

	"github.com/sekisoft/gogogo/model"
)

type WebHub struct {
	connections map[*WebConn]bool
	register    chan *WebConn
	unregister  chan *WebConn
	broadcast   chan *model.WebSocketEvent
	stop        chan string
}

var webHub = &WebHub{
	register:    make(chan *WebConn),
	unregister:  make(chan *WebConn),
	connections: make(map[*WebConn]bool),
	broadcast:   make(chan *model.WebSocketEvent),
	stop:        make(chan string),
}

func Publish(message *model.WebSocketEvent) {
	webHub.Broadcast(message)
}

func (h *WebHub) Register(webConn *WebConn) {
	h.register <- webConn

	msg := model.NewWebSocketEvent(webConn.PlayerId, "", model.WEBSOCKET_EVENT_HELLO_WORLD)
	msg.Add("server_version", fmt.Sprintf("%v", model.CurrentVersion))
	go Publish(msg)
}

func (h *WebHub) Unregister(webConn *WebConn) {
	h.unregister <- webConn
}

func (h *WebHub) Broadcast(message *model.WebSocketEvent) {
	if message != nil {
		h.broadcast <- message
	}
}

func (h *WebHub) Stop() {
	h.stop <- "all"
}

func (h *WebHub) Start() {
	go func() {
		for {
			select {
			case webCon := <-h.register:
				h.connections[webCon] = true

			case webCon := <-h.unregister:
				if _, ok := h.connections[webCon]; ok {
					delete(h.connections, webCon)
					close(webCon.Send)
				}

			case msg := <-h.broadcast:
				for webCon := range h.connections {
					if shouldSendEvent(webCon, msg) {
						select {
						case webCon.Send <- msg:
						default:
							close(webCon.Send)
							delete(h.connections, webCon)
						}
					}
				}

			case s := <-h.stop:
				l4g.Debug("Websocket stopped: %s", s)

				for webCon := range h.connections {
					webCon.WebSocket.Close()
				}

				return
			}
		}
	}()
}

func shouldSendEvent(webCon *WebConn, msg *model.WebSocketEvent) bool {
	return true
}
