package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1217C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	var s []byte
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &s)
		n, ans := len(s), 0
		pre1 := make([]int, n)
		p := -1
		for i, b := range s {
			pre1[i] = p
			if b == '1' {
				p = i
			}
			s[i] -= '0'
		}
		for i := range s {
			for j, l, v := i, 1, 0; j >= 0; j-- {
				v |= int(s[j]) << (l - 1)
				if v == l {
					ans++
				}
				if l > 20 {
					if l < v && v <= i-pre1[j] {
						ans++
					}
					break
				}
				l++
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1217C(os.Stdin, os.Stdout) }
