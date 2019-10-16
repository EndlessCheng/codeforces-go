package copypasta

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_segmentTree(t *testing.T) {
	assert := assert.New(t)

	n := 10
	st := make(segmentTree, 4*n)
	st.init([]int64{2, 4, 4, 2, 110, 30})
	_, maxPos := st.query2(1, 1)
	assert.EqualValues(0, maxPos)
	_, maxPos = st.query2(1, 4)
	assert.EqualValues(2, maxPos)
	_, maxPos = st.query2(1, n)
	assert.EqualValues(4, maxPos)
	st.update(3, 2)
	_, maxPos = st.query2(1, 4)
	assert.EqualValues(1, maxPos)
}

func Test_lazySegmentTree(t *testing.T) {
	assert := assert.New(t)

	n := 10
	st := make(lazySegmentTree, 4*n)
	st.init(make([]int64, n+1))
	st.update(1, 10, 100)
	st.update(2, 4, 100)
	assert.EqualValues(100, st.query(1, 1))
	assert.EqualValues(200, st.query(3, 3))
	assert.EqualValues(400, st.query(3, 4))
	assert.EqualValues(1300, st.query(1, 10))
}
