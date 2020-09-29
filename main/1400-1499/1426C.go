package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1426C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fprintln(out, sort.Search(63244, func(k int) bool { x := k / 2; return (1+x)*(1+k-x) >= n || (2+x)*(k-x) >= n }))
	}
}

//func main() { CF1426C(os.Stdin, os.Stdout) }
