思路类似 [347. 前 K 个高频元素](https://leetcode.cn/problems/top-k-frequent-elements/)。

维护前 $k$ 小元素，可以用**最大堆**。

遍历 $\textit{queries}$，计算点 $(x,y)$ 到原点的曼哈顿距离 $d=|x|+|y|$。

把 $d$ 入堆，如果堆大小超过 $k$，就弹出堆顶（最大的元素）。

当堆的大小等于 $k$ 时，堆顶就是第 $k$ 小的距离。

具体请看 [视频讲解](https://www.bilibili.com/video/BV142Hae7E5y/) 第二题，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def resultsArray(self, queries: List[List[int]], k: int) -> List[int]:
        ans = [-1] * len(queries)
        h = []
        for i, (x, y) in enumerate(queries):
            heappush(h, -abs(x) - abs(y))  # 加负号变成最大堆
            if len(h) > k:
                heappop(h)
            if len(h) == k:
                ans[i] = -h[0]
        return ans
```

```java [sol-Java]
class Solution {
    public int[] resultsArray(int[][] queries, int k) {
        PriorityQueue<Integer> pq = new PriorityQueue<>((a, b) -> b - a);
        int[] ans = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            pq.offer(Math.abs(queries[i][0]) + Math.abs(queries[i][1]));
            if (pq.size() > k) {
                pq.poll();
            }
            ans[i] = pq.size() == k ? pq.peek() : -1;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> resultsArray(vector<vector<int>>& queries, int k) {
        vector<int> ans(queries.size(), -1);
        priority_queue<int> pq;
        for (int i = 0; i < queries.size(); i++) {
            pq.push(abs(queries[i][0]) + abs(queries[i][1]));
            if (pq.size() > k) {
                pq.pop();
            }
            if (pq.size() == k) {
                ans[i] = pq.top();
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func resultsArray(queries [][]int, k int) []int {
	ans := make([]int, len(queries))
	h := hp{}
	for i, q := range queries {
		heap.Push(&h, abs(q[0])+abs(q[1]))
		if h.Len() > k {
			heap.Pop(&h)
		}
		if h.Len() < k {
			ans[i] = -1
		} else {
			ans[i] = h.IntSlice[0]
		}
	}
	return ans
}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\log k)$，其中 $m$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。返回值不计入。

## 优化

如果 $\textit{queries}$ 的长度 $m$ 小于 $k$，那么返回一个全为 $-1$ 的数组。

否则，答案的前 $k-1$ 项都是 $-1$。

先把前 $k$ 项入堆。

对于后面的数，如果其大于等于堆顶，则不操作堆，否则替换堆顶。

下面代码仅提供 Python 和 Go，这两个语言可以直接修改堆顶。

```py [sol-Python3]
class Solution:
    def resultsArray(self, queries: List[List[int]], k: int) -> List[int]:
        m = len(queries)
        ans = [-1] * m
        if m < k:
            return ans

        h = [-abs(x) - abs(y) for x, y in queries[:k]]
        heapify(h)
        ans[k - 1] = -h[0]

        for i in range(k, m):
            x, y = queries[i]
            d = -abs(x) - abs(y)
            if d > h[0]:
                heapreplace(h, d)
            ans[i] = -h[0]
        return ans
```

```go [sol-Go]
func resultsArray(queries [][]int, k int) []int {
	m := len(queries)
	ans := make([]int, m)
	if m < k {
		for i := range ans {
			ans[i] = -1
		}
		return ans
	}

	h := hp{make([]int, k)}
	for i, q := range queries[:k] {
		h.IntSlice[i] = abs(q[0]) + abs(q[1])
		ans[i] = -1
	}
	heap.Init(&h)
	ans[k-1] = h.IntSlice[0]

	for i := k; i < m; i++ {
		q := queries[i]
		d := abs(q[0]) + abs(q[1])
		if d < h.IntSlice[0] {
			h.IntSlice[0] = d
			heap.Fix(&h, 0)
		}
		ans[i] = h.IntSlice[0]
	}
	return ans
}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(k + (m-k)\log k)$，其中 $m$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。返回值不计入。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
