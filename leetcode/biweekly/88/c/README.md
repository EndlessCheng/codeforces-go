下午 2 点在 B 站直播讲周赛和双周赛的题目，[欢迎关注](https://space.bilibili.com/206214/dynamic)~

---

由于答案是一大堆数字的异或和，根据**贡献法**的思想，我们可以讨论每个数字在这一大堆数字中出现了多少次，对答案的贡献是多少。

设 $n$ 为 $\textit{nums}_1$ 的长度，$m$ 为 $\textit{nums}_2$ 的长度。

对于 $\textit{nums}_1[i]$，由于它要与 $\textit{nums}_2$ 的每个元素异或一次，因此它一共出现了 $m$ 次。由于一个元素异或它自己等于 $0$，因此如果 $m$ 是偶数，则 $\textit{nums}_1[i]$ 对答案的贡献是 $0$，否则是 $\textit{nums}_1[i]$。

对于 $\textit{nums}_2$ 的元素，分析同理。

```py [sol1-Python3]
class Solution:
    def xorAllNums(self, nums1: List[int], nums2: List[int]) -> int:
        ans = 0
        if len(nums2) % 2: ans ^= reduce(xor, nums1)
        if len(nums1) % 2: ans ^= reduce(xor, nums2)
        return ans
```

```go [sol1-Go]
func xorAllNums(nums1, nums2 []int) (ans int) {
	if len(nums2)%2 > 0 {
		for _, x := range nums1 {
			ans ^= x
		}
	}
	if len(nums1)%2 > 0 {
		for _, x := range nums2 {
			ans ^= x
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(n+m)$，其中 $n$ 为 $\textit{nums}_1$ 的长度，$m$ 为 $\textit{nums}_2$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干变量。
