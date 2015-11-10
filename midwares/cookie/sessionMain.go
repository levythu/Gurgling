package cookie

import (
    "github.com/levythu/gurgling/midwares/cookie/kvstore"
    "github.com/levythu/gurgling/midwares/cookie/uuid"
    . "github.com/levythu/gurgling"
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

func (this *Session)Handler(req Request, res Response) (bool, Request, Response) {
    var hackedRes=&resSession{}
    hackedRes.Response=res
    hackedRes.sessionInfo=this



    return true, req, hackedRes
}
