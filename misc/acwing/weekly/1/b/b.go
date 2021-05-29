package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	mid := a[n/2]
	Fprint(out, mid+sort.Search(1e10, func(x int) bool {
		x += mid
		s := 0
		for i := n / 2; i < n; i++ {
			if a[i] >= x {
				break
			}
			s += x - a[i]
		}
		return s > k
	})-1)
}

func main() { run(os.Stdin, os.Stdout) }
