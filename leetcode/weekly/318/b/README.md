一个大小为 $k$ 的**滑动窗口**。

用一个哈希表 $\textit{cnt}$ 维护窗口内的元素个数，以及窗口内的元素和 $\textit{sum}$，如果 $\textit{cnt}$ 的大小等于 $k$，说明找到了 $k$ 个互不相同的元素，用此时的 $\textit{sum}$ 更新答案的最大值。

```py [sol1-Python3]
class Solution:
    def maximumSubarraySum(self, nums: List[int], k: int) -> int:
        ans = 0
        cnt = Counter(nums[:k - 1])
        s = sum(nums[:k - 1])
        for in_, out in zip(nums[k - 1:], nums):
            cnt[in_] += 1  # 移入元素
            s += in_
            if len(cnt) == k:
                ans = max(ans, s)
            cnt[out] -= 1  # 移出元素
            if cnt[out] == 0:
                del cnt[out]  # 重要：及时移除个数为 0 的数据
            s -= out
        return ans
```

```go [sol1-Go]
func maximumSubarraySum(nums []int, k int) int64 {
	ans, sum := 0, 0
	cnt := map[int]int{}
	for _, x := range nums[:k-1] {
		cnt[x]++
		sum += x
	}
	for i := k - 1; i < len(nums); i++ {
		cnt[nums[i]]++ // 移入元素
		sum += nums[i]
		if len(cnt) == k && sum > ans {
			ans = sum
		}
		x := nums[i+1-k]
		cnt[x]-- // 移出元素
		if cnt[x] == 0 {
			delete(cnt, x) // 重要：及时移除个数为 0 的数据
		}
		sum -= x
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(k)$。
