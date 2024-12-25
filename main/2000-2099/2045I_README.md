这种整体求和题目，可以先从贡献法开始思考。

从左到右遍历 $a$，考虑 $a_i$ 的贡献：

- 如果 $a_i$ 是首次出现，那么：
  - $a_i$ 和一个不在 $a$ 中的数计算 $f(x,y)$，贡献是 $c$，其中 $c$ 是 $[1,m]$ 中的不在 $a$ 中的整数个数。
  - $a_i$ 和另一个在 $a$ 中且不等于 $a_i$ 的数计算 $f(x,y)$，这样的数有 $m-1-c$ 个。
      - 如果另一个数的首次出现在 $a_i$ 左边，那么把 $a_i$ 加到这个序列的末尾；
      - 如果另一个数的首次出现在 $a_i$ 右边，那么 $a_i$ 作为这个序列的第一个数。
      - 所以 $a_i$ 的贡献是 $m-1-c$。
  - 总的贡献是 $c + (m-1-c) = m-1$。
- 如果 $a_i$ 不是首次出现，设上一次出现的下标为 $\textit{pre}$：
  - 设下标在 $[\textit{pre}+1,i-1]$ 中的**不同的数**的个数有 $k$ 个，我们可以在这 $k$ 个序列的末尾都添加一个 $a_i$，所以 $a_i$ 的贡献是 $k$。

计算区间内不同的数的个数，方法同 [P1972. HH 的项链](https://www.luogu.com.cn/problem/P1972)，用树状数组维护每个数的最新下标，查询区间内下标的个数，即为区间内不同的数的个数。

```go
package main
import("bufio";."fmt";"os")

type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

// [1,i] 的元素和
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

// [l,r] 的元素和
func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m, x, ans int
	Fscan(in, &n, &m)
	pre := make([]int, m)
	tree := make(fenwick, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &x) // a[i]
		x--
		if pre[x] == 0 {
			ans += m - 1 // a[i] 的贡献
		} else {
			ans += tree.query(pre[x]+1, i) // a[i] 的贡献
			tree.update(pre[x], -1) // 维护 a[i] 的最新下标
		}
		tree.update(i, 1) // 维护 a[i] 的最新下标
		pre[x] = i
	}
	Print(ans)
}
```

**时间复杂度**：$\mathcal{O}(m + n\log n)$。如果把 $\textit{pre}$ 改成哈希表，则时间复杂度为 $\mathcal{O}(n\log n)$。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
