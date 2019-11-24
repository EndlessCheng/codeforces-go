package copypasta

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_trie(t_ *testing.T) {
	assert := assert.New(t_)

	t := newTrie()
	t.put("ab", 1)
	t.put("ab", 1)
	t.put("ab", 1)
	t.put("ab", 1)
	t.put("abc", 1)
	t.put("abc", 1)
	t.put("zz", 1)
	assert.Len(t.nodes, 6)
	assert.EqualValues(0, t.nodes[0].dupCnt)
	assert.EqualValues(0, t.nodes[1].dupCnt)
	assert.EqualValues(4, t.nodes[2].dupCnt)
	assert.EqualValues(2, t.nodes[3].dupCnt)
	assert.EqualValues(0, t.nodes[4].dupCnt)
	assert.EqualValues(1, t.nodes[5].dupCnt)
	_, found := t.get("a")
	assert.False(found)
	_, found = t.get("ab")
	assert.True(found)
	_, found = t.get("abc")
	assert.True(found)
	_, found = t.get("abcd")
	assert.False(found)
	s, _ := t.prefix("a")
	assert.Equal("ab", s)
	s, _ = t.prefix("ab")
	assert.Equal("ab", s)
	s, _ = t.prefix("abc")
	assert.Equal("abc", s)
	s, _ = t.prefix("abcd")
	assert.Equal("", s)
	s, _ = t.prefix("z")
	assert.Equal("zz", s)
	s, _ = t.prefix("zzz")
	assert.Equal("", s)
}
