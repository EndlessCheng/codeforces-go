考虑枚举其中一台机器 $i$，那么另一台机器的价格必须严格小于 $\textit{budget} - \textit{costs}[i]$。

如果把机器按照价格从小到大**排序**，我们就可以在 $[0,i-1]$ 中**二分查找**最后一台价格小于 $\textit{budget} - \textit{costs}[i]$ 的机器 $j$。在 $[0,i-1]$ 中二分是为了避免选同一台机器。

关于二分查找的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

然而，最后一台机器 $j$ 的容量并不是最大的，我们要求的是 $[0,j]$ 中的最大容量。

这启发我们维护一个关于机器容量的**前缀最大值**数组 $\textit{preMax}$。

具体地，定义 $\textit{preMax}[j+1]$ 表示 $[0,j]$ 中的最大容量。$\textit{preMax}[0]=0$ 当作哨兵。

那么有

$$
\textit{preMax}[j+1] = \max(\textit{preMax}[j], \textit{capacity}[j])
$$

这样二分之后，$\textit{preMax}[j+1]$ 就是另一台机器的最大容量了。如果这台机器不存在，那么我们会取 $\textit{preMax}[0] = 0$，相当于不买另一台机器。

#### 答疑

**问**：对于机器 $A$，如果另一台要买的机器 $B$ 在 $A$ 的右边呢？我们会漏掉这种情况吗？

**答**：继续往后遍历，遍历到 $B$ 时，在左边二分找到 $A$。所以不会漏掉最优解。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

## 写法一：前缀最大值

```py [sol-Python3]
class Solution:
    def maxCapacity(self, costs: List[int], capacity: List[int], budget: int) -> int:
        # 把 costs[i] 和 capacity[i] 绑在一起排序
        a = [(cost, cap) for cost, cap in zip(costs, capacity) if cost < budget]  # 太贵的机器直接忽略
        a.sort(key=lambda p: p[0])

        pre_max = [0] * (len(a) + 1)
        ans = 0
        for i, (cost, cap) in enumerate(a):
            # 二分第一台价格 >= budget-cost 的机器，下标减一，就是最后一台价格 < budget-cost 的机器
            j = bisect_left(range(i), budget - cost, key=lambda j: a[j][0])
            # (j - 1) + 1 == j
            ans = max(ans, cap + pre_max[j])  # j=0 的情况对应单选一台机器
            pre_max[i + 1] = max(pre_max[i], cap)
        return ans
```

```java [sol-Java]
class Solution {
    public int maxCapacity(int[] costs, int[] capacity, int budget) {
        int n = costs.length;
        Integer[] idx = new Integer[n];
        for (int i = 0; i < n; i++) {
            idx[i] = i;
        }
        Arrays.sort(idx, (i, j) -> costs[i] - costs[j]);

        int[] preMax = new int[n + 1];
        int ans = 0;
        for (int k = 0; k < n && costs[idx[k]] < budget; k++) { // 太贵的机器直接忽略
            int i = idx[k];
            // 二分找到第一台价格 >= budget-costs[i] 的机器，下标减一，就是最后一台价格 < budget-costs[i] 的机器
            int j = lowerBound(idx, k, costs, budget - costs[i]);
            // (j - 1) + 1 == j
            ans = Math.max(ans, capacity[i] + preMax[j]); // j=0 的情况对应单选一台机器
            preMax[k + 1] = Math.max(preMax[k], capacity[i]);
        }
        return ans;
    }

    // 原理见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(Integer[] idx, int right, int[] costs, int target) {
        int left = -1;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            if (costs[idx[mid]] >= target) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxCapacity(vector<int>& costs, vector<int>& capacity, int budget) {
        // 把 costs[i] 和 capacity[i] 绑在一起排序
        int n = costs.size();
        vector<pair<int, int>> a;
        for (int i = 0; i < n; i++) {
            if (costs[i] < budget) { // 太贵的机器直接忽略
                a.emplace_back(costs[i], capacity[i]);
            }
        }
        ranges::sort(a, {}, &pair<int, int>::first);

        vector<int> pre_max(a.size() + 1);
        int ans = 0;
        for (int i = 0; i < a.size(); i++) {
            auto& [cost, cap] = a[i];
            // 二分第一台价格 >= budget-cost 的机器，下标减一，就是最后一台价格 < budget-cost 的机器
            int j = lower_bound(a.begin(), a.begin() + i, pair(budget - cost, 0)) - a.begin();
            ans = max(ans, cap + pre_max[j]); // j=0 的情况对应单选一台机器
            pre_max[i + 1] = max(pre_max[i], cap);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxCapacity(costs []int, capacity []int, budget int) (ans int) {
	// 把 costs[i] 和 capacity[i] 绑在一起排序
	type pair struct{ cost, cap int }
	n := len(costs)
	a := make([]pair, 0, n)
	for i, cost := range costs {
		if cost < budget { // 太贵的机器直接忽略
			a = append(a, pair{cost, capacity[i]})
		}
	}
	slices.SortFunc(a, func(a, b pair) int { return a.cost - b.cost })

	preMax := make([]int, len(a)+1)
	for i, p := range a {
		// 二分第一台价格 >= budget-p.cost 的机器，下标减一，就是最后一台价格 < budget-p.cost 的机器
		j := sort.Search(i, func(j int) bool { return a[j].cost >= budget-p.cost })
		// (j - 1) + 1 == j
		ans = max(ans, p.cap+preMax[j]) // j=0 的情况对应单选一台机器
		preMax[i+1] = max(preMax[i], p.cap)
	}
	return
}
```

## 写法二：单调栈

对于两台机器 $A$ 和 $B$，如果机器 $B$ 又贵，容量又小，全方面不如机器 $A$，那么机器 $B$ 就是垃圾数据，直接忽略。

这启发我们在遍历的同时，用一个栈维护遍历过的机器，只有当新遍历到的机器的容量比栈顶大时，才入栈。注意价格已经从小到大排序了，无需比较。

```py [sol-Python3]
class Solution:
    def maxCapacity(self, costs: List[int], capacity: List[int], budget: int) -> int:
        a = [(cost, cap) for cost, cap in zip(costs, capacity) if cost < budget]
        a.sort(key=lambda p: p[0])

        st = [(0, 0)]  # 栈底加个哨兵
        ans = 0
        for cost, cap in a:
            j = bisect_left(st, (budget - cost,)) - 1
            ans = max(ans, cap + st[j][1])  # j=0 的情况对应单选一台机器
            if cap > st[-1][1]:
                st.append((cost, cap))
        return ans
```

```java [sol-Java]
class Solution {
    public int maxCapacity(int[] costs, int[] capacity, int budget) {
        int n = costs.length;
        Integer[] idx = new Integer[n];
        for (int i = 0; i < n; i++) {
            idx[i] = i;
        }
        Arrays.sort(idx, (i, j) -> costs[i] - costs[j]);

        List<int[]> st = new ArrayList<>();
        st.add(new int[]{0, 0}); // 栈底加个哨兵
        int ans = 0;
        for (int k = 0; k < n && costs[idx[k]] < budget; k++) { // 太贵的机器直接忽略
            int i = idx[k];
            int j = lowerBound(st, budget - costs[i]) - 1;
            ans = Math.max(ans, capacity[i] + st.get(j)[1]); // j=0 的情况对应单选一台机器
            if (capacity[i] > st.getLast()[1]) {
                st.add(new int[]{costs[i], capacity[i]});
            }
        }
        return ans;
    }

    // 原理见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(List<int[]> st, int target) {
        int left = -1;
        int right = st.size();
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            if (st.get(mid)[0] >= target) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxCapacity(vector<int>& costs, vector<int>& capacity, int budget) {
        // 把 costs[i] 和 capacity[i] 绑在一起排序
        int n = costs.size();
        vector<pair<int, int>> a;
        for (int i = 0; i < n; i++) {
            if (costs[i] < budget) { // 太贵的机器直接忽略
                a.emplace_back(costs[i], capacity[i]);
            }
        }
        ranges::sort(a, {}, &pair<int, int>::first);

        vector<pair<int, int>> st = {{0, 0}}; // 栈底加个哨兵
        int ans = 0;
        for (auto& [cost, cap] : a) {
            int j = ranges::lower_bound(st, pair(budget - cost, 0)) - st.begin() - 1;
            ans = max(ans, cap + st[j].second); // j=0 的情况对应单选一台机器
            if (cap > st.back().second) {
                st.emplace_back(cost, cap);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxCapacity(costs, capacity []int, budget int) (ans int) {
	type pair struct{ cost, cap int }
	n := len(costs)
	a := make([]pair, 0, n)
	for i, cost := range costs {
		if cost < budget {
			a = append(a, pair{cost, capacity[i]})
		}
	}
	slices.SortFunc(a, func(a, b pair) int { return a.cost - b.cost })

	st := []pair{{}} // 栈底加个哨兵
	for _, p := range a {
		j := sort.Search(len(st), func(j int) bool { return st[j].cost >= budget-p.cost }) - 1
		ans = max(ans, p.cap+st[j].cap) // j=0 的情况对应单选一台机器
		if p.cap > st[len(st)-1].cap {
			st = append(st, p)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{costs}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
