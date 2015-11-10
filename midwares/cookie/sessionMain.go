package cookie

import (
    "github.com/levythu/gurgling/midwares/cookie/kvstore"
    "github.com/levythu/gurgling/midwares/cookie/uuid"
    . "github.com/levythu/gurgling"
    . "github.com/levythu/gurgling/definition"
    "net/http"
)

type Session struct {
    StoreIO kvstore.KvStore
    Secret string
    UUIDGenerator func() string
    Resave bool
    Rolling bool
    Path string
    HttpOnly bool
    Secure bool     // Since gurgling does not support HTTPS currently, this value is preserved.
    // <=0 for Session Cookie
    MaxAge int

    uidgen *uuid.UUID
}

func ASession(secret string) *Session {
    var ret=&Session{
        StoreIO: &kvstore.MemKvStore{
            Map: make(map[string]map[string]string),
        },
        Secret: secret,
        Resave: false,
        Rolling: false,
        Path: "/",
        HttpOnly: true,
        Secure: false,
        MaxAge: 0,
        uidgen: uuid.AUUID(0),
    }
    ret.UUIDGenerator=func() string {
        return ret.uidgen.Get()
    }
    return ret
}

const sid_cookie_key="sid"

func (this *Session)Handler(req Request, res Response) (bool, Request, Response) {
    var hackedRes=&resSession{}
    hackedRes.Response=res
    hackedRes.sessionInfo=this

    var sessionid=getSignedCookieContent(req.R(), sid_cookie_key, this.Secret)
    if sessionid=="" {
        req.F()["session"]=make(map[string]string)
    } else {
        var sessionContent=this.StoreIO.Get(sessionid)
        if sessionContent==nil {
            req.F()["session"]=make(map[string]string)
        } else {
            req.F()["session"]=sessionContent
            req.F()["sessionid"]=sessionid
            res.F()["sessionid"]=sessionid
        }
    }
    var result=req.F()["session"].(map[string]string)
    var bakup=make(map[string]string)
    for k, v:=range result {
        bakup[k]=v
    }
    res.F()["session-old"]=bakup

    return true, req, hackedRes
}

func (this *Session)UpdateSession(res Response) {
    var reqR=res.F()[RKEY_LOW_LAYER_R].(map[string]Tout)

    var sessionid, ok=res.F()["sessionid"].(string)
    var session, ok2=reqR["session"].(map[string]string)
    var session_old, ok_old=res.F()["session-old"].(map[string]string)
    if !ok_old {
        // logical exception
        return
    }
    if !ok {
        // not setup a session yet
        if !ok2 || len(session)==0 {
            // not modified, do nothing then.
            return
        } else {
            // a new session should be establish
            var newSID=this.UUIDGenerator()
            if this.StoreIO.Set(newSID, session)!=nil {
                return
            }
            this.writeSession(res, newSID)
            return
        }
    } else {
        // already a session, check whether to update
        if !ok2 || len(session)==0 {
            // remove the session
            this.StoreIO.Remove(sessionid)
            this.writeSession(res, "")
            return
        }
        if this.Rolling {
            this.writeSession(res, sessionid)
        }
        if this.Resave || !compareMap(session, session_old) {
            this.StoreIO.Set(sessionid, session)
        }
    }
}

// sid="" for remove it now
func (this *Session)writeSession(res Response, sid string) {
    var toSet=http.Cookie{
        Name: sid_cookie_key,
        Value: sid,
        Path: this.Path,
        Secure: this.Secure,
        HttpOnly: this.HttpOnly,
    }
    if this.MaxAge>0 {
        toSet.MaxAge=this.MaxAge
    }
    if sid=="" {
        removeSignedCookie(res.R(), &sid)
        return
    }
    setSignedCookie(res.R(), &sid)
}

func compareMap(map1, map2 map[string]string) bool {
    if len(map1)!=len(map2) {
        return false
    }
    for k, v:=range map1 {
        var v2, ok=map2[k]
        if !ok || v2!=v {
            return false
        }
    }

    return true
}
