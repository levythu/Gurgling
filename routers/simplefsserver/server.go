package simplefsserver

import (
    . "github.com/levythu/gurgling"
    . "github.com/levythu/gurgling/definition"
    "github.com/levythu/gurgling/midwares/staticfs"
    "io/ioutil"
    "io"
)

// An simple and naive application of midwares/staticfs, indexing directory and handling 404.
// the rendered page will not be saved.
func renderDeirectory(req Request, res Response, directory string) {
    var fileList, err=ioutil.ReadDir(directory)
    if err!=nil {
        res.Status("Internal error while reading file", 500)
        return
    }

    var content="<!DOCTYPE html><html><head><title>Directory listing for "
    content+=req.Path()
    content+="</title><style>h2 {text-align: center;}table {max-width: 35em;margin: 0 auto;padding-top: 1em;font: 1em Arial, Helvetica, sans-serif;}a {text-decoration: none;color: #0070C0;}p {padding-top: 1em;}</style></head><body><h2>Directory listing for "
    content+=req.Path()
    content+="</h2><table>"

    for _, elem:=range fileList {
        content+="<tr><td><a href=\""+elem.Name()+"\">"+elem.Name()+"</a></td></tr>"
    }
    content+="<tr><td><p>by <a href=\"https://github.com/levythu/gurgling\">Gurgling "+Version+"</a></p></td></tr></table></body></html>"

    res.Set("Content-Type", "text/html; charset=utf-8")
    io.WriteString(res, content)
}

func ASimpleFSServer(basePath string) Router {
    var fs=staticfs.AStaticfs(basePath).(*staticfs.FsMidware)
    fs.DefaultRender=renderDeirectory

    return ARouter().Use(fs)
}
