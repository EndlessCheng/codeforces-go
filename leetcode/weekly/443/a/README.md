要想到达 $i$，我们可以先找一个 $[0,i]$ 中 $\textit{cost}$ 最小的人，和他交换，然后再和右边的人免费交换，直到到达 $i$。

所以 $\textit{answer}[i]$ 就是 $[0,i]$ 中 $\textit{cost}$ 的最小值（前缀最小值）。

这可以从左到右遍历 $\textit{cost}$ 算出来。

代码实现时，直接原地修改 $\textit{cost}$，这样可以做到 $\mathcal{O}(1)$ 空间。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def minCosts(self, cost: List[int]) -> List[int]:
        for i in range(1, len(cost)):
            cost[i] = min(cost[i], cost[i - 1])
        return cost
```

```py [sol-Python3 一行]
class Solution:
    def minCosts(self, cost: List[int]) -> List[int]:
        return list(accumulate(cost, min))
```

```java [sol-Java]
class Solution {
    public int[] minCosts(int[] cost) {
        for (int i = 1; i < cost.length; i++) {
            cost[i] = Math.min(cost[i], cost[i - 1]);
        }
        return cost;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> minCosts(vector<int>& cost) {
        for (int i = 1; i < cost.size(); i++) {
            cost[i] = min(cost[i], cost[i - 1]);
        }
        return cost;
    }
};
```

```go [sol-Go]
func minCosts(cost []int) []int {
	for i := 1; i < len(cost); i++ {
		cost[i] = min(cost[i], cost[i-1])
	}
	return cost
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{cost}$ 的长度。
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
