本题可以用线段树来做，时空复杂度均和字符集的大小（本题为 $26$）无关，且这种写法可以支持查询 $s$ 任意子串的最长重复子串的长度。（本题查询的是整个 $s$）

做法类似求动态最大子段和（洛谷 P4513 小白逛公园 https://www.luogu.com.cn/problem/P4513 ），线段树的每个节点维护对应区间的：

- 前缀最长连续字符个数 $\textit{pre}$；
- 后缀最长连续字符个数 $\textit{suf}$；
- 该区间最长连续字符个数 $\textit{max}$。

合并两个子区间时，如果前一个区间（记作 $a$）的末尾字符等于后一个区间（记作 $b$）的第一个字符，则可以合并这两个区间：

- 如果 $a$ 的 $\textit{suf}$ 等于 $a$ 的长度，那么就可以更新合并后的区间的 $\textit{pre}$ 值；
- 如果 $b$ 的 $\textit{pre}$ 等于 $b$ 的长度，那么就可以更新合并后的区间的 $\textit{suf}$ 值；
- 如果上面两个不成立，那么 $\textit{a.suf} + \textit{b.pre}$ 可以考虑成为合并后的区间的 $\textit{max}$。

为了编码的方便，下面代码额外维护了区间长度和区间首尾字符。大部分均为线段树模板，主要逻辑是 `merge` 的写法。

```go
type data struct {
	pre, suf, max, size int
	lch, rch            byte
}

type seg []struct {
	l, r int
	data
}

func (t seg) set(o int, ch byte) {
	t[o].lch = ch
	t[o].rch = ch
}

func (t seg) merge(a, b data) data {
	d := data{a.pre, b.suf, max(a.max, b.max), a.size + b.size, a.lch, b.rch}
	if a.rch == b.lch { // 中间字符相同，可以合并
		if a.suf == a.size {
			d.pre += b.pre
		}
		if b.pre == b.size {
			d.suf += a.suf
		}
		if a.suf != a.size && b.pre != b.size {
			d.max = max(d.max, a.suf+b.pre)
		}
		d.max = max(d.max, max(d.pre, d.suf))
	}
	return d
}

func (t seg) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].data = t.merge(lo.data, ro.data)
}

func (t seg) build(s string, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].pre = 1
		t[o].suf = 1
		t[o].max = 1
		t[o].size = 1
		t.set(o, s[l-1])
		return
	}
	m := (l + r) >> 1
	t.build(s, o<<1, l, m)
	t.build(s, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, i int, ch byte) {
	if t[o].l == t[o].r {
		t.set(o, ch)
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, ch)
	} else {
		t.update(o<<1|1, i, ch)
	}
	t.maintain(o)
}

// 由于题目查询的是整个字符串的最长重复子串的长度，下面这块代码是多余的，为了题解的完整性而保留
func (t seg) query(o, l, r int) data {
	if l <= t[o].l && t[o].r <= r {
		return t[o].data
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	lv := t.query(o<<1, l, r)
	rv := t.query(o<<1|1, l, r)
	return t.merge(lv, rv)
}

func longestRepeating(s, queryCharacters string, queryIndices []int) []int {
	n := len(s)
	t := make(seg, n*4)
	t.build(s, 1, 1, n)
	ans := make([]int, len(queryIndices))
	for i, index := range queryIndices {
		t.update(1, index+1, queryCharacters[i])
		ans[i] = t[1].max
	}
	return ans
}

func max(a, b int) int { if b > a { return b }; return a}
```
