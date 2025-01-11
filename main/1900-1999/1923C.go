package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1923C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, q, v, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		s := make([]int, n+1)
		for i := range n {
			Fscan(in, &v)
			if v == 1 {
				s[i+1] = s[i] - 1
			} else {
				s[i+1] = s[i] + v - 1
			}
		}
		for range q {
			Fscan(in, &l, &r)
			if l < r && s[l-1] <= s[r] {
				Fprintln(out, "YES")
			} else {
				Fprintln(out, "NO")
			}
		}
	}
}

//func main() { cf1923C(bufio.NewReader(os.Stdin), os.Stdout) }
