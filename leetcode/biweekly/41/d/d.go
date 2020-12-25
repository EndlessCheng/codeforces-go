package main

// github.com/EndlessCheng/codeforces-go
type mqData struct {
	v   int
	del int
}
type mq struct {
	data []mqData
	size int
}

func (q mq) less(a, b mqData) bool { return a.v <= b.v }
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
func (q mq) top() (v int) { return q.data[0].v }

func boxDelivering(a [][]int, _, maxBoxes, maxWeight int) (ans int) {
	n := len(a)
	sumDiff := make([]int, n+1)
	sumW := make([]int, n+1)
	for i, b := range a {
		sumDiff[i+1] = sumDiff[i]
		if i < n-1 && b[0] != a[i+1][0] {
			sumDiff[i+1]++
		}
		sumW[i+1] = sumW[i] + b[1]
	}

	q := mq{}
	q.push(0)
	for i := range a {
		for q.size > maxBoxes || sumW[i+1]-sumW[i+1-q.size] > maxWeight {
			q.pop()
		}
		ans = q.top() + sumDiff[i] + 2
		q.push(ans - sumDiff[i+1])
	}
	return
}
