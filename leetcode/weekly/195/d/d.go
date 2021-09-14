package main

func f(p []int) int { return p[1] - p[0] }

type mqData struct {
	p   []int
	del int
}
type mq struct {
	data []mqData
	size int
}

func (q mq) less(a, b mqData) bool { return f(a.p) >= f(b.p) }
func (q *mq) push(p []int) {
	q.size++
	d := mqData{p, 1}
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
func (q mq) top() []int { return q.data[0].p }

func findMaxValueOfEquation(a [][]int, k int) int {
	ans := int(-1e9)
	q := mq{}
	for _, p := range a {
		for q.size > 0 && p[0]-q.top()[0] > k {
			q.pop()
		}
		if q.size > 0 {
			ans = max(ans, p[0]+p[1]+f(q.top()))
		}
		q.push(p)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
