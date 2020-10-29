package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1151D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var b int
	var n, ans int64
	Fscan(in, &n)
	a := make(sort.IntSlice, n)
	for i := range a {
		Fscan(in, &a[i], &b)
		a[i] -= b
		ans += int64(b)
	}
	ans *= n - 1
	sort.Sort(sort.Reverse(a))
	for i, v := range a {
		ans += int64(i) * int64(v)
	}
	Fprint(out, ans)
}

//func main() { CF1151D(os.Stdin, os.Stdout) }
