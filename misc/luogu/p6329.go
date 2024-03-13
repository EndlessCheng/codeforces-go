package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://space.bilibili.com/206214
type fenwick6329 []int

func (f fenwick6329) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

func (f fenwick6329) pre(i int) (res int) {
	for i = min(i, len(f)-1); i > 0; i &= i - 1 {
		res += f[i]
	}
	return res
}

func p6329(_r io.Reader, out io.Writer) {
	var n, m, v, w, op, x, y, ans int
	in := bufio.NewReader(_r)
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	deleted := make([]bool, len(g))
	size := make([]int, len(g))
	var findCentroid func(int, int, int) (int, int, int)
	findCentroid = func(v, fa, compSize int) (minSize, ct, faCt int) {
		minSize = math.MaxInt
		maxSubSize := 0
		size[v] = 1
		for _, w := range g[v] {
			if w != fa && !deleted[w] {
				if minSizeW, ctW, faCtW := findCentroid(w, v, compSize); minSizeW < minSize {
					minSize, ct, faCt = minSizeW, ctW, faCtW
				}
				maxSubSize = max(maxSubSize, size[w])
				size[v] += size[w]
			}
		}
		maxSubSize = max(maxSubSize, compSize-size[v])
		if maxSubSize < minSize {
			minSize, ct, faCt = maxSubSize, v, fa
		}
		return
	}

	type disInfo struct{ ct, sonI, ctDis int }
	paCts := make([][]disInfo, len(g))
	mergeSonInfo := make([]fenwick6329, len(g))
	sonInfo := make([][]fenwick6329, len(g))

	var dfs func(int, int, int)
	dfs = func(v, fa, compSize int) {
		_, ct, faCt := findCentroid(v, fa, compSize)

		totA := make(fenwick6329, compSize+1)
		sonInfo[ct] = make([]fenwick6329, len(g[ct]))
		for i, w := range g[ct] {
			if deleted[w] {
				continue
			}
			var sizeW int
			if w != faCt {
				sizeW = size[w]
			} else {
				sizeW = compSize - size[ct]
			}
			sumA := make(fenwick6329, sizeW+1)
			var f func(int, int, int)
			f = func(v, fa, d int) {
				paCts[v] = append(paCts[v], disInfo{ct, i, d})
				sumA[d] += a[v]
				totA[d] += a[v]
				for _, w := range g[v] {
					if w != fa && !deleted[w] {
						f(w, v, d+1)
					}
				}
			}
			f(w, ct, 1)
			for i := 1; i <= sizeW; i++ {
				if j := i + i&-i; j <= sizeW {
					sumA[j] += sumA[i]
				}
			}
			sonInfo[ct][i] = sumA
		}
		for i := 1; i <= compSize; i++ {
			if j := i + i&-i; j <= compSize {
				totA[j] += totA[i]
			}
		}
		mergeSonInfo[ct] = totA

		deleted[ct] = true
		for _, w := range g[ct] {
			if !deleted[w] {
				if w != faCt {
					dfs(w, ct, size[w])
				} else {
					dfs(w, ct, compSize-size[ct])
				}
			}
		}
	}
	dfs(0, -1, len(g))

	for ; m > 0; m-- {
		Fscan(in, &op, &x, &y)
		x ^= ans
		y ^= ans
		x--
		if op == 1 {
			delta := y - a[x]
			a[x] = y
			for _, p := range paCts[x] {
				mergeSonInfo[p.ct].update(p.ctDis, delta)
				sonInfo[p.ct][p.sonI].update(p.ctDis, delta)
			}
		} else {
			ans = a[x] + mergeSonInfo[x].pre(y)
			for _, p := range paCts[x] {
				if p.ctDis > y {
					continue
				}
				ans += a[p.ct] + mergeSonInfo[p.ct].pre(y-p.ctDis) - sonInfo[p.ct][p.sonI].pre(y-p.ctDis)
			}
			Fprintln(out, ans)
		}
	}
}

//func main() { p6329(os.Stdin, os.Stdout) }
