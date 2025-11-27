package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"strings"
)

// https://github.com/EndlessCheng
func cf1679E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353
	var n, q int
	var s string
	Fscan(in, &n, &s, &q)
	cq := strings.Count(s, "?")

	const mx = 17
	pow := make([][mx + 1]int, cq+1)
	for sz := 1; sz <= mx; sz++ {
		pow[0][sz] = 1
		for j := 1; j <= cq; j++ {
			pow[j][sz] = pow[j-1][sz] * sz % mod
		}
	}

	f := [1 << mx][mx + 1]int{}
	for i := range 2*n - 1 {
		c := cq
		must := 0
		l, r := i/2, (i+1)/2
		for l >= 0 && r < n {
			if s[l] == '?' || s[r] == '?' {
				if l < r {
					c--
				}
				if s[l] != '?' {
					must |= 1 << (s[l] - 'a')
				} else if s[r] != '?' {
					must |= 1 << (s[r] - 'a')
				}
			} else if s[l] != s[r] {
				break
			}
			for sz := bits.OnesCount(uint(must)); sz <= mx; sz++ {
				f[must][sz] = (f[must][sz] + pow[c][sz]) % mod
			}
			l--
			r++
		}
	}

	for i := range mx {
		for s := 0; s < 1<<mx; s++ {
			s |= 1 << i
			for sz := bits.OnesCount(uint(s)); sz <= mx; sz++ {
				f[s][sz] = (f[s][sz] + f[s^1<<i][sz]) % mod
			}
		}
	}

	for range q {
		Fscan(in, &s)
		m := 0
		for _, b := range s {
			m |= 1 << (b - 'a')
		}
		Fprintln(out, f[m][bits.OnesCount(uint(m))])
	}
}

//func main() { cf1679E(bufio.NewReader(os.Stdin), os.Stdout) }
