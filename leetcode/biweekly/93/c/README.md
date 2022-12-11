[视频讲解](https://www.bilibili.com/video/BV1kR4y1r7Df/) 已出炉，欢迎点赞三连，在评论区分享你对这场双周赛的看法~

---

首先每个石头都需要选，这样可以尽量让最大距离最小。

如何跳呢？比如四块石头 $a-b-c-d$：

- 不采用间隔跳，路径为 $a-d-c-b-a$；
- 采用间隔跳，路径为 $a-c-d-b-a$；

由于 $a-d$ 必然比 $a-c$ 和 $b-d$ 大，从这点可以看出，间隔跳是最优的。

该结论可以推广至更多的石头的情况。

```py [sol1-Python3]
class Solution:
    def maxJump(self, stones: List[int]) -> int:
        ans = stones[1] - stones[0]
        for i in range(2, len(stones)):
            ans = max(ans, stones[i] - stones[i - 2])
        return ans
```

```go [sol1-Go]
func maxJump(stones []int) int {
	ans := stones[1] - stones[0]
	for i := 2; i < len(stones); i++ {
		ans = max(ans, stones[i]-stones[i-2])
	}
	return ans
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{stones}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干额外变量。
