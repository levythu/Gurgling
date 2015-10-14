package main

import (
    . "github.com/levythu/gurgling"
)

func getPageRouter() Router {
    var page=ARouter()
    page.Get("/", func(req Request, res Response) {
        res.Send("The list of pages.")
    })
    page.Get("/edit", func(req Request, res Response) {
        res.Send("Edit page.")
    })

    return page
}
