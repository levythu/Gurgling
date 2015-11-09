package cookie

import (
    "net/http"
    . "github.com/levythu/gurgling"
    "time"
)

type Cookier struct {
    // Implementing Midware
}

func ACookie() *Cookier {
    return &Cookier{}
}

func (this *Cookier)Handler(req Request, res Response) (bool, Request, Response) {
    req["Cookie"]=CookieHandler(&CookieHandler {
        req: req,
        res: res,
    })
    return true, req, res
}

type CookieHandler *cookieHandler
type cookieHandler struct {
    req Request
    res Response
}

// Get the cookie specified by name.
// If such cookie does not exist, an empty string will be returned.
func (CookieHandler this)Get(name string) string {
    var target, err=this.req.R().Cookie(name)
    if err!=nil {
        return ""
    }
    return target.Value
}

func (CookieHandler this)GetCookie(name string) *http.Cookie {
    var target, err=this.req.R().Cookie(name)
    if err!=nil {
        return nil
    }
    return target
}

// Set a session cookie
func (CookieHandler this)Set(name string, val string) {
    SetCookie(&http.Cookie{
        Name: name,
        Value: val,
    })
}

func (CookieHandler)SetWithAge(name string, val string, age int) {
    SetCookie(&http.Cookie{
        Name: name,
        Value: val,
        MaxAge: age,
    })
}
func (CookieHandler)SetWithExpire(name string, val string, expireAt time.Time) {
    SetCookie(&http.Cookie{
        Name: name,
        Value: val,
        Expires: expireAt,
    })
}
func (CookieHandler)Remove(name string) {
    SetCookie(&http.Cookie{
        Name: name,
        Value: "",
        MaxAge: -1,
    })
}

func (CookieHandler)SetCookie(val *http.Cookie) {
    http.SetCookie(this.req.R(), val)
}
