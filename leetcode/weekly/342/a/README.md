下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

```py [sol1-Python3]
class Solution:
    def findDelayedArrivalTime(self, arrivalTime: int, delayedTime: int) -> int:
        return (arrivalTime + delayedTime) % 24
```

```go [sol1-Go]
func findDelayedArrivalTime(arrivalTime, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
