[视频讲解](https://www.bilibili.com/video/BV1DM4y1x7bR/) 第三题。

**证明**：分割出的两个数组的支配元素就是原数组的支配元素。

设这两个数组的支配元素为 $y$（题目要求支配元素相同），那么对于第一个数组有

$$
\text{freq}_1(y) \cdot 2 > i+1
$$

对于第二个数组有

$$
\text{freq}_2(y) \cdot 2 > n-i-1
$$

由于这两个数组合并之后就是原数组，所以

$$
\text{freq}(y) \cdot 2 = \text{freq}_1(y) \cdot 2 + \text{freq}_2(y) \cdot 2 > (i+1) + (n-i-1) = n
$$

上式表明，$y$ 就是原数组的支配元素，证毕。

### 算法

首先求出 $\textit{nums}$ 的支配元素（绝对众数）$\textit{mode}$ 及其出现次数 $\textit{total}$。

然后枚举 $i$，一边枚举一边统计 $\text{freq}_1(\textit{mode})$，那么 $\text{freq}_2(\textit{mode}) =\textit{total} -\text{freq}_1(\textit{mode})$。

只要满足 $\text{freq}_1(\textit{mode}) \cdot 2 > i+1$ 且 $\text{freq}_2(\textit{mode}) \cdot 2 > n-i-1$，就返回 $i$。

如果没有这样的 $i$，返回 $-1$。

```py [sol-Python3]
class Solution:
    def minimumIndex(self, nums: List[int]) -> int:
        mode, total = Counter(nums).most_common(1)[0]
        freq1 = 0
        for i, x in enumerate(nums):
            freq1 += x == mode
            if freq1 * 2 > i + 1 and (total - freq1) * 2 > len(nums) - i - 1:
                return i
        return -1
```

```go [sol-Go]
func minimumIndex(nums []int) int {
	// 也可以用摩尔投票法实现
	freq := map[int]int{}
	mode := nums[0]
	for _, x := range nums {
		freq[x]++
		if freq[x] > freq[mode] {
			mode = x
		}
	}

	total := freq[mode]
	freq1 := 0
	for i, x := range nums {
		if x == mode {
			freq1++
		}
		if freq1*2 > i+1 && (total-freq1)*2 > len(nums)-i-1 {
			return i
		}
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。用摩尔投票法可以做到 $\mathcal{O}(1)$ 额外空间，具体见 [169. 多数元素](https://leetcode.cn/problems/majority-element/)。
