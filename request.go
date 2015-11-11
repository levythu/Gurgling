package gurgling

import (
    "net/http"
    "strings"
    . "github.com/levythu/gurgling/definition"
)

type Request interface {
    //  Path is the url not containing mount point.
    Path() string
    // pointer to path for modification
    p2Path() *string

    // BaseUrl is the mount point
    BaseUrl() string
    // pointer to BaseUrl for modification
    p2BaseUrl() *string

    // OriginalPath is the full URL requested.
    OriginalPath() string

    // the Hostname of the request
    Hostname() string

    // the query of requst
    // the original query processor is map[string][]string, this one just
    // reserve the first setting.
    Query() map[string]string

    // Preserved for midware. By default it will not return a ReadCloser for
    // raw read.
    Body() Tout

    // the method of the request
    Method() string

    // get data in the headers, if not specified, return ""
    Get(string) string
    GetAll(key string) []string

    // get the Original resonse, only use it for advanced purpose
    R() *http.Request

    // extra use for midwares. Most of time the value is a function.
    F() map[string]Tout

    // nonexist returns ""
    Referer() string
}

// Return a OriRequest, which acts every default behavior
func NewRequest(res *http.Request, mountpoint string) Request {
    var ret=&OriRequest{
        r: res,
        f: make(map[string]Tout),
        path: res.URL.Path,
        baseurl: mountpoint,
    }
    ret.parsedQuery=make(map[string]string)
    var tquery=res.URL.Query()
    for k, v:=range tquery {
        if len(v)>0 {
            ret.parsedQuery[k]=v[0]
        }
    }

    return ret
}

type OriRequest struct {
    r *http.Request
    f map[string]Tout

    path string
    baseurl string

    parsedQuery map[string]string
}
func (this *OriRequest)Path() string {
    return this.path
}
func (this *OriRequest)p2Path() *string {
    return &this.path
}
func (this *OriRequest)BaseUrl() string {
    return this.baseurl
}
func (this *OriRequest)p2BaseUrl() *string {
    return &this.baseurl
}
func (this *OriRequest)OriginalPath() string {
    return this.r.URL.Path
}
func (this *OriRequest)Hostname() string {
    return this.r.URL.Host
}
func (this *OriRequest)Query() map[string]string {
    return this.parsedQuery
}
func (this *OriRequest)Body() Tout {
    return this.F()["body"]
}
func (this *OriRequest)Method() string {
    return strings.ToUpper(this.r.Method)
}
func (this *OriRequest)Get(key string) string {
    return this.r.Header.Get(key)
}
func (this *OriRequest)GetAll(key string) []string {
    return this.r.Header[http.CanonicalHeaderKey(key)]
}
func (this *OriRequest)R() *http.Request {
    return this.r
}
func (this *OriRequest)F() map[string]Tout {
    return this.f
}
func (this *OriRequest)Referer() string {
    return this.r.Referer()
}
