package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveLagoonPrefix(t *testing.T) {

	s := removeLagoonPrefix(LAGOON_PREFIX + "aaa")
	assert.Equal(t, s, "aaa")

	s = removeLagoonPrefix(LAGOON_PREFIX)
	assert.Equal(t, s, "")

	s = removeLagoonPrefix("aaa")
	assert.Equal(t, s, "aaa")

}
