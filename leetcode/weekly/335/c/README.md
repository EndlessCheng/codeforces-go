下午两点【biIibiIi@灵茶山艾府】直播讲题，记得关注哦~

---

### 前置知识：分解质因数

见[【力扣周赛 324】](https://www.bilibili.com/video/BV1LW4y1T7if/) 第二题。

### 思路

对于一个质因子 $p$，设它在数组中的最左和最右的位置为 $\textit{left}$ 和 $\textit{right}$。

那么答案是不能在区间 $[\textit{left},\textit{right})$ 中的。注意区间右端点可能为答案。

因此这题本质上和 [55. 跳跃游戏](https://leetcode.cn/problems/jump-game/) 是类似的，找从 $0$ 出发，最远遇到的区间右端点，即为答案。

```py [sol1-Python3]
class Solution:
    def findValidSplit(self, nums: List[int]) -> int:
        n = len(nums)
        left = {}  # left[p] 表示质数 p 首次出现的下标
        right = [-1] * n  # right[i] 表示左端点为 i 的区间的右端点的最大值

        def f(p: int, i: int) -> None:
            if p in left:
                right[left[p]] = i  # 记录左端点 l 对应的右端点的最大值
            else:
                left[p] = i  # 第一次遇到质数 p

        for i, x in enumerate(nums):
            d = 2
            while d * d <= x:  # 分解质因数
                if x % d == 0:
                    f(d, i)
                    x //= d
                    while x % d == 0:
                        x //= d
                d += 1
            if x > 1: f(x, i)

        max_r = right[0]
        for l, r in enumerate(right):
            if l > max_r:  # 最远可以遇到 max_r
                return max_r  # 也可以写 l-1
            max_r = max(max_r, r)
        return -1
```

```go [sol1-Go]
func findValidSplit(nums []int) int {
	n := len(nums)
	left := map[int]int{}   // left[p] 表示质数 p 首次出现的下标
	right := make([]int, n) // right[i] 表示左端点为 i 的区间的右端点的最大值
	for i := range right {
		right[i] = -1
	}
	f := func(p, i int) {
		if l, ok := left[p]; ok {
			right[l] = i // 记录左端点 l 对应的右端点的最大值
		} else {
			left[p] = i // 第一次遇到质数 p
		}
	}
	for i, x := range nums {
		for d := 2; d*d <= x; d++ { // 分解质因数
			if x%d == 0 {
				f(d, i)
				for x /= d; x%d == 0; x /= d {
				}
			}
		}
		if x > 1 {
			f(x, i)
		}
	}

	maxR := right[0]
	for l, r := range right {
		if l > maxR { // 最远可以遇到 maxR
			return maxR // 也可以写 l-1
		}
		maxR = max(maxR, r)
	}
	return -1
}

func max(a, b int) int { if a < b { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n\sqrt U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$O\left(n + \dfrac{U}{\log U}\right)$。$U$ 范围内的质数个数有 $O\left(\dfrac{U}{\log U}\right)$ 个。

### 相似题目

- [55. 跳跃游戏](https://leetcode.cn/problems/jump-game/)
- [45. 跳跃游戏 II](https://leetcode.cn/problems/jump-game-ii/)
- [56. 合并区间](https://leetcode.cn/problems/merge-intervals/)
- [1024. 视频拼接](https://leetcode.cn/problems/video-stitching/)
- [1326. 灌溉花园的最少水龙头数目](https://leetcode.cn/problems/minimum-number-of-taps-to-open-to-water-a-garden/)
