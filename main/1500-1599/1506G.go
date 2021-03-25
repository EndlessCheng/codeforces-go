package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1506G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		cnt := ['z' + 1]int{}
		for _, b := range s {
			cnt[b]++
		}
		stk := []byte{}
		inStk := ['z' + 1]bool{}
		for _, b := range s {
			cnt[b]--
			if inStk[b] {
				continue
			}
			for len(stk) > 0 && stk[len(stk)-1] <= b && cnt[stk[len(stk)-1]] > 0 {
				inStk[stk[len(stk)-1]] = false
				stk = stk[:len(stk)-1]
			}
			stk = append(stk, b)
			inStk[b] = true
		}
		Fprintf(out, "%s\n", stk)
	}
}

//func main() { CF1506G(os.Stdin, os.Stdout) }
