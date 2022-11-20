[视频讲解](https://www.bilibili.com/video/BV1A3411f7H3/) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

# 方法一：排序 + 分组统计

对于 $x$，设：

- 小于 $x$ 的数有 $a$ 个；
- 等于 $x$ 的数有 $b$ 个；
- 大于 $x$ 的数有 $c$ 个。

那么 $x$ 对答案的贡献是 $abc$。

累加所有贡献，得到答案。

代码实现时，通过排序可以快速求出 $a\ b\ c$。

```py [sol1-Python3]
class Solution:
    def unequalTriplets(self, nums: List[int]) -> int:
        nums.sort()
        ans = start = 0
        for i, (x, y) in enumerate(pairwise(nums)):
            if x != y:
                ans += start * (i - start + 1) * (len(nums) - 1 - i)
                start = i + 1
        return ans
```

```go [sol1-Go]
func unequalTriplets(nums []int) (ans int) {
	sort.Ints(nums)
	start, n := 0, len(nums)
	for i, x := range nums[:n-1] {
		if x != nums[i+1] {
			ans += start * (i - start + 1) * (n - 1 - i)
			start = i + 1
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$，忽略排序时的栈开销，仅用到若干变量。

# 方法二：利用对称性

由于元素的位置是不重要的，我们可以直接用哈希表统计，方法一中的 $a\ b\ c$ 重定义为：

- 在 $x$ 之前遍历过的数有 $a$ 个；
- （当前遍历的）等于 $x$ 的数有 $b$ 个；
- 在 $x$ 之后遍历过的数有 $c$ 个。

```py [sol2-Python3]
class Solution:
    def unequalTriplets(self, nums: List[int]) -> int:
        ans, a, c = 0, 0, len(nums)
        for b in Counter(nums).values():
            c -= b
            ans += a * b * c
            a += b
        return ans
```

```go [sol2-Go]
func unequalTriplets(nums []int) (ans int) {
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}
	a, c := 0, len(nums)
	for _, b := range cnt {
		c -= b
		ans += a * b * c
		a += b
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。
