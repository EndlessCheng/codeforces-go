下午 2 点在 B 站直播讲周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

#### 前言

在求解「上升子序列」问题时，一般有两种优化方法：

1. 单调栈 + 二分优化；
2. 线段树、平衡树等数据结构优化。

这两种做法都可以用 $O(n\log n)$ 的时间解决 [300. 最长递增子序列](https://leetcode.cn/problems/longest-increasing-subsequence/)。

---

对于本题，由于额外有一个差值不超过 $k$ 的约束，用线段树更好处理。

具体来说，定义 $f[i][j]$ 表示 $\textit{nums}$ 的前 i 个元素中，以 $j$ 结尾的满足题目两个条件的最长子序列的长度。

我们可以从 $f[i-1][j']$ 转移过来，这里 $j-k\le j'<j$，取最大值，得

$$
f[i][j] = 1 + \max_{j'=j-k}^{j-1} f[i-1][j']
$$

上式有一个「区间求最大值」的过程，这非常适合用**线段树**计算，且由于 $f[i]$ 只会从 $f[i-1]$ 转移过来，我们可以把 $f$ 的第一个维度优化掉。这样我们可以**用线段树表示整个 $f$ 数组**，在上面查询和更新。

最后答案为 $\max(f[n-1])$，对应到线段树上就是根节点的值。

```py [sol1-Python3]
class Solution:
    def lengthOfLIS(self, nums: List[int], k: int) -> int:
        u = max(nums)
        mx = [0] * (4 * u)

        def modify(o: int, l: int, r: int, i: int, val: int) -> None:
            if l == r:
                mx[o] = val
                return
            m = (l + r) // 2
            if i <= m:
                modify(o * 2, l, m, i, val)
            else:
                modify(o * 2 + 1, m + 1, r, i, val)
            mx[o] = max(mx[o * 2], mx[o * 2 + 1])

        # 返回区间 [l,r] 内的最大值
        def query(o: int, l: int, r: int, L: int, R: int) -> int:  # L 和 R 在整个递归过程中均不变，将其大写，视作常量
            if L <= l and r <= R: return mx[o]
            res = 0
            m = (l + r) // 2
            if L <= m: res = query(o * 2, l, m, L, R)
            if R > m: res = max(res, query(o * 2 + 1, m + 1, r, L, R))
            return res

        for x in nums:
            if x == 1:
                modify(1, 1, u, 1, 1)
            else:
                res = 1 + query(1, 1, u, max(x - k, 1), x - 1)
                modify(1, 1, u, x, res)
        return mx[1]
```

```java [sol1-Java]
class Solution {
    int[] max;

    public int lengthOfLIS(int[] nums, int k) {
        var u = 0;
        for (var x : nums) u = Math.max(u, x);
        max = new int[u * 4];
        for (var x : nums) {
            if (x == 1) modify(1, 1, u, 1, 1);
            else {
                var res = 1 + query(1, 1, u, Math.max(x - k, 1), x - 1);
                modify(1, 1, u, x, res);
            }
        }
        return max[1];
    }

    private void modify(int o, int l, int r, int idx, int val) {
        if (l == r) {
            max[o] = val;
            return;
        }
        var m = (l + r) / 2;
        if (idx <= m) modify(o * 2, l, m, idx, val);
        else modify(o * 2 + 1, m + 1, r, idx, val);
        max[o] = Math.max(max[o * 2], max[o * 2 + 1]);
    }

    // 返回区间 [L,R] 内的最大值
    private int query(int o, int l, int r, int L, int R) { // L 和 R 在整个递归过程中均不变，将其大写，视作常量
        if (L <= l && r <= R) return max[o];
        var res = 0;
        var m = (l + r) / 2;
        if (L <= m) res = query(o * 2, l, m, L, R);
        if (R > m) res = Math.max(res, query(o * 2 + 1, m + 1, r, L, R));
        return res;
    }
}
```

```cpp [sol1-C++]
class Solution {
    vector<int> max;

    void modify(int o, int l, int r, int i, int val) {
        if (l == r) {
            max[o] = val;
            return;
        }
        int m = (l + r) / 2;
        if (i <= m) modify(o * 2, l, m, i, val);
        else modify(o * 2 + 1, m + 1, r, i, val);
        max[o] = std::max(max[o * 2], max[o * 2 + 1]);
    }

    // 返回区间 [L,R] 内的最大值
    int query(int o, int l, int r, int L, int R) { // L 和 R 在整个递归过程中均不变，将其大写，视作常量
        if (L <= l && r <= R) return max[o];
        int res = 0;
        int m = (l + r) / 2;
        if (L <= m) res = query(o * 2, l, m, L, R);
        if (R > m) res = std::max(res, query(o * 2 + 1, m + 1, r, L, R));
        return res;
    }

public:
    int lengthOfLIS(vector<int> &nums, int k) {
        int u = *max_element(nums.begin(), nums.end());
        max.resize(u * 4);
        for (int x: nums) {
            if (x == 1) modify(1, 1, u, 1, 1);
            else {
                int res = 1 + query(1, 1, u, std::max(x - k, 1), x - 1);
                modify(1, 1, u, x, res);
            }
        }
        return max[1];
    }
};
```

```go [sol1-Go]
type seg []struct{ l, r, max int }

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) modify(o, i, val int) {
	if t[o].l == t[o].r {
		t[o].max = val
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.modify(o<<1, i, val)
	} else {
		t.modify(o<<1|1, i, val)
	}
	t[o].max = max(t[o<<1].max, t[o<<1|1].max)
}

// 返回区间 [l,r] 内的最大值
func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].max
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return max(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func lengthOfLIS(nums []int, k int) int {
	mx := 0
	for _, x := range nums {
		mx = max(mx, x)
	}
	t := make(seg, mx*4)
	t.build(1, 1, mx)
	for _, x := range nums {
		if x == 1 {
			t.modify(1, 1, 1)
		} else {
			t.modify(1, x, 1+t.query(1, max(x-k, 1), x-1))
		}
	}
	return t[1].max
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$O(n\logU)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$O(U)$。
