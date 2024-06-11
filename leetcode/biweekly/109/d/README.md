把 $n$ 看成背包容量，$n_i^x$ 看成物品，本题就是一个 0-1 背包模板题，具体请看[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。如果这个视频对你有帮助，欢迎一键三连！

代码实现时，由于 $n=300,x=1$ 算出来的结果不超过 $64$ 位整数的最大值，所以可以在计算结束后再取模。

```py [sol-Python3]
MX_N, MX_X = 300, 5
f = [[1] + [0] * MX_N for _ in range(MX_X)]
for x in range(MX_X):
    for i in count(1):
        v = i ** (x + 1)
        if v > MX_N: break
        for s in range(MX_N, v - 1, -1):
            f[x][s] += f[x][s - v]

class Solution:
    def numberOfWays(self, n: int, x: int) -> int:
        return f[x - 1][n] % (10 ** 9 + 7)
```

```go [sol-Go]
func numberOfWays(n, x int) int {
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; pow(i, x) <= n; i++ {
		v := pow(i, x)
		for s := n; s >= v; s-- {
			f[s] += f[s-v]
		}
	}
	return f[n] % (1e9 + 7)
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x
		}
		x = x * x
	}
	return res
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2\log x)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

- [494. 目标和](https://leetcode.cn/problems/target-sum/)
- [879. 盈利计划](https://leetcode.cn/problems/profitable-schemes/)
