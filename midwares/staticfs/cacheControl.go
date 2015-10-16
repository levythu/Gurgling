package staticfs

// provide options for cache control

// CacheControl is the timespan (in seconds) for browser to keep the resource without sending GET request.
// max-age=xxx sets it.
// no-cache means max-age=0
// no-store force the browser to fetch the data all the time

import (
    "net/http"
    "strconv"
)

type CacheStrategy int64

func (this CacheStrategy)String() string {
    var val=int64(this)
    if val<=-2 {
        return "no-store"
    } else if val==-1 {
        return "no-cache"
    }
    return "max-age="+strconv.FormatInt(val, 10)
}
func (this CacheStrategy)Cached() bool {
    var val=int64(this)
    return val>=0
}

const (
    NO_CACHE=CacheStrategy(-1)
    NO_STORE=CacheStrategy(-2)
)

var (
    HEADER_CACHE_CONTROL=http.CanonicalHeaderKey("Cache-Control")
    HEADER_MODIFICATION_TIMESTAMP=http.CanonicalHeaderKey("If-Modified-Since")
    HEADER_LAST_MODIFIED=http.CanonicalHeaderKey("Last-Modified")
)
