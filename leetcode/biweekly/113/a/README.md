[视频讲解](https://www.bilibili.com/video/BV1PV411N76R/)

## 方法一：暴力

由于右移 $n$ 次就变回原数组了，所以答案至多为 $n-1$。

我们可以不断右移，每次右移前，先判断数组是否有序，如果有序就直接返回右移次数，否则就继续右移。

如果循环中途没有返回，最后返回 $-1$，表示无法得到递增数组。

```py [sol-Python3]
class Solution:
    def minimumRightShifts(self, nums: List[int]) -> int:
        for i in range(len(nums)):
            if all(x < y for x, y in pairwise(nums)):
                return i
            nums = [nums[-1]] + nums[:-1]
        return -1
```

```go [sol-Go]
func minimumRightShifts2(a []int) int {
	for i := 0; i < len(a); i++ {
		if sort.IntsAreSorted(a) {
			return i
		}
		a = append(a[len(a)-1:], a[:len(a)-1]...)
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：至多两段递增子数组

问题等价于：

- $\textit{nums}$ 有至多两段递增子数组
- 如果有两段，我们需要把第二段拼在第一段前面，所以还需要满足 $\textit{nums}[0] > \textit{nums}[n-1]$，否则无法构成递增数组。

算法：

1. 先循环找到第一段。
2. 如果第一段的长度就是 $n$，返回 $0$。
3. 否则至少有两段，如果 $\textit{nums}[0] < \textit{nums}[n-1]$ 就返回 $-1$。
4. 把第二段的起始下标记作 $\textit{mid}$。
5. 循环找第二段。
6. 如果超过两段，那么返回 $-1$。
7. 否则返回 $n-\textit{mid}$，因为第二段的长度就是右移次数。

```py [sol-Python3]
class Solution:
    def minimumRightShifts(self, nums: List[int]) -> int:
        i, n = 1, len(nums)
        while i < n and nums[i - 1] < nums[i]:
            i += 1
        if i == n:
            return 0
        if nums[0] < nums[-1]:
            return -1
        mid = i
        i += 1
        while i < n and nums[i - 1] < nums[i]:
            i += 1
        if i < n:
            return -1
        return n - mid
```

```go [sol-Go]
func minimumRightShifts(a []int) int {
	i, n := 1, len(a)
	for i < n && a[i-1] < a[i] {
		i++
	}
	if i == n {
		return 0
	}
	if a[0] < a[n-1] {
		return -1
	}
	mid := i
	i++
	for i < n && a[i-1] < a[i] {
		i++
	}
	if i < n {
		return -1
	}
	return n - mid
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。
