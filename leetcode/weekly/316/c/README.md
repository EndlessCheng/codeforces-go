[视频讲解](https://www.bilibili.com/video/BV1ne4y1e7nu) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

## 方法一：枚举 + 考察变化量

将 $\textit{nums}$ 和 $\textit{cost}$ 绑在一起排序。

首先计算使所有元素都等于 $\textit{nums}[0]$ 的总开销 $\textit{total}$，以及所有 $\textit{cost}$ 的和 $\textit{sumCost}$。

然后考虑要使所有元素都等于 $\textit{nums}[1]$，$\textit{total}$ 的**变化量**是多少：

- 有 $\textit{cost}[0]$ 这么多的开销要增加 $\textit{nums}[1]-\textit{nums}[0]$；
- 有 $\textit{sumCost}-\textit{cost}[0]$ 这么多的开销要减少 $\textit{nums}[1]-\textit{nums}[0]$。

因此 $\textit{total}$ 减少了

$$
(\textit{sumCost} - 2 \cdot \textit{cost}[0]) \cdot (\textit{nums}[1]-\textit{nums}[0])
$$

按照这个公式模拟后续 $\textit{nums}[i]$，取所有 $\textit{total}$ 最小值为答案。

```py [sol1-Python3]
class Solution:
    def minCost(self, nums: List[int], cost: List[int]) -> int:
        a = sorted(zip(nums, cost))
        ans = total = sum((x - a[0][0]) * c for x, c in a)
        sum_cost = sum(cost)
        for (x0, c), (x1, _) in pairwise(a):
            sum_cost -= c * 2
            total -= sum_cost * (x1 - x0)
            ans = min(ans, total)
        return ans
```

```go [sol1-Go]
func minCost(nums, cost []int) int64 {
	type pair struct{ x, c int }
	a := make([]pair, len(nums))
	for i, x := range nums {
		a[i] = pair{x, cost[i]}
	}
	slices.SortFunc(a, func(p, q pair) int { return p.x - q.x })

	var total, sumCost int64
	for _, p := range a {
		total += int64(p.c) * int64(p.x-a[0].x)
		sumCost += int64(p.c)
	}
	ans := total
	for i := 1; i < len(a); i++ {
		sumCost -= int64(a[i-1].c * 2)
		total -= sumCost * int64(a[i].x-a[i-1].x)
		ans = min(ans, total)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。

## 方法二：中位数贪心

把 $\textit{cost}[i]$ 理解成 $\textit{nums}[i]$ 的出现次数。

根据 [中位数贪心及其证明](https://leetcode.cn/problems/5TxKeK/solution/zhuan-huan-zhong-wei-shu-tan-xin-dui-din-7r9b/)，把所有数变成中位数是最优的。

代码实现时，仍然按照方法一那样排序，然后不断累加 $\textit{cost}[i]$，首次累加到 $\ge\dfrac{\textit{sumCost}}{2}$ 时就找到了中位数。

由于 $\textit{sumCost}$ 可能是奇数，所以要上取整，即首次累加到 $\ge\left\lceil\dfrac{\textit{sumCost}}{2}\right\rceil$ 时就找到了中位数。

```py [sol2-Python3]
class Solution:
    def minCost(self, nums: List[int], cost: List[int]) -> int:
        a = sorted(zip(nums, cost))
        s, mid = 0, (sum(cost) + 1) // 2
        for x, c in a:
            s += c
            if s >= mid:
                return sum(abs(y - x) * c for y, c in a)  # 把所有数都变成 x
```

```go [sol2-Go]
func minCost(nums, cost []int) (ans int64) {
	type pair struct{ x, c int }
	a := make([]pair, len(nums))
	sumCost := int64(0)
	for i, c := range cost {
		a[i] = pair{nums[i], c}
		sumCost += int64(c)
	}
	slices.SortFunc(a, func(p, q pair) int { return p.x - q.x })

	s, mid := int64(0), (sumCost+1)/2
	for _, p := range a {
		s += int64(p.c)
		if s >= mid {
			// 把所有数都变成 p.x
			for _, q := range a {
				ans += int64(abs(q.x-p.x)) * int64(q.c)
			}
			break
		}
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。
