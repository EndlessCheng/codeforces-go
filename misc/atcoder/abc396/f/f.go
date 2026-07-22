package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
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

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, v, ans int
	Fscan(in, &n, &m)
	t := make(fenwick, m+1)
	s := make([]int, m)
	for i := range n {
		Fscan(in, &v)
		ans += i - t.pre(v+1)
		t.add(v + 1)
		s[v] += i*2 - n + 1
	}

	for i := m - 1; i >= 0; i-- {
		Fprintln(out, ans)
		ans += s[i]
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
