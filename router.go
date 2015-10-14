package gurgling

import (
    "net/http"
    "github.com/levythu/gurgling/matcher"
    . "github.com/levythu/gurgling/definition"
)

// Indeed an interface.
type Router interface {
    Use(string, Tout) Router
    Get(string, Tout) Router
    Post(string, Tout) Router
    Put(string, Tout) Router
    Delete(string, Tout) Router
    UseSpecified(string, string/*=""*/, Tout, bool) Router
    ServeHTTP(http.ResponseWriter, *http.Request)
    Handler(Request, Response) (bool, Request, Response)
}

// Midware is a handler that could modify anything and make the data flow continue by
// returning manipulated Req, Res and true. otherwise return false and the Res/Req is
// not specified.
type Midware func(Request, Response) (bool, Request, Response)
// a interface version of midware
type IMidware interface {
    Handler(Request, Response) (bool, Request, Response)
}
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

// processor must be a IMidware(including Router)/Midware or Terminal, otherwise panic.
// MountPoint should be valid, otherwise panic.
func (this *router)Use(mountpoint string, processor Tout) Router {
    return this.UseSpecified(mountpoint, "", processor, false)
}

// a GET specified version for use
func (this *router)Get(mountpoint string, processor Tout) Router {
    return this.UseSpecified(mountpoint, "GET", processor, true)
}

// a POST specified version for use
func (this *router)Post(mountpoint string, processor Tout) Router {
    return this.UseSpecified(mountpoint, "GET", processor, true)
}

// a POST specified version for use
func (this *router)Put(mountpoint string, processor Tout) Router {
    return this.UseSpecified(mountpoint, "PUT", processor, true)
}

// a POST specified version for use
func (this *router)Delete(mountpoint string, processor Tout) Router {
    return this.UseSpecified(mountpoint, "DELETE", processor, true)
}

// The detailed version of use. Default method is WILDCARD.
func (this *router)UseSpecified(mountpoint string, method string/*=""*/, processor Tout, isStrict bool) Router {
    if !this.mat.CheckRuleValidity(&mountpoint) {
        panic(INVALID_RULE)
    }

    switch processor:=processor.(type) {
    case IMidware:
        // Always use Midware as storage.
        this.mat.AddRule(mountpoint, method, Midware(func(req Request, res Response) (bool, Request, Response) {
            return processor.Handler(req, res)
        }), isStrict)
    case func(Request, Response):
        // Always use Midware as storage.
        this.mat.AddRule(mountpoint, method, Midware(func(req Request, res Response) (bool, Request, Response) {
            processor(req, res)
            return false, nil, nil
        }), isStrict)
    case Terminal:
        // Always use Midware as storage.
        this.mat.AddRule(mountpoint, method, Midware(func(req Request, res Response) (bool, Request, Response) {
            processor(req, res)
            return false, nil, nil
        }), isStrict)
    case func(Request, Response) (bool, Request, Response):
        this.mat.AddRule(mountpoint, method, Midware(processor), isStrict)
    case Midware:
        this.mat.AddRule(mountpoint, method, processor, isStrict)
    default:
        panic(INVALID_INVALID_USE)
    }
    /*
    if des, ok:=processor.(IMidware); ok {
        // Always use Midware as storage.
        this.mat.AddRule(mountpoint, method, Midware(func(req Request, res Response) (bool, Request, Response) {
            return des.Handler(req, res)
        }))
    } else if des, ok:=processor.(Midware); ok {
        this.mat.AddRule(mountpoint, method, des)
    } else if des, ok:=processor.(Terminal); ok {
        // Always use Midware as storage.
        this.mat.AddRule(mountpoint, method, Midware(func(req Request, res Response) (bool, Request, Response) {
            des(req, res)
            return false, nil, nil
        }))
    } else {
        panic(INVALID_INVALID_USE)
    }
    */
    return this
}

func (this *router)ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var req=NewRequest(r, this.initMountPoint)
    var res=NewResponse(w)
    this.Handler(req, res)
}

func (this *router)Handler(req Request, res Response) (bool, Request, Response) {
    var workstatus Tout=nil
    var result Tout
    var oldPath=req.Path()
    var oldBase=req.BaseUrl()
    for {
        result, workstatus=this.mat.Match(req.p2Path(), req.p2BaseUrl(), req.F(), req.Method(), workstatus)
        if result==nil {
            // No any more match, return 404
            res.Status("Resource not found. \nby gurgling", 404)
            return false, req, res
        }
        var isContinue, newReq, newRes=(result.(Midware))(req, res)
        if !isContinue {
            // Match and is indicated not to continue. Exit.
            return false, nil, nil
        }
        *(newReq.p2Path())=oldPath
        *(newReq.p2BaseUrl())=oldBase
        // Refresh req/res, continue to match the next
        req=newReq
        res=newRes
    }
}
