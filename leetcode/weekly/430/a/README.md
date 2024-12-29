由于我们只能变大，不能变小，那么第一个数肯定不需要变。

对于后面的每个数 $x$：

- 如果它比前一个数 $\textit{pre}$ 大，那么不变。
- 否则，$x$ 至少要增大到 $\textit{pre}+1$，才能保持严格递增。增大到恰好等于 $\textit{pre}+1$ 是最优的（不然后面的数需要变得更大），操作 $\textit{pre}+1-x$ 次。

具体请看 [视频讲解](https://www.bilibili.com/video/BV13f68YjE7o/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minimumOperations(self, grid: List[List[int]]) -> int:
        ans = 0
        for col in zip(*grid):
            pre = -inf
            for x in col:
                ans += max(pre + 1 - x, 0)
                pre = max(pre + 1, x)
        return ans
```

```java [sol-Java]
class Solution {
    public int minimumOperations(int[][] grid) {
        int ans = 0;
        for (int j = 0; j < grid[0].length; j++) {
            int pre = -1;
            for (int[] row : grid) {
                int x = row[j];
                ans += Math.max(pre + 1 - x, 0);
                pre = Math.max(pre + 1, x);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumOperations(vector<vector<int>>& grid) {
        int ans = 0;
        for (int j = 0; j < grid[0].size(); j++) {
            int pre = -1;
            for (auto& row : grid) {
                int x = row[j];
                ans += max(pre + 1 - x, 0);
                pre = max(pre + 1, x);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumOperations(grid [][]int) (ans int) {
	for j := range grid[0] {
		pre := -1
		for _, row := range grid {
			x := row[j]
			ans += max(pre+1-x, 0)
			pre = max(pre+1, x)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。

## 变形题

1. 改成严格递减。
2. 改成操作是把一个数减少 $1$。
3. 改成操作是把一个数恰好增加 $k$。其中 $k$ 是一个非零整数。

更多相似题目，见下面贪心题单中的「**§1.4 从最左/最右开始贪心**」。

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
10. 【本题相关】[贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
