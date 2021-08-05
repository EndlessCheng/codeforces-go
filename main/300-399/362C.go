package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF362C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, tot, cnt int
	max := -1
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		a[i]++
	}
	for i, v := range a {
		tree := make([]int, n+1)
		sum := func(i int) (s int) {
			for ; i > 0; i &= i - 1 {
				s += tree[i]
			}
			return
		}
		for _, w := range a[i+1:] {
			if v <= w {
				continue
			}
			tot++
			mx := sum(v) - sum(w)
			if mx > max {
				max, cnt = mx, 1
			} else if mx == max {
				cnt++
			}
			for i := w; i <= n; i += i & -i {
				tree[i]++
			}
		}
	}
	Fprint(out, tot-(max*2+1), cnt)
}

//func main() { CF362C(os.Stdin, os.Stdout) }
