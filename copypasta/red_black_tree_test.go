package copypasta

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_rbt(t *testing.T) {
	assert := assert.New(t)

	rbt := NewRBTree()
	assert.True(rbt.Empty())
	assert.True(rbt.Begin().End())
	assert.True(rbt.RBegin().REnd())

	rbt.MultiPut(1)
	rbt.MultiPut(10)
	rbt.MultiPut(10)
	rbt.MultiPut(100)

	assert.False(rbt.Empty())
	assert.EqualValues(3, rbt.Size())

	assert.EqualValues([]keyType{1, 10, 10, 100}, rbt.MultiKeys())
	assert.EqualValues([]valueType{1, 2, 1}, rbt.Values())

	assert.EqualValues(1, rbt.Min().Key)
	assert.EqualValues(100, rbt.Max().Key)

	assert.Nil(rbt.Lookup(2))
	assert.NotNil(rbt.Lookup(10))
	assert.EqualValues(2, rbt.Lookup(10).Value)

	assert.EqualValues(1, rbt.Floor(9).Key)
	assert.EqualValues(10, rbt.Floor(10).Key)
	assert.EqualValues(10, rbt.Floor(11).Key)

	assert.EqualValues(10, rbt.Ceiling(9).Key)
	assert.EqualValues(10, rbt.Ceiling(10).Key)
	assert.EqualValues(100, rbt.Ceiling(11).Key)

	it := rbt.NewIterator(rbt.Lookup(10))
	assert.EqualValues(1, it.Prev().node.Key)
	it = rbt.NewIterator(rbt.Lookup(10))
	assert.EqualValues(100, it.Next().node.Key)

	rbt.MultiDelete(10)
	assert.NotNil(rbt.Lookup(10))
	assert.EqualValues(1, rbt.Lookup(10).Value)
	rbt.MultiDelete(10)
	assert.Nil(rbt.Lookup(10))
}
