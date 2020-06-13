package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF652B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	for i, j := 0, n-1; i < j; i++ {
		Fprint(out, a[i], a[j], " ")
		j--
	}
	if n&1 > 0 {
		Fprint(out, a[n/2])
	}
}

//func main() { CF652B(os.Stdin, os.Stdout) }
