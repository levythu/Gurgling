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
    req.F()["Cookie"]=&CookieHandler {
        req: req,
        res: res,
    }
    return true, req, res
}

type CookieHandler struct {
    req Request
    res Response
}

// Get the cookie specified by name.
// If such cookie does not exist, an empty string will be returned.
func (this *CookieHandler)Get(name string) string {
    var target, err=this.req.R().Cookie(name)
    if err!=nil {
        return ""
    }
    return target.Value
}

func (this *CookieHandler)GetCookie(name string) *http.Cookie {
    var target, err=this.req.R().Cookie(name)
    if err!=nil {
        return nil
    }
    return target
}

// Set a session cookie
func (this *CookieHandler)Set(name string, val string) {
    this.SetCookie(&http.Cookie{
        Name: name,
        Value: val,
    })
}

func (this *CookieHandler)SetWithAge(name string, val string, age int) {
    this.SetCookie(&http.Cookie{
        Name: name,
        Value: val,
        MaxAge: age,
    })
}
func (this *CookieHandler)SetWithExpire(name string, val string, expireAt time.Time) {
    this.SetCookie(&http.Cookie{
        Name: name,
        Value: val,
        Expires: expireAt,
    })
}
func (this *CookieHandler)Remove(name string) {
    this.SetCookie(&http.Cookie{
        Name: name,
        Value: "",
        MaxAge: -1,
    })
}

func (this *CookieHandler)SetCookie(val *http.Cookie) {
    http.SetCookie(this.res.R(), val)
}
