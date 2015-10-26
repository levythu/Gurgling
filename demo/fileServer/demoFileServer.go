package main

import (
    . "github.com/levythu/gurgling"
    "github.com/levythu/gurgling/routers/simplefsserver"
    "github.com/levythu/gurgling/midwares/bodyparser"
    "fmt"
)

func main() {
    // Use the directory of the program as root directory to the public
    SetGEnv("release")

    var router Router=simplefsserver.ASimpleFSServer("public").Use(bodyparser.ABodyParser()).Post(func(req Request, res Response){
        fmt.Println(req.Body())
    })

    fmt.Println("Running on port 8192...")
    router.Launch(":8192")
}
