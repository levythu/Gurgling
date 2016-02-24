package main

import (
    . "github.com/levythu/gurgling"
    "github.com/levythu/gurgling/routers/simplefsserver"
    "fmt"
)

func main() {
    // Use the directory of the program as root directory to the public
    SetGEnv("release")

    var router Router=simplefsserver.ASimpleFSServer(".")

    fmt.Println("Running on port 8192...")
    router.Launch(":8192")
}
