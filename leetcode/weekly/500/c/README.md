由于 $\textit{nums}$ 是严格递增的，对于 $i<j<k$，从 $i$ 直接跳到 $k$ 的代价，等于从 $i$ 跳到 $j$ 再跳到 $k$ 的代价之和。所以我们可以**一步步移动**。

> 用数学语言描述，就是 $\textit{nums}[k] - \textit{nums}[i] = (\textit{nums}[k] - \textit{nums}[j]) + (\textit{nums}[j] - \textit{nums}[i])$。

由于 $\textit{nums}$ 是严格递增的，所以使用方式一移动到相邻下标的代价至少是 $1$，那么贪心地，**能用方式二移动，就用方式二**。

此外，由于走回头路一定会包含往返的一段，这会花费多余的代价，所以我们不会走回头路。

计算从左到右一步步移动的代价的前缀和，以及从右到左一步步移动的代价的前缀和，即可 $\mathcal{O}(1)$ 回答询问。关于前缀和的原理，请看 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

[本题视频讲解](https://www.bilibili.com/video/BV1719oB4EWf/?t=7m34s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minCost(self, nums: list[int], queries: list[list[int]]) -> list[int]:
        n = len(nums)
        sum_l = [0] * n  # sum_l[i] 等于从 i 移动到 0 的代价和
        sum_r = [0] * n  # sum_r[i] 等于从 0 移动到 i 的代价和
        for i in range(1, n):
            # 往左走 i -> i-1
            if i < n - 1 and nums[i] - nums[i - 1] > nums[i + 1] - nums[i]:  # closest(i) = i+1
                cost = nums[i] - nums[i - 1]  # 只能用方式一往左走
            else:
                cost = 1
            sum_l[i] = sum_l[i - 1] + cost

            # 往右走 i-1 -> i
            if i > 1 and nums[i - 1] - nums[i - 2] <= nums[i] - nums[i - 1]:  # closest(i-1) = i-2
                cost = nums[i] - nums[i - 1]  # 只能用方式一往右走
            else:
                cost = 1
            sum_r[i] = sum_r[i - 1] + cost

        ans = [0] * len(queries)
        for i, q in enumerate(queries):
            l, r = q
            if l < r:
                # cost(0 -> r) - cost(0 -> l) = cost(l -> r)
                ans[i] = sum_r[r] - sum_r[l]
            else:
                # cost(l -> 0) - cost(r -> 0) = cost(l -> r)
                ans[i] = sum_l[l] - sum_l[r]
        return ans
```

```java [sol-Java]
class Solution {
    public int[] minCost(int[] nums, int[][] queries) {
        int n = nums.length;
        int[] sumL = new int[n]; // sumL[i] 等于从 i 移动到 0 的代价和
        int[] sumR = new int[n]; // sumR[i] 等于从 0 移动到 i 的代价和
        for (int i = 1, cost; i < n; i++) {
            // 往左走 i -> i-1
            if (i < n - 1 && nums[i] - nums[i - 1] > nums[i + 1] - nums[i]) { // closest(i) = i+1
                cost = nums[i] - nums[i - 1]; // 只能用方式一往左走
            } else {
                cost = 1;
            }
            sumL[i] = sumL[i - 1] + cost;

            // 往右走 i-1 -> i
            if (i > 1 && nums[i - 1] - nums[i - 2] <= nums[i] - nums[i - 1]) { // closest(i-1) = i-2
                cost = nums[i] - nums[i - 1]; // 只能用方式一往右走
            } else {
                cost = 1;
            }
            sumR[i] = sumR[i - 1] + cost;
        }

        int[] ans = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int l = queries[i][0];
            int r = queries[i][1];
            if (l < r) {
                // cost(0 -> r) - cost(0 -> l) = cost(l -> r)
                ans[i] = sumR[r] - sumR[l];
            } else {
                // cost(l -> 0) - cost(r -> 0) = cost(l -> r)
                ans[i] = sumL[l] - sumL[r];
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> minCost(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size();
        vector<int> sum_l(n); // sum_l[i] 等于从 i 移动到 0 的代价和
        vector<int> sum_r(n); // sum_r[i] 等于从 0 移动到 i 的代价和
        for (int i = 1, cost; i < n; i++) {
            // 往左走 i -> i-1
            if (i < n - 1 && nums[i] - nums[i - 1] > nums[i + 1] - nums[i]) { // closest(i) = i+1
                cost = nums[i] - nums[i - 1]; // 只能用方式一往左走
            } else {
                cost = 1;
            }
            sum_l[i] = sum_l[i - 1] + cost;

            // 往右走 i-1 -> i
            if (i > 1 && nums[i - 1] - nums[i - 2] <= nums[i] - nums[i - 1]) { // closest(i-1) = i-2
                cost = nums[i] - nums[i - 1]; // 只能用方式一往右走
            } else {
                cost = 1;
            }
            sum_r[i] = sum_r[i - 1] + cost;
        }

        vector<int> ans(queries.size());
        for (int i = 0; i < queries.size(); i++) {
            int l = queries[i][0], r = queries[i][1];
            if (l < r) {
                // cost(0 -> r) - cost(0 -> l) = cost(l -> r)
                ans[i] = sum_r[r] - sum_r[l];
            } else {
                // cost(l -> 0) - cost(r -> 0) = cost(l -> r)
                ans[i] = sum_l[l] - sum_l[r];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minCost(nums []int, queries [][]int) []int {
	n := len(nums)
	sumL := make([]int, n) // sumL[i] 等于从 i 移动到 0 的代价和
	sumR := make([]int, n) // sumR[i] 等于从 0 移动到 i 的代价和
	for i := 1; i < n; i++ {
		// 往左走 i -> i-1
		cost := 1
		if i < n-1 && nums[i]-nums[i-1] > nums[i+1]-nums[i] { // closest(i) = i+1
			cost = nums[i] - nums[i-1] // 只能用方式一往左走
		}
		sumL[i] = sumL[i-1] + cost

		// 往右走 i-1 -> i
		cost = 1
		if i > 1 && nums[i-1]-nums[i-2] <= nums[i]-nums[i-1] { // closest(i-1) = i-2
			cost = nums[i] - nums[i-1] // 只能用方式一往右走
		}
		sumR[i] = sumR[i-1] + cost
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		if l < r {
			// cost(0 -> r) - cost(0 -> l) = cost(l -> r)
			ans[i] = sumR[r] - sumR[l]
		} else {
			// cost(l -> 0) - cost(r -> 0) = cost(l -> r)
			ans[i] = sumL[l] - sumL[r]
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 专题训练

见下面数据结构题单的「**一、前缀和**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
