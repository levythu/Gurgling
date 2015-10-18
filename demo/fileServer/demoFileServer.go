package main

import (
    . "github.com/levythu/gurgling"
    "github.com/levythu/gurgling/routers/simplefsserver"
    "fmt"
)

func main() {
    var router=simplefsserver.ASimpleFSServer("public/")

    fmt.Println("Running on port 8192...")
    LaunchServer(":8192", router)
}
