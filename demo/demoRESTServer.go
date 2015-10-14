package main

import (
    "fmt"
    "github.com/levythu/gurgling"
    "net/http"
)

func main() {
    var router=gurgling.GetRouter("/")

    router.Get("/", func(req gurgling.Request, res gurgling.Response) {
        res.Send("hahahaha")
    })

    http.Handle("/", router)
    fmt.Println("Running...")
    http.ListenAndServe(":9144", nil)

}
