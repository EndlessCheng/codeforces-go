package copypasta

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_pSegmentTree(t *testing.T) {
	assert := assert.New(t)
	type pair struct{ v, i int }

	a := []int{100, 20, 50, 50, 30}
	n := len(a)

	st := make(pst, n+1)
	st.init(n)

	ps := make([]pair, n)
	for i, v := range a {
		ps[i] = pair{v, i}
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].v < ps[j].v })
	kthArr := make([]int, n)
	for i, p := range ps {
		kthArr[p.i] = i + 1
	}
	t.Log("kthArr", kthArr)
	t.Log("按照原数组元素的顺序插入该元素的名次", kthArr)
	for i, kth := range kthArr {
		st.update(i+1, i, kth, 1)
		t.Logf("插入元素 %d 的名次 %d", a[i], kth)
	}
	t.Log(ps[st.queryKth(1, 5, 1)-1].v)
	t.Log(ps[st.queryKth(1, 5, 2)-1].v)
	t.Log(ps[st.queryKth(1, 5, 3)-1].v)
	t.Log(ps[st.queryKth(1, 5, 4)-1].v)
	t.Log(ps[st.queryKth(1, 5, 5)-1].v)
	t.Log(ps[st.queryKth(3, 5, 1)-1].v)
	t.Log(ps[st.queryKth(3, 5, 2)-1].v)
	t.Log(ps[st.queryKth(3, 5, 3)-1].v)

	n = 4
	st = make(pst, n+1)
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
