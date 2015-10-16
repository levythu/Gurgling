package gurgling

import (
    "net/http"
    "github.com/levythu/gurgling/matcher"
    . "github.com/levythu/gurgling/definition"
    "io"
)

// Indeed an interface.
type Router interface {
    Use(paraList ...interface{}) Router
    Get(paraList ...interface{}) Router
    Post(paraList ...interface{}) Router
    Put(paraList ...interface{}) Router
    Delete(paraList ...interface{}) Router
    Last(Cattail) Router
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

// a midware fixing on the rear.
type Cattail func(Request, Response)

// one sandwich means mounting a midware and a cattail
type Sandwich interface {
    IMidware
    Final(Request, Response)
}

type router struct {
    // Implementing http.Handler
    mountMap map[string]*router
    initMountPoint string
    mat matcher.Matcher
    tailList []Cattail
}

// The error will be fatal.
func GetRouter(MountPoint string) Router {
    if !checkMountpointValidity(&MountPoint) {
        panic(INVALID_MOUNT_POINT)
    }
    return &router {
        mountMap: make(map[string]*router),
        initMountPoint: MountPoint,
        mat: matcher.NewBFMatcher(),
    }
}

// No mountpoint, used only for router not for connecting go http server
const not_init_mounter="$NOT_VALID$"
func ARouter() Router {
    return &router {
        mountMap: make(map[string]*router),
        initMountPoint: "",
        mat: matcher.NewBFMatcher(),
        tailList: []Cattail{},
    }
}

// processor must be a IMidware(including Router)/Midware or Terminal, otherwise panic.
// MountPoint should be valid, otherwise panic.
//func (this *router)Use(mountpoint string, processor Tout) Router
//func (this *router)Use(processor Tout) Router
func (this *router)Use(paraList ...interface{}) Router {
    var mountpoint string
    var processor Tout
    mountpoint, processor=extractParameters(paraList...)
    return this.UseSpecified(mountpoint, "", processor, false)
}

// a GET specified version for use
func (this *router)Get(paraList ...interface{}) Router {
    var mountpoint string
    var processor Tout
    mountpoint, processor=extractParameters(paraList...)
    return this.UseSpecified(mountpoint, "GET", processor, true)
}

func (this *router)Last(process Cattail) Router {
    this.tailList=append(this.tailList, process)
    return this
}

// a POST specified version for use
func (this *router)Post(paraList ...interface{}) Router {
    var mountpoint string
    var processor Tout
    mountpoint, processor=extractParameters(paraList...)
    return this.UseSpecified(mountpoint, "GET", processor, true)
}

// a POST specified version for use
func (this *router)Put(paraList ...interface{}) Router {
    var mountpoint string
    var processor Tout
    mountpoint, processor=extractParameters(paraList...)
    return this.UseSpecified(mountpoint, "PUT", processor, true)
}

// a POST specified version for use
func (this *router)Delete(paraList ...interface{}) Router {
    var mountpoint string
    var processor Tout
    mountpoint, processor=extractParameters(paraList...)
    return this.UseSpecified(mountpoint, "DELETE", processor, true)
}

// The detailed version of use. Default method is WILDCARD.
func (this *router)UseSpecified(mountpoint string, method string/*=""*/, processor Tout, isStrict bool) Router {
    if !this.mat.CheckRuleValidity(&mountpoint) {
        panic(INVALID_RULE)
    }

    switch processor:=processor.(type) {
    case Sandwich:
        // Can only be mounted to the root("/")
        if (mountpoint!="" && mountpoint!="/") || isStrict || method!="" {
            panic(SANDWICH_MOUNT_ERROR)
        }
        this.mat.AddRule(mountpoint, method, Midware(func(req Request, res Response) (bool, Request, Response) {
            return processor.Handler(req, res)
        }), isStrict)
        this.Last(func(req Request, res Response) {
            processor.Final(req, res)
        })
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
    if this.initMountPoint==not_init_mounter {
        panic(INVALID_MOUNT_POINT)
    }
    var req=NewRequest(r, this.initMountPoint)
    var res=NewResponse(w)
    this.Handler(req, res)
}

const content="<!DOCTYPE html><html><head><title>404 Not Found" +
    "</title><style>h2 {text-align: center;}table {max-width: 35em;margin: 0 auto;padding-top: 1em;font: 1em Arial, Helvetica, sans-serif;}a {text-decoration: none;color: #0070C0;}p {padding-top: 1em;}</style></head><body><h2>404 Not Found"+
    "</h2><table>"+
    "<tr><td><p>by <a href=\"https://github.com/levythu/gurgling\">Gurgling "+Version+"</a></p></td></tr></table></body></html>"
func (this *router)Handler(req Request, res Response) (bool, Request, Response) {
    defer func() {
        for _, elem:=range this.tailList {
            elem(req, res)
        }
    }()
    var workstatus Tout=nil
    var result Tout
    var oldPath=req.Path()
    var oldBase=req.BaseUrl()
    for {
        result, workstatus=this.mat.Match(req.p2Path(), req.p2BaseUrl(), req.F(), req.Method(), workstatus)
        if result==nil {
            // No any more match, return 404
            res.Set("Content-Type", "text/html; charset=utf-8")
            res.SendCode(404)
            io.WriteString(res, content)
            return false, nil, nil
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
