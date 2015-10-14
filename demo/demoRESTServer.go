package main

import (
    "fmt"
    . "github.com/levythu/gurgling"
    "net/http"
)

func main() {
    var router=GetRouter("/")

    router.Use("/", func(req Request, res Response) (bool, Request, Response) {
        req.F()["huahua"]="123"
        fmt.Println(req.Path())
        return true, req, res
    })
    router.Get("/233", func(req Request, res Response) {
        res.Send(req.F()["huahua"].(string))
    })
    router.Get("/", func(req Request, res Response) {
        res.Send(req.Path())
    })

    http.Handle("/", router)
    fmt.Println("Running...")
    http.ListenAndServe(":9144", nil)

}
