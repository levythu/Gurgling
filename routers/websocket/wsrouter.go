package websocket

import (
	"golang.org/x/net/websocket"
    . "github.com/levythu/gurgling"
)

type message struct {
    msgType byte    // 0: msg, 1: add, 2: leave, 100: error
    msgContent []byte
    msgSender *websocket.Conn

    err error
}

type WebsocketServerAsync struct {
    oriwsHandler websocket.Handler

    //catchPanic bool
    bufferSize int
    workers int
    msgChan chan *message

    msgHandler func(Session, []byte, error)
    errorHandler func(Session, error)
    connectHandler func(Session)
    leaveHandler func(Session)
}

func AWebsocketServerAsync(paraList ...int) *WebsocketServerAsync {
    var ret=&WebsocketServerAsync{}
    ret.msgChan=make(chan *message)
    if len(paraList)>=1 {
        ret.bufferSize=paraList[0]
    } else {
        ret.bufferSize=4096
    }

    if len(paraList)>=2 {
        ret.workers=paraList[1]
    } else {
        ret.workers=1
    }

    ret.oriwsHandler=func(c *websocket.Conn) {
        defer func() {
			err:=c.Close()
			if err!=nil {
				ret.msgChan<-&message {
                    msgType: 100,
                    msgContent: nil,
                    msgSender: c,
                    err: err,
                }
			}
            ret.msgChan<-&message {
                msgType: 2,
                msgContent: nil,
                msgSender: c,
                err: nil,
            }
		}()

        ret.msgChan<-&message {
            msgType: 1,
            msgContent: nil,
            msgSender: c,
            err: nil,
        }

        for {
            var buffer=make([]byte, ret.bufferSize)
            var readCount, err=c.Read(buffer)
            ret.msgChan<-&message {
                msgType: 0,
                msgContent: buffer[:readCount],
                msgSender: c,
                err: err,
            }
            if err!=nil {
                return
            }
        }
    }

    for i:=0; i<ret.workers; i++ {
        go ret.worker()
    }

    return ret
}

func (this *WebsocketServerAsync)OnConnect(f func(Session)) *WebsocketServerAsync {
    this.connectHandler=f;
    return this
}
func (this *WebsocketServerAsync)OnDisconnect(f func(Session)) *WebsocketServerAsync {
    this.leaveHandler=f;
    return this
}
func (this *WebsocketServerAsync)OnMessage(f func(Session, []byte, error)) *WebsocketServerAsync {
    this.msgHandler=f;
    return this
}
func (this *WebsocketServerAsync)OnError(f func(Session, error)) *WebsocketServerAsync {
    this.errorHandler=f;
    return this
}


func (this *WebsocketServerAsync)worker() {
    for msg:=range this.msgChan {
        var sender Session=&WebSocketSession {
            websocket.Conn: msg.msgSender,
        }
        if msg.msgType==0 {
            if this.msgHandler!=nil {
                this.msgHandler(sender, msg.msgContent, msg.err)
            }
        } else if msg.msgType==1 {
            if this.connectHandler!=nil {
                this.connectHandler(sender)
            }
        } else if msg.msgType==2 {
            if this.leaveHandler!=nil {
                this.leaveHandler(sender)
            }
        } else {
            if this.errorHandler!=nil {
                this.errorHandler(sender, msg.err)
            }
        }
    }
}
func (this *WebsocketServerAsync)Handler(req Request, res Response) (bool, Request, Response) {
    this.oriwsHandler.ServeHTTP(res.R(), req.R())

    return false, nil, nil
}
