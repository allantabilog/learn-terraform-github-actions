package matching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCaseInsensitiveMatch(t *testing.T) {
	assert.True(t, CaseInsensitiveMatch("Hello", "HELLO"))
}

func TestCaseInsensitiveSubstringMatch(t *testing.T) {
	assert.True(t, CaseInsensitiveSubstringMatch("Hello", "ELLO"))
	assert.True(t, CaseInsensitiveSubstringMatch("Dear Prudence", "DEAR-PRUDENCE"))
	assert.True(t, CaseInsensitiveSubstringMatch("Jurisprudence", "RIS-PRUDENCE"))
	assert.True(t, CaseInsensitiveSubstringMatch("Diamonds and pearls", "DIAMONDS-AND-PEARLS"))
}
