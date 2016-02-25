package auth

import (
    . "github.com/levythu/gurgling"
    . "github.com/levythu/gurgling/definition"
    "io"
)

type BasicAuthMidware struct {
    inspector func(string, string) bool
    realm string
    Action Terminal
}

const content401="<!DOCTYPE html><html><head><title>401 Unauthorized" +
    "</title><style>h2 {text-align: center;}table {max-width: 35em;margin: 0 auto;padding-top: 1em;font: 1em Arial, Helvetica, sans-serif;}a {text-decoration: none;color: #0070C0;}p {text-align:right;padding-top: 0.6em;}div {position: absolute;width: 100%;left: 0;border-bottom: 1px solid #CCC;padding-top: 0.5em;}</style></head><body><h2>401 Unauthorized"+
    "</h2><table>"+
    "<tr><td><div></div><p>by <a href=\"https://github.com/levythu/gurgling\">Gurgling "+Version+"</a></p></td></tr></table></body></html>"

func (this *BasicAuthMidware)Handler(req Request, res Response) (bool, Request, Response) {
    var user, pass, ok=req.R().BasicAuth()
    if this.inspector==nil || (ok && this.inspector(user, pass)) {
        // authentication passed.
        return true, req, res
    }
    // not passed, require insert password
    if this.Action!=nil {
        this.Action(req, res)
        return false, req, res
    }
    res.Set("WWW-Authenticate", "Basic realm=\""+this.realm+"\"")
    if CGurgling.ID==ID_FOR_PREDEFINED_DEBUG {
        res.Status("401 Unauthorized.", 401)
    } else if CGurgling.ID==ID_FOR_PREDEFINED_RELEASE {
        if res.Set("Content-Type", "text/html; charset=utf-8")==nil {
            if res.SendCode(401)==nil {
                io.WriteString(res, content401)
            }
        }
    } else {
        res.SendCode(401)
    }

    return false, req, res
}

func generateInspector(preUser string, prePass string) func(string, string) bool {
    return func(usr string, pass string) bool {
        return usr==preUser && pass==prePass
    }
}

// ABasicAuth(func(string, string) bool, [string]) returns a midware that implements customized authentication
// ABasicAuth(string, string, [string]) returns a midware that use fixed username/passwd as certification
func ABasicAuth(paraList ...interface{}) *BasicAuthMidware {
    var ret=BasicAuthMidware{
        realm: "The Site",
        Action: nil,
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
    } else if len(paraList)==3 {
        var e1, o1=paraList[0].(string)
        var e2, o2=paraList[1].(string)
        var e3, o3=paraList[2].(string)
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
