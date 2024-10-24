package matching

import (
	"sandbox/strings/normalise"
	"strings"
)

func CaseInsensitiveMatch(s string, t string) bool {
	return strings.EqualFold(s, t)
}

func CaseInsensitiveSubstringMatch(s string, t string) bool {
	return strings.Contains(normalise.Normalise(s), normalise.Normalise(t))
}
