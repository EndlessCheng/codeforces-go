package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1674E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n int
	Fscan(in, &n)
	fi, se := int(1e9), int(1e9)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] < fi {
			fi, se = a[i], fi
		} else if a[i] < se {
			se = a[i]
		}
	}
	ans := (fi+1)/2 + (se+1)/2
	for i := 1; i < n; i++ {
		v, w := a[i-1], a[i]
		if i < n-1 {
			m := min(v, a[i+1])
			ans = min(ans, m+(max(v, a[i+1])-m+1)/2)
		}
		if v < w {
			v, w = w, v
		}
		if v >= w*2 {
			ans = min(ans, (v+1)/2)
		} else {
			d := v - w
			w -= d
			ans = min(ans, d+w/3*2+w%3)
		}
	}
	Fprint(out, ans)
}

//func main() { CF1674E(os.Stdin, os.Stdout) }
