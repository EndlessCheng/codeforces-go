下午两点【bilibili@灵茶山艾府】直播讲题，记得关注哦~

---

### 前置知识：二分

见 [【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)，详细介绍了二分的原理及实现。

### 思路

看到「最大化最小值」或者「最小化最大值」就要想到**二分答案**，这是一个固定的套路。

为什么？一般来说，二分的值越大，越能/不能满足要求；二分的值越小，越不能/能满足要求，有单调性，可以二分。

类似的题目在先前的周赛中出现过多次，例如：

- [2439. 最小化数组中的最大值](https://leetcode.cn/problems/minimize-maximum-of-array/)
- [2513. 最小化两个数组中的最大值](https://leetcode.cn/problems/minimize-the-maximum-of-two-arrays/)
- [2517. 礼盒的最大甜蜜度](https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/)
- [2528. 最大化城市的最小供电站数目](https://leetcode.cn/problems/maximize-the-minimum-powered-city/)

然后就可以像 [198. 打家劫舍](https://leetcode.cn/problems/house-robber/) 那样写一个 DP 了：

设二分的最大金额为 $\textit{mx}$，定义 $f[i]$ 表示在前 $i$ 个房屋中窃取金额不超过 $\textit{mx}$ 的房屋的最大个数。

分类讨论：

- 不选第 $i$ 个房屋：$f[i] = f[i-1]$；
- 选第 $i$ 个房屋，前提是金额不超过 $\textit{mx}$：$f[i] = f[i-2]+1$。

这两取最大值，即

$$
f[i] = \max(f[i-1], f[i-2] + 1)
$$

代码实现时，可以用两个变量滚动计算。

```py [sol1-Python3]
class Solution:
    def minCapability(self, nums: List[int], k: int) -> int:
        def check(mx: int) -> int:
            f0 = f1 = 0
            for x in nums:
                if x > mx: f0 = f1
                else: f0, f1 = f1, max(f1, f0 + 1)
            return f1
        return bisect_left(range(max(nums)), k, key=check)
```

```go [sol1-Go]
func minCapability(nums []int, k int) int {
	return sort.Search(1e9, func(mx int) bool {
		f0, f1 := 0, 0
		for _, x := range nums {
			if x <= mx {
				f0, f1 = f1, max(f1, f0+1)
			} else {
				f0 = f1
			}
		}
		return f1 >= k
	})
}

func max(a, b int) int { if b > a { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=max(\textit{nums})$。
- 空间复杂度：$O(1)$。
