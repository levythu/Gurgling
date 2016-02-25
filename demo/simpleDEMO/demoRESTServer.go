package main

import (
    "fmt"
    . "github.com/levythu/gurgling"
    "github.com/levythu/gurgling/midwares/staticfs"
    "github.com/levythu/gurgling/midwares/analyzer"
)

func main() {
    var router=ARouter()
    var page=getPageRouter()

    router.Use(analyzer.ASimpleAnalyzer())
    router.Use(staticfs.AStaticfs("public/"))
    router.Get(func(res Response) {
        res.Send("This is index.")
    }).Use("/page", page)

    fmt.Println("Running...")
    router.Launch(":8192")
}
/*
func main() {
    var router=ARouter()
    var subrouter=ARouter()

    router.Use("/", func(req Request, res Response) (bool, Request, Response) {
        fmt.Println("---------", req.OriginalPath(), "---------")
        fmt.Println("<"+req.BaseUrl()+">", ",", "<"+req.Path()+">")
        // pass the request.
        return true, req, res
    })
    router.Use("/sub", subrouter)
    subrouter.Use("/", func(req Request, res Response) (bool, Request, Response) {
        fmt.Println("<"+req.BaseUrl()+">", ",", "<"+req.Path()+">")
        return true, req, res
    })

    LaunchServer(":8192", router)
}
*/
