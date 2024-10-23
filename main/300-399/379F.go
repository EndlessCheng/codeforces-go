package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf379F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var q int
	Fscan(in, &q)

	n := 5 + q*2
	const mx = 20
	pa := make([][mx]int, n)
	pa[2][0], pa[3][0], pa[4][0] = 1, 1, 1
	dep := make([]int, n)
	dep[2], dep[3], dep[4] = 1, 1, 1

	uptoDep := func(v, d int) int {
		for k := uint(dep[v] - d); k > 0; k &= k - 1 {
			v = pa[v][bits.TrailingZeros(k)]
		}
		return v
	}
	getLCA := func(v, w int) int {
		if dep[v] > dep[w] {
			v, w = w, v
		}
		w = uptoDep(w, dep[v])
		if w == v {
			return v
		}
		for i := mx - 1; i >= 0; i-- {
			if pv, pw := pa[v][i], pa[w][i]; pv != pw {
				v, w = pv, pw
			}
		}
		return pa[v][0]
	}
	getDis := func(v, w int) int { return dep[v] + dep[w] - dep[getLCA(v, w)]*2 }

	end1, end2, diameter := 2, 3, 2
	cur := 5
	for range q {
		var v int
		Fscan(in, &v)
		pa[cur][0] = v
		for i := 0; i+1 < mx; i++ {
			pa[cur][i+1] = pa[pa[cur][i]][i]
		}
		pa[cur+1] = pa[cur]
		dep[cur] = dep[v] + 1
		dep[cur+1] = dep[cur]

		if d := getDis(end1, cur); d > diameter {
			end2 = cur
			diameter = d
		} else if d = getDis(end2, cur); d > diameter {
			end1 = cur
			diameter = d
		}
		Fprintln(out, diameter)
		cur += 2
	}
}

//func main() { cf379F(bufio.NewReader(os.Stdin), os.Stdout) }
