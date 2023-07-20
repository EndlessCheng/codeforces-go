显然 $\textit{num}+2t$ 是可达成数字。但对于更大的数，就无法在 $t$ 次操作内和 $\textit{num}$ 相等了，所以答案为 $\textit{num}+2t$。

```py [sol-Python3]
class Solution:
    def theMaximumAchievableX(self, num: int, t: int) -> int:
        return num + t * 2
```

```go [sol-Go]
func theMaximumAchievableX(num, t int) int {
	return num + t*2
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。
