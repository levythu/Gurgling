package matcher

import (
    "testing"
    . "github.com/levythu/gurgling/definition"
)

func TestRegEXPRules(t *testing.T) {
    var regmatcher=ARegexpMatcher()
    assert(t, regmatcher.AddRule("/huahua/post/", "POST", 99, false)==true)
    assert(t, regmatcher.AddRule("/huahua", "", 1, false)==true)
    assert(t, regmatcher.AddRule("qq", "",2, false)==false)
    assert(t, regmatcher.AddRule("/huahua/asd/", "GET", 3, false)==true)

    var baseurl, path string
    // TODO: add self-assert test
    baseurl=""
    path="/huahua/hahaha"
    var q=make(map[string]Tout)
    t.Log(baseurl, path)
    t.Log(regmatcher.Match(&path, &baseurl, q, "GET", nil))
    t.Log(baseurl, path)

    baseurl="/ex2"
    path="/huahua/post/123/456"
    t.Log(baseurl, path)
    t.Log(regmatcher.Match(&path, &baseurl, q, "HEAD", nil))
    t.Log(baseurl, path)

    baseurl="/ex2"
    path="/huahua/post/123/456"
    t.Log(baseurl, path)
    t.Log(regmatcher.Match(&path, &baseurl, q, "POST", nil))
    t.Log(baseurl, path)

    baseurl="/ex2/qq"
    path="/huahua/asd/123/456"
    t.Log(baseurl, path)
    t.Log(regmatcher.Match(&path, &baseurl, q, "GET", nil))
    t.Log(baseurl, path)

    baseurl=""
    path="/qq"
    t.Log(baseurl, path)
    t.Log(regmatcher.Match(&path, &baseurl, q, "GET", nil))
    t.Log(baseurl, path)
}
