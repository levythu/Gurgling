package gurgling

// Simple wrapper of http

import (
    "net/http"
)

func LaunchServer(addr string, r Router) error {
    return http.ListenAndServe(addr, r)
}
