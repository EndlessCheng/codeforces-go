示例 2 的 $\textit{nums}=[4,2,5,1]$，我们把所有子数组的值算出来，得到一个矩阵 $M$，其中 $M_{l,r}$ 表示子数组 $[l,r]$ 的值。规定 $l>r$ 时值为 $0$。

$$
M = \begin{bmatrix}
0 & 2 & 3 & 4    \\
0 & 0 & 3 & 4    \\
0 & 0 & 0 & 4    \\
0 & 0 & 0 & 0    \\
\end{bmatrix}
$$

当左端点固定时，右端点越大，子数组的最小值越小，最大值越大，所以子数组的值也就越大。

所以矩阵的每一行都是**递增**的。问题相当于：

- 合并 $n$ 个递增列表，计算前 $k$ 大元素之和。

根据 [23. 合并 K 个升序链表](https://leetcode.cn/problems/merge-k-sorted-lists/) 的 [做法](https://leetcode.cn/problems/merge-k-sorted-lists/solutions/2384305/liang-chong-fang-fa-zui-xiao-dui-fen-zhi-zbzx/)：

1. 把矩阵每一行的最后一个数 $M_{l,n-1}$ 加到最大堆中。
2. 循环 $k$ 次。
3. 每次循环，弹出堆顶，把堆顶 $M_{l,r}$ 加入答案，然后把左边元素 $M_{l,r-1}$ 入堆。

我们不能直接把整个 $M$ 算出来（太慢了），而是在入堆的时候计算，这需要一个数据结构，支持查询区间最小值和区间最大值。可以用线段树，或者 ST 表。下面用的 ST 表。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
# 手写 min max 更快
min = lambda a, b: b if b < a else a
max = lambda a, b: b if b > a else a

def op(a: Tuple[int, int], b: Tuple[int, int]) -> Tuple[int, int]:
    return min(a[0], b[0]), max(a[1], b[1])

class ST:
    def __init__(self, a: List[int]):
        n = len(a)
        sz = n.bit_length()
        st = [[None] * sz for _ in range(n)]
        for i, x in enumerate(a):
            st[i][0] = (x, x)
        for j in range(1, sz):
            for i in range(n - (1 << j) + 1):
                st[i][j] = op(st[i][j - 1], st[i + (1 << (j - 1))][j - 1])
        self.st = st

    # [l, r) 左闭右开
    def query(self, l: int, r: int) -> int:
        k = (r - l).bit_length() - 1
        mn, mx = op(self.st[l][k], self.st[r - (1 << k)][k])
        return mx - mn

class Solution:
    def maxTotalValue(self, nums: List[int], k: int) -> int:
        n = len(nums)
        st = ST(nums)
        # 取负号变成最大堆
        h = [(-st.query(i, n), i, n) for i in range(n)]
        heapify(h)

        ans = 0
        for _ in range(k):  # 题目保证 k 不超过所有子数组的个数
            d, l, r = heappop(h)
            ans -= d
            if l < r - 1:
                heappush(h, (-st.query(l, r - 1), l, r - 1))
        return ans
```

```java [sol-Java]
class Solution {
    private int[] op(int[] a, int[] b) {
        return new int[]{Math.min(a[0], b[0]), Math.max(a[1], b[1])};
    }

    private int[][][] newSt(int[] a) {
        int n = a.length;
        int sz = 32 - Integer.numberOfLeadingZeros(n);
        int[][][] st = new int[n][sz][2];
        for (int i = 0; i < n; i++) {
            st[i][0][0] = a[i];
            st[i][0][1] = a[i];
        }
        for (int j = 1; j < sz; j++) {
            for (int i = 0; i + (1 << j) <= n; i++) {
                st[i][j] = op(st[i][j - 1], st[i + (1 << (j - 1))][j - 1]);
            }
        }
        return st;
    }

    // [l,r) 左闭右开
    private int query(int[][][] st, int l, int r) {
        int k = 31 - Integer.numberOfLeadingZeros(r - l);
        int[] t = op(st[l][k], st[r - (1 << k)][k]);
        return t[1] - t[0];
    }

    public long maxTotalValue(int[] nums, int k) {
        int n = nums.length;
        int[][][] st = newSt(nums);

        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> b[0] - a[0]);
        for (int i = 0; i < n; i++) {
            pq.add(new int[]{query(st, i, n), i, n});
        }

        long ans = 0;
        while (k-- > 0) {
            int[] top = pq.poll();
            ans += top[0];
            int l = top[1], r = top[2];
            if (l < r - 1) {
                pq.add(new int[]{query(st, l, r - 1), l, r - 1});
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    pair<int, int> op(const pair<int, int>& a, const pair<int, int>& b) {
        return {min(a.first, b.first), max(a.second, b.second)};
    }

    vector<vector<pair<int, int>>> new_st(vector<int>& a) {
        size_t n = a.size();
        int sz = bit_width(n);
        vector st(n, vector<pair<int, int>>(sz));
        for (int i = 0; i < n; i++) {
            st[i][0] = {a[i], a[i]};
        }
        for (int j = 1; j < sz; j++) {
            for (int i = 0; i + (1 << j) <= n; i++) {
                st[i][j] = op(st[i][j - 1], st[i + (1 << (j - 1))][j - 1]);
            }
        }
        return st;
    }

    // [l,r) 左闭右开
    int query(const vector<vector<pair<int, int>>>& st, int l, int r) {
        int k = bit_width((uint32_t) r - l) - 1;
        auto [mn, mx] = op(st[l][k], st[r - (1 << k)][k]);
        return mx - mn;
    }

public:
    long long maxTotalValue(vector<int>& nums, int k) {
        int n = nums.size();
        auto st = new_st(nums);

        priority_queue<tuple<int, int, int>> pq;
        for (int i = 0; i < n; i++) {
            pq.emplace(query(st, i, n), i, n);
        }

        long long ans = 0;
        while (k--) {
            auto [d, l, r] = pq.top();
            pq.pop();
            ans += d;
            if (l < r - 1) {
                pq.emplace(query(st, l, r - 1), l, r - 1);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
type pair struct{ min, max int }

func op(a, b pair) pair {
	return pair{min(a.min, b.min), max(a.max, b.max)}
}

type ST [][]pair

func newST(a []int) ST {
	n := len(a)
	sz := bits.Len(uint(n))
	st := make(ST, n)
	for i, x := range a {
		st[i] = make([]pair, sz)
		st[i][0] = pair{x, x}
	}
	for j := 1; j < sz; j++ {
		for i := range n - 1<<j + 1 {
			st[i][j] = op(st[i][j-1], st[i+1<<(j-1)][j-1])
		}
	}
	return st
}

// [l,r) 左闭右开
func (st ST) query(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	p := op(st[l][k], st[r-1<<k][k])
	return p.max - p.min
}

func maxTotalValue(nums []int, k int) (ans int64) {
	n := len(nums)
	st := newST(nums)
	h := make(hp, n)
	for i := range h {
		h[i] = tuple{st.query(i, n), i, n}
	}
	heap.Init(&h)

	for range k { // 题目保证 k 不超过所有子数组的个数
		t := heap.Pop(&h).(tuple)
		ans += int64(t.d)
		if t.l < t.r-1 {
			heap.Push(&h, tuple{st.query(t.l, t.r-1), t.l, t.r - 1})
		}
	}
	return
}

type tuple struct{ d, l, r int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].d > h[j].d }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+k)\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n\log n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
