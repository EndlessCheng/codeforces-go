package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
type fenwick []int

func (t fenwick) add(i int) {
	for ; i < len(t); i += i & -i {
		t[i]++
	}
}

func (t fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

func cf1725L(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		s[i] += s[i-1]
	}
	if s[n] < 0 {
		Fprint(out, -1)
		return
	}
	for i := 1; i < n; i++ {
		if s[i] < 0 || s[i] > s[n] {
			Fprint(out, -1)
			return
		}
	}

	b := slices.Clone(s[1:])
	slices.Sort(b)

	ans := 0
	t := make(fenwick, n+1)
	for i := 1; i < n; i++ {
		x := sort.SearchInts(b, s[i]) + 1
		ans += i - 1 - t.pre(x)
		t.add(x)
	}
	Fprint(out, ans)
}

//func main() { cf1725L(bufio.NewReader(os.Stdin), os.Stdout) }
