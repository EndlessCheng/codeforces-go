package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1620D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		has := map[int]bool{}
		c := [3]int{}
		mx := 0
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			has[v] = true
			c[v%3]++
			if v > mx {
				mx = v
			}
		}

		ans := mx / 3
		if mx%3 == 0 {
			if c[1] > 0 || c[2] > 0 {
				ans++ // 1 $1 + 1 $2 + mx/3-1 $3
			}
		} else if mx%3 == 1 {
			ans++ // $1, or 2 $2 + mx/3-1 $3
			if c[2] > 0 && (has[1] || has[mx-1]) {
				ans++ // $1
			}
		} else {
			ans++ // $2
			if c[1] > 0 {
				ans++ // $1
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1620D(os.Stdin, os.Stdout) }
