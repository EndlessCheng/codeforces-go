package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/big"
	"os"
	"sort"
)

// https://www.luogu.com.cn/problem/P1080

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, x, y int64
	Fscan(in, &n, &x, &y)
	a := make([]struct{ x, y int64 }, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
	}
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.x*a.y < b.x*b.y })

	ans := big.NewInt(0)
	m := big.NewInt(x)
	for _, p := range a {
		if x := new(big.Int).Quo(m, big.NewInt(p.y)); x.Cmp(ans) > 0 {
			ans = x
		}
		m.Mul(m, big.NewInt(p.x))
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
