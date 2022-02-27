由于答案不可能超过让最慢的车跑 $\textit{totalTrips}$ 趟所花费的时间，我们可以将其作为二分的上界。

二分答案 $x$，则我们可以完成 $\sum\limits_{i=0}^{n-1} \Big\lfloor\dfrac{x}{\textit{time}[i]}\Big\rfloor$ 趟旅途，将其和 $\textit{totalTrips}$ 比较来做二分。

Python3.10 新增带 `key` 的二分，从而可以一行解决此题。

```Python [sol1-Python3]
class Solution:
    def minimumTime(self, time: List[int], totalTrips: int) -> int:
        return bisect_left(range(totalTrips * max(time)), totalTrips, key=lambda x: sum(x // t for t in time))
```

```go [sol1-Go]
func minimumTime(time []int, totalTrips int) int64 {
	return int64(sort.Search(totalTrips*1e7, func(x int) bool {
		tot := 0
		for _, t := range time {
			tot += x / t
			if tot >= totalTrips {
				return true
			}
		}
		return false
	}))
}
```
