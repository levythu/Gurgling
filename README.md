# Gurgling
An extremely-light framework for Golang to build restful API and Website.

## Quick Start

### Install

```sh
go get github.com/levythu/gurgling
```
### Setup a HTTP server

```go
package main
    
import (
    . "github.com/levythu/gurgling"
    "net/http"
)

func main() {
    // Create a gate router.
    const mountPoint="/"
    var router=GetRouter(mountPoint)

    // Mount one handler
    router.Get("/", func(req Request, res Response) {
        res.Send("Hello, World!")
    })

    // Mount the gate to net/http and run the server
    http.Handle(mountPoint, router)
    fmt.Println("Running...")
    http.ListenAndServe(":8080", nil)
}
```	

### Create sub-routers and mount them

```go

// Create a non-gate router.
var pageRouter=ARouter()

// Mount handler and midware
pageRouter.Use("/", func(req Request, res Response) (bool, Request, Response) {
    fmt.Println(req.Path())
	return true, req, res
})
pageRouter.Get("/editor", func(req Request, res Response) {
    res.Send("Editors")
})

// Mount the router to the previous one
router.Use(pageRouter)
```	

