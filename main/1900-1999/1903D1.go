package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1903D1(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, k int
	Fscan(in, &n, &q)
	const mx = 20
	low := n << mx
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		low -= a[i]
	}

	for range q {
		Fscan(in, &k)
		if k >= low {
			Fprintln(out, 1<<mx+(k-low)/n)
			continue
		}
		ans := 0
		for i := mx - 1; i >= 0; i-- {
			nxt := ans | 1<<i
			c := 0
			for _, v := range a {
				mask := 1<<bits.Len(uint(nxt&^v)) - 1
				c += nxt&mask - v&mask
			}
			if c <= k {
				ans = nxt
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1903D1(bufio.NewReader(os.Stdin), os.Stdout) }
