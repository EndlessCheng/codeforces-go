package copypasta

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

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

	st := newPST(n, n, n)
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
	st = newPST(n, n, 2*n)
	st.init(n)
	st.update(1, 0, 2, 10)
	st.update(1, 1, 1, -5)
	st.update(2, 1, 1, -100)
	assert.EqualValues(5, st.query(1, 1, 2))
	assert.EqualValues(10, st.query(1, 2, 2))
	assert.EqualValues(-95, st.query(2, 1, 2))
	assert.EqualValues(-105, st.query(2, 1, 1))
	assert.EqualValues(10, st.query(2, 2, 2))
}
