package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1060D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n int
	Fscan(in, &n)
	l := make([]int, n)
	r := make([]int, n)
	for i := range l {
		Fscan(in, &l[i], &r[i])
	}
	sort.Ints(l)
	sort.Ints(r)
	ans := int64(n)
	for i, v := range l {
		ans += int64(max(v, r[i]))
	}
	Fprint(out, ans)
}

//func main() { CF1060D(os.Stdin, os.Stdout) }
