package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1491B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var T, n, ud, lr int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &ud, &lr)
		a := make([]int, n)
		d := 0
		for i := range a {
			if Fscan(in, &a[i]); i > 0 && abs(a[i]-a[i-1]) > d {
				d = abs(a[i] - a[i-1])
			}
		}
		if d > 1 {
			Fprintln(out, 0)
		} else if d == 1 {
			Fprintln(out, min(ud, lr))
		} else {
			Fprintln(out, lr+min(ud, lr))
		}
	}
}

//func main() { CF1491B(os.Stdin, os.Stdout) }
