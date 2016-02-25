package main

import (
    . "github.com/levythu/gurgling"
    u "github.com/levythu/gurgling/utils"
    "fmt"
)

const targetHostname="http://www.levy.at"

func main() {

    var router=ARouter()
    router.Use(func(req Request, res Response) {
        var oriURL=req.R().URL
        var target=targetHostname+req.Path()
        if oriURL.RawQuery!="" {
            target+="?"+oriURL.RawQuery
        }
        if oriURL.Fragment!="" {
            target+="#"+oriURL.Fragment
        }
        if u.Pipe(req, res, target)==nil {
            fmt.Println("[SUCC] Piping to "+target)
        } else {
            fmt.Println("[FAIL] Piping to "+target)
        }
    })

    fmt.Println("Running on port 8192...")
    router.Launch(":8192")
}
