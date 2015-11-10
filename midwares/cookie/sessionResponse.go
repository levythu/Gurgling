package cookie

import (
    . "github.com/levythu/gurgling"
    "github.com/levythu/gurgling/encoding"
)

type resSession struct {
    // implementing Response
    Response
    sessionInfo *Session
}

func (this *resSession)Send(cont string) error {
    this.sessionInfo.UpdateSession(this)
    return this.Response.Send(cont)
}

func (this *resSession)SendFileEx(f string, d string, e encoding.Encoder, c int) error {
    this.sessionInfo.UpdateSession(this)
    return this.Response.SendFileEx(f, d, e, c)
}

func (this *resSession)SendFile(f string) error {
    this.sessionInfo.UpdateSession(this)
    return this.Response.SendFile(f)
}

func (this *resSession)Status(d string, c int) error {
    this.sessionInfo.UpdateSession(this)
    return this.Response.Status(d, c)
}

func (this *resSession)SendCode(c int) error {
    this.sessionInfo.UpdateSession(this)
    return this.Response.SendCode(c)
}

func (this *resSession)Redirect(d string) error {
    this.sessionInfo.UpdateSession(this)
    return this.Response.Redirect(d)
}

func (this *resSession)RedirectEX(d string, c int) error {
    this.sessionInfo.UpdateSession(this)
    return this.Response.RedirectEX(d, c)
}

func (this *resSession)JSON(obj interface{}) error {
    this.sessionInfo.UpdateSession(this)
    return this.Response.JSON(obj)
}

func (this *resSession)JSONEx(obj interface{}, c int) error {
    this.sessionInfo.UpdateSession(this)
    return this.Response.JSONEx(obj, c)
}
