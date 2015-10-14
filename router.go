package gurgling

import (
    "net/http"
)

type Router *router
// Implementing http.Handler
type router struct {
    mountMap map[string]*router
    initMountPoint string
}

func (this *router)ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (this *router)mainHandler(w http.ResponseWriter, r *http.Request) {

}
