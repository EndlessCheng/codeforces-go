package main

import (
	. "fmt"
	"io"
	"math/rand"
)

// https://github.com/EndlessCheng
func cf1305F(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	b := a
	const mx = 30 // 20 也可以
	if n > mx {
		rand.Shuffle(n, func(i, j int) { a[i], a[j] = a[j], a[i] })
		b = a[:mx]
	}
	ps := map[int]int8{}
	for _, x := range b {
		for i := x - 1; i < x+2; i++ {
			v := i
			for p := 2; p*p <= v; p++ {
				if v%p > 0 {
					continue
				}
				for v /= p; v%p == 0; v /= p {
				}
				ps[p] = 0
			}
			if v > 1 {
				ps[v] = 0
			}
		}
	}

	ans := n
o:
	for p := range ps {
		s := 0
		for _, v := range a {
			if v < p {
				s += p - v
			} else {
				s += min(v%p, p-v%p)
			}
			if s >= ans { // 最优性剪枝
				continue o
			}
		}
		ans = s
	}
	Fprint(out, ans)
}

//func main() { cf1305F(bufio.NewReader(os.Stdin), os.Stdout) }
