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
	s, _ := t.minPrefix("a")
	assert.Equal("ab", s)
	s, _ = t.minPrefix("ab")
	assert.Equal("ab", s)
	s, _ = t.minPrefix("abc")
	assert.Equal("abc", s)
	s, _ = t.minPrefix("abcd")
	assert.Equal("", s)
	s, _ = t.minPrefix("z")
	assert.Equal("zz", s)
	s, _ = t.minPrefix("zzz")
	assert.Equal("", s)
}
