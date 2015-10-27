package analyzer

import (
    . "github.com/levythu/gurgling"
    . "github.com/levythu/gurgling/definition"
    "github.com/levythu/gurgling/encoding"
    "net/http"
)

// A wrapper for simple Response to Hack the data, recording some events

type logResponse struct {
    // implement response
    o Response

    // triggered when the header is successfully sent. Except redirect.
    OnHeadSent func(Response, int)
}

func (this *logResponse)Send(content string) error {
    var err=this.o.Send(content)
    if err==nil || err==SENT_BUT_ABORT {
        this.OnHeadSent(this.o, 200)
    }
    return err
}
func (this *logResponse)Set(k string, v string) error {
    return this.o.Set(k, v)
}
func (this *logResponse)Get(k string) string {
    return this.o.Get(k)
}
func (this *logResponse)Write(c []byte) (int, error) {
    return this.o.Write(c)
}
func (this *logResponse)SendFileEx(f string, d string, e encoding.Encoder, c int) error {
    var err=this.o.SendFileEx(f, d, e, c)
    if err==nil || err==SENT_BUT_ABORT {
        this.OnHeadSent(this.o, c)
    }
    return err
}
func (this *logResponse)SendFile(f string) error {
    var err=this.o.SendFile(f)
    if err==nil || err==SENT_BUT_ABORT {
        this.OnHeadSent(this.o, 200)
    }
    return err
}
func (this *logResponse)Status(d string, c int) error {
    var err=this.o.Status(d, c)
    if err==nil || err==SENT_BUT_ABORT {
        this.OnHeadSent(this.o, c)
    }
    return err
}
func (this *logResponse)SendCode(c int) error {
    var err=this.o.SendCode(c)
    if err==nil || err==SENT_BUT_ABORT {
        this.OnHeadSent(this.o, c)
    }
    return err
}
func (this *logResponse)Redirect(d string) error {
    var err=this.o.Redirect(d)
    if err==nil || err==SENT_BUT_ABORT {
        this.OnHeadSent(this.o, 307)
    }
    return err
}
func (this *logResponse)RedirectEX(d string, c int) error {
    var err=this.o.RedirectEX(d, c)
    if err==nil || err==SENT_BUT_ABORT {
        this.OnHeadSent(this.o, c)
    }
    return err
}
func (this *logResponse)R() http.ResponseWriter {
    return this.o.R()
}
func (this *logResponse)F() map[string]Tout {
    return this.o.F()
}
func (this *logResponse)JSON(obj interface{}) error {
    var err=this.o.JSON(obj)
    if err==nil || err==SENT_BUT_ABORT {
        this.OnHeadSent(this.o, 200)
    }
    return err
}
func (this *logResponse)JSONEx(obj interface{}, c int) error {
    var err=this.o.JSONEx(obj, c)
    if err==nil || err==SENT_BUT_ABORT {
        this.OnHeadSent(this.o, c)
    }
    return err
}
