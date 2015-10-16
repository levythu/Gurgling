package analyzer

import (
    . "github.com/levythu/gurgling"
    "time"
)

// An analyzer only logging the response code, response time in ms and request method.

struct SimpleAnalyzer {
    // nothing but implements Sandwich
}

func ASimpleAnalyzer() Sandwich {
    return &SimpleAnalyzer{}
}

const token_returncode="SimpleAnalyzer-Status-Code"
const token_starttime="SimpleAnalyzer-Start-Time"

func logCode(res Response, c int) {
    res.F()[token]=c
}

func (this *SimpleAnalyzer)Handler(req Request, res Response) (bool, Request, Response) {
    var newRes=&logResponse {
        o: res,
        OnHeadSent: logCode
    }
    newRes.F()[token_starttime]=time.Now().UnixNano()
    return true, req, newRes
}
func (this *SimpleAnalyzer)Final(req Request, res Response) {
    
}
