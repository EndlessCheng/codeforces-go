package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF274A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, ans int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	ban := map[int]bool{}
	for i := n - 1; i >= 0; i-- {
		if v := a[i]; !ban[v] {
			ans++
			if v%k == 0 {
				ban[v/k] = true
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF274A(os.Stdin, os.Stdout) }
