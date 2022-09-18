下午 2 点在 B 站直播讲周赛和双周赛的题目，[欢迎关注](https://space.bilibili.com/206214/dynamic)~

---

#### 提示 1

考虑最坏情况，即先亏钱（$\textit{cost}>\textit{cashback}$），再赚钱。

记 $\textit{totalLose}$ 为亏钱下的所有 $\textit{cost}-\textit{cashback}$ 之和。

#### 提示 2

如何最大化初始 $\textit{money}$？

枚举所有交易，分类讨论：

- 对于 $\textit{cost}_i\le\textit{cashback}_i$ 的交易，这笔交易可以发生在亏钱后，此时初始 $\textit{money}=\textit{totalLose}+\textit{cost}_i$；
- 对于 $\textit{cost}_i>\textit{cashback}_i$ 的交易，由于已经考虑在 $\textit{totalLose}$ 中了，需要从 $\textit{totalLose}$ 中减去 $\textit{cost}_i-\textit{cashback}_i$，再加上 $\textit{cost}_i$，化简得到初始 $\textit{money}=\textit{totalLose}+\textit{cashback}_i$。

取所有初始 $\textit{money}$ 的最大值，即为答案。

```py [sol1-Python3]
class Solution:
    def minimumMoney(self, transactions: List[List[int]]) -> int:
        total_lose = mx = 0
        for cost, cashback in transactions:
            total_lose += max(cost - cashback, 0)
            mx = max(mx, min(cost, cashback))
        return total_lose + mx
```

```java [sol1-Java]
class Solution {
    public long minimumMoney(int[][] transactions) {
        var totalLose = 0L;
        var mx = 0;
        for (var t : transactions) {
            totalLose += Math.max(t[0] - t[1], 0);
            mx = Math.max(mx, Math.min(t[0], t[1]));
        }
        return totalLose + mx;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    long long minimumMoney(vector<vector<int>> &transactions) {
        long total_lose = 0L;
        int mx = 0;
        for (auto &t : transactions) {
            total_lose += max(t[0] - t[1], 0);
            mx = max(mx, min(t[0], t[1]));
        }
        return total_lose + mx;
    }
};
```

```go [sol1-Go]
func minimumMoney(transactions [][]int) int64 {
	totalLose, mx := 0, 0
	for _, t := range transactions {
		totalLose += max(t[0]-t[1], 0)
		mx = max(mx, min(t[0], t[1]))
	}
	return int64(totalLose + mx)
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{transactions}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干变量。

#### 思考题

如果把题干的「任意一种」改成「至少一种」要怎么做？
