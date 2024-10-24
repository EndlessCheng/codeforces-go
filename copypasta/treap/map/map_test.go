package multiset

import (
	"cmp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt(tt *testing.T) {
	assert := assert.New(tt)

	t := newTreap[int, string]()
	assert.Equal(0, t.size())
	assert.True(t.empty())
	assert.Nil(t.min())
	assert.Nil(t.max())
	assert.Nil(t.kth(0))
	assert.Nil(t.kth(1))
	assert.Nil(t.kth(-1))
	assert.Equal(0, t.lowerBoundIndex(0))
	assert.Equal(0, t.upperBoundIndex(0))
	assert.Nil(t.prev(0))
	assert.Nil(t.next(0))
	assert.Nil(t.find(0))

	t.put(1, "a")
	assert.Equal(1, t.size())
	assert.False(t.empty())
	assert.Equal(1, t.min().key)
	assert.Equal(1, t.max().key)
	assert.Equal(1, t.find(1).key)
	assert.Equal("a", t.find(1).value)

	t.put(10, "c")
	t.put(2, "b")
	t.put(1, "aa")
	assert.Equal("aa", t.find(1).value)
	t.put(2, "bb")
	assert.Equal("bb", t.find(2).value)
	assert.Equal(1, t.min().key)
	assert.Equal(10, t.max().key)
	assert.Nil(t.find(0))
	assert.NotNil(t.find(1))
	assert.NotNil(t.find(2))
	assert.Nil(t.find(3))

	assert.Equal(1, t.kth(0).key)
	assert.Equal(2, t.kth(1).key)
	assert.Equal(10, t.kth(2).key)
	assert.Nil(t.kth(-1))
	assert.Nil(t.kth(3))

	assert.Equal(0, t.lowerBoundIndex(0))
	assert.Equal(0, t.lowerBoundIndex(1))
	assert.Equal(1, t.lowerBoundIndex(2))
	assert.Equal(2, t.lowerBoundIndex(3))
	assert.Equal(2, t.lowerBoundIndex(10))
	assert.Equal(3, t.lowerBoundIndex(11))

	assert.Equal(0, t.upperBoundIndex(0))
	assert.Equal(1, t.upperBoundIndex(1))
	assert.Equal(2, t.upperBoundIndex(2))
	assert.Equal(2, t.upperBoundIndex(9))
	assert.Equal(3, t.upperBoundIndex(10))

	assert.Nil(t.prev(1))
	assert.Equal(1, t.prev(2).key)
	assert.Equal(2, t.prev(9).key)
	assert.Equal(2, t.prev(10).key)
	assert.Equal(10, t.prev(11).key)

	assert.Equal(1, t.next(0).key)
	assert.Equal(2, t.next(1).key)
	assert.Equal(10, t.next(2).key)
	assert.Equal(10, t.next(9).key)
	assert.Nil(t.next(10))

	t.delete(1) // 只删除一个 1
	assert.Equal(2, t.size())

	t.delete(1) // 无效
	assert.Equal(2, t.size())

	t.delete(2)
	assert.Equal(1, t.size())

	t.delete(2) // 无效
	assert.Equal(1, t.size())

	t.delete(10)
	assert.Equal(0, t.size())
	assert.True(t.empty())
	assert.Nil(t.min())
	assert.Nil(t.max())
	assert.Nil(t.kth(0))
	assert.Nil(t.kth(1))
	assert.Nil(t.kth(-1))
	assert.Equal(0, t.lowerBoundIndex(0))
	assert.Equal(0, t.upperBoundIndex(0))
	assert.Nil(t.prev(0))
	assert.Nil(t.next(0))
}

func TestPair(tt *testing.T) {
	assert := assert.New(tt)

	type pair struct{ x, y int }
	t := newTreapWith[pair, string](func(a, b pair) int { return cmp.Or(a.x-b.x, a.y-b.y) })
	assert.Equal(0, t.size())
	assert.True(t.empty())
	assert.Nil(t.min())
	assert.Nil(t.max())

	t.put(pair{1, 2}, "b")
	assert.Equal(1, t.size())
	assert.False(t.empty())
	assert.Equal(pair{1, 2}, t.min().key)
	assert.Equal(pair{1, 2}, t.max().key)
	assert.Equal(pair{1, 2}, t.find(pair{1, 2}).key)
	assert.Equal("b", t.find(pair{1, 2}).value)

	t.put(pair{10, 20}, "c")
	t.put(pair{1, 1}, "a")
	t.put(pair{1, 2}, "bb")
	assert.Equal("bb", t.find(pair{1, 2}).value)
	t.put(pair{1, 1}, "aa")
	assert.Equal("aa", t.find(pair{1, 1}).value)
	assert.Equal(pair{1, 1}, t.min().key)
	assert.Equal(pair{10, 20}, t.max().key)
	assert.NotNil(t.find(pair{1, 1}))
	assert.Nil(t.find(pair{1, 0}))
	assert.Nil(t.find(pair{1, 3}))

	assert.Equal(pair{1, 1}, t.kth(0).key)
	assert.Equal(pair{1, 2}, t.kth(1).key)
	assert.Equal(pair{10, 20}, t.kth(2).key)
	assert.Nil(t.kth(-1))
	assert.Nil(t.kth(3))

	assert.Equal(0, t.lowerBoundIndex(pair{}))
	assert.Equal(0, t.lowerBoundIndex(pair{1, 1}))
	assert.Equal(1, t.lowerBoundIndex(pair{1, 2}))
	assert.Equal(2, t.lowerBoundIndex(pair{1, 3}))
	assert.Equal(2, t.lowerBoundIndex(pair{10, 20}))
	assert.Equal(3, t.lowerBoundIndex(pair{10, 21}))

	assert.Equal(pair{1, 1}, t.prev(pair{1, 2}).key)
	assert.Equal(pair{1, 2}, t.next(pair{1, 1}).key)

	t.delete(pair{10, 20})
	t.delete(pair{1, 2})
	t.delete(pair{1, 2})
	t.delete(pair{1, 1})
	t.delete(pair{1, 1})
	assert.Equal(0, t.size())
	assert.True(t.empty())
	assert.Nil(t.min())
	assert.Nil(t.max())
}
