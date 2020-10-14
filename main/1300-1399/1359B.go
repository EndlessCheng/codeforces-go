package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1359B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, x, y int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &x, &y)
		if y > 2*x {
			y = 2 * x
		}
		ans := 0
		for ; n > 0; n-- {
			Fscan(in, &s)
			for _, s := range strings.Split(s, "*") {
				if s != "" {
					ans += len(s)/2*y + len(s)&1*x
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1359B(os.Stdin, os.Stdout) }
