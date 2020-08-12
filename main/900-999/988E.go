package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF988E(in io.Reader, out io.Writer) {
	var s []byte
	Fscan(in, &s)
	n, ans := len(s), 99
	if n < 2 {
		Fprint(out, -1)
		return
	}
	for i := range s {
		for j := range s {
			if i == j {
				continue
			}
			t := make([]byte, n)
			copy(t, s)
			c := 0
			for k := i; k < n-1; k++ {
				t[k], t[k+1] = t[k+1], t[k]
				c++
			}
			k := j
			if j > i {
				k--
			}
			for ; k < n-2; k++ {
				t[k], t[k+1] = t[k+1], t[k]
				c++
			}
			for k = 0; t[k] == '0'; k++ {
			}
			for ; k > 0; k-- {
				t[k], t[k-1] = t[k-1], t[k]
				c++
			}
			if c < ans && (t[n-2]&15*10+t[n-1]&15)%25 == 0 {
				ans = c
			}
		}
	}
	if ans == 99 {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { CF988E(os.Stdin, os.Stdout) }
