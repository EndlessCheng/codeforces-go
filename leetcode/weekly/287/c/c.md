#### 提示 1

力扣周赛的套路：看到第三题 + 求最大/最小，就先往二分答案上想。

---

由于每堆的糖果越多，能分出来的子堆也就越少，因此答案可以二分。

二分的上界为 $\Big\lfloor\dfrac{\sum\textit{candies}}{k}\Big\rfloor$。

```Python [sol1-Python3]
class Solution:
    def maximumCandies(self, candies: List[int], k: int) -> int:
        # 由于 x 越大 sum 越小，因此取负号从而满足二分条件
        return bisect_right(range(sum(candies) // k), -k, key=lambda x: -sum(v // (x + 1) for v in candies))
```

```go [sol1-Go]
func maximumCandies(candies []int, k int64) int {
	return sort.Search(1e7, func(size int) bool {
		size++
		cnt := int64(0)
		for _, candy := range candies {
			cnt += int64(candy / size)
		}
		return cnt < k
	})
}
```
