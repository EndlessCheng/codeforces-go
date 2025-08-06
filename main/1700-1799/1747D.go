package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1747D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, l, r int
	Fscan(in, &n, &q)
	a := make([]int, n)
	sp := make([]int, n+1)
	s := make([]int, n+1)
	pre := make([]int, n+1)
	pos := map[int]int{0: 1} // 偏移一位，这样默认值 0 就表示不合法
	for i := range a {
		Fscan(in, &a[i])
		sp[i+1] = sp[i] + a[i]
		s[i+1] = s[i] ^ a[i]
		v := s[i+1]<<1 | (i+1)&1
		pre[i+1] = pos[v^1]
		pos[v] = i + 2 // 偏移一位
	}
	for range q {
		Fscan(in, &l, &r)
		l--
		if sp[r] == sp[l] {
			Fprintln(out, 0)
		} else if s[r] != s[l] || r-l < 3 {
			Fprintln(out, -1)
		} else if (r-l)%2 > 0 || a[l] == 0 || a[r-1] == 0 {
			Fprintln(out, 1)
		} else if pre[r] > l {
			Fprintln(out, 2)
		} else {
			Fprintln(out, -1)
		}
	}
}

//func main() { cf1747D(bufio.NewReader(os.Stdin), os.Stdout) }
