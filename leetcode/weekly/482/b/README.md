### 情况一：各买各的

不用 $\textit{costBoth}$，花费

$$
\textit{cost}_1\cdot \textit{need}_1 + \textit{cost}_2\cdot \textit{need}_2
$$

### 情况二：我包了

用 $\textit{costBoth}$ 买 $\max(\textit{need}_1,\textit{need}_2)$ 个，花费

$$
\textit{costBoth}\cdot \max(\textit{need}_1,\textit{need}_2)
$$

### 情况三：混合策略

如果 $\textit{costBoth} < \textit{cost}_1 + \textit{cost}_2$，但又比其中一个大，那么可以先买 $\min(\textit{need}_1,\textit{need}_2)$ 个，剩余的再单独买。

不妨设 $\textit{need}_1 \le \textit{need}_2$（不满足则交换），花费

$$
\textit{costBoth}\cdot \textit{need}_1 + \textit{cost}_2\cdot(\textit{need}_2 - \textit{need}_1)
$$

三种情况取最小值。

代码实现时，无需判断 $\textit{costBoth} < \textit{cost}_1 + \textit{cost}_2$。如果不满足，那么情况一（各买各的）更优。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def minimumCost(self, cost1: int, cost2: int, costBoth: int, need1: int, need2: int) -> int:
        res1 = cost1 * need1 + cost2 * need2  # 各买各的
        if need1 > need2:
            need1, need2 = need2, need1
            cost2 = cost1
        res2 = costBoth * need2  # 我包了
        res3 = costBoth * need1 + cost2 * (need2 - need1)  # 混合策略
        return min(res1, res2, res3)
```

```java [sol-Java]
class Solution {
    public long minimumCost(int cost1, int cost2, int costBoth, int need1, int need2) {
        long res1 = (long) cost1 * need1 + (long) cost2 * need2; // 各买各的
        if (need1 > need2) {
            int tmp = need1;
            need1 = need2;
            need2 = tmp;
            cost2 = cost1;
        }
        long res2 = (long) costBoth * need2; // 我包了
        long res3 = (long) costBoth * need1 + (long) cost2 * (need2 - need1); // 混合策略
        return Math.min(res1, Math.min(res2, res3));
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumCost(int cost1, int cost2, int costBoth, int need1, int need2) {
        long long res1 = 1LL * cost1 * need1 + 1LL * cost2 * need2; // 各买各的
        if (need1 > need2) {
            swap(need1, need2);
            cost2 = cost1;
        }
        long long res2 = 1LL * costBoth * need2; // 我包了
        long long res3 = 1LL * costBoth * need1 + 1LL * cost2 * (need2 - need1); // 混合策略
        return min({res1, res2, res3});
    }
};
```

```go [sol-Go]
func minimumCost(cost1, cost2, costBoth, need1, need2 int) int64 {
	res1 := cost1*need1 + cost2*need2 // 各买各的
	if need1 > need2 {
		need1, need2 = need2, need1
		cost2 = cost1
	}
	res2 := costBoth * need2 // 我包了
	res3 := costBoth*need1 + cost2*(need2-need1) // 混合策略
	return int64(min(res1, res2, res3))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

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
