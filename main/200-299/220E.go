package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
type fenwick20 []int

func (f fenwick20) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

// [1,i] 的和
func (f fenwick20) sum(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func cf220E(in io.Reader, out io.Writer) {
	var n, k, ans int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	b := slices.Clone(a)
	slices.Sort(b)
	b = slices.Compact(b)
	m := len(b)

	// 计算不删除时的逆序对（直接从 k 中减掉）
	suf := make(fenwick20, m+1)
	for i := n - 1; i >= 0; i-- {
		a[i] = sort.SearchInts(b, a[i]) + 1 // 离散化
		k -= suf.sum(a[i] - 1)
		suf.update(a[i], 1)
	}

	pre := make(fenwick20, m+1)
	l := 0
	for r := 1; r < n; r++ {
		// 从后缀中删除 a[r-1]，撤销逆序对（a[r-1] 与 pre 和 suf 的逆序对）
		suf.update(a[r-1], -1)
		k += l - pre.sum(a[r-1]) + suf.sum(a[r-1]-1)
		for l < r {
			// 尝试往前缀添加 a[l]
			inv := l - pre.sum(a[l]) + suf.sum(a[l]-1)
			if inv > k { // 逆序对太多了，无法添加
				break
			}
			// 添加后，总逆序对个数 <= k，说明 (l,r) 满足要求
			k -= inv
			pre.update(a[l], 1)
			l++
		}
		// 右端点为 r 时，左端点可以是 0,1,...,l-1，一共 l 个
		ans += l
	}
	Fprint(out, ans)
}

//func main() { cf220E(bufio.NewReader(os.Stdin), os.Stdout) }
