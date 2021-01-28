package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type pair struct {
		g int64
		i int
	}
	unique := func(a []pair) []pair {
		j := 0
		for i := 1; i < len(a); i++ {
			if a[j].g != a[i].g {
				j++
			}
			a[j] = a[i]
		}
		return a[:j+1]
	}

	n, v, ps, ans := 0, int64(0), []pair{}, int(1e9)
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		for j, p := range ps {
			ps[j].g = gcd(p.g, v)
		}
		ps = unique(append(ps, pair{v, i}))
		if ps[0].g == 1 && i-ps[0].i+1 < ans {
			ans = i - ps[0].i + 1
		}
	}
	if ans == 1e9 {
		ans = -1
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func gcd(a, b int64) int64 {
	for a > 0 {
		a, b = b%a, a
	}
	return b
}
