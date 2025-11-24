package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1616H(in io.Reader, out io.Writer) {
	type node struct {
		son [2]*node
		sz  int
	}
	size := func(o *node) int {
		if o == nil {
			return 0
		}
		return o.sz
	}
	root := &node{}

	var n, x, v int
	Fscan(in, &n, &x)
	for range n {
		Fscan(in, &v)
		o := root
		for i := 29; i >= 0; i-- {
			b := v >> i & 1
			if o.son[b] == nil {
				o.son[b] = &node{}
			}
			o = o.son[b]
			o.sz++
		}
	}

	const mod = 998244353
	pow2 := make([]int, n+1)
	pow2[0] = 1
	for i := 1; i <= n; i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}

	var dfs func(*node, *node, int) int
	dfs = func(p, q *node, d int) int {
		if p == nil || q == nil {
			return pow2[size(p)+size(q)]
		}
		if d < 0 {
			if p == q {
				return pow2[p.sz]
			}
			return pow2[p.sz+q.sz]
		}

		bit := x >> d & 1
		// 用数学归纳法可以证明，对于固定的 p.son[i]，第二个参数是固定的
		// 所以相当于在 dfs(p.son[i])，这等价于 dfs 一棵二叉树
		res1 := dfs(p.son[0], q.son[bit], d-1)
		res2 := dfs(p.son[1], q.son[bit^1], d-1)

		if bit > 0 {
			if p == q {
				return res1
			}
			return res1 * res2 % mod
		}

		if p == q {
			return (res1 + res2 - 1) % mod
		}
		// 注意我们算的是子序列（子集）的个数，不是数对的个数
		// bit == 0 时，res1 和 res2 不能有交，所以是加法而不是乘法
		// 考虑上面是 1，这里是 0 的情形，那么同一组子树内部可以随意选，这解释了末两行代码
		return (res1 + res2 - 1 +
			(pow2[size(p.son[0])]-1)*(pow2[size(p.son[1])]-1) +
			(pow2[size(q.son[0])]-1)*(pow2[size(q.son[1])]-1)) % mod
	}
	Fprint(out, (dfs(root, root, 29)+mod*2-1)%mod)
}

//func main() { debug.SetGCPercent(-1); cf1616H(bufio.NewReader(os.Stdin), os.Stdout) }
