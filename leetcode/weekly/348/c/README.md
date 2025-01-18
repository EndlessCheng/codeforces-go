## 视频讲解

见[【周赛 348】](https://www.bilibili.com/video/BV1do4y1K7Wq/)第三题，欢迎点赞投币！

## 提示 1

如果对同一行反复操作，那么只有最后一次对这行的操作会计入答案。列同理。

## 提示 2

正难则反，倒序操作 $\textit{queries}$。

## 提示 3

以行为例。如果 $\textit{queries}[i]$ 操作的是行，那么需要知道：

1. 这一行之前有没有操作过（这里「之前」指大于 $i$ 的操作）。对此可以用哈希表 $\textit{visRow}$ 记录被操作过的行号，哈希表 $\textit{visCol}$ 记录被操作过的列号。
2. 这一行有多少列之前操作过，这就是 $\textit{visCol}$ 的长度 $m$。那么剩余可以填入的格子为 $n-m$，答案增加了 $(n-m)\cdot \textit{val}_i$。

这样可以做到 $\mathcal{O}(q)$ 的时间复杂度（与 $n$ 无关）。

代码实现时，可以把 $\textit{visRow}$ 和 $\textit{visCol}$ 放到一个长为 $2$ 的数组中，简化代码逻辑。

```py [sol-Python3]
class Solution:
    def matrixSumQueries(self, n: int, queries: List[List[int]]) -> int:
        ans = 0
        vis = [set(), set()]
        for type, index, val in reversed(queries):
            if index not in vis[type]:  # 后面（>i）没有对这一行/列的操作
                # 这一行/列还剩下 n-len(vis[type^1]) 个可以填入的格子
                ans += (n - len(vis[type ^ 1])) * val
                vis[type].add(index)  # 标记操作过
        return ans
```

```java [sol-Java]
class Solution {
    public long matrixSumQueries(int n, int[][] queries) {
        long ans = 0;
        Set<Integer>[] vis = new Set[]{new HashSet<>(), new HashSet<>()};
        for (int i = queries.length - 1; i >= 0; i--) {
            var q = queries[i];
            int type = q[0], index = q[1], val = q[2];
            if (!vis[type].contains(index)) { // 后面（>i）没有对这一行/列的操作
                // 这一行/列还剩下 n-vis[type^1].size() 个可以填入的格子
                ans += (long) (n - vis[type ^ 1].size()) * val;
                vis[type].add(index);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long matrixSumQueries(int n, vector<vector<int>> &queries) {
        long long ans = 0;
        unordered_set<int> vis[2];
        for (int i = queries.size() - 1; i >= 0; i--) {
            auto &q = queries[i];
            int type = q[0], index = q[1], val = q[2];
            if (!vis[type].count(index)) { // 后面（>i）没有对这一行/列的操作
                // 这一行/列还剩下 n-vis[type^1].size() 个可以填入的格子
                ans += (long long) (n - vis[type ^ 1].size()) * val;
                vis[type].insert(index);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func matrixSumQueries(n int, queries [][]int) (ans int64) {
	vis := [2]map[int]bool{{}, {}}
	for i := len(queries) - 1; i >= 0; i-- {
		q := queries[i]
		tp, index, val := q[0], q[1], q[2]
		if !vis[tp][index] { // 后面（>i）没有对这一行/列的操作
			// 这一行/列还剩下 n-len(vis[tp^1]) 个可以填入的格子
			ans += int64(n-len(vis[tp^1])) * int64(val)
			vis[tp][index] = true // 标记操作过
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(q)$，其中 $q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(\min\{q,n\})$。哈希表中至多有 $\mathcal{O}(n)$ 个数。

## 相似题目

- [2382. 删除操作后的最大子段和](https://leetcode.cn/problems/maximum-segment-sum-after-removals/)

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
