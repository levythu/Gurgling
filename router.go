package gurgling

import (
    "net/http"
)

// Indeed an interface.
type Router interface {
    //TODO
}

// The callback is a function provided in Midware to pass the process to the next.
type MidwareCallback func(Request, Response)

// must call MidwareCallback to continue process
type Midware func(Request, Response, MidwareCallback)
// a midware that will never callback
type Terminal func(Request, Response)
// Implementing http.Handler
type router struct {
    mountMap map[string]*router
    initMountPoint string
}

// The error will be fatal.
func GetRouter(MountPoint string) Router {
    if !CheckMountpointValidity(MountPoint) {
        panic(INVALID_MOUNT_POINT)
    }
    return &router {
        mountMap: make(map[string]*router),
        initMountPoint: MountPoint,
    }
}

// processor must be a Router/Midware or Terminal, otherwise panic.
// MountPoint should be valid, otherwise panic.
func (this *router)Use(mountpoint string, processor Tout) {
    if !CheckMountpointValidity(MountPoint) {
        panic(INVALID_MOUNT_POINT)
    }

}

func (this *router)ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (this *router)Handler(req Request, res Response, next MidwareCallback) {

}
