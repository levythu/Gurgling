package main

import (
    . "github.com/levythu/gurgling"
    "fmt"
)

func getPageRouter() Router {
    var page=ARouter()
    page.Get("/", func(req Request, res Response) {
        res.Send("The list of pages.")
    })
    page.Get("/edit", func(req Request, res Response) {
        res.Send("Edit page.")
    })
    page.Get("/redirect", func(req Request, res Response) {
        res.Redirect("/edit")
    })
    page.Use("/file", func(req Request, res Response) {
        if req.Path()=="" {
            res.Send("Specify the path please.")
            return
        }
        res.SendFile(req.Path()[1:])
    })
    page.Last(func(req Request, res Response) {
        fmt.Println("Page end.")
    })

    return page
}
