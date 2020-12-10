package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1156C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, d, ans int
	Fscan(in, &n, &d)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	for i, j := 0, n/2; i < n/2 && j < n; i++ {
		for ; j < n && a[i]+d > a[j]; j++ {
		}
		if j < n {
			ans++
			j++
		}
	}
	Fprint(out, ans)
}

//func main() { CF1156C(os.Stdin, os.Stdout) }
