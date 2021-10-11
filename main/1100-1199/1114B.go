package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1114B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, k int
	Fscan(in, &n, &m, &k)
	a := make([]int, n)
	id := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return a[id[i]] > a[id[j]] })
	id = id[:m*k]

	sum := int64(0)
	for _, i := range id {
		sum += int64(a[i])
	}
	Fprintln(out, sum)
	sort.Ints(id)
	for i := m - 1; i < len(id)-1; i += m {
		Fprint(out, id[i]+1, " ")
	}
}

//func main() { CF1114B(os.Stdin, os.Stdout) }
