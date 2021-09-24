package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1339B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Ints(a)
		ans := make([]interface{}, n)
		for i := n - 2 + n&1; i >= 0; i -= 2 {
			ans[i] = a[0]
			a = a[1:]
		}
		for i := 1; i < n; i += 2 {
			ans[i] = a[0]
			a = a[1:]
		}
		Fprintln(out, ans...)
	}
}

//func main() { CF1339B(os.Stdin, os.Stdout) }
