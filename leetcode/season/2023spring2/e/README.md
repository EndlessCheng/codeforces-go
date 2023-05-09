晚上 8:30[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

先简单记录一下思路，直播结束后继续更新题解和其它语言。

1. 拆位。
2. 想象成 $0$ 或者 $1$ 从左到右「穿过」数组 $\textit{arr}$，用线段树维护穿过某段区间后，这个比特如何变化。 
3. 线段树叶子：$0$ 和 $1$ 穿过（NAND）$0$ 都变成 $1$；$0$ 和 $1$ 穿过 $1$ 分别变成 $1$ 和 $0$。
4. 线段树合并区间：先变成穿过左儿子的结果，然后把这个结果穿过右儿子，得到穿过这段区间的结果。
5. $\textit{type}=0$ 就是线段树的单点修改。
6. $\textit{type}=1$ 可以拆位，对于每个比特位上的数字，分类讨论：
   1. 如果穿过整个数组，数字没变，那么 $n$ 次后仍然没变。
   2. 否则，如果 $x=1$，那么变成穿过一次数组后的结果。
   3. 否则，如果穿过两次和穿过一次的结果一样，那么 $n$ 次后的结果就是穿过一次的结果。
   4. 否则，穿过两次回到原数字，那么分 $x$ 的奇偶性讨论，奇变偶不变。

```go [sol1-Go]
var k int

type seg []struct {
	l, r int
	to   [2]int
}

func (t seg) set(o, val int) {
	t[o].to[1] = t[o].to[0] ^ val
}

func (t seg) maintain(o int) {
	a, b, c := t[o<<1].to, t[o<<1|1].to, [2]int{}
	for i := 0; i < k; i++ {
		c[0] |= b[a[0]>>i&1] >> i & 1 << i
		c[1] |= b[a[1]>>i&1] >> i & 1 << i
	}
	t[o].to = c
}

func (t seg) build(a []int, k, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].to[0] = 1<<k - 1
	if l == r {
		t.set(o, a[l-1])
		return
	}
	m := (l + r) >> 1
	t.build(a, k, o<<1, l, m)
	t.build(a, k, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, i, val int) {
	if t[o].l == t[o].r {
		t.set(o, val)
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

func getNandResult(K int, arr []int, operations [][]int) (ans int) {
	k = K
	t := make(seg, len(arr)*4)
	t.build(arr, k, 1, 1, len(arr))
	for _, op := range operations {
		if op[0] == 0 {
			t.update(1, op[1]+1, op[2])
			continue
		}
		to := t[1].to
		x, y := op[1], op[2]
		for i := 0; i < k; i++ {
			var res int
			y := y >> i & 1
			y1 := to[y] >> i & 1 // 穿过 arr 一次
			if y1 == y { // 不变
				res = y
			} else if x == 1 || to[y1]>>i&1 == y1 {
				// 只穿过一次，或者穿过两次和穿过一次相同
				res = y1
			} else {
				res = y ^ x%2 // 奇变偶不变
			}
			ans ^= res << i
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(k(n+q\log n))$，其中 $n$ 为 $\textit{arr}$ 的长度，$q$ 为 $\textit{operations}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
