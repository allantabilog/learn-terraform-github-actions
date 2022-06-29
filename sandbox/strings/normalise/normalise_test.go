package normalise

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalise(t *testing.T) {
	assert.Equal(t, Normalise("Transfer"), "transfer")
	assert.Equal(t, Normalise("Non-Transfer"), "nontransfer")
	assert.Equal(t, Normalise("TRANSFER"), "transfer")
	assert.Equal(t, Normalise("NON_TRANSFER"), "nontransfer")
}
