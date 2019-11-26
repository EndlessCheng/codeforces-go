package copypasta

import (
	"github.com/stretchr/testify/assert"
	"sort"
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

func Test_pSegmentTree(t *testing.T) {
	assert := assert.New(t)

	arr := []int{100, 20, 50, 23}
	n := len(arr)

	type pair struct{ val, i int }
	ps := make([]pair, n)
	for i, v := range arr {
		ps[i] = pair{v, i}
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].val < ps[j].val })
	kthArr := make([]int, n)
	for i, p := range ps {
		kthArr[p.i] = i + 1
	}

	st := newPST(n, n)
	st.init(n)
	for i, kth := range kthArr {
		st.update(i+1, i, kth, 1)
		t.Log("insert", kth)
	}
	t.Log(ps[st.queryKth(1, 4, 1)-1].val)
	t.Log(ps[st.queryKth(1, 4, 2)-1].val)
	t.Log(ps[st.queryKth(1, 4, 3)-1].val)
	t.Log(ps[st.queryKth(1, 4, 4)-1].val)

	n = 4
	st = newPST(n, 2*n)
	st.init(n)
	st.update(1, 0, 2, 10)
	st.update(1, 1, 1, -5)
	st.update(2, 1, 1, -100)
	assert.Equal(5, st.query(1, 1, 2))
	assert.Equal(10, st.query(1, 2, 2))
	assert.Equal(-95, st.query(2, 1, 2))
	assert.Equal(-105, st.query(2, 1, 1))
	assert.Equal(10, st.query(2, 2, 2))
}
