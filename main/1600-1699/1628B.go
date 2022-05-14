package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1628B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		ok, s2, s3 := false, [26]int{}, [26][26]int{}
		for Fscan(in, &n); n > 0; n-- {
			if Fscan(in, &s); ok {
			} else if len(s) == 1 || s[0] == s[len(s)-1] {
				ok = true
			} else if len(s) == 2 {
				x, y := s[0]-'a', s[1]-'a'
				if s2[y]>>x&1 > 0 || s3[y][x] > 0 {
					ok = true
				}
				s2[x] |= 1 << y
			} else {
				x, y, z := s[0]-'a', s[1]-'a', s[2]-'a'
				if s2[z]>>y&1 > 0 || s3[z][y]>>x&1 > 0 {
					ok = true
				}
				s3[x][y] |= 1 << z
			}
		}
		if ok {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1628B(os.Stdin, os.Stdout) }
