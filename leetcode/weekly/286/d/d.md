#### 提示 1

对于一个栈，我们只能移除其前缀。

#### 提示 2

对每个栈求其前缀和 $\textit{sum}$，$\textit{sum}$ 的第 $j$ 个元素视作一个体积为 $j$，价值为 $\textit{sum}[j]$ 的物品。

问题转化成从 $n$ 个物品组里面取物品体积和为 $k$ 的物品时的物品价值最大和，即分组背包模型。

---

定义 $f[i][j]$ 表示从前 $i$ 个组取体积之和为 $j$ 的物品时，物品价值之和的最大值。

枚举第 $i$ 个组的所有物品，设当前物品体积为 $w$，价值为 $v$，则有

$$
f[i][j] = \max(f[i][j], f[i-1][j-w]+v)
$$

答案为 $f[n][k]$。

代码实现时，可以仿照 01 背包的写法，将第一维压缩掉。

**时间复杂度**：将外层循环与最内层循环合并，即每个栈的大小之和，记作 $\textit{s}$，算上中间这层的循环，时间复杂度即为 $O(ks)$。
**空间复杂度**：$O(k)$。

```Python [sol1-Python3]
class Solution:
    def maxValueOfCoins(self, piles: List[List[int]], k: int) -> int:
        f = [0] * (k + 1)
        sum_n = 0
        for pile in piles:
            n = len(pile)
            for i in range(1, n):
                pile[i] += pile[i - 1]  # pile 前缀和
            sum_n = min(sum_n + n, k)  # 优化：j 从前 i 个栈的大小之和开始枚举（不超过 k）
            for j in range(sum_n, 0, -1):
                f[j] = max(f[j], max(f[j - w - 1] + pile[w] for w in range(min(n, j))))  # w 从 0 开始，物品体积为 w+1
        return f[k]
```

```go [sol1-Go]
func maxValueOfCoins(piles [][]int, k int) int {
	f := make([]int, k+1)
	sumN := 0
	for _, pile := range piles {
		n := len(pile)
		for i := 1; i < n; i++ {
			pile[i] += pile[i-1] // pile 前缀和
		}
		sumN = min(sumN+n, k) // 优化：j 从前 i 个栈的大小之和开始枚举（不超过 k）
		for j := sumN; j > 0; j-- {
			for w, v := range pile[:min(n, j)] {
				f[j] = max(f[j], f[j-w-1]+v) // w 从 0 开始，物品体积为 w+1
			}
		}
	}
	return f[k]
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
```

```C++ [sol1-C++]
class Solution {
public:
    int maxValueOfCoins(vector<vector<int>> &piles, int k) {
        vector<int> f(k + 1);
        int sumN = 0;
        for (auto &pile: piles) {
            int n = pile.size();
            for (int i = 1; i < n; ++i)
                pile[i] += pile[i - 1]; // pile 前缀和
            sumN = min(sumN + n, k); // 优化：j 从前 i 个栈的大小之和开始枚举（不超过 k）
            for (int j = sumN; j; --j)
                for (int w = 0; w < min(n, j); ++w)
                    f[j] = max(f[j], f[j - w - 1] + pile[w]); // w 从 0 开始，物品体积为 w+1
        }
        return f[k];
    }
};
```

```java [sol1-Java]
class Solution {
    public int maxValueOfCoins(List<List<Integer>> piles, int k) {
        var f = new int[k + 1];
        var sumN = 0;
        for (var pile : piles) {
            var n = pile.size();
            for (var i = 1; i < n; ++i)
                pile.set(i, pile.get(i) + pile.get(i - 1)); // pile 前缀和
            sumN = Math.min(sumN + n, k); // 优化：j 从前 i 个栈的大小之和开始枚举（不超过 k）
            for (var j = sumN; j > 0; --j)
                for (var w = 0; w < Math.min(n, j); ++w)
                    f[j] = Math.max(f[j], f[j - w - 1] + pile.get(w)); // w 从 0 开始，物品体积为 w+1
        }
        return f[k];
    }
}
```
