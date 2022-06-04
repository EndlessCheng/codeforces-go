package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1598C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		s := int64(0)
		for i := range a {
			Fscan(in, &a[i])
			s += int64(a[i])
		}
		s *= 2
		if s%int64(n) > 0 {
			Fprintln(out, 0)
			continue
		}
		s /= int64(n)
		ans := int64(0)
		c := map[int]int{}
		for _, v := range a {
			ans += int64(c[int(s)-v])
			c[v]++
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1598C(os.Stdin, os.Stdout) }
