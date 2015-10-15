package staticfs

import (
    . "github.com/levythu/gurgling"
)

// StaticFS provide packages for static file io and basic caching.

type FsMidware struct {
    // implements IMidware
    basePath string
    cacheControl CacheStrategy
}

// ignoring details about the class itself
// making a 120-seconds caching fs-midware
func AStaticfs(basePath string) IMidware {
    return &FsMidware {
        basePath: basePath,
        cacheControl: CacheStrategy(120),
    }
}

func (this *FsMidware)Handler(req Request, res Response) (bool, Request, Response) {

}
