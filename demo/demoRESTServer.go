package main

import (
    "fmt"
    . "github.com/levythu/gurgling"
    "net/http"
)

func main() {
    var router=GetRouter("/")
    var page=getPageRouter()

    router.Get("/", func(req Request, res Response) {
        res.Send("This is index.")
    })
    router.Use("/page", page)

    http.Handle("/", router)
    fmt.Println("Running...")
    http.ListenAndServe(":8192", nil)
}
