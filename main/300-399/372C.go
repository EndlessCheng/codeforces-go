package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type data72 struct {
	v   int64
	del int
}
type mq72 struct {
	data []data72
	size int
}

func (q mq72) less(a, b data72) bool { return a.v >= b.v }
func (q *mq72) push(v int64) {
	q.size++
	d := data72{v, 1}
	for len(q.data) > 0 && q.less(d, q.data[len(q.data)-1]) {
		d.del += q.data[len(q.data)-1].del
		q.data = q.data[:len(q.data)-1]
	}
	q.data = append(q.data, d)
}
func (q *mq72) pop() {
	q.size--
	if q.data[0].del > 1 {
		q.data[0].del--
	} else {
		q.data = q.data[1:]
	}
}
func (q mq72) top() (v int64) { return q.data[0].v }

func CF372C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var n, m, d, pos, hp, t, preT int
	Fscan(in, &n, &m, &d)
	pre := make([]int64, n)
	dp := make([]int64, n)
	for ; m > 0; m-- {
		Fscan(in, &pos, &hp, &t)
		pos--
		q := mq72{}
		half := n
		if int64(d)*int64(t-preT) < int64(n) {
			half = d * (t - preT)
		}
		for _, v := range pre[:half] {
			q.push(v)
		}
		for i := range dp {
			if i+half < n {
				q.push(pre[i+half])
			}
			if i > half {
				q.pop()
			}
			dp[i] = q.top() + int64(hp-abs(pos-i))
		}
		pre, dp = dp, pre
		preT = t
	}
	ans := pre[0]
	for _, v := range pre[1:] {
		if v > ans {
			ans = v
		}
	}
	Fprint(out, ans)
}

//func main() { CF372C(os.Stdin, os.Stdout) }
