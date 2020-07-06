package libstring

import (
	"regexp"
	"strings"
)

func curlyToRegex(curlyPattern string) string {
	out := strings.Replace(curlyPattern, "{", "(?P<", -1)
	out = strings.Replace(out, "}", ">.+)", -1)
	out = strings.Replace(out, "*", ".*?", -1)
	return out
}

// Match evaluates if the pattern and the incoming string matches.
func Match(curlyPattern, in string) (bool, map[string]string, error) {
	regexExpression := curlyToRegex(curlyPattern)

	isMatched := false

	params := make(map[string]string)

	r, err := regexp.Compile(regexExpression)
	if err != nil {
		return isMatched, params, err
	}

	match := r.FindStringSubmatch(in)

	for i, name := range r.SubexpNames() {
		if i > 0 && i <= len(match) {
			params[name] = match[i]
		}
	}

	if len(params) > 0 {
		isMatched = true

	} else if len(match) == 1 && match[0] == in { // exact match use-case.
		isMatched = true
	}

	return isMatched, params, nil
}
