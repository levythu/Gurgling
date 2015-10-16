package gurgling

import (
    "strings"
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
