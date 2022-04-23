package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1671B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var T, n int
O:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
	o:
		for l := a[0] - 1; l < a[0]+2; l++ {
			for i, v := range a {
				if abs(l+i-v) > 1 {
					continue o
				}
			}
			Fprintln(out, "YES")
			continue O
		}
		Fprintln(out, "NO")
	}
}

//func main() { CF1671B(os.Stdin, os.Stdout) }
