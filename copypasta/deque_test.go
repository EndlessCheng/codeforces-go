package copypasta

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deque(t *testing.T) {
	assert := assert.New(t)

	q := &deque{}
	q.pushL(1)
	assert.Equal(1, q.popR())
	assert.Equal(0, q.len())

	q.pushL(1)
	assert.Equal(1, q.popL())
	assert.Equal(0, q.len())

	q.pushR(1)
	assert.Equal(1, q.popL())
	assert.Equal(0, q.len())

	q.pushR(1)
	assert.Equal(1, q.popR())
	assert.Equal(0, q.len())
}
