package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1358F(_r io.Reader, out io.Writer) {
	eq := func(a, b []int64) bool {
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}
	eqR := func(a, b []int64) bool {
		for i, v := range a {
			if v != b[len(a)-1-i] {
				return false
			}
		}
		return true
	}

	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int64, n)
	for i := range b {
		Fscan(in, &b[i])
	}

	cntP := int64(0)
	ans := []byte{}
	if n == 1 {
		if a[0] != b[0] {
			Fprint(out, "IMPOSSIBLE")
		} else {
			Fprint(out, "SMALL\n0")
		}
		return
	}
	x, y := a[0], a[1]
	if x > y {
		x, y = y, x
	}
	for !eq(b, a) {
		if b[0] < 1 {
			Fprint(out, "IMPOSSIBLE")
			return
		}
		if eqR(b, a) {
			ans = append(ans, 'R')
			break
		}
		if sort.SliceIsSorted(b, func(i, j int) bool { return b[i] >= b[j] }) {
			for i, j := 0, n-1; i < j; i++ {
				b[i], b[j] = b[j], b[i]
				j--
			}
			ans = append(ans, 'R')
		} else if sort.SliceIsSorted(b, func(i, j int) bool { return b[i] <= b[j] }) {
			if n == 2 {
				c := b[1] / b[0]
				if b[0] == x {
					if (b[1]-y)%x != 0 {
						Fprint(out, "IMPOSSIBLE")
						return
					}
					c = (b[1] - y) / x
					b[1] = y
				} else {
					b[1] %= b[0]
				}
				cntP += c
				if cntP <= 2e5 {
					ans = append(ans, bytes.Repeat([]byte{'P'}, int(c))...)
				}
				continue
			}
			for i := len(b) - 1; i > 0; i-- {
				b[i] -= b[i-1]
			}
			cntP++
			if cntP <= 2e5 {
				ans = append(ans, 'P')
			}
		} else {
			Fprint(out, "IMPOSSIBLE")
			return
		}
	}
	if cntP > 2e5 {
		Fprintln(out, "BIG")
		Fprint(out, cntP)
		return
	}
	for i, j := 0, len(ans)-1; i < j; i++ {
		ans[i], ans[j] = ans[j], ans[i]
		j--
	}
	Fprintln(out, "SMALL")
	Fprintln(out, len(ans))
	Fprint(out, string(ans))
}

func main() { CF1358F(os.Stdin, os.Stdout) }
