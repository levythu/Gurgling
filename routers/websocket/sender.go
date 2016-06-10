package websocket

import (
	. "golang.org/x/net/websocket"
    "net/http"
)

type Session interface {
    Close() error
    Write(msg []byte) (n int, err error)
    Request() *http.Request
    R() *Conn
}


type WebSocketSession struct {
    *Conn
}
func (this *WebSocketSession)R() *Conn {
    return this.Conn
}
