package copypasta

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.True(t, isRectangleAnyOrder(vec{1, 1}, vec{6, 1}, vec{1, 0}, vec{6, 0}))
}
