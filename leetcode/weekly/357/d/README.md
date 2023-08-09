请看 [视频讲解](https://www.bilibili.com/video/BV1Yr4y1o7aP/) 第四题。

**按照利润从大到小排序**。先把前 $k$ 个项目选上。

考虑选第 $k+1$ 个项目，为了选它，我们必须从前 $k$ 个项目中**移除**一个项目。

由于已经按照利润从大到小排序，选这个项目不会让 $\textit{total\_profit}$ 变大，所以我们重点考虑能否让 $\textit{distinct\_categories}$ 变大。

分类讨论：

- 如果第 $k+1$ 个项目和前面某个已选项目的类别相同，那么无论怎么移除都不会让 $\textit{distinct\_categories}$ 变大，所以无需选择这个项目。
- 如果第 $k+1$ 个项目和前面任何已选项目的类别都不一样，考虑移除前面已选项目中的哪一个：
   - 如果移除的项目的类别只出现一次，那么选第 $k+1$ 个项目后，$\textit{distinct\_categories}$ 一减一增，保持不变，所以不考虑这种情况。
   - 如果移除的项目的类别重复出现多次，那么选第 $k+1$ 个项目后，$\textit{distinct\_categories}$ 会增加一，此时有可能会让优雅度变大，**一定**要选择这个项目。为什么说「一定」呢？因为 $\textit{total\_profit}$ 只会变小，我们现在的目标就是让 $\textit{total\_profit}$ 保持尽量大，同时让 $\textit{distinct\_categories}$ 增加，那么能让 $\textit{distinct\_categories}$ 增加就立刻选上！因为后面的利润更小，现在不选的话将来 $\textit{total\_profit}$ 只会更小。

按照这个过程，继续考虑选择后面的项目。计算优雅度，取最大值，即为答案。

代码实现时，我们应当移除已选项目中类别和前面重复且利润最小的项目，这可以用一个栈 $\textit{duplicate}$ 来维护，由于利润从大到小排序，所以栈顶就是最小的利润。注意对于后面的项目，由于我们只考虑之前没出现过的类别，也就是说这个后面的项目的类别只出现了一次，所以不应当加到 $\textit{duplicate}$ 中。

注：这个算法叫做**反悔贪心**。

```py [sol-Python3]
class Solution:
    def findMaximumElegance(self, items: List[List[int]], k: int) -> int:
        items.sort(key=lambda i: -i[0])  # 把利润从大到小排序
        ans = total_profit = 0
        vis = set()
        duplicate = []  # 重复类别的利润
        for i, (profit, category) in enumerate(items):
            if i < k:
                total_profit += profit
                if category not in vis:
                    vis.add(category)
                else:  # 重复类别
                    duplicate.append(profit)
            elif duplicate and category not in vis:
                vis.add(category)
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
        long ans = 0, totalProfit = 0;
        var vis = new HashSet<Integer>();
        var duplicate = new ArrayDeque<Integer>(); // 重复类别的利润
        for (int i = 0; i < items.length; i++) {
            int profit = items[i][0], category = items[i][1];
            if (i < k) {
                totalProfit += profit;
                if (!vis.add(category)) // 重复类别
                    duplicate.push(profit);
            } else if (!duplicate.isEmpty() && vis.add(category)) {
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
    long long findMaximumElegance(vector<vector<int>> &items, int k) {
        // 把利润从大到小排序
        sort(items.begin(), items.end(), [](const auto &a, const auto &b) {
            return a[0] > b[0];
        });
        long long ans = 0, total_profit = 0;
        unordered_set<int> vis;
        stack<int> duplicate; // 重复类别的利润
        for (int i = 0; i < items.size(); i++) {
            int profit = items[i][0], category = items[i][1];
            if (i < k) {
                total_profit += profit;
                if (!vis.insert(category).second) // 重复类别
                    duplicate.push(profit);
            } else if (!duplicate.empty() && vis.insert(category).second) {
                total_profit += profit - duplicate.top(); // 选一个重复类别中的最小利润替换
                duplicate.pop();
            } // else：比前面的利润小，而且类别还重复了，选它只会让 totalProfit 变小，vis.size() 不变，优雅度不会变大
            ans = max(ans, total_profit + (long long) vis.size() * (long long) vis.size());
        }
        return ans;
    }
};
```

```go [sol-Go]
func findMaximumElegance(items [][]int, k int) int64 {
	// 把利润从大到小排序
	sort.Slice(items, func(i, j int) bool { return items[i][0] > items[j][0] })
	ans, totalProfit := 0, 0
	vis := map[int]bool{}
	duplicate := []int{} // 重复类别的利润
	for i, p := range items {
		profit, category := p[0], p[1]
		if i < k {
			totalProfit += profit
			if !vis[category] {
				vis[category] = true
			} else { // 重复类别
				duplicate = append(duplicate, profit)
			}
		} else if len(duplicate) > 0 && !vis[category] {
			vis[category] = true
			totalProfit += profit - duplicate[len(duplicate)-1] // 选一个重复类别中的最小利润替换
			duplicate = duplicate[:len(duplicate)-1]
		} // else 比前面的利润小，而且类别还重复了，选它只会让 totalProfit 变小，len(vis) 不变，优雅度不会变大
		ans = max(ans, totalProfit+len(vis)*len(vis))
	}
	return int64(ans)
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{items}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$。
