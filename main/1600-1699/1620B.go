package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1620B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}

	var T, w, h, n, l, r int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &w, &h)
		ans := int64(0)
		for i := 0; i < 4; i++ {
			for Fscan(in, &n, &l); n > 1; n-- {
				Fscan(in, &r)
			}
			if i < 2 {
				ans = max(ans, (r-l)*h)
			} else {
				ans = max(ans, (r-l)*w)
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1620B(os.Stdin, os.Stdout) }
