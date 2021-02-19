package main

import (
	. "fmt"
	"os"
)

type (
	input86 struct{ n int }
	req86   struct{ l, r int }
	resp86  struct{ v int }
	guess86 struct{ ans int }
)

// github.com/EndlessCheng/codeforces-go
func CF1486C2(in input86, Q func(req86) resp86) (gs guess86) {
	q := func(l, r int) int { return Q(req86{l, r}).v }
	n := in.n
	ans := 0
	defer func() { gs.ans = ans }()

	p := q(1, n)
	l, r := 1, n+1
	ll, rr := l, r
	for {
		if r-l == 2 {
			ans = l + r - 1 - q(l, r-1)
			return
		}

		if r-l == 3 {
			p, pl, pr := q(l, r-1), q(l, l+1), q(l+1, l+2)
			if pl == pr {
				if p > pl {
					ans = pl - 1
					return
				}
				ans = pl + 1
				return
			}
			if pl == p {
				ans = 2*l + 1 - p
				return
			}
			ans = 2*l + 3 - p
			return
		}

		m := (l + r) >> 1
		if l == ll && r == rr {
			if p < m {
				if pp := q(l, m-1); pp == p {
					r, rr = m, m
				} else {
					l, ll = m, p
				}
			} else {
				if pp := q(m, r-1); pp == p {
					l, ll = m, m
				} else {
					r, rr = m, p+1
				}
			}
		} else if r == rr {
			if pp := q(ll, m-1); pp == p {
				r, rr = m, m
			} else {
				l = m
			}
		} else {
			if pp := q(m, rr-1); pp == p {
				l, ll = m, m
			} else {
				r = m
			}
		}
	}
}

func ioq86() {
	in := os.Stdin
	out := os.Stdout

	Q := func(req req86) (resp resp86) {
		Fprintln(out, "?", req.l, req.r)
		//out.Flush()
		Fscan(in, &resp.v)
		return
	}

	d := input86{}
	Fscan(in, &d.n)
	gs := CF1486C2(d, Q)
	ans := gs.ans
	Fprintln(out, "!", ans)
	//out.Flush()
}

//func main() { ioq86() }
