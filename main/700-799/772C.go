package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF772C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var exgcd func(int, int) (g, x, y int)
	exgcd = func(a, b int) (g, x, y int) {
		if b == 0 {
			return a, 1, 0
		}
		g, y, x = exgcd(b, a%b)
		y -= a / b * x
		return
	}

	var k, n, v int
	Fscan(in, &k, &n)
	ban := make([]bool, n)
	for ; k > 0; k-- {
		Fscan(in, &v)
		ban[v] = true
	}

	gs := make([][]int, n+1)
	for i, b := range ban {
		if !b {
			gs[gcd(i, n)] = append(gs[gcd(i, n)], i) // 缩点
		}
	}
	dp := make([]int, n+1)
	from := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] += len(gs[i])
		for j := i * 2; j <= n; j += i { // 视作一个 DAG，跑最长路
			if dp[i] > dp[j] {
				dp[j] = dp[i]
				from[j] = i
			}
		}
	}
	ans := make([]int, 0, dp[n])
	for i := n; i > 0; i = from[i] {
		for _, v := range gs[i] {
			ans = append(ans, v)
		}
	}

	sz := len(ans)
	Fprintln(out, sz)
	Fprint(out, ans[sz-1], " ")
	for i := sz - 1; i > 0; i-- {
		g, x, _ := exgcd(ans[i], n)
		Fprint(out, (int(int64(x)*int64(ans[i-1]/g)%int64(n))+n)%n, " ")
	}
}

//func main() { CF772C(os.Stdin, os.Stdout) }
