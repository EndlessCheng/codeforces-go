阅读理解：

- 第 $0$ 位员工的用时为 $\textit{logs}[0][1]$。
- 第 $i$（$i>0$）位员工的用时为 $\textit{logs}[i][1] - \textit{logs}[i-1][1]$。

计算用时最大的员工编号，若有多个用时最大的，取其中编号最小的。

```py [sol1-Python3]
class Solution:
    def hardestWorker(self, _: int, logs: List[List[int]]) -> int:
        ans, max_t = logs[0]
        for (_, t1), (i, t) in pairwise(logs):
            t -= t1
            if t > max_t or t == max_t and i < ans:
                ans, max_t = i, t
        return ans
```

```java [sol1-Java]
class Solution {
    public int hardestWorker(int n, int[][] logs) {
        int ans = logs[0][0], maxT = logs[0][1];
        for (int i = 1; i < logs.length; i++) {
            int t = logs[i][1] - logs[i - 1][1];
            if (t > maxT || t == maxT && logs[i][0] < ans) {
                ans = logs[i][0];
                maxT = t;
            }
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int hardestWorker(int _, vector<vector<int>> &logs) {
        int ans = logs[0][0], max_t = logs[0][1];
        for (int i = 1; i < logs.size(); ++i) {
            int t = logs[i][1] - logs[i - 1][1];
            if (t > max_t || t == max_t && logs[i][0] < ans) {
                ans = logs[i][0];
                max_t = t;
            }
        }
        return ans;
    }
};
```

```go [sol1-Go]
func hardestWorker(_ int, logs [][]int) int {
	ans, maxT := logs[0][0], logs[0][1]
	for i := 1; i < len(logs); i++ {
		t := logs[i][1] - logs[i-1][1]
		if t > maxT || t == maxT && logs[i][0] < ans {
			ans, maxT = logs[i][0], t
		}
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(m)$，其中 $m$ 为 $\textit{logs}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

---

[往期每日一题题解](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注[ biIibiIi@灵茶山艾府](https://space.bilibili.com/206214)，高质量算法教学，持续输出中~
