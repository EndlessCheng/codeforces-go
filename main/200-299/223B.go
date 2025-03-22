package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf223B(in io.Reader, out io.Writer) {
	var s, t []byte
	Fscan(in, &s, &t)
	n, m := len(s), len(t)
	preCnt := make([]int, n)
	cnt := 0
	last := [26]int{}
	for i, b := range s {
		if cnt < m && t[cnt] == b {
			cnt++
			last[b-'a'] = cnt
		}
		if c := last[b-'a']; c > 0 {
			preCnt[i] = c
		}
	}

	cnt = m
	last = [26]int{}
	for i := n - 1; i >= 0; i-- {
		b := s[i]
		if cnt > 0 && t[cnt-1] == b {
			last[b-'a'] = cnt
			cnt--
		}
		sufCnt := int(1e9)
		if c := last[b-'a']; c > 0 {
			sufCnt = c
		}
		if preCnt[i] < sufCnt {
			Fprint(out, "No")
			return
		}
	}
	Fprint(out, "Yes")
}

//func main() { cf223B(bufio.NewReader(os.Stdin), os.Stdout) }
