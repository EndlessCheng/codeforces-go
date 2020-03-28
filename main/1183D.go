package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1183D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, tp int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		types := make([]int, n+1)
		for i := 0; i < n; i++ {
			Fscan(in, &tp)
			types[tp]++
		}
		cnts := make([]int, n+1)
		for _, t := range types {
			cnts[t]++
		}
		ans := 0
		cnt := 0
		for i := n; i > 0; i-- {
			cnt += cnts[i]
			if cnt > 0 {
				ans += i
				cnt--
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1183D(os.Stdin, os.Stdout) }
