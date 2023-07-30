下午两点[【b站@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，欢迎关注！

---

经典滑窗，视频讲解可以看[【基础算法精讲 01】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

```py [sol-Python3]
class Solution:
    def countCompleteSubarrays(self, nums: List[int]) -> int:
        m = len(set(nums))
        cnt = Counter()
        ans = left = 0
        for v in nums:  # 枚举子数组右端点 v=nums[i]
            ans += left  # 子数组左端点 < left 的都是合法的
            cnt[v] += 1
            while len(cnt) == m:
                ans += 1  # 子数组左端点等于 left 是合法的
                x = nums[left]
                cnt[x] -= 1
                if cnt[x] == 0:
                    del cnt[x]
                left += 1
        return ans
```

```go [sol-Go]
func countCompleteSubarrays(nums []int) (ans int) {
	set := map[int]struct{}{}
	for _, v := range nums {
		set[v] = struct{}{}
	}
	m := len(set)

	cnt := map[int]int{}
	left := 0
	for _, v := range nums { // 枚举子数组右端点 v=nums[i]
		ans += left // 子数组左端点 < left 的都是合法的
		cnt[v]++
		for len(cnt) == m {
			ans++ // 子数组左端点等于 left 是合法的
			x := nums[left]
			cnt[x]--
			if cnt[x] == 0 {
				delete(cnt, x)
			}
			left++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

#### 相似题目

- [992. K 个不同整数的子数组](https://leetcode.cn/problems/subarrays-with-k-different-integers/)
