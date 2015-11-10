package main

import (
    "fmt"
    . "github.com/levythu/gurgling"
    "github.com/levythu/gurgling/midwares/cookie"
)

func main() {
    var router=ARouter().Use(cookie.ASession("salt"))

    router.Get(func(req Request, res Response) {
        res.Send("This is index.")
    })

    fmt.Println("Running...")
    router.Launch(":8192")
}
