package main

import (
    . "github.com/levythu/gurgling"
    "fmt"
)

func main() {
    SetGEnv("release")

    var router Router=ARouter().Get(func(req Request, res Response) {
        res.Send("Hello World!")
    })

    fmt.Println("Running on port 8023...")
    router.Launch(":8023", "public/server.crt", "public/server.key")
}
