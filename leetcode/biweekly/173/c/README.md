## 核心思路

每个 $a[i]$ 的最大值会受到三个约束：

- $i$ 处的 $\textit{maxVal}$（如果有）。
- 来自左侧元素的约束。需要**从左到右**扫描一遍算出来。
- 来自右侧元素的约束。需要**从右到左**扫描一遍算出来。

## 具体思路

设 $i$ 处的值不得超过 $\textit{maxVal}[i]$。如果此处没有限制，那么 $\textit{maxVal}[i]=\infty$。

第一次扫描，从左到右。如果没有 $\textit{maxVal}[i]$ 的限制，那么拉满，即 $a[i+1] = a[i] + \textit{diff}[i]$ 是最优的。但这不能超过 $\textit{maxVal}[i+1]$，所以有

$$
a[i+1] = \min(a[i] + \textit{diff}[i], \textit{maxVal}[i+1])
$$

第一次扫描后，每个 $a[i]$ 都尽量拉满。但和 $\textit{maxVal}$ 计算 $\min$ 后，相邻元素可能不满足 $\textit{diff}$ 的约束，可能前一个数比较大，后一个数被 $\textit{maxVal}$ 突然拉低，导致两个相邻元素之差超过 $\textit{diff}$ 值。怎么办？拉低后，再反向修正一下。但先别急，遍历到最后一个元素再说。

注意到，对于最后一个数 $a[n-1]$，它的最大值只取决于左边的元素大小，以及 $\textit{maxVal}[n-1]$。第一次扫描已经把 $0$ 到 $n-2$ 的所有限制传播到了 $n-1$，所以第一次扫描结束时，$a[n-1]$ 就已经是它能达到的最大值了。

第二次扫描，从右到左。从已确定的值 $a[n-1]$ 出发：

- 用 $a[n-1] + \textit{diff}[n-2]$ 更新 $a[n-2]$ 的最小值。
- 用 $a[n-2] + \textit{diff}[n-3]$ 更新 $a[n-3]$ 的最小值。
- ……
- 用 $a[2] + \textit{diff}[1]$ 更新 $a[1]$ 的最小值。

注意 $a[0]=0$ 是固定值，无需修正。 

⚠**注意**：这里取的是 $\min$，只会把 $a[i]$ 变小，没有破坏在第一次扫描时建立的 $a[i+1]\le a[i] + \textit{diff}[i]$ 的约束。所以两次扫描后，每个 $a[i]$ 都满足 $\textit{diff}$ 的约束。我们得到了最终的 $a$。

最后返回 $\max(a)$。

[本题视频讲解](https://www.bilibili.com/video/BV1mUijBnEda/?t=4m32s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def findMaxVal(self, n: int, restrictions: List[List[int]], diff: List[int]) -> int:
        max_val = [inf] * n
        for i, mx in restrictions:
            max_val[i] = mx

        a = [0] * n
        for i, d in enumerate(diff):
            a[i + 1] = min(a[i] + d, max_val[i + 1])
        for i in range(n - 2, 0, -1):
            a[i] = min(a[i], a[i + 1] + diff[i])
        return max(a)
```

```java [sol-Java]
class Solution {
    public int findMaxVal(int n, int[][] restrictions, int[] diff) {
        int[] maxVal = new int[n];
        Arrays.fill(maxVal, Integer.MAX_VALUE);
        for (int[] r : restrictions) {
            maxVal[r[0]] = r[1];
        }

        int[] a = new int[n];
        for (int i = 0; i < diff.length; i++) {
            a[i + 1] = Math.min(a[i] + diff[i], maxVal[i + 1]);
        }
        int ans = a[n - 1];
        for (int i = n - 2; i > 0; i--) {
            a[i] = Math.min(a[i], a[i + 1] + diff[i]);
            ans = Math.max(ans, a[i]);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findMaxVal(int n, vector<vector<int>>& restrictions, vector<int>& diff) {
        vector<int> max_val(n, INT_MAX);
        for (auto& r : restrictions) {
            max_val[r[0]] = r[1];
        }

        vector<int> a(n);
        for (int i = 0; i < n - 1; i++) {
            a[i + 1] = min(a[i] + diff[i], max_val[i + 1]);
        }
        for (int i = n - 2; i > 0; i--) {
            a[i] = min(a[i], a[i + 1] + diff[i]);
        }
        return ranges::max(a);
    }
};
```

```go [sol-Go]
func findMaxVal(n int, restrictions [][]int, diff []int) int {
	maxVal := make([]int, n)
	for i := range maxVal {
		maxVal[i] = math.MaxInt
	}
	for _, r := range restrictions {
		maxVal[r[0]] = r[1]
	}

	a := make([]int, n)
	for i, d := range diff {
		a[i+1] = min(a[i]+d, maxVal[i+1])
	}
	for i := n - 2; i > 0; i-- {
		a[i] = min(a[i], a[i+1]+diff[i])
	}
	return slices.Max(a)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

如果改成 $\textit{diff}[i] = 1$，但 $n$ 最大是 $10^9$，怎么做？

这题是 [1840. 最高建筑高度](https://leetcode.cn/problems/maximum-building-height/)。

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
