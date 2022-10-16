一行写法（由于 Python 在处理反转时调用了底层的 C 库，这种写法实际比不用字符串的写法还要快）：

```py
class Solution:
    def sumOfNumberAndReverse(self, num: int) -> bool:
        return any(i + int(str(i)[::-1]) == num for i in range(0, num + 1))
```

不用字符串的写法：

```py [sol1-Python3]
class Solution:
    def sumOfNumberAndReverse(self, num: int) -> bool:
        for i in range(num + 1):
            rev, x = 0, i
            while x:
                rev = rev * 10 + x % 10
                x //= 10
            if i + rev == num:
                return True
        return False
```

```go [sol1-Go]
func sumOfNumberAndReverse(num int) bool {
	for i := 0; i <= num; i++ {
		rev := 0
		for x := i; x > 0; x /= 10 {
			rev = rev*10 + x%10
		}
		if i+rev == num {
			return true
		}
	}
	return false
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n=\textit{num}$。
- 空间复杂度：$O(1)$，仅用到若干变量。
