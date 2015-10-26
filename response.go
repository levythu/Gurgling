package gurgling

import (
    "net/http"
    "sync"
    "io"
    . "github.com/levythu/gurgling/definition"
    "github.com/levythu/gurgling/encoding"
    fp "path/filepath"
    MIME "mime"
    "os"
    "encoding/json"
)

// Depended by: gurgling/midwares/analyzer
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

    // Send files without any extra headers except contenttype and encrypt.
    // if contenttype is "", it will be inferred from file extension.
    // if encoder is nil, no encoder is used
    SendFileEx(string, string, encoding.Encoder, int) error
    // Shorthand for SendFileEx, infer mime, using gzip and return 200.
    SendFile(string) error

    // While done, any other operation except Write is not allowed anymore.
    Status(string, int) error

    // While done, any other operation except Write is not allowed anymore.
    SendCode(int) error

    // While done, any other operation except Write is not allowed anymore.
    Redirect(string) error
    RedirectEX(string, int) error

    // While done, any other operation except Write is not allowed anymore.
    JSON(interface{}) error
    JSONEx(interface{}, int) error

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

func (this *OriResponse)SendCode(code int) error {
    this.lock.Lock()
    defer this.lock.Unlock()

    if (this.haveSent) {
        return RES_HEAD_ALREADY_SENT
    }
    this.r.WriteHeader(code)
    this.haveSent=true

    return nil
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

func (this *OriResponse)RedirectEX(newAddr string, code int) error {
    this.lock.Lock()
    defer this.lock.Unlock()
    if (this.haveSent) {
        return RES_HEAD_ALREADY_SENT
    }

    this.haveSent=true

    this.r.Header().Set(LOCATION_HEADER, newAddr)
    this.r.WriteHeader(code) // moved temporarily

    return nil
}
func (this *OriResponse)Redirect(newAddr string) error {
    return this.RedirectEX(newAddr, 307)
}
func (this *OriResponse)JSONEx(obj interface{}, code int) error {
    this.lock.Lock()
    defer this.lock.Unlock()
    if (this.haveSent) {
        return RES_HEAD_ALREADY_SENT
    }

    var result, err=json.Marshal(obj)
    if err!=nil {
        return JSON_STRINGIFY_ERROR
    }

    this.haveSent=true
    this.r.WriteHeader(code)
    _, err=this.r.Write(result)
    return err
}
func (this *OriResponse)JSON(obj interface{}) error {
    return this.JSONEx(obj, 200)
}

func (this *OriResponse)SendFile(filepath string) error {
    return this.SendFileEx(filepath, "", encoding.GZipEncoder, 200)
}

func (this *OriResponse)SendFileEx(filepath string, mime string, encoder encoding.Encoder, httpCode int) error {
    this.lock.Lock()
    defer this.lock.Unlock()
    if (this.haveSent) {
        return RES_HEAD_ALREADY_SENT
    }

    // prepare file
    var fileHandler, fileErr=os.Open(filepath)
    if fileErr!=nil {
        return SENDFILE_FILEPATH_ERROR
    }

    // init mime
    if mime=="" {
        // infer the content type
        mime=MIME.TypeByExtension(fp.Ext(filepath))
        if mime=="" {
            mime=DEFAULT_CONTENT_TYPE
        }
    }
    // init encoder
    if encoder==nil {
        encoder=encoding.NOEncoder
    }
    var desWriter=encoder.WriterWrapper(this.r)
    if desWriter==nil {
        // fail to create. return a safe encoder or error? Now returns error.
        return SENDFILE_ENCODER_NOT_READY
    }

    this.haveSent=true

    // set headers
    this.r.Header().Set(CONTENT_TYPE_KEY, mime)
    if encoder.ContentEncoding()!="" {
        this.r.Header().Set(CONTENT_ENCODING, encoder.ContentEncoding())
    }

    this.r.WriteHeader(httpCode)

    // trnsfer file
    _, copyError:=io.Copy(desWriter, fileHandler)
    desWriter.Close()
    fileHandler.Close()
    // No need to close request writer. There's no such an interface.
    if copyError!=nil {
        // Attetez: seems not so accurate
        return SENT_BUT_ABORT
    }

    return nil
}
