package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1321C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, ans int
	var s []byte
	Fscan(in, &n, &s)
	s = append([]byte{0}, s...)
	s = append(s, 0)
	for curB := byte('z'); curB >= 'a'; curB-- {
	outer:
		for {
			for i, b := range s {
				if b == curB && (b-1 == s[i-1] || b-1 == s[i+1]) {
					ans++
					s = append(s[:i], s[i+1:]...)
					continue outer
				}
			}
			break
		}
	}
	Fprint(out, ans)
}

//func main() { CF1321C(os.Stdin, os.Stdout) }
