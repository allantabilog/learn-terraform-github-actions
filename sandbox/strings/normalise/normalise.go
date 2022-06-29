package normalise

import "strings"

func Normalise(s string) string {
	// replace all occurrences of - and _ with nothing
	// and then lowercase them
	transformed := s
	transformed = strings.ReplaceAll(transformed, "_", "")
	transformed = strings.ReplaceAll(transformed, "-", "")
	transformed = strings.ReplaceAll(transformed, " ", "")
	transformed = strings.ToLower(transformed)

	return transformed
}
