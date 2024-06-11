[视频讲解](https://www.bilibili.com/video/BV1AM4y1x7r4/)

首先，一旦有超过 $n$ 的数，直接返回 `false`。

那么剩下的数都是不超过 $n$ 的。

对于小于 $n$ 的数，不能超过一次；对于等于 $n$ 的数，不能超过两次。

在这些约束下，对于小于 $n$ 的数不可能有出现 $0$ 次的（否则某个数的出现次数会超过限制）；同理对于 $n$ 不可能只出现一次（否则某个小于 $n$ 的数会出现超过一次）。

```py [sol-Python3]
class Solution:
    def isGood(self, nums: List[int]) -> bool:
        n = len(nums) - 1
        cnt = [0] * (n + 1)
        for v in nums:
            if v > n or v == n and cnt[v] > 1 or v < n and cnt[v]:
                return False
            cnt[v] += 1
        return True
```

```go [sol-Go]
func isGood(nums []int) bool {
	n := len(nums) - 1
	cnt := make([]int, n+1)
	for _, v := range nums {
		if v > n || v == n && cnt[v] > 1 || v < n && cnt[v] > 0 {
			return false
		}
		cnt[v]++
	}
	return true
}
```

其余写法：

```py
class Solution:
    def isGood(self, nums: List[int]) -> bool:
        n = len(nums) - 1
        cnt = Counter(nums)
        return cnt[n] == 2 and all(cnt[i] == 1 for i in range(1, n))
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

