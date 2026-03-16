package main

import (
	. "fmt"
	"io"
	"math/rand"
	"time"
)

// https://github.com/EndlessCheng
func cf1198F(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	idx := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		idx[i] = i
	}

	ans := make([]any, n)
	t0 := time.Now()
	for time.Since(t0) < time.Second*4/10 {
		rand.Shuffle(n, func(i, j int) { idx[i], idx[j] = idx[j], idx[i] })
		p, q := 0, 0
		for _, i := range idx {
			if gcd(p, a[i]) != p {
				p = gcd(p, a[i])
				ans[i] = 1
			} else {
				q = gcd(q, a[i])
				ans[i] = 2
			}
		}
		if p == 1 && q == 1 {
			Fprintln(out, "YES")
			Fprintln(out, ans...)
			return
		}
	}
	Fprint(out, "NO")
}

//func main() { cf1198F(bufio.NewReader(os.Stdin), os.Stdout) }
