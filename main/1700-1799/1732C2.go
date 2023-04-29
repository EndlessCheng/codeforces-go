package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1732C2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &n)
		a := make([]int, n)
		xor := make([]int, n+1)
		l1, l := make([]int, n), -1
		for i := range a {
			Fscan(in, &a[i])
			xor[i+1] = xor[i] ^ a[i]
			if a[i] > 0 {
				l = i
			}
			l1[i] = l
		}
		r1, r := make([]int, n+1), n
		r1[n] = n
		for i := n - 1; i >= 0; i-- {
			if a[i] > 0 {
				r = i
			}
			r1[i] = r
		}
		for ; n > 0; n-- {
			Fscan(in, &l, &r)
			l--
			r--
			s := xor[r+1] ^ xor[l]
			resL, resR := l, r
			i := r1[l]
			for ; i <= r; i = r1[i+1] {
				s2 := s
				for j := l1[r]; ; j = l1[j-1] {
					if j < i {
						j = i
					}
					if j == i || s2|a[j] != s2 {
						if j-i < resR-resL {
							resL, resR = i, j
						}
						break
					}
					s2 ^= a[j]
				}
				if s|a[i] != s {
					break
				}
				s ^= a[i]
			}
			if i > r {
				Fprintln(out, l+1, l+1)
			} else {
				Fprintln(out, resL+1, resR+1)
			}
		}
	}
}

//func main() { CF1732C2(os.Stdin, os.Stdout) }
