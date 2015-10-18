package matcher

// A matcher that supports regexp

import (
    "strings"
    "regexp"
    . "github.com/levythu/gurgling/definition"
)

type regexpRecord struct {
    rulePattern *regexp.Regexp
    methodPattern string
    storage Tout
}

type RegexpMatcher struct {
    rules []*regexpRecord
}

func ARegexpMatcher() Matcher {
    return &RegexpMatcher {
        rules: []*regexpRecord{},
    }
}

func (this *RegexpMatcher)CheckRuleValidity(rule *string) bool {
    if *rule=="" {
        return true
    }
    if strings.HasPrefix(*rule, "^") {
        // no force-head, the rule will deal with it
        return false
    }
    if strings.HasSuffix(*rule, "$") {
        // no force-end, the rule will deal with it
        return false
    }
    if !strings.HasPrefix(*rule, "/") {
        return false
    }
    var ends=len(*rule)
    for ends>0 && (*rule)[ends-1]=='/' {
        ends--
    }
    *rule=(*rule)[:ends]
    return true
}
func (this *RegexpMatcher)AddRule(rulePattern string, methodPattern string, storage Tout, isStrict bool) bool {
    if !this.CheckRuleValidity(&rulePattern) {
        return false
    }
    rulePattern="^"+rulePattern // Match prefix
    if isStrict {
        rulePattern+="$"    // Match whole word
    }
    if rulePattern=="^$" {
        rulePattern="^/?$"  // exact "" should also match "/"
    }

    var targetEXP, err=regexp.Compile(rulePattern)
    if err!=nil {
        return false
    }

    this.rules=append(this.rules, &regexpRecord{
        rulePattern: targetEXP,
        methodPattern: methodPattern,
        storage: storage,
    })
    return true
}
const REGEXP_MATCHER_RESULT="regexp-matcher-result"
func (this *RegexpMatcher)Match(path *string, baseUrl *string, reqF map[string]Tout, method string/*=""*/, prevPoint Tout) (Tout, Tout) {
    var startpoint int
    if prevPoint==nil {
        startpoint=0
    } else {
        // forcing type assertion could cause panic, and is expected.
        startpoint=prevPoint.(int)
    }
    var length=len(this.rules)
    for startpoint<length {
        var res=this.rules[startpoint].rulePattern.FindStringSubmatch(*path)
        if len(res)>0 && (this.rules[startpoint].methodPattern=="" ||  this.rules[startpoint].methodPattern==method) {
            // Matched!
            *path=strings.TrimPrefix(*path, res[0])
            *baseUrl=*baseUrl+res[0]
            reqF[REGEXP_MATCHER_RESULT]=res
            return this.rules[startpoint].storage, startpoint+1
        }
        startpoint++
    }
    return nil, startpoint
}
