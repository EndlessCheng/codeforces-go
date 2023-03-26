下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

贪心，按照标记为 $1,0,-1$ 的顺序选。

```py [sol1-Python3]
class Solution:
    def kItemsWithMaximumSum(self, numOnes: int, numZeros: int, numNegOnes: int, k: int) -> int:
        if k <= numOnes + numZeros:
            return min(k, numOnes)
        return numOnes * 2 + numZeros - k
```

```go [sol1-Go]
func kItemsWithMaximumSum(numOnes, numZeros, _, k int) int {
	if k <= numOnes {
		return k
	}
	if k <= numOnes+numZeros {
		return numOnes
	}
	return numOnes*2 + numZeros - k
}
```

### 复杂度分析

- 时间复杂度：$O(1)$。
- 空间复杂度：$O(1)$。仅用到若干额外变量。
