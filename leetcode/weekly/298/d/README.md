## 寻找子问题

看示例 1，对于一个高为 $3$ 宽为 $5$ 的木块，**第一步**一共有 $6$ 种切割方案：

- 竖着切开，有 $4$ 种切法。
- 横着切开，有 $2$ 种切法。

比如横着切开，第一步可以分成一个高为 $2$ 宽为 $5$ 的木块，和一个高为 $1$ 宽为 $5$ 的木块。

这俩都是更小的木块，可以分别处理，接着切割（比如第一个横切，第二个竖切），这意味着我们要处理的问题都是「高为 $i$ 宽为 $j$ 的木块」。

## 状态定义

定义 $f[i][j]$ 表示切割一块高 $i$ 宽 $j$ 的木块，能得到的最多钱数。

分类讨论：

- 如果直接售卖，则收益为对应的 $\textit{price}$（如果存在的话）。
- 如果竖着切开，枚举切割位置（宽度）$k$，得到两个高为 $i$，宽分别为 $k$ 和 $j-k$ 的木块，最大收益为
    $$
    \max_{k=1}^{j-1} f[i][k]+f[i][j-k]
    $$
- 如果横着切开，枚举切割位置（高度）$k$，得到两个宽为 $j$，高分别为 $k$ 和 $i-k$ 的木块，最大收益为
    $$
    \max_{k=1}^{i-1} f[k][j]+f[i-k][j]
    $$

取上述三种情况的最大值，即为 $f[i][j]$。

答案：$f[m][n]$。

代码实现时，为了方便查询木块价格，可以用一个哈希表或数组记录高宽对应的木块价格。当然，这个做法是不必要的，后面会优化。

```Python [sol-Python3]
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

```java [sol-Java]
class Solution {
    public long sellingWood(int m, int n, int[][] prices) {
        int[][] pr = new int[m + 1][n + 1];
        for (int[] p : prices) {
            pr[p[0]][p[1]] = p[2];
        }

        long[][] f = new long[m + 1][n + 1];
        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                f[i][j] = pr[i][j];
                for (int k = 1; k < j; k++) f[i][j] = Math.max(f[i][j], f[i][k] + f[i][j - k]); // 垂直切割
                for (int k = 1; k < i; k++) f[i][j] = Math.max(f[i][j], f[k][j] + f[i - k][j]); // 水平切割
            }
        }
        return f[m][n];
    }
}
```

```C++ [sol-C++]
class Solution {
public:
    long long sellingWood(int m, int n, vector<vector<int>> &prices) {
        vector<vector<int>> pr(m + 1, vector<int>(n + 1));
        for (auto &p: prices) {
            pr[p[0]][p[1]] = p[2];
        }

        vector<vector<long long>> f(m + 1, vector<long long>(n + 1));
        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                f[i][j] = pr[i][j];
                for (int k = 1; k < j; k++) f[i][j] = max(f[i][j], f[i][k] + f[i][j - k]); // 垂直切割
                for (int k = 1; k < i; k++) f[i][j] = max(f[i][j], f[k][j] + f[i - k][j]); // 水平切割
            }
        }
        return f[m][n];
    }
};
```

```go [sol-Go]
func sellingWood(m, n int, prices [][]int) int64 {
	pr := make([][]int, m+1)
	for i := range pr {
		pr[i] = make([]int, n+1)
	}
	for _, price := range prices {
		pr[price[0]][price[1]] = price[2]
	}

	f := make([][]int64, m+1)
	for i := 1; i <= m; i++ {
		f[i] = make([]int64, n+1)
		for j := 1; j <= n; j++ {
			f[i][j] = int64(pr[i][j])
			for k := 1; k < j; k++ { // 垂直切割，枚举宽度 k
				f[i][j] = max(f[i][j], f[i][k]+f[i][j-k])
			}
			for k := 1; k < i; k++ { // 水平切割，枚举高度 k
				f[i][j] = max(f[i][j], f[k][j]+f[i-k][j])
			}
		}
	}
	return f[m][n]
}
```

```js [sol-JavaScript]
var sellingWood = function(m, n, prices) {
    const pr = Array.from({length: m + 1}, () => Array(n + 1).fill(0));
    for (const [w, h, p] of prices) {
        pr[w][h] = p;
    }

    const f = Array.from({length: m + 1}, () => Array(n + 1).fill(0));
    for (let i = 1; i <= m; i++) {
        for (let j = 1; j <= n; j++) {
            f[i][j] = pr[i][j];
            for (let k = 1; k < j; k++) f[i][j] = Math.max(f[i][j], f[i][k] + f[i][j - k]); // 垂直切割
            for (let k = 1; k < i; k++) f[i][j] = Math.max(f[i][j], f[k][j] + f[i - k][j]); // 水平切割
        }
    }
    return f[m][n];
};
```

```rust [sol-Rust]
impl Solution {
    pub fn selling_wood(m: i32, n: i32, prices: Vec<Vec<i32>>) -> i64 {
        let m = m as usize;
        let n = n as usize;
        let mut pr = vec![vec![0; n + 1]; m + 1];
        for p in &prices {
            pr[p[0] as usize][p[1] as usize] = p[2];
        }

        let mut f = vec![vec![0; n + 1]; m + 1];
        for i in 1..=m {
            for j in 1..=n {
                f[i][j] = pr[i][j] as i64;
                for k in 1..j { // 垂直切割，枚举宽度 k
                    f[i][j] = f[i][j].max(f[i][k] + f[i][j - k]);
                }
                for k in 1..i { // 水平切割，枚举高度 k
                    f[i][j] = f[i][j].max(f[k][j] + f[i - k][j]);
                }
            }
        }
        f[m][n]
    }
}
```

## 优化

回顾一下，对于一个高为 $3$ 宽为 $5$ 的木块，**第一步**一共有 $6$ 种切割方案：

- 竖着切开，有 $4$ 种切法。
- 横着切开，有 $2$ 种切法。

但实际上，横着切开，虽然位置不同，但得到的结果是相同的，即一个高为 $2$ 宽为 $5$ 的木块，和一个高为 $1$ 宽为 $5$ 的木块，所以本质上只有 $1$ 种切法。对于竖切也同理，本质上只有 $2$ 种切法。

因此，枚举 $k$ 的时候，只需要枚举到一半的位置，宽度至多为 $\left\lfloor\dfrac{j}{2}\right\rfloor$，高度至多为 $\left\lfloor\dfrac{i}{2}\right\rfloor$。

此外，在计算递推之前，可以直接将 $\textit{prices}$ 记录到 $f$ 数组中。

```Python [sol-Python3]
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

```java [sol-Java]
class Solution {
    public long sellingWood(int m, int n, int[][] prices) {
        long[][] f = new long[m + 1][n + 1];
        for (int[] p : prices) {
            f[p[0]][p[1]] = p[2];
        }
        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                for (int k = 1; k <= j / 2; k++) f[i][j] = Math.max(f[i][j], f[i][k] + f[i][j - k]); // 垂直切割
                for (int k = 1; k <= i / 2; k++) f[i][j] = Math.max(f[i][j], f[k][j] + f[i - k][j]); // 水平切割
            }
        }
        return f[m][n];
    }
}
```

```C++ [sol-C++]
class Solution {
public:
    long long sellingWood(int m, int n, vector<vector<int>> &prices) {
        vector<vector<long long>> f(m + 1, vector<long long>(n + 1));
        for (auto &p : prices) {
            f[p[0]][p[1]] = p[2];
        }
        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                for (int k = 1; k <= j / 2; k++) f[i][j] = max(f[i][j], f[i][k] + f[i][j - k]); // 垂直切割
                for (int k = 1; k <= i / 2; k++) f[i][j] = max(f[i][j], f[k][j] + f[i - k][j]); // 水平切割
            }
        }
        return f[m][n];
    }
};
```

```go [sol-Go]
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
			for k := 1; k <= j/2; k++ { // 垂直切割，枚举宽度 k
				f[i][j] = max(f[i][j], f[i][k]+f[i][j-k])
			}
			for k := 1; k <= i/2; k++ { // 水平切割，枚举高度 k
				f[i][j] = max(f[i][j], f[k][j]+f[i-k][j])
			}
		}
	}
	return f[m][n]
}
```

```js [sol-JavaScript]
var sellingWood = function(m, n, prices) {
    const f = Array.from({length: m + 1}, () => Array(n + 1).fill(0));
    for (const [w, h, p] of prices) {
        f[w][h] = p;
    }
    for (let i = 1; i <= m; i++) {
        for (let j = 1; j <= n; j++) {
            for (let k = 1; k < j; k++) f[i][j] = Math.max(f[i][j], f[i][k] + f[i][j - k]); // 垂直切割
            for (let k = 1; k < i; k++) f[i][j] = Math.max(f[i][j], f[k][j] + f[i - k][j]); // 水平切割
        }
    }
    return f[m][n];
};
```

```rust [sol-Rust]
impl Solution {
    pub fn selling_wood(m: i32, n: i32, prices: Vec<Vec<i32>>) -> i64 {
        let m = m as usize;
        let n = n as usize;
        let mut f = vec![vec![0; n + 1]; m + 1];
        for p in &prices {
            f[p[0] as usize][p[1] as usize] = p[2] as i64;
        }
        for i in 1..=m {
            for j in 1..=n {
                for k in 1..j { // 垂直切割，枚举宽度 k
                    f[i][j] = f[i][j].max(f[i][k] + f[i][j - k]);
                }
                for k in 1..i { // 水平切割，枚举高度 k
                    f[i][j] = f[i][j].max(f[k][j] + f[i - k][j]);
                }
            }
        }
        f[m][n]
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn(m+n))$。
- 空间复杂度：$\mathcal{O}(mn)$。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
