package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF374D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var q, n int
	Fscan(in, &q, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	t := make([]int, q+1)
	add := func(i, val int) {
		for ; i <= q; i += i & -i {
			t[i] += val
		}
	}
	sum := func(i int) (res int) {
		for ; i > 0; i &= i - 1 {
			res += t[i]
		}
		return
	}
	qs := make([]int8, q)
	for i := range qs {
		if Fscan(in, &qs[i]); qs[i] >= 0 {
			add(i+1, 1)
			continue
		}
		todo := []int{}
		for _, p := range a {
			j := sort.Search(q, func(j int) bool { return sum(j+1) >= p })
			if j == q {
				break
			}
			todo = append(todo, j)
		}
		for _, p := range todo {
			add(p+1, -1)
		}
	}
	ans := []byte{}
	for i, v := range qs {
		if sum(i+1) > sum(i) {
			ans = append(ans, '0'+byte(v))
		}
	}
	if len(ans) > 0 {
		Fprint(out, string(ans))
	} else {
		Fprint(out, "Poor stack!")
	}
}

//func main() { CF374D(os.Stdin, os.Stdout) }
