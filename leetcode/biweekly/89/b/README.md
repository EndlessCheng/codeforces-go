# 方法一：暴力

拆分 $n$，把二的幂次记到一个数组 $a$ 中，对每个询问，直接遍历 $a$。

```py [sol1-Python3]
MOD = 10 ** 9 + 7

class Solution:
    def productQueries(self, n: int, queries: List[List[int]]) -> List[int]:
        a = []
        while n:
            lb = n & -n
            a.append(lb)
            n ^= lb
        return [reduce(lambda x, y: x * y % MOD, a[l: r + 1]) for l, r in queries]
```

```go [sol1-Go]
const mod int = 1e9 + 7

func productQueries(n int, queries [][]int) []int {
	a := []int{}
	for n > 0 {
		lb := n & -n
		a = append(a, lb)
		n ^= lb
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		mul := 1
		for _, x := range a[q[0] : q[1]+1] {
			mul = mul * x % mod
		}
		ans[i] = mul
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$O(m\log n)$，其中 $m$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$O(\log n)$。返回值不计入。

# 方法二：预处理

注意到 $a$ 的长度很小，我们可以直接预处理所有询问。

```py [sol2-Python3]
MOD = 10 ** 9 + 7

class Solution:
    def productQueries(self, n: int, queries: List[List[int]]) -> List[int]:
        a = []
        while n:
            lb = n & -n
            a.append(lb)
            n ^= lb
        na = len(a)
        res = [[0] * na for _ in range(na)]
        for i, x in enumerate(a):
            res[i][i] = x
            for j in range(i + 1, na):
                res[i][j] = res[i][j - 1] * a[j] % MOD
        return [res[l][r] for l, r in queries]
```

```go [sol2-Go]
const mod int = 1e9 + 7

func productQueries(n int, queries [][]int) []int {
	a := []int{}
	for n > 0 {
		lb := n & -n
		a = append(a, lb)
		n ^= lb
	}
	na := len(a)
	res := make([][]int, na)
	for i, x := range a {
		res[i] = make([]int, na)
		res[i][i] = x
		for j := i + 1; j < na; j++ {
			res[i][j] = res[i][j-1] * a[j] % mod
		}
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = res[q[0]][q[1]]
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$O(m + \log^2 n)$，其中 $m$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$O(\log^2 n)$。返回值不计入。
