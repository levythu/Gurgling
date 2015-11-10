package cookie

import (
    . "github.com/levythu/gurgling"
    "fmt"
)

type resSession struct {
    // implementing Response
    Response
    sessionInfo *Session
}

func (this *resSession)Send(cont string) error {
    fmt.Println(cont)
    return this.Response.Send(cont)
}
