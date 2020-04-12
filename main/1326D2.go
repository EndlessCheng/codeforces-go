package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1326D2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var maxLen []int
	manacher := func(origin string) {
		n := len(origin)
		s := make([]byte, 2*n+3)
		s[0] = '^'
		for i := range origin {
			s[2*i+1] = '#'
			s[2*i+2] = origin[i]
		}
		s[2*n+1] = '#'
		s[2*n+2] = '$'
		maxLen = make([]int, 2*n+3)
		var mid, right int
		for i := 1; i < 2*n+2; i++ {
			if i < right {
				maxLen[i] = min(maxLen[2*mid-i], right-i)
			} else {
				maxLen[i] = 1
			}
			for s[i+maxLen[i]] == s[i-maxLen[i]] {
				maxLen[i]++
			}
			if right < i+maxLen[i] {
				mid = i
				right = i + maxLen[i]
			}
		}
	}
	q := func(l, r int) bool { return maxLen[l+r+2]-1 >= r-l+1 }

	var t int
	var s string
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &s)
		i, j := 0, len(s)-1
		for i < j && s[i] == s[j] {
			i++
			j--
		}
		if i >= j {
			Fprintln(out, s)
			continue
		}
		l, r := i, j
		manacher(s)
		for ; r > i && !q(i, r); r-- {
		}
		for ; l < j && !q(l, j); l++ {
		}
		if r-i > j-l {
			Fprintln(out, s[:r+1]+s[j+1:])
		} else {
			Fprintln(out, s[:i]+s[l:])
		}
	}
}

//func main() { CF1326D2(os.Stdin, os.Stdout) }
