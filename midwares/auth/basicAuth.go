package auth

import (
    . "github.com/levythu/gurgling"
)

type BasicAuthMidware struct {
    inspector func(string, string) bool
    realm string
}

func (this *BasicAuthMidware)Handler(req Request, res Response) (bool, Request, Response) {
    var user, pass, ok=req.R().BasicAuth()
    if this.inspector==nil || (ok && this.inspector(user, pass)) {
        // authentication passed.
        return true, req, res
    }
    // not passed, require insert password
    res.Set("WWW-Authenticate", "Basic realm=\""+this.realm+"\"")
    res.SendCode(401)
    return false, req, res
}

func generateInspector(preUser, prePass) func(string, string) bool {
    return func(usr string, pass string) bool {
        return usr==preUser && pass=prePass
    }
}

// ABasicAuth(func(string, string) bool, [string]) returns a midware that implements customized authentication
// ABasicAuth(string, string, [string]) returns a midware that use fixed username/passwd as certification
func ABasicAuth(paraList ...interface{}) *BasicAuthMidware {
    var ret=BasicAuthMidware{
        realm: "The Site",
    }
    if len(paraList)==1 {
        var tmp, ok=paraList[0].(func(string, string) bool)
        if !ok {
            panic(INVALID_PARAMETER)
        }
        ret.inspector=tmp
    } else if len(paraList)==2 {
        var tmp, ok=paraList[1].(string)
        if !ok {
            panic(INVALID_PARAMETER)
        }
        switch elem1:=paraList[0].(type) {
        case func(string, string) bool:
            ret.inspector=elem1
            ret.realm=tmp
        case string:
            ret.inspector=generateInspector(elem1, tmp)
        default:
            panic(INVALID_PARAMETER)
        }
    } else if len(paraList==3) {
        var e1, o1=paraList[0].(string)
        var e2, o2=paraList[0].(string)
        var e3, o3=paraList[0].(string)
        if o1 && o2 && o3 {
            ret.inspector=generateInspector(e1, e2)
            ret.realm=e3
        } else {
            panic(INVALID_PARAMETER)
        }
    } else {
        panic(INVALID_PARAMETER)
    }

    return &ret
}
