package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1430B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Ints(a)
		s := int64(0)
		for _, v := range a[n-k-1:] {
			s += int64(v)
		}
		Fprintln(out, s)
	}
}

//func main() { CF1430B(os.Stdin, os.Stdout) }
