题目要求 $\textit{copy}[i] - \textit{copy}[i - 1] = \textit{original}[i] - \textit{original}[i - 1]$，那么有

$$
\begin{aligned}
\textit{copy}[1] - \textit{copy}[0] &= \textit{original}[1] - \textit{original}[0]     \\
\textit{copy}[2] - \textit{copy}[1] &= \textit{original}[2] - \textit{original}[1]     \\
\textit{copy}[3] - \textit{copy}[2] &= \textit{original}[3] - \textit{original}[2]     \\
&\ \ \vdots \\
\textit{copy}[i] - \textit{copy}[i - 1] &= \textit{original}[i] - \textit{original}[i - 1] \\
\end{aligned}
$$

左右两边累加得

$$
\textit{copy}[i] - \textit{copy}[0] = \textit{original}[i] - \textit{original}[0] = 定值
$$

换句话说，确定了 $\textit{copy}[0]$，那么整个数组也就确定了。

题目要求

$$
u_i\le \textit{copy}[i] \le v_i
$$

不等式各项全部加上 $\textit{copy}[0] - \textit{copy}[i]$，得

$$
u_i + \textit{copy}[0] - \textit{copy}[i] \le \textit{copy}[0] \le v_i + \textit{copy}[0] - \textit{copy}[i]
$$

根据恒等式

$$
\textit{copy}[i] - \textit{copy}[0] = \textit{original}[i] - \textit{original}[0] = d_i
$$

替换，得

$$
u_i - d_i \le \textit{copy}[0] \le v_i - d_i
$$

所以我们可以得到 $n$ 个关于 $\textit{copy}[0]$ 的不等式，或者说区间：

$$
\begin{aligned}
& [u_0,v_0] \\
& [u_1 - d_1 ,v_1-d_1] \\
& [u_2 - d_2 ,v_2-d_2] \\
& \vdots \\
& [u_{n-1} - d_{n-1} ,v_{n-1}-d_{n-1}] \\
& \end{aligned}
$$

这些区间的**交集**，即为 $\textit{copy}[0]$ 能取到的值。

交集的大小即为答案。如果交集为空，返回 $0$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1m39bYiEVV/?t=1m2s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countArrays(self, original: List[int], bounds: List[List[int]]) -> int:
        mn, mx = -inf, inf
        for x, (u, v) in zip(original, bounds):
            d = x - original[0]
            mn = max(mn, u - d)  # 计算区间交集
            mx = min(mx, v - d)
        return max(mx - mn + 1, 0)  # 注意交集可能是空的
```

```java [sol-Java]
class Solution {
    public int countArrays(int[] original, int[][] bounds) {
        int mn = bounds[0][0], mx = bounds[0][1];
        for (int i = 1; i < bounds.length; i++) {
            int d = original[i] - original[0];
            mn = Math.max(mn, bounds[i][0] - d); // 计算区间交集
            mx = Math.min(mx, bounds[i][1] - d);
        }
        return Math.max(mx - mn + 1, 0); // 注意交集可能是空的
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countArrays(vector<int>& original, vector<vector<int>>& bounds) {
        int mn = bounds[0][0], mx = bounds[0][1];
        for (int i = 1; i < bounds.size(); i++) {
            int d = original[i] - original[0];
            mn = max(mn, bounds[i][0] - d); // 计算区间交集
            mx = min(mx, bounds[i][1] - d);
        }
        return max(mx - mn + 1, 0); // 注意交集可能是空的
    }
};
```

```go [sol-Go]
func countArrays(original []int, bounds [][]int) int {
	mn, mx := math.MinInt, math.MaxInt
	for i, b := range bounds {
		d := original[i] - original[0]
		mn = max(mn, b[0]-d) // 计算区间交集
		mx = min(mx, b[1]-d)
	}
	return max(mx-mn+1, 0) // 注意交集可能是空的
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{original}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
