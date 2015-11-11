package main

import (
    "fmt"
    . "github.com/levythu/gurgling"
    "github.com/levythu/gurgling/midwares/cookie"
)

func main() {
    var router=ARouter().Use(cookie.ASession("salt-like-this"))

    router.Get(func(req Request, res Response) {
        var session=req.F()["session"].(map[string]string)
        if v, ok:=session["user"]; !ok {
            res.Send("Please login first.")
        } else {
            res.Send("Hello, "+v+"!")
        }
    })
    router.Get("/logout", func(req Request, res Response) {
        req.F()["session"]=nil
        res.Send("Logouted.")
    })
    router.Use("/login", func(req Request, res Response) {
        var session=req.F()["session"].(map[string]string)
        if len(req.Path())<=1 {
            res.Send("Please specify the user.")
        } else {
            session["user"]=req.Path()[1:]
            res.Send("Logined.")
        }
    })

    fmt.Println("Running...")
    router.Launch(":8192")
}
