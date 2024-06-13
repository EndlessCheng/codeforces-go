按照利润从大到小排序。先把前 $k$ 个项目选上。

考虑第 $k+1$ 个项目，如果要选它，我们必须从前 $k$ 个项目中**移除**一个项目。

由于已经按照利润从大到小排序，选这个项目不会让 $\textit{totalProfit}$ 变大，所以**重点考虑能否让** $\textit{distinctCategories}$ **变大**。

分类讨论：

- 如果第 $k+1$ 个项目和前面某个已选项目的类别相同，那么无论怎么移除都不会让 $\textit{distinctCategories}$ 变大，所以无需选择这个项目。
- 如果第 $k+1$ 个项目和前面任何已选项目的类别都不一样，考虑移除前面已选项目中的哪一个：
   - 如果移除的项目的类别只出现一次，那么选第 $k+1$ 个项目后，$\textit{distinctCategories}$ 一减一增，保持不变，所以不考虑这种情况。
   - 如果移除的项目的类别重复出现多次，那么选第 $k+1$ 个项目后，$\textit{distinctCategories}$ 会增加一，此时有可能会让优雅度变大，**一定**要选择这个项目。为什么说「一定」呢？因为 $\textit{totalProfit}$ 只会变小，我们现在的目标就是让 $\textit{totalProfit}$ 保持尽量大，同时让 $\textit{distinctCategories}$ 增加，那么能让 $\textit{distinctCategories}$ 增加就立刻选上！因为后面的利润更小，现在不选的话将来 $\textit{totalProfit}$ 只会更小。

按照这个过程，继续考虑选择后面的项目。计算优雅度，取最大值，即为答案。

代码实现时，我们应当移除已选项目中类别和前面重复且利润最小的项目，这可以用一个栈 $\textit{duplicate}$ 来维护，由于利润从大到小排序，所以栈顶就是最小的利润。入栈前判断 $\textit{category}$ 之前是否遇到过，遇到则入栈。

注：这个算法叫做**反悔贪心**。

[本题视频讲解](https://www.bilibili.com/video/BV1Yr4y1o7aP/) 第四题。

```py [sol-Python3]
class Solution:
    def findMaximumElegance(self, items: List[List[int]], k: int) -> int:
        items.sort(key=lambda p: -p[0])  # 把利润从大到小排序
        ans = total_profit = 0
        vis = set()
        duplicate = []  # 重复类别的利润
        for i, (profit, category) in enumerate(items):
            if i < k:
                total_profit += profit  # 累加前 k 个项目的利润
                if category not in vis:
                    vis.add(category)
                else:  # 重复类别
                    duplicate.append(profit)
            elif duplicate and category not in vis:  # 之前没有的类别
                vis.add(category)  # len(vis) 变大
                total_profit += profit - duplicate.pop()  # 用最小利润替换
            # else: 比前面的利润小，而且类别还重复了，选它只会让 total_profit 变小，len(vis) 不变，优雅度不会变大
            ans = max(ans, total_profit + len(vis) * len(vis))
        return ans
```

```java [sol-Java]
class Solution {
    public long findMaximumElegance(int[][] items, int k) {
        // 把利润从大到小排序
        Arrays.sort(items, (a, b) -> b[0] - a[0]);
        long ans = 0;
        long totalProfit = 0;
        Set<Integer> vis = new HashSet<>();
        Deque<Integer> duplicate = new ArrayDeque<>(); // 重复类别的利润
        for (int i = 0; i < items.length; i++) {
            int profit = items[i][0];
            int category = items[i][1];
            if (i < k) {
                totalProfit += profit; // 累加前 k 个项目的利润
                if (!vis.add(category)) { // 重复类别
                    duplicate.push(profit);
                }
            } else if (!duplicate.isEmpty() && vis.add(category)) { // 之前没有的类别
                totalProfit += profit - duplicate.pop(); // 选一个重复类别中的最小利润替换
            } // else：比前面的利润小，而且类别还重复了，选它只会让 totalProfit 变小，vis.size() 不变，优雅度不会变大
            ans = Math.max(ans, totalProfit + (long) vis.size() * vis.size()); // 注意 1e5*1e5 会溢出
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long findMaximumElegance(vector<vector<int>>& items, int k) {
        // 把利润从大到小排序
        ranges::sort(items, [](const auto &a, const auto &b) { return a[0] > b[0]; });
        long long ans = 0, total_profit = 0;
        unordered_set<int> vis;
        stack<int> duplicate; // 重复类别的利润
        for (int i = 0; i < items.size(); i++) {
            int profit = items[i][0], category = items[i][1];
            if (i < k) {
                total_profit += profit; // 累加前 k 个项目的利润
                if (!vis.insert(category).second) { // 重复类别
                    duplicate.push(profit);
                }
            } else if (!duplicate.empty() && vis.insert(category).second) { // 之前没有的类别
                total_profit += profit - duplicate.top(); // 选一个重复类别中的最小利润替换
                duplicate.pop();
            } // else：比前面的利润小，而且类别还重复了，选它只会让 total_profit 变小，vis.size() 不变，优雅度不会变大
            ans = max(ans, total_profit + (long long) vis.size() * (long long) vis.size());
        }
        return ans;
    }
};
```

```go [sol-Go]
func findMaximumElegance(items [][]int, k int) int64 {
	// 把利润从大到小排序
	slices.SortFunc(items, func(a, b []int) int { return b[0] - a[0] })
	ans, totalProfit := 0, 0
	vis := map[int]bool{}
	duplicate := []int{} // 重复类别的利润
	for i, p := range items {
		profit, category := p[0], p[1]
		if i < k {
			totalProfit += profit // 累加前 k 个项目的利润
			if !vis[category] {
				vis[category] = true
			} else { // 重复类别
				duplicate = append(duplicate, profit)
			}
		} else if len(duplicate) > 0 && !vis[category] { // 之前没有的类别
			vis[category] = true // len(vis) 变大
			totalProfit += profit - duplicate[len(duplicate)-1] // 选一个重复类别中的最小利润替换
			duplicate = duplicate[:len(duplicate)-1]
		} // else：比前面的利润小，而且类别还重复了，选它只会让 totalProfit 变小，len(vis) 不变，优雅度不会变大
		ans = max(ans, totalProfit+len(vis)*len(vis))
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{items}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

计算 $k$ 个数：$\textit{distinctCategories}$ 恰好等于 $1,2,3,\cdots,k$ 时的最大优雅度。

## 相似题目

见 [数据结构题单](https://leetcode.cn/circle/discuss/mOr1u6/) 中的「§5.5 反悔堆」。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
