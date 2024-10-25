package multiset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt(tt *testing.T) {
	assert := assert.New(tt)

	t := newTreap()
	t.put(1, 2)
	t.put(3, 2)
	t.put(10, 1)

	assert.Equal(0, t.preSum(0))
	assert.Equal(1, t.preSum(1))
	assert.Equal(2, t.preSum(2))
	assert.Equal(5, t.preSum(3))
	assert.Equal(8, t.preSum(4))
	assert.Equal(18, t.preSum(5))

	assert.Equal(2, t.sumLessEqual(2))
	assert.Equal(8, t.sumLessEqual(3))
	assert.Equal(18, t.sumLessEqual(10))
	assert.Equal(18, t.sumGreaterEqual(1))
	assert.Equal(16, t.sumGreaterEqual(2))
	assert.Equal(10, t.sumGreaterEqual(4))
}
