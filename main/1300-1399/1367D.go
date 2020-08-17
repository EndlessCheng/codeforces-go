package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1367D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var t, n int
	var s []byte
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &s, &n)
		cnt := [26]int{}
		for _, b := range s {
			cnt[b-'a']++
		}
		p := byte(25)
		q := []int{}
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			if a[i] == 0 {
				q = append(q, i)
			}
		}

		ans := make([]byte, n)
		for ; len(q) > 0; p-- {
			for ; cnt[p] < len(q); p-- {
			}
			qq := q
			q = []int{}
			for _, i := range qq {
				ans[i] = 'a' + p
				for j := range a {
					if j != i {
						if a[j] -= abs(j - i); a[j] == 0 {
							q = append(q, j)
						}
					}
				}
			}
		}
		Fprintln(out, string(ans))
	}
}

//func main() { CF1367D(os.Stdin, os.Stdout) }
