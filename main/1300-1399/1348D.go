package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1348D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		ans := []int{}
		for i := 1; i <= n; i <<= 1 {
			ans = append(ans, i)
			n -= i
		}
		if n > 0 {
			ans = append(ans, n)
		}
		sort.Ints(ans)
		for i := len(ans) - 1; i > 0; i-- {
			ans[i] -= ans[i-1]
		}
		Fprintln(out, len(ans)-1)
		for _, v := range ans[1:] {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1348D(os.Stdin, os.Stdout) }
