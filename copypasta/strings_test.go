package copypasta

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_trie(t_ *testing.T) {
	assert := assert.New(t_)

	t := newTrie()
	t.insert("ab", 1)
	t.insert("ab", 1)
	t.insert("ab", 1)
	t.insert("ab", 1)
	t.insert("abc", 1)
	t.insert("abc", 1)
	t.insert("z", 1)
	assert.Len(t.nodes, 5)
	assert.EqualValues(0, t.nodes[0].dupCnt)
	assert.EqualValues(0, t.nodes[1].dupCnt)
	assert.EqualValues(4, t.nodes[2].dupCnt)
	assert.EqualValues(2, t.nodes[3].dupCnt)
	assert.EqualValues(1, t.nodes[4].dupCnt)
}
