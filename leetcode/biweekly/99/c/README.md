本题和 [56. 合并区间](https://leetcode.cn/problems/merge-intervals/) 是类似的。

我们需要把有交集的区间都归在同一个集合中，假设最后分成了 $m$ 个集合，那么每个集合都可以决定要在第一个组还是第二个组，所以方案数为 $2^m$。

怎么求出 $m$ 呢？

初始化 $m=1$。按照左端点排序，遍历数组，同时维护区间右端点的最大值 $\textit{maxR}$：

- 如果当前区间的左端点大于 $\textit{maxR}$，由于已经按照左端点排序了，那么后面任何区间都不会和之前的区间有交集，换句话说，产生了一个新的集合，$m$ 加一。
- 否则，当前区间与上一个区间在同一个集合。

附：[视频讲解](https://www.bilibili.com/video/BV1dY4y1C77x/)

```py [sol1-Python3]
class Solution:
    def countWays(self, ranges: List[List[int]]) -> int:
        ranges.sort(key=lambda p: p[0])
        m, max_r = 1, ranges[0][1]
        for l, r in ranges:
            m += l > max_r  # 产生了一个新的集合
            max_r = max(max_r, r)
        return pow(2, m, 10 ** 9 + 7)
```

```java [sol1-Java]
class Solution {
    private static final int MOD = (int) 1e9 + 7;

    public int countWays(int[][] ranges) {
        Arrays.sort(ranges, (a, b) -> a[0] - b[0]);
        int ans = 2, maxR = ranges[0][1];
        for (var p : ranges) {
            if (p[0] > maxR) // 产生了一个新的集合
                ans = ans * 2 % MOD;
            maxR = Math.max(maxR, p[1]);
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
    const int MOD = 1e9 + 7;
public:
    int countWays(vector<vector<int>> &ranges) {
        sort(ranges.begin(), ranges.end(), [](auto &a, auto &b) {
            return a[0] < b[0];
        });
        int ans = 2, max_r = ranges[0][1];
        for (auto &p : ranges) {
            if (p[0] > max_r) // 产生了一个新的集合
                ans = ans * 2 % MOD;
            max_r = max(max_r, p[1]);
        }
        return ans;
    }
};
```

```go [sol1-Go]
func countWays(ranges [][]int) int {
	const mod int = 1e9 + 7
	sort.Slice(ranges, func(i, j int) bool { return ranges[i][0] < ranges[j][0] })
	ans, maxR := 2, ranges[0][1]
	for _, p := range ranges[1:] {
		if p[0] > maxR { // 产生了一个新的集合
			ans = ans * 2 % mod
		}
		maxR = max(maxR, p[1])
	}
	return ans
}

func max(a, b int) int { if a < b { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$。忽略排序时的栈开销，仅用到若干额外变量。

### 相似题目

- [56. 合并区间](https://leetcode.cn/problems/merge-intervals/)
- [55. 跳跃游戏](https://leetcode.cn/problems/jump-game/)
- [45. 跳跃游戏 II](https://leetcode.cn/problems/jump-game-ii/)
- [1024. 视频拼接](https://leetcode.cn/problems/video-stitching/)
- [1326. 灌溉花园的最少水龙头数目](https://leetcode.cn/problems/minimum-number-of-taps-to-open-to-water-a-garden/)
