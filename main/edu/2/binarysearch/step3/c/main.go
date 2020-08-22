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
	if k == 1 {
		Fprint(out, a[n-1]-1)
		return
	}
	Fprint(out, sort.Search(1e9+1, func(d int) bool {
		p, c := a[0], 1
		for _, v := range a {
			if v-p >= d {
				p = v
				c++
			}
		}
		return c < k
	})-1)
}

func main() { run(os.Stdin, os.Stdout) }
