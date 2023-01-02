[视频讲解](https://www.bilibili.com/video/BV1FV4y1F7v7/) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

由于随着甜蜜度的增大，能选择的糖果数量变小，有单调性，所以可以用二分答案来做。

对 $\textit{price}$ 从小到大排序，二分答案 $d$。最小的数 $x$ 可以选，下一个可以选的数是第一个 $\ge x+d$ 的数，依此类推。

如果可以选的数 $< k$，说明 $d$ 取大了，否则说明 $d$ 取小了，根据这一点来二分。

二分上界可以取 $\left\lfloor\dfrac{\max(\textit{price})-\min(\textit{price})}{k-1}\right\rfloor$。这是因为**最小值不会超过平均值**。

有关二分的写法，可以看我的 [【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/) 这期视频。

```py [sol1-Python3]
class Solution:
    def maximumTastiness(self, price: List[int], k: int) -> int:
        price.sort()
        def check(d: int) -> bool:
            cnt, x0 = 1, price[0]
            for x in price:
                if x >= x0 + d:
                    cnt += 1
                    x0 = x
            return cnt >= k

        left, right = 0, (price[-1] - price[0]) // (k - 1) + 1  # 开区间
        while left + 1 < right:
            mid = (left + right) // 2
            if check(mid): left = mid
            else: right = mid
        return left
```

```go [sol1-Go]
func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	return sort.Search((price[len(price)-1]-price[0])/(k-1), func(d int) bool {
		d++
		cnt, x0 := 1, price[0]
		for _, x := range price[1:] {
			if x >= x0+d {
				cnt++
				x0 = x
			}
		}
		return cnt < k
	})
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log n + n\log U)$，其中 $n$ 为 $\textit{price}$ 的长度，$U=\max(\textit{price})$。
- 空间复杂度：$O(1)$，忽略排序的空间，仅用到若干额外变量。

#### 相同题目

- [1552. 两球之间的磁力](https://leetcode.cn/problems/magnetic-force-between-two-balls/)
