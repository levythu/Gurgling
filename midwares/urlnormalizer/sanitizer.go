package urlnormalizer

import (
    "strings"
    . "github.com/levythu/gurgling"
)

type Sanitizer struct {
    // implements IMidware
}

func ASanitizer() IMidware {
    return &Sanitizer{}
}

func sanitize(ori string) (string, bool) {
    if ori=="" || ori=="/" {
        return ori, false
    }
    var resStack=[]string{}
    var oriStack=strings.Split(ori, "/")
    if strings.HasPrefix(ori, "/") {
        oriStack=oriStack[1:]
    }
    if strings.HasSuffix(ori, "/") {
        oriStack=oriStack[:len(oriStack)-1]
    }
    var modified=false

    for _, str:=range oriStack {
        if str==".." {
            if len(resStack)>0 {
                resStack=resStack[:len(resStack)-1]
            }
            modified=true
        } else if str=="." {
            modified=true
        } else if str=="" {
            modified=true
        } else {
            resStack=append(resStack, str)
        }
    }

    var resURL=""
    for _, str:=range resStack {
        resURL+="/"+str
    }

    if strings.HasSuffix(ori, "/") {
        resURL+="/"
    }

    return resURL, modified
}

func (this *Sanitizer)Handler(req Request, res Response) (bool, Request, Response) {
    var result, isModified=sanitize(req.Path())
    if !isModified {
        return true, req, res
    }
    // redirect permanantly
    res.RedirectEX(result, 301)
    return false, nil, nil
}
