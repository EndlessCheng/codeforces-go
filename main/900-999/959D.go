package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF959D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mx int = 1435243
	primes := []int{}
	lpf := [mx + 1]int{}
	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			primes = append(primes, i)
			for j := i; j <= mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}

	var n, v, cur int
	lim := true
	vis := [mx + 1]bool{}
o:
	for Fscan(in, &n); n > 0; n-- {
		if lim {
			Fscan(in, &v)
			ps := []int{}
			for x := v; x > 1; {
				p := lpf[x]
				if vis[p] {
					lim = false
				o2:
					for v++; ; v++ {
						ps := []int{}
						for x := v; x > 1; {
							p := lpf[x]
							if vis[p] {
								continue o2
							}
							ps = append(ps, p)
							for x /= p; lpf[x] == p; x /= p {
							}
						}
						for _, p := range ps {
							vis[p] = true
						}
						Fprint(out, v, " ")
						continue o
					}
				}
				ps = append(ps, p)
				for x /= p; lpf[x] == p; x /= p {
				}
			}
			for _, p := range ps {
				vis[p] = true
			}
			Fprint(out, v, " ")
		} else {
			for vis[primes[cur]] {
				cur++
			}
			Fprint(out, primes[cur], " ")
			cur++
		}
	}
}

//func main() { CF959D(os.Stdin, os.Stdout) }
