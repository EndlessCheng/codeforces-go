package main

/*
从 O(n^2) 到 O(n)：单调栈+计算每个元素对答案的贡献

#### 方法一：暴力枚举所有子数组

写一个二重循环，外层循环枚举子数组的左边界，内层循环枚举子数组的右边界，同时维护当前子数组的最小值和最大值。

```go
func subArrayRanges(nums []int) (ans int64) {
	for i, num := range nums {
		min, max := num, num
		for _, v := range nums[i+1:] {
			if v < min {
				min = v
			} else if v > max {
				max = v
			}
			ans += int64(max - min)
		}
	}
	return
}
```

**复杂度分析**

- 时间复杂度：$O(n^2)$，其中 $n$ 是数组 $\textit{nums}$ 的长度。

- 空间复杂度：$O(1)$，我们只需要常数的空间保存若干变量。

#### 方法二：单调栈 + 计算每个元素对答案的贡献

不了解单调栈的同学可以看一下 496 题。

我们可以考虑每个元素作为最大值出现在了多少子数组中，以及作为最小值出现在了多少子数组中。

以最大值为例。我们可以求出 $\textit{nums}[i]$ 左侧**严格大于**它的最近元素位置 $\textit{left}[i]$，以及右侧**大于等于**它的元素位置 $\textit{right}[i]$。注意 $\textit{nums}$ 中可能有重复元素，所以这里右侧取大于等于，这样可以避免在有重复元素的情况下，重复统计相同的子数组。

设以 $\textit{nums}[i]$ 为最大值的子数组为 $\textit{nums}[l..r]$，则有

- $\textit{left}[i]<l\le i$
- $i\le r<\textit{right}[i]$

所以 $\textit{nums}[i]$ 可以作为最大值出现在

$$
(i-\textit{left}[i])\cdot (\textit{right}[i]-i)
$$

个子数组中，这对答案产生的贡献是

$$
(i-\textit{left}[i])\cdot(\textit{right}[i]-i)\cdot \textit{nums}[i]
$$

最小值的做法同理（注意贡献为负数）。

累加所有贡献即为答案。

```go

```

**复杂度分析**

- 时间复杂度：$O(n)$，其中 $n$ 是数组 $\textit{nums}$ 的长度。

- 空间复杂度：$O(n)$。

*/

// github.com/EndlessCheng/codeforces-go
func subArrayRanges2(nums []int) (ans int64) {
	for i, num := range nums {
		min, max := num, num
		for _, v := range nums[i+1:] {
			if v < min {
				min = v
			} else if v > max {
				max = v
			}
			ans += int64(max - min)
		}
	}
	return
}

// github.com/EndlessCheng/codeforces-go
func subArrayRanges(nums []int) int64 {
	n := len(nums)
	left := make([]int, n) // left[i] 为左侧严格大于 num[i] 的最近元素位置（不存在时为 -1）
	type pair struct{ v, i int }
	s := []pair{{2e9, -1}} // 哨兵
	for i, v := range nums {
		for s[len(s)-1].v <= v { s = s[:len(s)-1] }
		left[i] = s[len(s)-1].i
		s = append(s, pair{v, i})
	}

	right := make([]int, n) // right[i] 为右侧大于等于 num[i] 的最近元素位置（不存在时为 n）
	s = []pair{{2e9, n}}
	for i := n - 1; i >= 0; i-- {
		v := nums[i]
		for s[len(s)-1].v < v { s = s[:len(s)-1] }
		right[i] = s[len(s)-1].i
		s = append(s, pair{v, i})
	}

	ans := 0
	for i, v := range nums {
		ans += (i - left[i]) * (right[i] - i) * v
	}

	// 求左侧严格小于
	left = make([]int, n) // left[i] 为左侧严格小于 num[i] 的最近元素位置（不存在时为 -1）
	s = []pair{{-2e9, -1}}
	for i, v := range nums {
		for s[len(s)-1].v >= v { s = s[:len(s)-1] }
		left[i] = s[len(s)-1].i
		s = append(s, pair{v, i})
	}

	right = make([]int, n) // right[i] 为右侧小于等于 num[i] 的最近元素位置（不存在时为 n）
	s = []pair{{-2e9, n}}
	for i := n - 1; i >= 0; i-- {
		v := nums[i]
		for s[len(s)-1].v > v { s = s[:len(s)-1] }
		right[i] = s[len(s)-1].i
		s = append(s, pair{v, i})
	}

	for i, v := range nums {
		ans -= (i - left[i]) * (right[i] - i) * v
	}

	return int64(ans)
}
