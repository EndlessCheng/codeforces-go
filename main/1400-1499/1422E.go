package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1422E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var s string
	Fscan(in, &s)
	n := len(s)
	ans := make([]struct{sz int; s string}, n)
	st := make([]byte, n+1)
	j := n
	for i := n - 1; i >= 0; i-- {
		c := s[i]
		if c == st[j] && c > st[j+1] {
			j++
		} else {
			j--
			st[j] = c
		}
		ans[i].sz = n - j
		if n-j > 10 {
			ans[i].s = Sprintf("%s...%s", st[j:j+5], st[n-2:n])
		} else {
			ans[i].s = string(st[j:n])
		}
	}
	for _, p := range ans {
		Fprintln(out, p.sz, p.s)
	}
}

//func main() { cf1422E(bufio.NewReader(os.Stdin), os.Stdout) }
