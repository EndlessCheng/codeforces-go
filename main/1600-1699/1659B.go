package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1659B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s)
		pos := [2][]int{}
		for i, b := range s {
			pos[b&1] = append(pos[b&1], i)
		}

		ans := make([]int, n)
		p := pos[k&1]
		m := len(p)
		for i := 1; i < m && k > k&1; i += 2 {
			s[p[i-1]] ^= 1
			s[p[i]] ^= 1
			ans[p[i-1]]++
			ans[p[i]]++
			k -= 2
		}

		if k&1 > 0 {
			i := bytes.IndexByte(s, '1')
			if i < 0 {
				i = n - 1
			}
			for j := range s {
				if j != i {
					s[j] ^= 1
				}
			}
			ans[i]++
			k--
		} else if k > 0 && m&1 > 0 {
			s[p[m-1]] ^= 1
			s[n-1] ^= 1
			ans[p[m-1]]++
			ans[n-1]++
			k -= 2
		}
		ans[0] += k

		Fprintf(out, "%s\n", s)
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1659B(os.Stdin, os.Stdout) }
