package main

import (
	. "fmt"
	"io"
	"slices"
)

// github.com/EndlessCheng/codeforces-go
func CF1175D(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		s[i] += s[i-1]
	}

	ans := k * s[n]
	s = s[1:n]
	slices.Sort(s)
	for _, v := range s[:k-1] {
		ans -= v
	}
	Fprint(out, ans)
}

//func main() { CF1175D(bufio.NewReader(os.Stdin), os.Stdout) }
