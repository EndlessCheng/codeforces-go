package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type data87 struct{ val, del int }
type mq87 struct {
	data []data87
	size int
	less func(a, b data87) bool
}

func (q *mq87) push(v int) {
	q.size++
	d := data87{v, 1}
	for len(q.data) > 0 && q.less(d, q.data[len(q.data)-1]) {
		d.del += q.data[len(q.data)-1].del
		q.data = q.data[:len(q.data)-1]
	}
	q.data = append(q.data, d)
}

func (q *mq87) pop() {
	q.size--
	if q.data[0].del > 1 {
		q.data[0].del--
	} else {
		q.data = q.data[1:]
	}
}

func (q mq87) top() int {
	return q.data[0].val
}

func CF487B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, maxD, minL, v int
	Fscan(in, &n, &maxD, &minL)
	small := mq87{less: func(a, b data87) bool { return a.val <= b.val }}
	big := mq87{less: func(a, b data87) bool { return a.val >= b.val }}
	dpMQ := mq87{less: func(a, b data87) bool { return a.val <= b.val }}
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = 1e9
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		small.push(v)
		big.push(v)
		if i < minL {
			continue
		}
		dpMQ.push(dp[i-minL])
		for dpMQ.size > 0 && big.top()-small.top() > maxD {
			small.pop()
			big.pop()
			dpMQ.pop()
		}
		if dpMQ.size > 0 {
			dp[i] = dpMQ.top() + 1
		}
	}
	if dp[n] >= 1e9 { // 反思：写成 == WA 了一次
		dp[n] = -1
	}
	Fprint(out, dp[n])
}

//func main() { CF487B(os.Stdin, os.Stdout) }
