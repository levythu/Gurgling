package gurgling

import (
    "net/http"
    "sync"
    "io"
    . "github.com/levythu/gurgling/definition"
)

type Response interface {
    // Quick send with code 200. While done, any other operation except Write is not allowed anymore.
    // However, due to the framework of net/http, the response will not be closed until
    // the function returns. So it is suggested to return immediately.
    Send(string) error

    // Set headers.
    Set(string, string) error

    // Get the value in the headers. If nothing, returns "".
    Get(string) string

    // Write data to response body. It allows any corresponding operation.
    Write([]byte) (int, error)

    // While done, any other operation except Write is not allowed anymore.
    Status(string, int) error

    // get the Original resonse, only use it for advanced purpose
    R() http.ResponseWriter

    // extra use for midwares. Most of time the value is a function.
    F() map[string]Tout
}
func NewResponse(w http.ResponseWriter) Response {
    return &OriResponse{
        r: w,
        haveSent: false,
        lock: &sync.Mutex{},
        f: make(map[string]Tout),
    }
}

type OriResponse struct {
    // the Original resonse, only use it for advanced purpose
    r http.ResponseWriter
    // to guarantee the send action is only triggered once
    haveSent bool
    f map[string]Tout

    lock *sync.Mutex
}

func (this *OriResponse)Send(content string) error {
    return this.Status(content, 200)
}
func (this *OriResponse)Status(content string, code int) error {
    this.lock.Lock()
    defer this.lock.Unlock()

    if (this.haveSent) {
        return RES_HEAD_ALREADY_SENT
    }
    this.r.Header().Set("Content-Type", "text/plain; charset=utf-8")
    this.r.WriteHeader(code)
    this.haveSent=true
    _, err:=io.WriteString(this.r, content)

    return err
}
func (this *OriResponse)Set(key string, val string) error {
    this.lock.Lock()
    defer this.lock.Unlock()

    if (this.haveSent) {
        return RES_HEAD_ALREADY_SENT
    }
    this.r.Header().Set(key, val)
    return nil
}
func (this *OriResponse)Get(key string) string {
    return this.r.Header().Get(key)
}
func (this *OriResponse)Write(content []byte) (int, error) {
    return this.r.Write(content)
}
func (this *OriResponse)R() http.ResponseWriter {
    return this.r
}
func (this *OriResponse)F() map[string]Tout {
    return this.f
}
