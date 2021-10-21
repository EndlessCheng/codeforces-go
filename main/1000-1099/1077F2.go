package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type mqData77 struct {
	v   int64
	del int
}
type mq77 struct {
	data []mqData77
	size int
}

func (q *mq77) push(v int64) {
	q.size++
	d := mqData77{v, 1}
	for len(q.data) > 0 && q.data[len(q.data)-1].v <= v {
		d.del += q.data[len(q.data)-1].del
		q.data = q.data[:len(q.data)-1]
	}
	q.data = append(q.data, d)
}
func (q *mq77) pop() {
	q.size--
	if q.data[0].del > 1 {
		q.data[0].del--
	} else {
		q.data = q.data[1:]
	}
}
func (q mq77) top() (v int64) { return q.data[0].v }

func CF1077F2(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, lim, choose int
	Fscan(in, &n, &lim, &choose)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	dp := make([][]int64, choose)
	for i := range dp {
		dp[i] = make([]int64, n)
	}
	copy(dp[0], a[:lim])
	for i := 1; i < choose; i++ {
		q := mq77{}
		q.push(dp[i-1][i-1])
		up := (i + 1) * lim
		if up > n {
			up = n
		}
		for j := i; j < up; j++ {
			if q.size > lim {
				q.pop()
			}
			dp[i][j] = q.top() + a[j]
			q.push(dp[i-1][j])
		}
	}

	ans := int64(0)
	for _, v := range dp[choose-1][n-lim:] {
		if v > ans {
			ans = v
		}
	}
	if ans == 0 {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { CF1077F2(os.Stdin, os.Stdout) }
