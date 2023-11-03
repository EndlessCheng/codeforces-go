定义 $f[i][j]$ 表示用前 $i$ 种题目恰好组成 $j$ 分的方案数。

对于第 $i$ 种题目，枚举做 $k$ 道题目，则子问题为「前 $i-1$ 种题目恰好组成 $j-k\cdot \textit{marks}_i$ 分的方案数」，因此有

$$
f[i][j] = \sum\limits_{k=0} f[i-1][j-k\cdot \textit{marks}_i]
$$

注意 $k$ 不能超过 $\textit{count}_i$，且 $j-k\cdot \textit{marks}_i\ge 0$。

代码实现时可以像 0-1 背包那样，压缩成一维，具体可以看[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

> 注：滚动优化后，$k=0$ 就是 $f[j]$，无需计算。

- [本题视频讲解](https://www.bilibili.com/video/BV1SN411c7eD/)
- 更快的做法请看 [2902. 和带限制的子多重集合的数目](https://leetcode.cn/problems/count-of-sub-multisets-with-bounded-sum/solution/duo-zhong-bei-bao-fang-an-shu-cong-po-su-f5ay/)

```py [sol1-Python3]
class Solution:
    def waysToReachTarget(self, target: int, types: List[List[int]]) -> int:
        MOD = 10 ** 9 + 7
        f = [1] + [0] * target
        for count, marks in types:
            for j in range(target, 0, -1):
                for k in range(1, min(count, j // marks) + 1):
                    f[j] += f[j - k * marks]
                f[j] %= MOD
        return f[-1]
```

```java [sol1-Java]
class Solution {
    private static final int MOD = (int) 1e9 + 7;

    public int waysToReachTarget(int target, int[][] types) {
        var f = new int[target + 1];
        f[0] = 1;
        for (var p : types) {
            int count = p[0], marks = p[1];
            for (int j = target; j > 0; --j)
                for (int k = 1; k <= count && k <= j / marks; ++k)
                    f[j] = (f[j] + f[j - k * marks]) % MOD;
        }
        return f[target];
    }
}
```

```cpp [sol1-C++]
class Solution {
    const int MOD = 1e9 + 7;
public:
    int waysToReachTarget(int target, vector<vector<int>> &types) {
        int f[target + 1];
        memset(f, 0, sizeof(f));
        f[0] = 1;
        for (auto &p : types) {
            int count = p[0], marks = p[1];
            for (int j = target; j > 0; --j)
                for (int k = 1; k <= min(count, j / marks); ++k)
                    f[j] = (f[j] + f[j - k * marks]) % MOD;
        }
        return f[target];
    }
};
```

```go [sol1-Go]
func waysToReachTarget(target int, types [][]int) int {
	const mod int = 1e9 + 7
	f := make([]int, target+1)
	f[0] = 1
	for _, p := range types {
		count, marks := p[0], p[1]
		for j := target; j > 0; j-- {
			for k := 1; k <= count && k <= j/marks; k++ {
				f[j] += f[j-k*marks]
			}
			f[j] %= mod
		}
	}
	return f[target]
}
```

### 复杂度分析

- 时间复杂度：$O(TS)$，其中 $T$ 为 $\textit{target}$，$S$ 为所有 $\textit{count}_i$ 之和。
- 空间复杂度：$O(T)$。

### 相似题目

- [1981. 最小化目标值与所选元素的差](https://leetcode.cn/problems/minimize-the-difference-between-target-and-chosen-elements/)
- [2218. 从栈中取出 K 个硬币的最大面值和](https://leetcode.cn/problems/maximum-value-of-k-coins-from-piles/)

### 思考题

如果同类型题目需要区分，要怎么做呢？
