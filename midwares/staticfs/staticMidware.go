package staticfs

import (
    . "github.com/levythu/gurgling"
    "github.com/levythu/gurgling/midwares/urlnormalizer"
    "os"
    "path"
    "strings"
    "time"
)

// StaticFS provide packages for static file io and basic caching.
// If the file exist, just send the file. Otherwise pass it to the next handler.

type FsMidware struct {
    // implements IMidware
    basePath string
    cacheControl CacheStrategy
    // if non-empty, send the file representing a directory.
    RenderIndex string
    // if RenderIndex=="" or the file does not exist, use the default render.
    DefaultRender func(Request, Response, string)
}

var sanitizer=urlnormalizer.ASanitizer()

// ignoring details about the class itself
// making a 120-seconds caching fs-midware
func AStaticfs(basePath string) IMidware {
    return &FsMidware {
        basePath: basePath,
        cacheControl: CacheStrategy(120),
        RenderIndex: "index.html",
        DefaultRender: nil,
    }
}

func assert(err error) {
    if err!=nil {
        panic(err)
    }
}

func (this *FsMidware)Handler(req Request, res Response) (bool, Request, Response) {
    var isContinue bool
    isContinue, req, res=sanitizer.Handler(req, res)
    if !isContinue {
        return false, nil, nil
    }

    var targetFile=path.Join(this.basePath, req.Path())
    var fileMeta, err=os.Stat(targetFile)
    if err!=nil {
        // file does not exist or other errors.
        return true, req, res
    }
    if fileMeta.IsDir() && (this.DefaultRender!=nil || this.RenderIndex!="") {
        // is direction, check the last "/"
        if !strings.HasSuffix(req.Path(), "/") {
            // add "/"
            res.Redirect(req.OriginalPath()+"/")
            return false, nil, nil
        }
        if this.RenderIndex!="" {
            var indexToResolve=path.Join(this.basePath, req.Path(), this.RenderIndex)
            var itrMeta, err=os.Stat(indexToResolve)
            if err==nil && !itrMeta.IsDir() {
                this.handleFile(req, res, indexToResolve, itrMeta)
                return false, nil, nil
            }
        }
        if this.DefaultRender!=nil {
            this.DefaultRender(req, res, targetFile)
            return false, nil, nil
        }
        return true, req, res
    }
    // is a file
    this.handleFile(req, res, targetFile, fileMeta)
    return false, nil, nil
}

// handle the cache and manage data transmission
func (this *FsMidware)handleFile(req Request, res Response, filename string, fileinfo os.FileInfo) {
    const timeFormat="Mon, 19 Nov 2012 08:38:01 GMT"
    assert(res.Set(HEADER_CACHE_CONTROL, this.cacheControl.String()))
    assert(res.Set(HEADER_LAST_MODIFIED, fileinfo.ModTime().Format(timeFormat)))

    var strategy=req.GetAll(HEADER_CACHE_CONTROL)
    if strategy!=nil {
        for _, e:=range strategy {
            var t=strings.ToLower(e)
            if t=="no-cache" || t=="no-store" {
                // do not check cache.
                var err=res.SendFile(filename)
                if err!=nil && err!=SENDFILE_SENT_BUT_ABORT {
                    res.Status("Internal error while reading file", 500)
                }
                return
            }
        }
    }

    if mtime:=req.Get(HEADER_MODIFICATION_TIMESTAMP); mtime!="" {
        if ts, err:=time.Parse(timeFormat, mtime); err==nil {
            if !ts.Before(fileinfo.ModTime()) {
                // File not modified. return 304
                assert(res.SendCode(304))
                return
            }
        }
    }

    var err=res.SendFile(filename)
    if err!=nil && err!=SENDFILE_SENT_BUT_ABORT {
        res.Status("Internal error while reading file", 500)
    }
}
