package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1268A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	var s string
	Fscan(in, &n, &k, &s)
	base := s[:k]
o:
	for i := k; i < n; i += k {
		t := s[i:]
		if i+k <= n {
			t = s[i : i+k]
		}
		for j := range t {
			if t[j] != base[j] {
				if t[j] > base[j] {
					b := []byte(base)
					for j := k - 1; j >= 0; j-- {
						if b[j] < '9' {
							b[j]++
							break
						}
						b[j] = '0'
					}
					base = string(b)
				}
				break o
			}
		}
	}
	Fprintln(out, n)
	Fprintln(out, strings.Repeat(base, n/k)+base[:n%k])
}

//func main() { CF1268A(os.Stdin, os.Stdout) }
