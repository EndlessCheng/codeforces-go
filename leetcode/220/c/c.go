package main

// github.com/EndlessCheng/codeforces-go

type mqData struct{ val, del int }
type mq struct {
	data []mqData
	size int
}

func (q mq) less(a, b mqData) bool { return a.val >= b.val }
func (q *mq) push(v int) {
	q.size++
	d := mqData{v, 1}
	for len(q.data) > 0 && q.less(d, q.data[len(q.data)-1]) {
		d.del += q.data[len(q.data)-1].del
		q.data = q.data[:len(q.data)-1]
	}
	q.data = append(q.data, d)
}
func (q *mq) pop() {
	q.size--
	if q.data[0].del > 1 {
		q.data[0].del--
	} else {
		q.data = q.data[1:]
	}
}
func (q mq) top() int { return q.data[0].val }

func maxResult(a []int, k int) int {
	q := mq{}
	ans := a[0]
	q.push(ans)
	for i := 1; i < len(a); i++ {
		if q.size > k {
			q.pop()
		}
		ans = q.top() + a[i]
		q.push(ans)
	}
	return ans
}
