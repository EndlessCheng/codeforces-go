package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1312C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, k, v int64
o:
	for Fscan(in, &t); t > 0; t-- {
		cnts := [55]int64{}
		for Fscan(in, &n, &k); n > 0; n-- {
			i := 0
			for Fscan(in, &v); v > 0; v /= k {
				cnts[i] += v % k
				i++
			}
		}
		for _, v := range cnts {
			if v > 1 {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1312C(os.Stdin, os.Stdout) }
