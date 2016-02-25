package piperouter

import (
    . "github.com/levythu/gurgling"
)

type PipeRouter struct {
    Router
}

func pipeToUrl(req Request, res Response) {

}

func APipeRouter() Router {
    var trouter=ARouter().Use(pipeToUrl)
    return PipeRouter{
        Router: trouter,
    }
}
