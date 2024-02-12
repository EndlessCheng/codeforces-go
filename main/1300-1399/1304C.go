package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1304C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ t, down, up int64 }

	solve := func() (ans bool) {
		var n, initT, t, down, up int64
		Fscan(in, &n, &initT)
		a := make([]pair, n)
		b := make([]pair, 0, n)
		for i := range a {
			Fscan(in, &t, &down, &up)
			a[i] = pair{t, down, up}
			if i == 0 || t > a[i-1].t {
				b = append(b, a[i])
			} else {
				b[len(b)-1].down = max(b[len(b)-1].down, down)
				b[len(b)-1].up = min(b[len(b)-1].up, up)
			}
		}

		down, up = initT, initT
		for i, p := range b {
			if p.down > p.up {
				return false
			}
			if i == 0 {
				down -= p.t
				up += p.t
			} else {
				down -= p.t - b[i-1].t
				up += p.t - b[i-1].t
			}
			if down > p.up || up < p.down {
				return false
			}
			down = max(down, p.down)
			up = min(up, p.up)
		}
		return true
	}

	var t int
	Fscan(in, &t)
	for _case := 1; _case <= t; _case++ {
		Fprintln(out, map[bool]string{true: "YES", false: "NO"} [solve()])
	}
}

//func main() { CF1304C(os.Stdin, os.Stdout) }
