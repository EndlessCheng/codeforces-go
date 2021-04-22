package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
	"sort"
)

// https://www.luogu.com.cn/problem/P1325

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const eps = 1e-8
	var n, d, x, y, ans int
	Fscan(in, &n, &d)
	a := make([]struct{ l, r float64 }, n)
	for i := range a {
		Fscan(in, &x, &y)
		if y > d {
			Fprint(out, -1)
			return
		}
		v := math.Sqrt(float64(d*d - y*y))
		a[i].l = float64(x) - v
		a[i].r = float64(x) + v
	}
	sort.Slice(a, func(i, j int) bool { return a[i].r < a[j].r })
	pre := a[0].r
	for i := 0; i < n; {
		ans++
		for ; i < n && a[i].l-eps < pre; i++ {
		}
		if i < n {
			pre = a[i].r
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
