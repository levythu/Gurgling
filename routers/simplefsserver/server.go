package simplefsserver

import (
    . "github.com/levythu/gurgling"
    . "github.com/levythu/gurgling/definition"
    "github.com/levythu/gurgling/midwares/staticfs"
    "github.com/levythu/gurgling/midwares/analyzer"
    "io/ioutil"
    "io"
)

// An simple and naive application of midwares/staticfs, indexing directory and handling 404.
// the rendered page will not be saved.
func renderDirectory(req Request, res Response, directory string) {
    var fileList, err=ioutil.ReadDir(directory)
    if err!=nil {
        res.Status("Internal error while reading file", 500)
        return
    }

    var content="<!DOCTYPE html><html><head><title>Directory listing for "
    content+=req.Path()
    content+="</title><style>h2 {text-align: center;}table {max-width: 35em;margin: 0 auto;padding-top: 1em;font: 1em Arial, Helvetica, sans-serif;}a {text-decoration: none;color: #0070C0;}p {text-align:right;padding-top:0.6em;}div {position: absolute;width: 100%;left:0;border-bottom:1px solid #CCC;padding-top:0.5em;}</style></head><body><h2>Directory listing for "
    content+=req.Path()
    content+="</h2><table>"

    var filenames=[]string{"./", "../"}
    for _, elem:=range fileList {
        var tN=elem.Name()
        if tN==".." || tN=="." {
            continue
        }
        if elem.IsDir() {
            filenames=append(filenames, tN+"/")
        } else {
            filenames=append(filenames, tN)
        }
    }
    for _, elem:=range filenames {
        content+="<tr><td><a href=\""+elem+"\">"+elem+"</a></td></tr>"
    }
    content+="<tr><td><div></div><p>by <a href=\"https://github.com/levythu/gurgling\">Gurgling "+Version+"</a></p></td></tr></table></body></html>"

    res.Set("Content-Type", "text/html; charset=utf-8")
    res.SendCode(200)
    io.WriteString(res, content)
}

const content="<!DOCTYPE html><html><head><title>404 Not Found" +
    "</title><style>h2 {text-align: center;}table {max-width: 35em;margin: 0 auto;padding-top: 1em;font: 1em Arial, Helvetica, sans-serif;}a {text-decoration: none;color: #0070C0;}p {text-align:right;padding-top: 0.6em;}div {position: absolute;width: 100%;left: 0;border-bottom: 1px solid #CCC;padding-top: 0.5em;}</style></head><body><h2>404 Not Found"+
    "</h2><table>"+
    "<tr><td><div></div><p>by <a href=\"https://github.com/levythu/gurgling\">Gurgling "+Version+"</a></p></td></tr></table></body></html>"
func notFound(req Request, res Response) {
    res.Set("Content-Type", "text/html; charset=utf-8")
    res.SendCode(404)
    io.WriteString(res, content)
}

func ASimpleFSServer(basePath string) Router {
    var fs=staticfs.AStaticfs(basePath).(*staticfs.FsMidware)
    var ay=analyzer.ASimpleAnalyzer()
    fs.DefaultRender=renderDirectory

    return ARouter().Use(ay).Use(fs)
}
func NewSimpleFSServer(basePath string, cacheStrategy staticfs.CacheStrategy, autoIndexing bool) Router {
    var fs=staticfs.AStaticfs(basePath).(*staticfs.FsMidware)
    if autoIndexing {
        fs.DefaultRender=renderDirectory
    }
    var ay=analyzer.ASimpleAnalyzer()
    fs.CacheControl=cacheStrategy

    return ARouter().Use(ay).Use(fs)
}
