package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveEkaraPrefix(t *testing.T) {

	s := RemoveEkaraPrefix(EKARA_PREFIX + "aaa")
	assert.Equal(t, s, "aaa")

	s = RemoveEkaraPrefix(EKARA_PREFIX)
	assert.Equal(t, s, "")

	s = RemoveEkaraPrefix("aaa")
	assert.Equal(t, s, "aaa")

}
