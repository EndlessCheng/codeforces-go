package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1396B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		sum, max := 0, 0
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			sum += v
			if v > max {
				max = v
			}
		}
		if max > sum-max || sum&1 > 0 {
			Fprintln(out, "T")
		} else {
			Fprintln(out, "HL")
		}
	}
}

//func main() { CF1396B(os.Stdin, os.Stdout) }
