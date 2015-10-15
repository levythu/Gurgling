package staticfs

// provide options for cache control

// CacheControl is the timespan (in seconds) for browser to keep the resource without sending GET request.
// max-age=xxx sets it.
// no-cache means max-age=0
// no-store force the browser to fetch the data all the time

type CacheStrategy int64

const (
    NO_CACHE=CacheStrategy(-1)
    NO_STORE=CacheStrategy(-2)
)
