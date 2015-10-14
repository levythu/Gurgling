package main

import (
    "fmt"
    "github.com/levythu/gurgling"
    "net/http"
)

func main() {
    var router=gurgling.GetRouter("/")

    router.Use("/", func(req gurgling.Request, res gurgling.Response) (bool, gurgling.Request, gurgling.Response) {
        req.F()["huahua"]="123"
        return true, req, res
    })
    router.Get("/233", func(req gurgling.Request, res gurgling.Response) {
        res.Send(req.F()["huahua"].(string))
    })
    router.Get("/", func(req gurgling.Request, res gurgling.Response) {
        res.Send(req.Path())
    })

    http.Handle("/", router)
    fmt.Println("Running...")
    http.ListenAndServe(":9144", nil)

}
