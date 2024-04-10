package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf296B(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n int
	var s, t string
	var le, gr bool
	Fscan(bufio.NewReader(in), &n, &s, &t)
	all, same := 1, 1
	for i := range s {
		v, w := s[i], t[i]
		if v == '?' || w == '?' {
			if v == '?' {
				all = all * 10 % mod
			}
			if w == '?' {
				all = all * 10 % mod
			}
			if v == '?' && w == '?' {
				same = same * 10 % mod
			}
		} else if v < w {
			le = true
		} else if v > w {
			gr = true
		}
	}
	if le && gr {
		Fprint(out, all)
		return
	}

	res1, res2 := 1, 1
	for i, v := range s {
		w := t[i]
		if v == '?' {
			if w == '?' {
				res1 = res1 * 55 % mod
				res2 = res2 * 55 % mod
			} else {
				res1 = res1 * int(w-'0'+1) % mod
				res2 = res2 * int('9'-w+1) % mod
			}
		} else if w == '?' {
			res1 = res1 * int('9'-v+1) % mod
			res2 = res2 * int(v-'0'+1) % mod
		}
	}

	if !gr {
		all -= res1
	}
	if !le {
		all -= res2
	}
	if !le && !gr {
		all += same
	}
	Fprint(out, (all%mod+mod)%mod)
}

//func main() { cf296B(os.Stdin, os.Stdout) }
