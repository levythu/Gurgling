package unittest

import (
    "github.com/levythu/gurgling"
    "testing"
)

func assert(t *testing.T, cond bool) {
    if !cond {
        t.Fail()
    }
}

func TestCheckMountpointValidity(t *testing.T) {
    var perftest=func(casestr string, expres string, expret bool) {
        assert(t, gurgling.CheckMountpointValidity(&casestr)==expret)
        if expret==true {
            assert(t, casestr==expres)
        }
    }
    perftest("/a", "/a", true)
    perftest("/a/", "/a", true)
    perftest("a/", "", false)
    perftest("//a/", "//a", true)
    perftest("/", "/", true)
    perftest("//////", "/", true)
    perftest("/ab/asd/12d///", "/ab/asd/12d", true)
}
