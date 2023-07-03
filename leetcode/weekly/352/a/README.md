下午两点直播讲题，记得关注[【b站@灵茶山艾府】](https://space.bilibili.com/206214)哦~

---

题目的约束实际上把数组划分成了若干段，每段都满足要求，且互不相交。

那么遍历一遍，计算每一段的长度，取最大值，即为答案。

```py [sol-Python3]
class Solution:
    def longestAlternatingSubarray(self, a: List[int], threshold: int) -> int:
        ans, i, n = 0, 0, len(a)
        while i < n:
            if a[i] % 2 or a[i] > threshold:
                i += 1
            else:
                i0 = i
                i += 1
                while i < n and a[i] <= threshold and a[i] % 2 != a[i - 1] % 2:
                    i += 1  # i 是全局变量，二重循环 i+=1 至多执行 O(n) 次
                ans = max(ans, i - i0)
        return ans
```

```go [sol-Go]
func longestAlternatingSubarray(a []int, threshold int) (ans int) {
	for i, n := 0, len(a); i < n; {
		if a[i]%2 > 0 || a[i] > threshold {
			i++
		} else {
			i0 := i
			for i++; i < n && a[i] <= threshold && a[i]%2 != a[i-1]%2; i++ {}
			ans = max(ans, i-i0)
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。注意 $i$ 是全局变量，只会增加，不会减少。所以二重循环至多执行 $\mathcal{O}(n)$ 次。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
