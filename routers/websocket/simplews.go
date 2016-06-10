package websocket

import (
	"golang.org/x/net/websocket"
    . "github.com/levythu/gurgling"
)

type WebsocketServer struct {
    oriwsHandler websocket.Handler
}

func AWebsocketServer(SessionHandler func(*WebSocketSession)) *WebsocketServer {
    var ret=&WebsocketServer{}

    ret.oriwsHandler=func(c *websocket.Conn) {
        SessionHandler(&WebSocketSession {
            websocket.Conn: c,
        })
    }

    return ret
}

func (this *WebsocketServer)Handler(req Request, res Response) (bool, Request, Response) {
    this.oriwsHandler.ServeHTTP(res.R(), req.R())

    return false, nil, nil
}
