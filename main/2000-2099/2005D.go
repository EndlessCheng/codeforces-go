package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2005D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type pair struct{ x, y int }
		a := make([]pair, n)
		for i := range a {
			Fscan(in, &a[i].x)
		}
		for i := range a {
			Fscan(in, &a[i].y)
		}

		suf := make([]pair, n+1)
		for i := n - 1; i >= 0; i-- {
			suf[i].x = gcd(suf[i+1].x, a[i].x)
			suf[i].y = gcd(suf[i+1].y, a[i].y)
		}

		mx, cnt := 0, 0
		type result struct{ g1, g2, l int }
		res := []result{}
		pre := pair{}
		for i, v := range a {
			for j, p := range res {
				res[j].g1 = gcd(p.g1, v.y)
				res[j].g2 = gcd(p.g2, v.x)
			}
			res = append(res, result{gcd(pre.x, v.y), gcd(pre.y, v.x), i})

			j := 1
			for k := 1; k < len(res); k++ {
				if res[k].g1 != res[k-1].g1 || res[k].g2 != res[k-1].g2 {
					res[j] = res[k]
					j++
				}
			}
			res = res[:j]

			prePos := i + 1
			for k := len(res) - 1; k >= 0; k-- {
				p := res[k]
				s := gcd(p.g1, suf[i+1].x) + gcd(p.g2, suf[i+1].y)
				if s > mx {
					mx = s
					cnt = prePos - p.l
				} else if s == mx {
					cnt += prePos - p.l
				}
				prePos = p.l
			}

			pre.x = gcd(pre.x, v.x)
			pre.y = gcd(pre.y, v.y)
		}
		Fprintln(out, mx, cnt)
	}
}

//func main() { cf2005D(bufio.NewReader(os.Stdin), os.Stdout) }
