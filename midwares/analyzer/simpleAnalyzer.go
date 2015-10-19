package analyzer

import (
    . "github.com/levythu/gurgling"
    "time"
    "fmt"
    "strconv"
)

// An analyzer only logging the response code, response time in ms and request method.

type SimpleAnalyzer struct {
    // nothing but implements Sandwich
}

func ASimpleAnalyzer() Sandwich {
    return &SimpleAnalyzer{}
}

const token_returncode="SimpleAnalyzer-Status-Code"
const token_starttime="SimpleAnalyzer-Start-Time"

func logCode(res Response, c int) {
    res.F()[token_returncode]=c
}

func (this *SimpleAnalyzer)Handler(req Request, res Response) (bool, Request, Response) {
    var newRes=&logResponse {
        o: res,
        OnHeadSent: logCode,
    }
    newRes.F()[token_starttime]=time.Now().UnixNano()
    return true, req, newRes
}
func (this *SimpleAnalyzer)Final(req Request, res Response) {
    var timeStart, ok=res.F()[token_starttime].(int64)
    var timeElpase string
    if ok {
        var t=time.Now().UnixNano()
        timeElpase=strconv.FormatInt((t-timeStart)/1000000, 10)+"ms"
    } else {
        timeElpase="xxxx"
    }

    var statusCode, ok2=res.F()[token_returncode].(int)
    var codeStr string
    if ok2 {
        codeStr=strconv.Itoa(statusCode)
    } else {
        codeStr="---"
    }

    var url=req.R().URL
    fmt.Println("- "+timeElpase+"\t\t"+codeStr+"\t"+req.Method()+"\t"+url.Path+"?"+url.RawQuery)
}
