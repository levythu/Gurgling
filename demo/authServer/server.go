package main

import (
    . "github.com/levythu/gurgling"
    "github.com/levythu/gurgling/midwares/auth"
    "fmt"
)

func main() {
    SetGEnv("release")
    var root=ARouter()
    //root.Use(auth.ABasicAuth("levy", "levythu", "This site only"))
    root.Use(auth.ABasicAuth(func(u string, p string) bool {
        if u+"00"==p {
            return true
        }
        return false
    }, "Testify"))
    root.Get(func(req Request, res Response) {
        res.Send("Greetings!")
    })

    fmt.Println("Running on port 8192...")
    root.Launch(":8192")
}
