package main

import (
    "fmt"
    . "github.com/levythu/gurgling"
)

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
