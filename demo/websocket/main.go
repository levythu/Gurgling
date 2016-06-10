package main

import (
    . "github.com/levythu/gurgling"
    ws "github.com/levythu/gurgling/routers/websocket"
)

func main() {
    var r=ARouter()
    var echo=ws.AWebsocketServerAsync().OnConnect(func(c ws.Session) {
        c.Write([]byte("ALOHA!"))
    }).OnMessage(func(c ws.Session, content []byte, err error) {
        c.Write(content)
    })

    var hello=ws.AWebsocketServer(func(c *ws.WebSocketSession) {
        c.Write([]byte("ALOHA!"))
        c.Close()
    })

    r.Use("/echo", echo).Use("/hello", hello)

    r.Launch(":8080")
}
