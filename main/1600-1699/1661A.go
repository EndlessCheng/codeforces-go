package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1661A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int64) int64 {
		if x < 0 {
			return -x
		}
		return x
	}
	min := func(a, b int64) int64 {
		if a > b {
			return b
		}
		return a
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int64, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int64, n)
		for i := range b {
			Fscan(in, &b[i])
		}
		ans := int64(0)
		for i := 1; i < n; i++ {
			ans += min(abs(a[i]-a[i-1])+abs(b[i]-b[i-1]), abs(a[i]-b[i-1])+abs(b[i]-a[i-1]))
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1661A(os.Stdin, os.Stdout) }
