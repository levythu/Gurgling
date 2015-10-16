package gurgling

import (
    "strings"
    "io"
)

// Check whether the mountpoint is valid, which must start with "/"
// If ends with "/", the function will remove it and return TRUE.
func checkMountpointValidity(mountpoint *string) bool {
    if !strings.HasPrefix(*mountpoint, "/") {
        return false
    }
    var ends=len(*mountpoint)
    for ends>0 && (*mountpoint)[ends-1]=='/' {
        ends--
    }
    *mountpoint=(*mountpoint)[:ends]
    return true
}

// used for quick invoke of use/get/...
// if one parameter, returns "/", p1
// if two, returns p1, p2
func extractParameters(paraList ...interface{}) (string, interface{}) {
    if len(paraList)==0 || len(paraList)>2 {
        panic(INVALID_PARAMETER)
    }
    if len(paraList)==1 {
        return "/", paraList[0]
    }
    var tmp, ok=paraList[0].(string)
    if !ok {
        panic(INVALID_PARAMETER)
    }
    return tmp, paraList[1]
}

// wrapper for write-only response
type writeOnly struct {
    // implement io.Writer
    w io.Writer
}
func (this *writeOnly)Write(p []byte) (n int, err error) {
    return this.w.Write(p)
}
func newWO(w io.Writer) io.Writer {
    return &writeOnly {
        w: w,
    }
}
