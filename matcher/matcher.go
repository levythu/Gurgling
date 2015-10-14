package matcher

// The package is a built-in accessory to match regs for router.

import (
    . "github.com/levythu/gurgling/definition"
)

// Attentez: the AddRule does not guarantee thread-safety, so avoid to exec Match() when adding rules.
type Matcher interface {
    // Check whether valid, and do any possible correction.
    CheckRuleValidity(rule *string) bool

    // returning value indicates whether successfully
    // The default method is "", which means WILDCARD.
    // isStrict==true: must strictly match the pattern (whole word), "/" should match "".
    AddRule(rulePattern string, methodPattern string/*=""*/, storage Tout, isStrict bool) bool

    // path and baseUrl is used for modification and matching. prevPoint is
    // used for acceleration and states storing.
    // returns the storage and new statepoint, modifying the path and baseUrl as matched.
    // reqF is req.F(), store data if needed.
    // If match failure, storage should be nil and statepoint is not specified, path,
    // base URL and reqF should remain the unchanged.
    Match(path *string, baseUrl *string, reqF map[string]Tout, method string, prevPoint Tout) (/*storage*/Tout, /*statepoint*/Tout)
}
