package gurgling

import (
    "net/http"
    "github.com/levythu/gurgling/matcher"
)

// Indeed an interface.
type Router interface {
    Use(string, Tout)
    Get(string, Tout)
    Post(string, Tout)
    Put(string, Tout)
    Delete(string, Tout)
    UseSpecified(string, string/*=""*/, Tout)
    ServeHTTP(http.ResponseWriter, *http.Request)
    Handler(Request, Response) (bool, Request, Response)
}

// Midware is a handler that could modify anything and make the data flow continue by
// returning manipulated Req, Res and true. otherwise return false and the Res/Req is
// not specified.
type Midware func(Request, Response) (bool, Request, Response)
// a midware that will always return (false, nil, nil)
type Terminal func(Request, Response)
// Implementing http.Handler
type router struct {
    mountMap map[string]*router
    initMountPoint string
    mat matcher.Matcher
}

// The error will be fatal.
func GetRouter(MountPoint string) Router {
    if !CheckMountpointValidity(&MountPoint) {
        panic(INVALID_MOUNT_POINT)
    }
    return &router {
        mountMap: make(map[string]*router),
        initMountPoint: MountPoint,
        mat: matcher.NewBFMatcher(),
    }
}

// processor must be a Router/Midware or Terminal, otherwise panic.
// MountPoint should be valid, otherwise panic.
func (this *router)Use(mountpoint string, processor Tout) {
    this.UseSpecified(mountpoint, "", processor)
}

// a GET specified version for use
func (this *router)Get(mountpoint string, processor Tout) {
    this.UseSpecified(mountpoint, "GET", processor)
}

// a POST specified version for use
func (this *router)Post(mountpoint string, processor Tout) {
    this.UseSpecified(mountpoint, "GET", processor)
}

// a POST specified version for use
func (this *router)Put(mountpoint string, processor Tout) {
    this.UseSpecified(mountpoint, "PUT", processor)
}

// a POST specified version for use
func (this *router)Delete(mountpoint string, processor Tout) {
    this.UseSpecified(mountpoint, "DELETE", processor)
}

// The detailed version of use. Default method is WILDCARD.
func (this *router)UseSpecified(mountpoint string, method string/*=""*/, processor Tout) {
    if !this.mat.CheckRuleValidity(&mountpoint) {
        panic(INVALID_RULE)
    }
    switch processor.(type) {
    case Terminal:
        // Always use Midware as storage.
        this.mat.AddRule(mountpoint, method, func(req Request, res Response) (bool, Request, Response) {
            (processor.(Terminal))(req, res)
            return false, nil, nil
        })
    case Midware:
        this.mat.AddRule(mountpoint, method, processor)
    case Router:
        // Always use Midware as storage.
        this.mat.AddRule(mountpoint, method, func(req Request, res Response) (bool, Request, Response) {
            return (processor.(Router)).Handler(req, res)
        })
    default:
        panic(INVALID_INVALID_USE)
    }
}

func (this *router)ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var req=NewRequest(r, this.initMountPoint)
    var res=NewResponse(w)
    this.Handler(req, res)
}

func (this *router)Handler(req Request, res Response) (bool, Request, Response) {
    var workstatus Tout=nil
    var result Tout
    for {
        result, workstatus=this.mat.Match(req.P2Path(), req.P2BaseUrl(), req.Method(), workstatus)
        if result==nil {
            // No any more match, return 404
            res.Status("Resource not found. \nby gurgling", 404)
            return false, req, res
        }
        var isContinue, newReq, newRes=(result.(Midware))(req, res)
        if !isContinue {
            // Match and is indicated not to continue. Exit.
            return false, newReq, newRes
        }
        // Refresh req/res, continue to match the next
        req=newReq
        res=newRes
    }
}
