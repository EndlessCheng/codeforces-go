package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1415D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v int
	Fscan(in, &n)
	if n > 60 {
		Fprint(out, 1)
		return
	}
	s := make([]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		s[i+1] = s[i] ^ v
	}
	ans := n
	for m := 1; m < n; m++ {
	o:
		for r := m + 1; r <= n; r++ {
			for l := m - 1; l >= 0 && r-l-2 < ans; l-- {
				if s[r]^s[m] < s[m]^s[l] {
					ans = r - l - 2
					continue o
				}
			}
		}
	}
	if ans == n {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { CF1415D(os.Stdin, os.Stdout) }
