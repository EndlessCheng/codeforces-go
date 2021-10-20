package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1384B2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k, lim, d int
	for Fscan(in, &T); T > 0; T-- {
		bad, t := false, 0
		for Fscan(in, &n, &k, &lim); n > 0; n-- {
			Fscan(in, &d)
			if bad {
				continue
			}
			up := lim - d
			if up < 0 {
				bad = true
				continue
			}
			if up >= k { // safe
				t = 0
				continue
			}
			t++
			if t > k {
				if t-k > up {
					bad = true
				}
			} else if k-up > t {
				t = k - up
			}
		}
		if bad {
			Fprintln(out, "No")
		} else {
			Fprintln(out, "Yes")
		}
	}
}

//func main() { CF1384B2(os.Stdin, os.Stdout) }
