「任意一种交易顺序下，都能完成所有交易」意味着我们要考虑**最坏情况**需要多少初始 $\textit{money}$。

什么情况是最坏情况？

**先亏钱**（$\textit{cost}>\textit{cashback}$），**再赚钱**。主打一个欲扬先抑。

记 $\textit{totalLose}$ 为亏钱时的所有 $\textit{cost}-\textit{cashback}$ 之和。

遍历 $\textit{transactions}$，分类讨论：

- 对于不亏钱（$\textit{cost}\le\textit{cashback}$）的交易，本着欲扬先抑的精神，**这笔交易发生在亏钱后**。根据题意，为了完成这笔交易，此时的钱至少得是 $\textit{cost}$，相当于初始 $\textit{money}=\textit{totalLose}+\textit{cost}$。
- 对于亏钱（$\textit{cost}>\textit{cashback}$）的交易，为了确保我们能在任意顺序下完成交易，**假设这笔交易是最后一笔亏钱的交易**，由于 $\textit{cost}-\textit{cashback}$ 已经计入 $\textit{totalLose}$ 中，需要从 $\textit{totalLose}$ 中减去 $\textit{cost}-\textit{cashback}$，再加上 $\textit{cost}$，化简得到初始 $\textit{money}=\textit{totalLose}+\textit{cashback}$。

取所有初始 $\textit{money}$ 的最大值，即为答案。相当于计算的是 $\textit{totalLose}$ 加上 $\min(\textit{cost}_i,\textit{cashback}_i)$ 的最大值。

请看 [视频讲解](https://www.bilibili.com/video/BV1MT411u7fW) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minimumMoney(self, transactions: List[List[int]]) -> int:
        total_lose = mx = 0
        for cost, cashback in transactions:
            total_lose += max(cost - cashback, 0)
            mx = max(mx, min(cost, cashback))
        return total_lose + mx
```

```java [sol-Java]
class Solution {
    public long minimumMoney(int[][] transactions) {
        long totalLose = 0;
        int mx = 0;
        for (var t : transactions) {
            totalLose += Math.max(t[0] - t[1], 0);
            mx = Math.max(mx, Math.min(t[0], t[1]));
        }
        return totalLose + mx;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumMoney(vector<vector<int>>& transactions) {
        long long total_lose = 0;
        int mx = 0;
        for (auto& t : transactions) {
            total_lose += max(t[0] - t[1], 0);
            mx = max(mx, min(t[0], t[1]));
        }
        return total_lose + mx;
    }
};
```

```go [sol-Go]
func minimumMoney(transactions [][]int) int64 {
	totalLose, mx := 0, 0
	for _, t := range transactions {
		totalLose += max(t[0]-t[1], 0)
		mx = max(mx, min(t[0], t[1]))
	}
	return int64(totalLose + mx)
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{transactions}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干变量。

## 思考题

如果把题干的「任意一种」改成「至少一种」要怎么做？

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
