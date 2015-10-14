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
