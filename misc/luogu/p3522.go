package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p3522(in io.Reader, out io.Writer) {
	var n, l, r, left, ans int
	Fscan(in, &n)
	type pair struct{ v, i int }
	q := []pair{}
	for i := 0; i < n; i++ {
		Fscan(in, &l, &r)
		for len(q) > 0 && q[0].v > r {
			left = q[0].i + 1
			q = q[1:]
		}
		ans = max(ans, i-left+1)
		for len(q) > 0 && q[len(q)-1].v <= l {
			q = q[:len(q)-1]
		}
		q = append(q, pair{l, i})
	}
	Fprint(out, ans)
}

//func main() { p3522(bufio.NewReader(os.Stdin), os.Stdout) }
