package utils

import (
    . "github.com/levythu/gurgling"
    "net/http"
    "crypto/tls"
    "strings"
    "io"
)

var pipeClient=&http.Client{}
var piptClientSSL=&http.Client{Transport: &http.Transport{
	TLSClientConfig: &tls.Config{},
}}
func Pipe(req Request, res Response, targetURL string) error {
    if strings.HasPrefix(strings.ToLower(targetURL), "https://") {
        return PipeX(req, res, targetURL, piptClientSSL)
    }
    return PipeX(req, res, targetURL, pipeClient)
}

func deepCopyHeader(src http.Header) http.Header {
    var srch=map[string][]string(src)
    var target=map[string][]string{}
    for k, v:=range srch {
        var tmp=make([]string, len(v))
        copy(tmp, v)
        target[k]=tmp
    }

    return target
}
func deepCopyHeaderIn(src http.Header, des http.Header) {
    var srch=map[string][]string(src)
    var target=map[string][]string(des)
    for k, v:=range srch {
        var tmp=make([]string, len(v))
        copy(tmp, v)
        target[k]=tmp
    }
}

func PipeX(req Request, res Response, targetURL string, client *http.Client) error {
    var oReq=req.R();
    var oRes=res.R();
    var proxyRequest, err=http.NewRequest(oReq.Method, targetURL, oReq.Body)
    if err!=nil {
        return err
    }

    proxyRequest.Header=deepCopyHeader(oReq.Header)
    var proxyResponse, err2=client.Do(proxyRequest)
    if err2!=nil {
        return err2
    }

    deepCopyHeaderIn(proxyResponse.Header, oRes.Header())
    oRes.WriteHeader(proxyResponse.StatusCode)
    io.Copy(oRes, proxyResponse.Body)
    proxyResponse.Body.Close()

    return nil
}
