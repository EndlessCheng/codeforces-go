package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1553B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var T int
	var txt, tar []byte
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &txt, &tar)
		n := len(tar)
		halfLen := make([]int, n)
		for i, mid, r := 0, 0, 0; i < n; i++ {
			hl := 1
			if i < r {
				hl = min(halfLen[mid*2-i], r-i)
			}
			for ; i >= hl && i+hl < n && tar[i-hl] == tar[i+hl]; hl++ {
			}
			if i+hl > r {
				mid, r = i, i+hl
			}
			halfLen[i] = hl
		}
		isP := func(l, r int) bool { return halfLen[(l+r)/2]*2 > r-l+1 }

		maxMatchPrefix := func(text, pattern []byte) int {
			match := make([]int, len(pattern))
			for i, c := 1, 0; i < len(pattern); i++ {
				v := pattern[i]
				for c > 0 && pattern[c] != v {
					c = match[c-1]
				}
				if pattern[c] == v {
					c++
				}
				match[i] = c
			}
			lenP := len(pattern)
			c, mx := 0, 0
			for _, v := range text {
				for c > 0 && pattern[c] != v {
					c = match[c-1]
				}
				if pattern[c] == v {
					if c++; c > mx {
						mx = c
					}
				}
				if c == lenP {
					return lenP
				}
			}
			return mx
		}

		preMax := maxMatchPrefix(txt, tar)
		for i := 0; i < n/2; i++ {
			tar[i], tar[n-1-i] = tar[n-1-i], tar[i]
		}
		sufMax := maxMatchPrefix(txt, tar)
		for i := 0; i < n; i += 2 {
			if isP(0, i) && n-i/2 <= sufMax || isP(n-1-i, n-1) && n-i/2 <= preMax {
				Fprintln(out, "YES")
				continue o
			}
		}
		Fprintln(out, "NO")
	}
}

//func main() { CF1553B(os.Stdin, os.Stdout) }
