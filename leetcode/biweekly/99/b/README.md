下午两点【biIibiIi@灵茶山艾府】直播讲题，记得关注哦~

---

设答案为 $f(n)$，那么 $f(n)$ 相当于在 $f(n-1)$ 的基础上多了 $4$ 组 $n-1$ 的格子。

![c2c93618d2a71c5c56054c276f7f89c.png](https://pic.leetcode.cn/1677945101-eulEPr-c2c93618d2a71c5c56054c276f7f89c.png)

所以 

$$
f(n) =
\begin{cases} 
1,&n=1\\
f(n-1) + 4(n-1),&n \ge 2
\end{cases}
$$

化简得

$$
f(n) = 1 + 4(1+2+\cdots n-1) = 1 + 2n(n-1)
$$

```py [sol1-Python3]
class Solution:
    def coloredCells(self, n: int) -> int:
        return 1 + 2 * n * (n - 1)
```

```go [sol1-Go]
func coloredCells(n int) int64 {
	return 1 + 2*int64(n)*int64(n-1)
}
```

### 复杂度分析

- 时间复杂度：$O(1)$。
- 空间复杂度：$O(1)$。仅用到若干变量。
