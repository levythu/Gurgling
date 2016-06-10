# Gurgling 0.6.0
An extremely-light framework for Golang to build restful API and Website.

**0.6.0 NEWS! Websocket is AVAILABLE!**

**Special Thanks to [Express](http://expressjs.com/), which provides API samples for this project.**

## Quick Start

### Install it

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
    router.Get(func(res Response) {
        res.Send("Hello, World!")
    })

    // Launch the server
    fmt.Println("Running...")
    router.Launch(":8080")
}
```

### Or simply chain them

```go
package main

import (
    . "github.com/levythu/gurgling"
)

func main() {
    ARouter( ).Get(func(res Response) {
                res.Send("Hello, World!")
            }).Launch(":8080")
}
```

### Then, create sub-routers and mount them

```go

// Create a non-gate router.
var pageRouter=ARouter()

// Mount handler and midware
pageRouter.Use(func(req Request, res Response) bool {
    fmt.Println(req.Path())
	return true
})
pageRouter.Get("/editor", func(req Request, res Response) {
    res.Send("Here's the editor.")
})

// Mount the router to the previous one
router.Use("/page", pageRouter)
```
### Want some regular expressions? Here you are!
```go
package main

import (
    "fmt"
    . "github.com/levythu/gurgling"
)

func main() {
    var router=ARegexpRouter()

    router.Get(`/(\d+)`, func(req Request, res Response) {
        // the submatches are stored in req.F()["RR"] as []string
        var matchResult=req.F()["RR"].([]string)
        res.Send("The digits are "+matchResult[1])
    })
    router.Get(`/.*`, func(req Request, res Response) {
        res.Send("Please visit paths consisting of digits.")
    })

    fmt.Println("Running...")
    router.Launch(":8192")
}
```

### A HTTPS server is simple as well
```go
package main

import (
    . "github.com/levythu/gurgling"
)

func main() {

    var router=ARouter()
    router.Get(func(res Response) {
        res.Send("Hello, World!")
    })

    // Launch the server, using prepared certificate and private key
    fmt.Println("Running...")
    router.Launch(":8192", "public/server.crt", "private/server.key")
}
```

## API Docs
### Router
The core module of gurgling. It is indeed an interface, and is implemented by `router`, its default and original version. The interface is designed for extensions.

#### `func GetRouter(MountPoint string, matchHandler matcher.Matcher) Router`
Creates and returns one default router for gateway. The mountpoint it is mounted to in `http.Handle()` function should be specified here.  
The matcher needs to be specified. It could be either `BFMatcher` or `RegexpMatcher`.

#### `func ARouter() Router`
Creates and returns one default router with `mountpoint="/"`, which is the default mountpoint for `Router.Launch()`

#### `func ARegexpRouter() Router`
Creates and returns one default router with `mountpoint="/"`. It differs from `ARouter` in that its returned router supports regexp rule. Of course, this leads to longer time to match. Thus, for non-regexp use, create router with `ARouter()` instead.

#### `func (Router)Launch(addr string) error`
Invokes `net/http` to launch the HTTP server at `addr`. This function is supposed to keep running unless an error is encountered.

#### `func (Router)Launch(addr string, certFile, keyFile string) error`
Invokes `net/http` to launch the HTTPS server at `addr`. `certFile` and `keyFile` are the filepaths of certificate and private key file required.

#### `func (Router)Launch([certFile, keyFile string, [setAHTTPRedirector bool]]) error`
If all the parameters are omitted, a HTTP server is launched at port 80.
If `certFile` and `keyFile` are given, a HTTPS server is launched at port 443. When `setAHTTPRedirector` is given and set to `true`, a HTTP server is also launched at port 80 automatically to redirect all http traffic to https.

#### `func (Router)Use([mountpoint string], processor Tout) Router`
Mounts a mountable to the router at mountpoint. Mountpoint must start with `/`. Create a router with matcher `RegexpMatcher` could make regexp supported, but **DO NOT** use first-char or last-char matcher (`^` or `$`). It will try to match the mountpoint by prefix.   
Returns the router itself for method chaining.  
**NOTE: the first parameter could be omitted and will be set to "/" by default**

Mountable includes the following items:

##### **`Midware`** (`type Midware func(Request, Response) (bool, Request, Response)`)  
A function, receiving `Request` and `Response` as parameter.  
Returns three value, two of which are modified res&req (if no modification just return the original, or use `Hopper`) and the boolean indicates whether to pass the request to the next handler (`false` for not).

```go
router.Use("/", func(req Request, res Response) (bool, Request, Response) {
    fmt.Println(req.Path())
	// PASS the request to the next handler.
	return true, req, res
})
```

##### **`IMidware`**  
An interface which implement `Midware` function as `.Handler()`.  
Since `Router` also implement the function, `Router` is a special IMidware, which never passes request to the next.

```go
var anotherRouter=ARouter()
router.Use("/", anotherRouter)
```

##### **`Hopper`** (`type Hopper func(Request, Response) bool`)  
Simplified version of `Midware`. It will not modify original `res` and `req`.

```go
router.Use("/", func(req Request, res Response) bool {
    fmt.Println(req.Path())
	// Stop passing it.
	return false
})
```

##### **`Terminal`** (`type Terminal func(Request, Response)`)
A function, receiving `Request` and `Response` as parameter.  
It is a simplified form of Midware and will never pass request. So it does not have return value and quite easy to code.

```go
router.Use("/", func(req Request, res Response) {
    res.Send("Hello, World!")
})
```

##### **`Dump`** (`type Terminal func(Response)`)
A function, receiving only `Response` as parameter, which means the codes in the function do not need request anymore.
It is a simplified form of Terminal.

```go
router.Use("/", func(req Request, res Response) {
    res.Send("Hello, World!")
})
```

#### `func (Router)Last(processor Cattail) Router`
Mount a cattail to the end of the router.  
`type Cattail func(Request, Response)` is a function that will always get executed after the request is handled by normal routers and midwares. In such circumstance, the response header is certainly to be sent. So most methods of `Response` are invalid. However, it is still provided for appending data, although not recommended.  
Like `Router.Use()`, all the Cattails will get executed in the order they are mounted in codes.

#### `func (Router)Get([mountpoint string,] processor Tout) Router`
Similar to `Router.Use()` but differs in two points:

- Mountpoint must match the whole path, not the prefix to trigger the rule.
- Only GET method will trigger the rule.

#### `func (Router)Post([mountpoint string,] processor Tout) Router`
Similar to `Router.Get()` but triggered by POST request.

#### `func (Router)Put([mountpoint string,] processor Tout) Router`
Similar to `Router.Get()` but triggered by PUT request.

#### `func (Router)Delete([mountpoint string,] processor Tout) Router`
Similar to `Router.Get()` but triggered by DELETE request.

#### `func (Router)UseSpecified(mountpoint string, method string, processor Tout, isStrict bool) Router`
General version of `Router.Use()`/`Router.Get()`/`Router.Put()`/`Router.Delete()`/`Router.Post()`.  

- `method` specifies the trigger method. Empty string means WILDCARD.  
- `isStrict` indicates whether the match is performed strictly. (Matches whole path or prefix)

#### `func (Router)SetErrorHandler(handler RouterErrorCatcher) Router`
Set the runtime error handler to recover from panic. The handler is `func(Request, Response, interface{})`, the third parameter of which is the panic content. Note that if there is any panic in the handler, the whole program will not avoid suffering.  
The default handler is `nil`, which will not recover from panic, thus good for debugging. However, it is recommended to use `DefaultCacher` for the root router, which is to render a `500 Internal Error` page to client.

#### `func (Router)Set404Handler(handler Terminal) Router`
Set the 404 handler. When no rules matches the request, the handler will get executed. If set to `nil`, the server will return a string "404 NOT FOUND" by default.

### Response
The interface provided in Handler callback wrapping functions for quick response, in Express format. Since it is an interface, further hack by midwares is possible.

#### `func (Response)Status(content string, code int) error`
Quick send a message and specify the return code. Note that if any head-sending operation like `Response.Status`, `Response.Send`, `Response.SendFile`, `Response.SendFileEx`, `Rsponse.SendCode`, `Response.Write`, `Response.Redirect`, `Response.JSON`, `Response.JSONEx` has been invoked before, this one will fail and returns `RES_HEAD_ALREADY_SENT`.

It will set `Content-Type: text/plain; charset=utf-8`.

#### `func (Response)Send(content string) error`
Quick invocation for `Response.Status`, using code 200. Can only be invoked successfully without any preceding head-sending invoking.

#### `func (Response)RedirectEX(newAddr string, code int) error`
Redirect to newAddr by returning `code` as status code. Can only be invoked successfully without any preceding head-sending invoking.

#### `func (Response)Redirect(newAddr string) error`
Redirect to newAddr by returning 307 (Moved Temporarily). Can only be invoked successfully without any preceding head-sending invoking.

#### `func (Response)JSON(obj interface{}) error`
Stringify the obj and send it with status code 200. Can only be invoked successfully without any preceding head-sending invoking.

#### `func (Response)JSONEx(obj interface{}, code int) error`
Stringify the obj and send it with status code `code`. Can only be invoked successfully without any preceding head-sending invoking.

#### `func (Response)SendCode(code int) error`
Sending the code without any body content. Can only be invoked successfully without any preceding head-sending invoking.

#### `func (Response)SendFileEx(filepath string, mime string, encoder encoding.Encoder, code int) error`
Sending a local file with content encoding specified by `encoder`, http status code `code` and MIME type `mime`.

- `filepath`: the path for the file to be sent.
- `mime`: if left empty, gurgling will try to infer it according to file extension.
- `encoder`: if nil, a default encoder will be used, which sends the data as they are. For more encoders refer to `github.com/levythu/gurgling/encoding`
- `code`: http status code.

Can only be invoked successfully without any preceding head-sending invoking. The errors are the following:

- `SENDFILE_ENCODER_NOT_READY`: encoder fails to manipulate data.
- `SENDFILE_FILEPATH_ERROR`: fail to open target file.
- `SENT_BUT_ABORT`: start to send but then abort. **Note that in such case the header was sent, so many operation which require no preceding head-sending invoking will fail.**

#### `func (Response)SendFile(filepath string) error`
Quick invocation for `Response.SendFileEx`, using code 200 and gzip compressor. MIME will be inferred. Can only be invoked successfully without any preceding head-sending invoking.

Example:

```go
var page=ARouter()
page.Use("/file", func(req Request, res Response) {
    if req.Path()=="" {
        res.Send("Specify the path please.")
        return
    }
    // Eliminating the prefix "/"
    res.SendFile(req.Path()[1:])
})
```

#### `func (Response)Set(key string, val string) error`
Sets the header. Can only be invoked successfully without any preceding head-sending invoking.

#### `func (Response)Get(key string) string`
Looks up the header. If the key does not exist, an empty string will be returned.

#### `func (Response)Write(data []byte) (int, error)`
The original interface of `ResponseWriter`. Can be invoked any time but when the header has not been sent, it will sent the header with code=200 automatically.

#### `func (Response)R() http.ResponseWriter`
Returns the wrapped original `ResponseWriter` for advanced use.

#### `func (Response)F() map[string]Tout`
`Tout` is `interface{}`. It is preserved for any use by midwares, e.g., storing extracted data or adding functions.

### Request
The interface provided in Handler callback wrapping functions for convenient access to request information. Note that most of the data are read-only.

#### `func (Request)Path() string`
Returns the current path the Mountable is working on. It usually starts with "/". Note that the mountpoint is excluded and can be found in baseUrl.

Example:
```go
var router=ARouter()
var subrouter=ARouter()

router.Use("/", func(req Request, res Response) (bool, Request, Response) {
    fmt.Println("<"+req.BaseUrl()+">", ",", "<"+req.Path()+">")
    // pass the request.
    return true, req, res
})
router.Use("/sub", subrouter)
subrouter.Use("/", func(req Request, res Response) (bool, Request, Response) {
    fmt.Println("<"+req.BaseUrl()+">", ",", "<"+req.Path()+">")
    return true, req, res
})

LaunchServer(":8192", router)
```
Run it:
> GET http://localhost:8192
```
<> , </>
```
> GET http://localhost:8192/sub
```
<> , </sub>
</sub> , <>
```
> GET http://localhost:8192/sub/
```
<> , </sub/>
</sub> , </>
```
> GET http://localhost:8192/sub/another
```
<> , </sub/another>
</sub> , </another>
```

#### `func (Request)BaseUrl() string`
Returns the base Url. See the examples above.

#### `func (Request)OriginalPath() string`
Returns the original path.

#### `func (Request)Hostname() string`
Returns the requesting hostname, including port number if specified.

#### `func (Request)Query() map[string]string`
Returns the query map. If one key is set multiple times, only the first setting will be valid.

#### `func (Request)Method() string`
Returns the method of the request, e.g, `GET` or `POST`

#### `func (Request)Get(key string) string`
Looks up values in HEADER. If not specified, an empty string will be returned.

#### `func (Request)R() *http.Request`
Returns the wrapped original `*Request` for advanced use. Please note that modifying its URL member will not lead modification in `Path`/`BaseUrl`, but `OriginalPath`, `Hostname` will be affected.

#### `func (Request)F() map[string]Tout`
`Tout` is `interface{}`. It is preserved for any use by midwares, e.g., storing extracted data or adding functions.

## Attachment #1: Components List
### Midwares
Midwares only modify the request & response to assist implementing some functions. Typically, subsequent handler is needed to mount on the router which a midware is mounted on.

- `github.com/levythu/gurgling/midwares/analyzer`: A set of analyzer for logging and timing each request.
- `github.com/levythu/gurgling/midwares/bodyparser`: A midware trying to parse the body of any kind to JSON, Forms or simply Bytes buffer, typically mounted to a POST API.
- `github.com/levythu/gurgling/midwares/cookie`: A midware to parse cookies from the request and provide setting API. It also provides signed cookie and session components.
- `github.com/levythu/gurgling/midwares/staticfs`: A midware to look up the path of each GET request from one specified directory and return the file if existing. Otherwise, pass the request to the next handler. Cache control is implemented.
- `github.com/levythu/gurgling/midwares/urlnormalizer`: Normalizing the url, e.g., sanitizing it to remove redundant slashes and dots.

### Routers
Routers are terminals, all the other components (except `LAST`) mounted on the router will never get triggered.

- `github.com/levythu/gurgling/routers/simplefsserver`: A simple static file server, which mainly relies on `staticfs` midware. Additionally, it provides 404 page and auto indexing option.
- `github.com/levythu/gurgling/routers/websocket`: Router implementing websocket protocol and providing simple interface for full-duplex commnunication between server and client.
