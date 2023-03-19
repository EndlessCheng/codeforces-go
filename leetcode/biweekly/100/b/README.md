下午两点【biIibiIi@灵茶山艾府】直播讲题，记得关注哦~

---

### 提示 1

田忌赛马。

### 提示 2-1

想一想，$\textit{nums}$ 的最小值是否要参与贡献伟大值？要和谁匹配？

### 提示 2-2

$\textit{nums}$ 的最小值要参与匹配，否则更大的数字更难匹配上。

$\textit{nums}$ 的最小值要与次小值匹配，这样后面的数字才能取匹配更大的数。

### 提示 3

为了方便实现，对 $\textit{nums}$ 从小到大排序。（为什么可以排序？因为只在乎匹配关系，与下标无关。）

例如示例 1 排序后为 $[1,1,1,2,3,3,5]$。那么前三个 $1$ 分别与 $2,3,3$ 匹配，$2$ 与 $5$ 匹配，后面就没有数字能匹配了。

```py [sol1-Python3]
class Solution:
    def maximizeGreatness(self, nums: List[int]) -> int:
        nums.sort()
        i = 0
        for x in nums:
            if x > nums[i]:
                i += 1
        return i
```

```go [sol1-Go]
func maximizeGreatness(nums []int) (i int) {
	sort.Ints(nums)
	for _, x := range nums {
		if x > nums[i] {
			i++
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$。忽略排序时的栈开销，仅用到若干额外变量。
