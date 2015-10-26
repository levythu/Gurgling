package gurgling

import (
    . "github.com/levythu/gurgling/definition"
)

// Public global configurations for gurgling.

type Gurgling_Config struct {
    H500 func(req Request, res Response, thePanic interface{})
    H404 Terminal

    F map[string]Tout   // for further extension
}

var CGurgling=initPreCGurgling(CGurgling_Predefined_forDebug)

func SetPreCGurgling(obj *Gurgling_Config) {
    // exec a deep copy
    CGurgling.H500=obj.H500
    CGurgling.H404=obj.H404

    CGurgling.F=make(map[string]Tout)
    for k, v:=range obj.F {
        CGurgling.F[k]=v
    }
}
func SetGEnv(theme string) {
    if v, ok:=_CGurgling_Predefined_map[theme];ok {
        SetPreCGurgling(v)
    }
}

func initPreCGurgling(obj *Gurgling_Config) Gurgling_Config {
    var ret Gurgling_Config

    // exec a deep copy
    ret.H500=obj.H500
    ret.H404=obj.H404

    ret.F=make(map[string]Tout)
    for k, v:=range obj.F {
        ret.F[k]=v
    }

    return ret
}

var (
    CGurgling_Predefined_forDebug=&Gurgling_Config{
        F: map[string]Tout{},
        H500: nil,
        H404: nil,
    }
    CGurgling_Predefined_forRelease=&Gurgling_Config{
        F: map[string]Tout{},
        H500: DefaultCacher,
        H404: Default404Cacher,
    }

    _CGurgling_Predefined_map=map[string]*Gurgling_Config {
        "debug":    CGurgling_Predefined_forDebug,
        "release":  CGurgling_Predefined_forRelease,
    }
)
