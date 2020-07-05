package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1375D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		a := make([]int, n)
		cnt := map[int]int{}
		for i := range a {
			Fscan(in, &a[i])
			cnt[a[i]]++
		}
		ans := []interface{}{}
		f := func(mex int) bool {
			for i, v := range a {
				if i != v {
					cnt[v]--
					cnt[mex]++
					a[i] = mex
					ans = append(ans, i+1)
					return false
				}
			}
			return true
		}
		for {
			mex := 0
			for ; cnt[mex] > 0; mex++ {
			}
			if mex >= n {
				if f(mex) {
					break
				}
			} else if a[mex] == mex {
				f(mex)
			} else {
				cnt[a[mex]]--
				cnt[mex]++
				a[mex] = mex
				ans = append(ans, mex+1)
			}
		}
		Fprintln(out, len(ans))
		Fprintln(out, ans...)
	}
}

//func main() { CF1375D(os.Stdin, os.Stdout) }
