package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1323C(_r io.Reader, _w io.Writer) {
	var n, cntL, ans int
	var s []byte
	Fscan(bufio.NewReader(_r), &n, &s)
	for _, b := range s {
		if b == '(' {
			cntL++
		}
	}
	if 2*cntL != n {
		Fprint(_w, -1)
		return
	}

	cntL = 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			cntL++
			continue
		}
		cntL--
		if cntL >= 0 {
			continue
		}
		st := i
		for i++; i < n; i++ {
			if s[i] == '(' {
				cntL++
				if cntL == 0 {
					break
				}
			} else {
				cntL--
			}
		}
		ans += i - st + 1
	}
	Fprint(_w, ans)
}

//func main() { CF1323C(os.Stdin, os.Stdout) }
