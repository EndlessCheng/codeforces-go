本题 [视频讲解](https://www.bilibili.com/video/BV1CW4y1k7B3) 已出炉，欢迎点赞三连~

---

#### 提示 1

垂直方向或水平方向切一刀后，我们得到了两块更小的木块，也得到了两个更小的子问题。

#### 提示 2

枚举垂直方向切割的位置和水平方向切割的位置。

#### 提示 3

定义 $f[i][j]$ 表示对一块高 $i$ 宽 $j$ 的木块，切割后能得到的最多钱数。

如果直接售卖，则收益为对应的 $\textit{price}_i$（如果存在的话）。

如果垂直切割，则最大收益为

$$
\max_{k=1}^{j-1} f[i][k]+f[i][j-k]
$$

如果水平切割，则最大收益为

$$
\max_{k=1}^{i-1} f[k][j]+f[i-k][j]
$$

取上述三种情况的最大值，即为最终的 $f[i][j]$。

答案为 $f[m][n]$。

#### 复杂度分析

- 时间复杂度：$O(mn(m+n))$。
- 空间复杂度：$O(mn)$。

```Python [sol1-Python3]
class Solution:
    def sellingWood(self, m: int, n: int, prices: List[List[int]]) -> int:
        pr = {(h, w): p for h, w, p in prices}
        f = [[0] * (n + 1) for _ in range(m + 1)]
        for i in range(1, m + 1):
            for j in range(1, n + 1):
                f[i][j] = max(pr.get((i, j), 0),
                              max((f[i][k] + f[i][j - k] for k in range(1, j)), default=0),  # 垂直切割
                              max((f[k][j] + f[i - k][j] for k in range(1, i)), default=0))  # 水平切割
        return f[m][n]
```

```java [sol1-Java]
class Solution {
    public long sellingWood(int m, int n, int[][] prices) {
        var pr = new int[m + 1][n + 1];
        for (var p : prices) pr[p[0]][p[1]] = p[2];

        var f = new long[m + 1][n + 1];
        for (var i = 1; i <= m; i++)
            for (var j = 1; j <= n; j++) {
                f[i][j] = pr[i][j];
                for (var k = 1; k < j; k++) f[i][j] = Math.max(f[i][j], f[i][k] + f[i][j - k]); // 垂直切割
                for (var k = 1; k < i; k++) f[i][j] = Math.max(f[i][j], f[k][j] + f[i - k][j]); // 水平切割
            }
        return f[m][n];
    }
}
```

```C++ [sol1-C++]
class Solution {
public:
    long long sellingWood(int m, int n, vector<vector<int>> &prices) {
        int pr[m + 1][n + 1]; memset(pr, 0, sizeof(pr));
        for (auto &p : prices) pr[p[0]][p[1]] = p[2];

        long f[m + 1][n + 1];
        for (int i = 1; i <= m; i++)
            for (int j = 1; j <= n; j++) {
                f[i][j] = pr[i][j];
                for (int k = 1; k < j; k++) f[i][j] = max(f[i][j], f[i][k] + f[i][j - k]); // 垂直切割
                for (int k = 1; k < i; k++) f[i][j] = max(f[i][j], f[k][j] + f[i - k][j]); // 水平切割
            }
        return f[m][n];
    }
};
```

```go [sol1-Go]
func sellingWood(m, n int, prices [][]int) int64 {
	pr := make([][]int, m+1)
	for i := range pr {
		pr[i] = make([]int, n+1)
	}
	for _, price := range prices {
		pr[price[0]][price[1]] = price[2]
	}

	f := make([][]int, m+1)
	for i := 1; i <= m; i++ {
		f[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			f[i][j] = pr[i][j]
			for k := 1; k < j; k++ { // 垂直切割
				f[i][j] = max(f[i][j], f[i][k]+f[i][j-k])
			}
			for k := 1; k < i; k++ { // 水平切割
				f[i][j] = max(f[i][j], f[k][j]+f[i-k][j])
			}
		}
	}
	return int64(f[m][n])
}

func max(a, b int) int { if b > a { return b }; return a }
```

上述代码有两处优化点：

1. 根据对称性，内层循环枚举到一半的位置即可；
2. 注意到我们是从小往大计算 $f$ 的，我们可以直接将 $\textit{prices}$ 记录到 $f$ 中，而不会影响每个 $f[i][j]$ 的计算。

```Python [sol2-Python3]
class Solution:
    def sellingWood(self, m: int, n: int, prices: List[List[int]]) -> int:
        f = [[0] * (n + 1) for _ in range(m + 1)]
        for h, w, p in prices:
            f[h][w] = p
        for i in range(1, m + 1):
            for j in range(1, n + 1):
                f[i][j] = max(f[i][j],
                              max((f[i][k] + f[i][j - k] for k in range(1, j // 2 + 1)), default=0),  # 垂直切割
                              max((f[k][j] + f[i - k][j] for k in range(1, i // 2 + 1)), default=0))  # 水平切割
        return f[m][n]
```

```java [sol2-Java]
class Solution {
    public long sellingWood(int m, int n, int[][] prices) {
        var f = new long[m + 1][n + 1];
        for (var p : prices) f[p[0]][p[1]] = p[2];
        for (var i = 1; i <= m; i++)
            for (var j = 1; j <= n; j++) {
                for (var k = 1; k <= j / 2; k++) f[i][j] = Math.max(f[i][j], f[i][k] + f[i][j - k]); // 垂直切割
                for (var k = 1; k <= i / 2; k++) f[i][j] = Math.max(f[i][j], f[k][j] + f[i - k][j]); // 水平切割
            }
        return f[m][n];
    }
}
```

```C++ [sol2-C++]
class Solution {
public:
    long long sellingWood(int m, int n, vector<vector<int>> &prices) {
        long f[m + 1][n + 1]; memset(f, 0, sizeof(f));
        for (auto &p : prices) f[p[0]][p[1]] = p[2];
        for (int i = 1; i <= m; i++)
            for (int j = 1; j <= n; j++) {
                for (int k = 1; k <= j / 2; k++) f[i][j] = max(f[i][j], f[i][k] + f[i][j - k]); // 垂直切割
                for (int k = 1; k <= i / 2; k++) f[i][j] = max(f[i][j], f[k][j] + f[i - k][j]); // 水平切割
            }
        return f[m][n];
    }
};
```

```go [sol2-Go]
func sellingWood(m, n int, prices [][]int) int64 {
	f := make([][]int64, m+1)
	for i := range f {
		f[i] = make([]int64, n+1)
	}
	for _, price := range prices {
		f[price[0]][price[1]] = int64(price[2])
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k <= j/2; k++ { // 垂直切割
				f[i][j] = max(f[i][j], f[i][k]+f[i][j-k])
			}
			for k := 1; k <= i/2; k++ { // 水平切割
				f[i][j] = max(f[i][j], f[k][j]+f[i-k][j])
			}
		}
	}
	return f[m][n]
}

func max(a, b int64) int64 { if b > a { return b }; return a }
```



