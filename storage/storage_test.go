package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveEkaraPrefix(t *testing.T) {

	s := removeEkarePrefix(EKARA_PREFIX + "aaa")
	assert.Equal(t, s, "aaa")

	s = removeEkaraPrefix(EKARA_PREFIX)
	assert.Equal(t, s, "")

	s = removeEkaraPrefix("aaa")
	assert.Equal(t, s, "aaa")

}
