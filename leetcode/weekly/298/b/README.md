本题 [视频讲解](https://www.bilibili.com/video/BV1CW4y1k7B3) 已出炉，欢迎点赞三连~

---

枚举答案 $n$。

把个位数单独拆开看，每个数可以表示成 $10$ 的倍数加上 $k$ 的形式。

由于这 $n$ 个数都以 $k$ 结尾，那么 $\textit{num}-nk$ 必须是 $10$ 的倍数。

从小到大枚举 $n$，找到第一个满足 $\textit{num}-nk$ 是 $10$ 的倍数的 $n$。

由于 $n$ 不会超过 $\textit{num}$，我们至多枚举到 $\textit{num}$ 时停止。

注意特判 $\textit{num}=0$ 的情况，此时返回 $0$。

```Python [sol1-Python3]
class Solution:
    def minimumNumbers(self, num: int, k: int) -> int:
        if num == 0: return 0
        for n in range(1, num + 1):
            if num - n * k < 0: break
            if (num - n * k) % 10 == 0: return n
        return -1
```

```go [sol1-Go]
func minimumNumbers(num, k int) int {
	if num == 0 {
		return 0
	}
	for n := 1; n <= num && num-n*k >= 0; n++ {
		if (num-n*k)%10 == 0 {
			return n
		}
	}
	return -1
}
```

特判 $k=0$ 的情况，还可以减少一部分循环次数：

```Python [sol2-Python3]
class Solution:
    def minimumNumbers(self, num: int, k: int) -> int:
        if num == 0: return 0
        if k == 0: return -1 if num % 10 else 1
        return next((n for n in range(1, num // k + 1) if (num - n * k) % 10 == 0), -1)
```

```go [sol2-Go]
func minimumNumbers(num, k int) int {
	if num == 0 {
		return 0
	}
	if k == 0 {
		if num%10 == 0 {
			return 1
		}
		return -1
	}
	for n := 1; n*k <= num; n++ {
		if (num-n*k)%10 == 0 {
			return n
		}
	}
	return -1
}
```

进一步地，由于

$$
n\cdot k\equiv(n\bmod 10)\cdot k \pmod{10}
$$

枚举到 $n=11$ 时，$(\textit{num}-nk)\bmod 10$ 的结果会和 $n=1$ 时相同，对于更大的 $n$ 也同样。

因此至多枚举到 $n=10$ 就行了。

#### 复杂度分析

- 时间复杂度：$O(1)$。枚举的次数至多为 $10$。
- 空间复杂度：$O(1)$，仅用到若干变量。

```Python [sol3-Python3]
class Solution:
    def minimumNumbers(self, num, k):
        if num == 0: return 0
        if k == 0: return -1 if num % 10 else 1
        return next((n for n in range(1, min(num // k + 1, 11)) if ((num - n * k) % 10 == 0)), -1)
```

```go [sol3-Go]
func minimumNumbers(num, k int) int {
	if num == 0 {
		return 0
	}
	for n := 1; n <= 10 && n*k <= num; n++ {
		if (num-n*k)%10 == 0 {
			return n
		}
	}
	return -1
}
```
