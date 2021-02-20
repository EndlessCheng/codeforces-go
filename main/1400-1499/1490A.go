package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1490A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := 0
		for i := 1; i < n; i++ {
			v, w := a[i-1], a[i]
			if v < w {
				v, w = w, v
			}
			for ; v > 2*w; v = (v + 1) / 2 {
				ans++
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1490A(os.Stdin, os.Stdout) }
