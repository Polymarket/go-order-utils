package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomSalt(t *testing.T) {
	r1 := GenerateRandomSalt()
	r2 := GenerateRandomSalt()
	r3 := GenerateRandomSalt()

	assert.Positive(t, r1)
	assert.Positive(t, r2)
	assert.Positive(t, r3)

	assert.NotEqual(t, r1, r2)
	assert.NotEqual(t, r2, r3)
}
