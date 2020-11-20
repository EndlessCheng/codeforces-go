package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1440B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n*k)
		for i := range a {
			Fscan(in, &a[i])
		}
		s := int64(0)
		for i, c := n*k, 0; c < k; c++ {
			i -= n/2 + 1
			s += int64(a[i])
		}
		Fprintln(out, s)
	}
}

//func main() { CF1440B(os.Stdin, os.Stdout) }
