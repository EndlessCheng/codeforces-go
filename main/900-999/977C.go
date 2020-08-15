package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF977C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a) // 此题用作测试 Go 的快排性能，见 https://codeforces.com/contest/977/submission/75301978
	if k == n {
		Fprint(_w, int(1e9))
	} else if k > 0 && a[k-1] != a[k] {
		Fprint(_w, a[k]-1)
	} else if k == 0 && a[0] != 1 {
		Fprint(_w, 1)
	} else {
		Fprint(_w, -1)
	}
}

//func main() { CF977C(os.Stdin, os.Stdout) }
