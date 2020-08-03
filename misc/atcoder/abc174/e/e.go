package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Fprint(out, sort.Search(1e9, func(x int) bool {
		if x == 0 {
			return false
		}
		cnt := 0
		for _, v := range a {
			cnt += (v - 1) / x
		}
		return cnt <= k
	}))
}

func main() { run(os.Stdin, os.Stdout) }
