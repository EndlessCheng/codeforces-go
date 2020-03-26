package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1292B(_r io.Reader, _w io.Writer) {
	type point struct{ x, y int64 }
	abs := func(x int64) int64 {
		if x < 0 {
			return -x
		}
		return x
	}
	dis := func(a, b point) int64 { return abs(a.x-b.x) + abs(a.y-b.y) }

	var ax, ay, bx, by, tt int64
	var p, st point
	Fscan(_r, &p.x, &p.y, &ax, &ay, &bx, &by, &st.x, &st.y, &tt)
	ps := []point{}
	// 没注意范围 WA 了一发
	for ; p.x < 9e16 && p.y < 9e16; p.x, p.y = ax*p.x+bx, ay*p.y+by {
		ps = append(ps, p)
	}

	ans := 0
	for i, p := range ps {
		t := tt
		// 注：因为算的不是欧几里得距离，这样写复杂了，只需要算两段距离就行，但是这样写适用性更广。
		if d := dis(st, p); d <= t {
			t -= d
			c := 1
			for j := i; j > 0; j-- {
				if d := dis(ps[j], ps[j-1]); d <= t {
					t -= d
					c++
				} else {
					break
				}
			}
			if i+1 < len(ps) {
				if d := dis(ps[0], ps[i+1]); d <= t {
					t -= d
					c++
					for j := i + 2; j < len(ps); j++ {
						if d := dis(ps[j], ps[j-1]); d <= t {
							t -= d
							c++
						} else {
							break
						}
					}
				}
			}
			if c > ans {
				ans = c
			}
		}
	}
	Fprint(_w, ans)
}

//func main() { CF1292B(os.Stdin, os.Stdout) }
