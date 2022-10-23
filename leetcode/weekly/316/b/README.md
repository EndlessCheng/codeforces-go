暴力枚举即可。下午两点在B站会讲到另外一种复杂度更优的做法。

```py [sol1-Python3]
class Solution:
    def subarrayGCD(self, nums: List[int], k: int) -> int:
        ans = 0
        for i in range(len(nums)):
            g = 0
            for j in range(i, len(nums)):
                g = gcd(g, nums[j])
                if g < k: break
                if g == k: ans += 1
        return ans
```

```go [sol1-Go]
func subarrayGCD(nums []int, k int) (ans int) {
	for i := range nums {
		g := 0
		for _, x := range nums[i:] {
			g = gcd(g, x)
			if g < k {
				break
			}
			if g == k {
				ans++
			}
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$O(n^2\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=max(\textit{nums})$。
- 空间复杂度：$O(1)$，仅用到若干变量。
