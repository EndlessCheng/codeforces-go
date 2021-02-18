package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1487D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	a := []int{}
	for i := 1; 2*i*(i+1)+1 < 1e9; i++ {
		a = append(a, 2*i*(i+1)+1)
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fprintln(out, sort.SearchInts(a, n+1))
	}
}

//func main() { CF1487D(os.Stdin, os.Stdout) }
