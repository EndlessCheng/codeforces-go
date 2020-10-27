package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1175C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		if k == 0 {
			Fprintln(out, a[0])
			continue
		}
		min := int(1e9)
		for i, v := range a[k:] {
			if d := (v - a[i] + 1) / 2; d < min {
				min = d
				x = (v + a[i]) / 2
			}
		}
		Fprintln(out, x)
	}
}

//func main() { CF1175C(os.Stdin, os.Stdout) }
