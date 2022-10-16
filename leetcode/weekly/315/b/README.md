一行写法（由于 Python 在处理反转时调用了底层的 C 库，这种写法实际比不用字符串的写法还要快）：

```py
class Solution:
    def countDistinctIntegers(self, nums: List[int]) -> int:
        return len(set(nums) | set(int(str(x)[::-1]) for x in nums))
```

不用字符串 + 一次遍历：

```py [sol1-Python3]
class Solution:
    def countDistinctIntegers(self, nums: List[int]) -> int:
        s = set()
        for x in nums:
            s.add(x)
            rev = 0
            while x:
                rev = rev * 10 + x % 10
                x //= 10
            s.add(rev)
        return len(s)
```

```go [sol1-Go]
func countDistinctIntegers(nums []int) int {
	set := map[int]struct{}{}
	for _, x := range nums {
		set[x] = struct{}{}
		rev := 0
		for ; x > 0; x /= 10 {
			rev = rev*10 + x%10
		}
		set[rev] = struct{}{}
	}
	return len(set)
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=max(\textit{nums})$。
- 空间复杂度：$O(n)$。
