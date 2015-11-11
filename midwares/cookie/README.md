# Gurgling - Cookie
Cookie and Session handler midware for Gurgling. Introduced in Gurgling 0.4.0.

## Quick Start

### Import it

```go
import "github.com/levythu/gurgling/midwares/cookie"
```
### Use it to manipulate cookie

```go
func main() {
    var router=ARouter().Use(cookie.ACookie())

    // Mount one handler
    router.Get("/get", func(req Request, res Response) {
        var c=req.F()["cookie"].(*cookie.CookieHandler)
        if result:=c.Get("key"); result=="" {
            res.Send("No cookie found.")
        } else {
            res.Send("The cookie is "result)
        }
    })
    router.Get("/set", func(req Request, res Response) {
        var c=req.F()["cookie"].(*cookie.CookieHandler)
        c.Set("key", "DEMO")
        res.Send("Cookie is set.")
    })

    // Launch the server
    fmt.Println("Running...")
    router.Launch(":8080")
}
```

### Or handle session

```go
func main() {
    var router=ARouter().Use(cookie.ASession("do-not-tell-others"))

    router.Get(func(req Request, res Response) {
        var session=req.F()["session"].(map[string]string)
        if v, ok:=session["user"]; !ok {
            res.Send("Please login first.")
        } else {
            res.Send("Hello, "+v+"!")
        }
    })
    router.Get("/logout", func(req Request, res Response) {
        req.F()["session"]=nil
        res.Send("Logouted.")
    })
    router.Use("/login", func(req Request, res Response) {
        var session=req.F()["session"].(map[string]string)
        if len(req.Path())<=1 {
            res.Send("Please specify the user.")
        } else {
            session["user"]=req.Path()[1:]
            res.Send("Logined.")
        }
    })

    fmt.Println("Running...")
    router.Launch(":8080")
}
```

## API Docs

To be detailed...
