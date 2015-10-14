# Gurgling
An extremely-light framework for Golang to build restful API and Website.

**Special Thanks to [Express](http://expressjs.com/), which provides API samples for this project.**

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
)

func main() {
    // Create a root router.
    var router=ARouter()

    // Mount one handler
    router.Get("/", func(req Request, res Response) {
        res.Send("Hello, World!")
    })

    // Mount the gate to net/http and run the server
    fmt.Println("Running...")
    LaunchServer(":8080", router)
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
    res.Send("Here's the editor.")
})

// Mount the router to the previous one
router.Use("/page", pageRouter)
```

## API Docs
### Router
The core module of gurgling. It is indeed an interface, and is implemented by `router`, its default and original version. The interface is designed for extensions.

#### `func GetRouter(MountPoint string) Router`
Creates and returns one default router for gateway. The mountpoint it is mounted to in `http.Handle()` function should be specified here.

#### `func ARouter() Router`
Creates and returns one default router with `mountpoint="/"`.

#### `func (*router)Use(mountpoint string, processor Tout) Router`
Mounts a mountable to the router at mountpoint. Mountpoint must start with `/`, regexp is currently unsupported. It will try to match the mountpoint by prefix.   
Returns the router itself for method chaining.

Mountable includes the following items:

##### **`Midware`** (`type Midware func(Request, Response) (bool, Request, Response)`)  
A function, receiving `Request` and `Response` as parameter.  
Returns three value, two of which are modified res&req (if no modification just return the original) and the boolean indicates whether to pass the request to the next handler (`false` for no).

```go
router.Use("/", func(req Request, res Response) (bool, Request, Response) {
    fmt.Println(req.Path())
	// PASS the request to the next handler.
	return true, req, res
})
```

##### **`IMidware`**  
An interface which implement Midware function as `.Handler()`.  
Since Router also implement the function, Router is a special IMidware. It will never pass request to the next.

```go
var anotherRouter=ARouter()
router.Use("/", anotherRouter)
```

##### **`Terminal`** (`type Terminal func(Request, Response)`)
A function, receiving `Request` and `Response` as parameter.  
It is a short form of Midware and will never pass request. So it does not have return value and quiet easy to code.

```go
router.Use("/", func(req Request, res Response) {
    res.Send("Hello, World!")
})
```

#### `func (*router)Get(mountpoint string, processor Tout) Router`
Similar to `Router.Use()` but differs in two points:

- Mountpoint must match the whole word, not the prefix to trigger the rule.
- Only GET method will trigger the rule.

#### `func (*router)Post(mountpoint string, processor Tout) Router`
Similar to `Router.Get()` but triggered by POST request.

#### `func (*router)Put(mountpoint string, processor Tout) Router`
Similar to `Router.Get()` but triggered by PUT request.

#### `func (*router)Delete(mountpoint string, processor Tout) Router`
Similar to `Router.Get()` but triggered by DELETE request.

#### `func (*router)UseSpecified(mountpoint string, method string, processor Tout, isStrict bool) Router`
General version of `Router.Use()`/`Router.Get()`/`Router.Put()`/`Router.Delete()`/`Router.Post()`.  

- `method` specifies the trigger method. Empty string means WILDCARD.  
- `isStrict` indicates whether the match is performed strictly.
