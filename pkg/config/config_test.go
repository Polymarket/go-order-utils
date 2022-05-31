package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContracts(t *testing.T) {
	c2, err := GetContracts(80001)
	assert.NotNil(t, c2)
	assert.Nil(t, err)
	assert.Equal(t, c2, MUMBAI_CONTRACTS)

	c3, err := GetContracts(137)
	assert.NotNil(t, c3)
	assert.Nil(t, err)
	assert.Equal(t, c3, MATIC_CONTRACTS)

	cn, err := GetContracts(100000)
	assert.Nil(t, cn)
	assert.NotNil(t, err)
}
