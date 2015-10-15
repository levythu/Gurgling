package urlnormalizer

import (
    "testing"
)

func assert(t *testing.T, cond bool) {
    if !cond {
        t.Fail()
    }
}

func TestRules(t *testing.T) {
    t.Log(sanitize(""))
    t.Log(sanitize("/"))
    t.Log(sanitize("////"))
    t.Log(sanitize("/./"))
    t.Log(sanitize("/../asd"))
    t.Log(sanitize("/asd/../"))
    t.Log(sanitize("/asd/.."))
    t.Log(sanitize("/asd/../../../.."))
    t.Log(sanitize("/asd/./123"))
    t.Log(sanitize("/asd///../123/.//999//"))
}
