package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://www.luogu.com.cn/problem/P5785

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, s, t, c int
	Fscan(in, &n, &s)
	sumT := make([]int, n+1)
	sumC := make([]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &t, &c)
		sumT[i+1] = sumT[i] + t
		sumC[i+1] = sumC[i] + c
	}
	dp := make([]int, n+1)
	q := []int{0}
	Y := func(i int) int { return dp[i] }
	X := func(i int) int { return sumC[i] }
	// 下凸包：i0-i1 的斜率 < i1-i2 的斜率
	// 上凸包：i0-i1 的斜率 > i1-i2 的斜率
	less := func(i0, i1, i2 int) bool {
		y0, y1, y2 := Y(i0), Y(i1), Y(i2)
		x0, x1, x2 := X(i0), X(i1), X(i2)
		return float64(y1-y0)/float64(x1-x0) < float64(y2-y1)/float64(x2-x1)
	}
	push := func(i int) {
		for len(q) > 1 && !less(q[len(q)-2], q[len(q)-1], i) {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	for i := 1; i <= n; i++ {
		k := s + sumT[i]
		// 在队列中二分第一个斜率大于（小于）k 的位置
		// 下凸包：>
		// 上凸包：<
		j := sort.Search(len(q)-1, func(j int) bool { return float64(Y(q[j+1])-Y(q[j]))/float64(X(q[j+1])-X(q[j])) > float64(k) })
		dp[i] = Y(q[j]) - k*X(q[j]) + sumT[i]*sumC[i] + s*sumC[n]
		push(i)
	}
	Fprint(out, dp[n])
}

func main() { run(os.Stdin, os.Stdout) }
