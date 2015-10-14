package matcher

import (
    "testing"
)

func assert(t *testing.T, cond bool) {
    if !cond {
        t.Fail()
    }
}

func TestRules(t *testing.T) {
    var bfmatcher=NewBFMatcher()
    assert(t, bfmatcher.AddRule("/huahua/post/", "POST", 99)==true)
    assert(t, bfmatcher.AddRule("/huahua", "", 1)==true)
    assert(t, bfmatcher.AddRule("qq", "",2)==false)
    assert(t, bfmatcher.AddRule("/huahua/asd/", "GET", 3)==true)

    var baseurl, path string
    // TODO: add self-assert test
    baseurl=""
    path="/huahua/hahaha"
    t.Log(baseurl, path)
    t.Log(bfmatcher.Match(&path, &baseurl, nil, "GET", nil))
    t.Log(baseurl, path)

    baseurl="/ex2"
    path="/huahua/post/123/456"
    t.Log(baseurl, path)
    t.Log(bfmatcher.Match(&path, &baseurl, nil, "HEAD", nil))
    t.Log(baseurl, path)

    baseurl="/ex2"
    path="/huahua/post/123/456"
    t.Log(baseurl, path)
    t.Log(bfmatcher.Match(&path, &baseurl, nil, "POST", nil))
    t.Log(baseurl, path)

    baseurl="/ex2/qq"
    path="/huahua/asd/123/456"
    t.Log(baseurl, path)
    t.Log(bfmatcher.Match(&path, &baseurl, nil, "GET", nil))
    t.Log(baseurl, path)

    baseurl=""
    path="/qq"
    t.Log(baseurl, path)
    t.Log(bfmatcher.Match(&path, &baseurl, nil, "GET", nil))
    t.Log(baseurl, path)
}
