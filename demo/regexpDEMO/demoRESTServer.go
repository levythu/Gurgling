package main

import (
    "fmt"
    . "github.com/levythu/gurgling"
    "github.com/levythu/gurgling/matcher"
    "github.com/levythu/gurgling/midwares/analyzer"
)

func main() {
    var router=ARegexpRouter()

    router.Use(analyzer.ASimpleAnalyzer())
    router.Get(`/(\d+)`, func(req Request, res Response) {
        var matchResult=req.F()[matcher.REGEXP_RESULT].([]string)
        res.Send("The digits are "+matchResult[1])
    })
    router.Get(`/.*` ,func(req Request, res Response) {
        res.Send("Please visit paths consisting of digits.")
    })

    fmt.Println("Running...")
    router.Launch(":8192")
}
