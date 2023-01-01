遍历每个数位，判断能否整除 $\textit{num}$。

```py [sol1-Python3]
class Solution:
    def countDigits(self, num: int) -> int:
        ans, x = 0, num
        while x:
            ans += num % (x % 10) == 0
            x //= 10
        return ans
```

```go [sol1-Go]
func countDigits(num int) (ans int) {
	for x := num; x > 0; x /= 10 {
		if num%(x%10) == 0 {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(\log \textit{num})$。
- 空间复杂度：$O(1)$，仅用到若干变量。
