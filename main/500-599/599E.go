package main

import (
	. "fmt"
	"io"
	"math/bits"
)

func cf599E(in io.Reader, out io.Writer) {
	var n, m, q, x, y, lca int
	Fscan(in, &n, &m, &q)
	type tuple struct{ tp, x, y, lca int }
	req := make([]tuple, 0, m+q)
	for ; m > 0; m-- {
		Fscan(in, &x, &y)
		req = append(req, tuple{0, x - 1, y - 1, 0})
	}
	for ; q > 0; q-- {
		Fscan(in, &x, &y, &lca)
		if x == y && x != lca {
			Fprint(out, 0)
			return
		}
		req = append(req, tuple{1, x - 1, y - 1, lca - 1})
	}

	check := func(mask, root, subMask, child int) bool {
		for _, r := range req {
			if mask>>r.x&1 == 0 || mask>>r.y&1 == 0 ||
				subMask>>r.x&1 == subMask>>r.y&1 { // 不在考虑范围内 or 已经考虑过了
				continue
			}
			if r.tp == 0 { // 树边
				if r.y == root { // 保证 a 上 b 下
					r.x, r.y = r.y, r.x
				}
				if r.x != root || r.y != child {
					return false
				}
			} else if r.lca != root { // a 和 b 在不同的子树/连通块中，所以 rt 必须是 LCA
				return false
			}
		}
		return true
	}

	f := make([][13]int, 1<<n)
	for i := 0; i < n; i++ {
		f[1<<i][i] = 1
	}
	for mask := 2; mask < 1<<n; mask++ {
		if mask&(mask-1) == 0 {
			continue
		}
		for t := uint(mask); t > 0; t &= t - 1 {
			root := bits.TrailingZeros(t) // mask 子树的根
			m := mask ^ 1<<root
			lb := m & -m
			m ^= lb
			for sub, ok := m, true; ok; ok = sub != m {
				subMask := sub | lb
				for t2 := uint(subMask); t2 > 0; t2 &= t2 - 1 {
					child := bits.TrailingZeros(t2) // subMask 子树的根，同时也是 root 的儿子
					if check(mask, root, subMask, child) {
						f[mask][root] += f[mask^subMask][root] * f[subMask][child]
					}
				}
				sub = (sub - 1) & m
			}
		}
	}
	Fprint(out, f[1<<n-1][0])
}

//func main() { cf599E(os.Stdin, os.Stdout) }
