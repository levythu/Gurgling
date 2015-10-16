package gurgling

// define utilities relevant to error handling

import (
    "io"
    . "github.com/levythu/gurgling/definition"
)

type RouterErrorCatcher func(Request, Response, interface{})

const content500="<!DOCTYPE html><html><head><title>500 Internal Error" +
    "</title><style>h2 {text-align: center;}table {max-width: 35em;margin: 0 auto;padding-top: 1em;font: 1em Arial, Helvetica, sans-serif;}a {text-decoration: none;color: #0070C0;}p {text-align:right;padding-top: 0.6em;}div {position: absolute;width: 100%;left: 0;border-bottom: 1px solid #CCC;padding-top: 0.5em;}</style></head><body><h2>500 Internal Error"+
    "</h2><table>"+
    "<tr><td><div></div><p>by <a href=\"https://github.com/levythu/gurgling\">Gurgling "+Version+"</a></p></td></tr></table></body></html>"
func DefaultCacher(req Request, res Response, thePanic interface{}) {
    if res.Set("Content-Type", "text/html; charset=utf-8")==nil {
        if res.SendCode(500)==nil {
            io.WriteString(res, content500)
        }
    }
}
