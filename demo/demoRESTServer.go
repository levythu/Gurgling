package main

import (
    "fmt"
    . "github.com/levythu/gurgling"
)
/*
func main() {
    var router=ARouter()
    var page=getPageRouter()

    router.Get("/", func(req Request, res Response) {
        res.Send("This is index.")
    })
    router.Use("/page", page)

    fmt.Println("Running...")
    LaunchServer(":8192", router)
}
*/
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
