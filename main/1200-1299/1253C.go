package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1253C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	s := make([]int64, m)
	ans := int64(0)
	for i, v := range a {
		s[i%m] += int64(v)
		ans += s[i%m]
		Fprint(out, ans, " ")
	}
}

//func main() { CF1253C(os.Stdin, os.Stdout) }
