package bodyparser

import (
    . "github.com/levythu/gurgling"
    . "github.com/levythu/gurgling/definition"
    "strings"
    "encoding/json"
    "io/ioutil"
)

// Force to fetch data and store it in req.F["body"]

type BodyParser struct {
    // implement IMidware

    // if nil, no filter.
    MethodFilter map[string]bool
    MaxMemoryUse int64
}

func ABodyParser() IMidware {
    return &BodyParser {
        MethodFilter: map[string]bool {
            "POST": true,
            "PUT": true,
        },
        MaxMemoryUse: 1024*1024*4,   //4MB cache for multipart form.
    }
}

func parseContentType(raw string) string {
    // ACCORDING to rfc1341, the contentType should contain type in its head and optical parameters
    // divided by semicolons.
    var i=strings.Index(raw, ";")
    if i<0 {
        return raw
    }
    return raw[:i]
}

func (this *BodyParser)Handler(req Request, res Response) (isCont bool, nReq Request, nRes Response) {
    isCont=true
    nReq=req
    nRes=res

    if !(this.MethodFilter==nil || this.MethodFilter[req.Method()]) {
        return
    }

    var contentType=strings.ToLower(parseContentType(req.Get(CONTENT_TYPE_KEY)))
    if contentType=="application/x-www-form-urlencoded" {
        // Parse it as key-value.
        // in the case the body is url.Values
        var err=req.R().ParseForm()
        if err==nil {
            req.F()["body"]=req.R().PostForm
        }
    } else if contentType=="multipart/form-data" {
        // Parse it as multipart form.
        // in the case the body is *multipart.Form
        var err=req.R().ParseMultipartForm(this.MaxMemoryUse)
        if err==nil {
            req.F()["body"]=req.R().MultipartForm
        }
    } else if contentType=="application/json" {
        // Parse it as JSON.
        // in the case the body is map[string]Tout
        var rawData, err=ioutil.ReadAll(req.R().Body)
        if err!=nil {
            return
        }
        var ret map[string]Tout
        err=json.Unmarshal(rawData, &ret)
        if err==nil {
            req.F()["body"]=ret
        }
    } else {
        // Fetch it but do not parse
        // in the case the body is []byte
        var rawData, err=ioutil.ReadAll(req.R().Body)
        if err==nil {
            req.F()["body"]=rawData
        }
    }
    return
}
