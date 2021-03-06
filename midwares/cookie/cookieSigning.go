package cookie

// low-layer functions for cookie signature.

import (
    "net/http"
    "crypto/sha256"
    "fmt"
)

const sign_suffix="-gsign"
const mid_salt="~github.com/levythu/gurgling/midwares/cookie::cookieSigning~"

func _hash(src string) string {
    return fmt.Sprintf("%x", sha256.Sum256([]byte(src)))
}

func setSignedCookie(w http.ResponseWriter, cookie *http.Cookie, secret string) {
    var signCookie=*cookie
    signCookie.Name+=sign_suffix
    signCookie.Value=_hash(secret+mid_salt+signCookie.Value)
    http.SetCookie(w, cookie)
    http.SetCookie(w, &signCookie)
}

func removeSignedCookie(w http.ResponseWriter, cookie *http.Cookie) {
    cookie.MaxAge=-1
    var signCookie=*cookie
    signCookie.Name+=sign_suffix
    http.SetCookie(w, cookie)
    http.SetCookie(w, &signCookie)
}

// If not valid, returns nil
func getSignedCookie(r *http.Request, name string, secret string) *http.Cookie {
    var target, err=r.Cookie(name)
    if err!=nil {
        return nil
    }
    var targetSigned, err2=r.Cookie(name+sign_suffix)
    if err2!=nil {
        return nil
    }

    if _hash(secret+mid_salt+target.Value)!=targetSigned.Value {
        return nil
    }
    return target
}

func getSignedCookieContent(r *http.Request, name string, secret string) string {
    var res=getSignedCookie(r, name, secret)
    if res==nil {
        return ""
    }
    return res.Value
}
