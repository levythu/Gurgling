package matcher

// Match by prefix sequently.

import (
    "strings"
    . "github.com/levythu/gurgling/definition"
)

type ruleRecord struct {
    rulePattern string
    methodPattern string
    storage Tout
}

type BruteforceMatcher struct {
    rules []*ruleRecord
}

func NewBFMatcher() Matcher {
    return &BruteforceMatcher {
        rules: []*ruleRecord{},
    }
}

func (this *BruteforceMatcher)CheckRuleValidity(rule *string) bool {
    if *rule=="" {
        return true
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
func (this *BruteforceMatcher)AddRule(rulePattern string, methodPattern string, storage Tout) bool {
    if !this.CheckRuleValidity(&rulePattern) {
        return false
    }
    this.rules=append(this.rules, &ruleRecord{
        rulePattern: rulePattern,
        methodPattern: methodPattern,
        storage: storage,
    })
    return true
}
func (this *BruteforceMatcher)Match(path *string, baseUrl *string, reqF map[string]Tout, method string/*=""*/, prevPoint Tout) (Tout, Tout) {
    var startpoint int
    if prevPoint==nil {
        startpoint=0
    } else {
        // forcing type assertion could cause panic, and is expected.
        startpoint=prevPoint.(int)
    }
    var length=len(this.rules)
    for startpoint<length {
        if strings.HasPrefix(*path, this.rules[startpoint].rulePattern) && (this.rules[startpoint].methodPattern=="" ||  this.rules[startpoint].methodPattern==method) {
            // Matched!
            *path=strings.TrimPrefix(*path, this.rules[startpoint].rulePattern)
            *baseUrl=*baseUrl+this.rules[startpoint].rulePattern
            return this.rules[startpoint].storage, startpoint+1
        }
        startpoint++
    }
    return nil, startpoint
}
