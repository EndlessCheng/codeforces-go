本题可以用线段树做，时空复杂度均与字符集的大小（本题为 $26$）无关，且这种写法可以支持查询 $s$ 的任意子串的最长重复子串的长度（本题查询的是整个 $s$）。

类似求动态最大子段和（[洛谷 P4513 小白逛公园](https://www.luogu.com.cn/problem/P4513)），线段树的每个节点维护对应区间的：

- 前缀最长连续字符个数 $\textit{pre}$；
- 后缀最长连续字符个数 $\textit{suf}$；
- 该区间最长连续字符个数 $\textit{max}$。

合并两个子区间时，如果前一个区间（记作 $a$）的末尾字符等于后一个区间（记作 $b$）的第一个字符，则可以合并这两个区间：

- 如果 $a$ 的 $\textit{suf}$ 等于 $a$ 的长度，那么就可以更新合并后的区间的 $\textit{pre}$ 值；
- 如果 $b$ 的 $\textit{pre}$ 等于 $b$ 的长度，那么就可以更新合并后的区间的 $\textit{suf}$ 值；
- 如果上面两个不成立，那么 $\textit{a.suf} + \textit{b.pre}$ 可以考虑成为合并后的区间的 $\textit{max}$。

具体见代码实现，大部分为线段树模板，主要逻辑是 `maintain` 的写法。

```py [sol-Python3]
class Solution:
    def longestRepeating(self, s: str, queryCharacters: str, queryIndices: List[int]) -> List[int]:
        s = list(s)
        n = len(s)
        pre = [0] * (4 * n)
        suf = [0] * (4 * n)
        mx = [0] * (4 * n)

        def maintain(o: int, l: int, r: int) -> None:
            pre[o] = pre[o * 2]
            suf[o] = suf[o * 2 + 1]
            mx[o] = max(mx[o * 2], mx[o * 2 + 1])
            m = (l + r) // 2
            if s[m - 1] == s[m]:  # 中间字符相同，可以合并
                if suf[o * 2] == m - l + 1:
                    pre[o] += pre[o * 2 + 1]
                if pre[o * 2 + 1] == r - m:
                    suf[o] += suf[o * 2]
                mx[o] = max(mx[o], suf[o * 2] + pre[o * 2 + 1])

        def build(o: int, l: int, r: int) -> None:
            if l == r:
                pre[o] = suf[o] = mx[o] = 1
                return
            m = (l + r) // 2
            build(o * 2, l, m)
            build(o * 2 + 1, m + 1, r)
            maintain(o, l, r)

        def update(o: int, l: int, r: int, i: int) -> None:
            if l == r: return
            m = (l + r) // 2
            if i <= m:
                update(o * 2, l, m, i)
            else:
                update(o * 2 + 1, m + 1, r, i)
            maintain(o, l, r)

        build(1, 1, n)
        ans = []
        for c, i in zip(queryCharacters, queryIndices):
            s[i] = c
            update(1, 1, n, i + 1)
            ans.append(mx[1])
        return ans
```

```java [sol-Java]
class Solution {
    char[] s;
    int[] pre, suf, max;

    void maintain(int o, int l, int r) {
        pre[o] = pre[o << 1];
        suf[o] = suf[o << 1 | 1];
        max[o] = Math.max(max[o << 1], max[o << 1 | 1]);
        var m = (l + r) >> 1;
        if (s[m - 1] == s[m]) { // 中间字符相同，可以合并
            if (suf[o << 1] == m - l + 1) pre[o] += pre[o << 1 | 1];
            if (pre[o << 1 | 1] == r - m) suf[o] += suf[o << 1];
            max[o] = Math.max(max[o], suf[o << 1] + pre[o << 1 | 1]);
        }
    }

    void build(int o, int l, int r) {
        if (l == r) {
            pre[o] = suf[o] = max[o] = 1;
            return;
        }
        var m = (l + r) / 2;
        build(o << 1, l, m);
        build(o << 1 | 1, m + 1, r);
        maintain(o, l, r);
    }

    void update(int o, int l, int r, int i) {
        if (l == r) return;
        var m = (l + r) / 2;
        if (i <= m) update(o << 1, l, m, i);
        else update(o << 1 | 1, m + 1, r, i);
        maintain(o, l, r);
    }

    public int[] longestRepeating(String s, String queryCharacters, int[] queryIndices) {
        this.s = s.toCharArray();
        int n = this.s.length, m = queryIndices.length;
        pre = new int[n << 2];
        suf = new int[n << 2];
        max = new int[n << 2];
        build(1, 1, n);
        var ans = new int[m];
        for (var i = 0; i < m; ++i) {
            this.s[queryIndices[i]] = queryCharacters.charAt(i);
            update(1, 1, n, queryIndices[i] + 1);
            ans[i] = max[1];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    string s;
    vector<int> pre, suf, max;

    void maintain(int o, int l, int r) {
        pre[o] = pre[o << 1];
        suf[o] = suf[o << 1 | 1];
        max[o] = std::max(max[o << 1], max[o << 1 | 1]);
        int m = (l + r) >> 1;
        if (s[m - 1] == s[m]) { // 中间字符相同，可以合并
            if (suf[o << 1] == m - l + 1) pre[o] += pre[o << 1 | 1];
            if (pre[o << 1 | 1] == r - m) suf[o] += suf[o << 1];
            max[o] = std::max(max[o], suf[o << 1] + pre[o << 1 | 1]);
        }
    }

    void build(int o, int l, int r) {
        if (l == r) {
            pre[o] = suf[o] = max[o] = 1;
            return;
        }
        int m = (l + r) / 2;
        build(o << 1, l, m);
        build(o << 1 | 1, m + 1, r);
        maintain(o, l, r);
    }

    void update(int o, int l, int r, int i) {
        if (l == r) return;
        int m = (l + r) / 2;
        if (i <= m) update(o << 1, l, m, i);
        else update(o << 1 | 1, m + 1, r, i);
        maintain(o, l, r);
    }

public:
    vector<int> longestRepeating(string &s, string &queryCharacters, vector<int> &queryIndices) {
        this->s = s;
        int n = s.length(), m = queryIndices.size();
        pre.resize(n << 2);
        suf.resize(n << 2);
        max.resize(n << 2);
        build(1, 1, n);
        vector<int> ans(m);
        for (int i = 0; i < m; ++i) {
            this->s[queryIndices[i]] = queryCharacters[i];
            update(1, 1, n, queryIndices[i] + 1);
            ans[i] = max[1];
        }
        return ans;
    }
};
```

```go [sol-Go]
var s []byte

type seg []struct{ l, r, pre, suf, max int }

func (t seg) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].pre = lo.pre
	t[o].suf = ro.suf
	t[o].max = max(lo.max, ro.max)
	if s[lo.r-1] == s[lo.r] { // 中间字符相同，可以合并
		if lo.suf == lo.r-lo.l+1 {
			t[o].pre += ro.pre
		}
		if ro.pre == ro.r-ro.l+1 {
			t[o].suf += lo.suf
		}
		t[o].max = max(t[o].max, lo.suf+ro.pre)
	}
}

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].pre = 1
		t[o].suf = 1
		t[o].max = 1
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, i int) {
	if t[o].l == t[o].r {
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i)
	} else {
		t.update(o<<1|1, i)
	}
	t.maintain(o)
}

func longestRepeating(S, queryCharacters string, queryIndices []int) []int {
	s = []byte(S)
	n := len(s)
	t := make(seg, n*4)
	t.build(1, 1, n)
	ans := make([]int, len(queryIndices))
	for i, index := range queryIndices {
		s[index] = queryCharacters[i]
		t.update(1, index+1)
		ans[i] = t[1].max
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+k\log n)$，其中 $n$ 是 $s$ 的长度，$k$ 是 $\textit{queryCharacters}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
