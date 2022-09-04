下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

上联：[return true](https://leetcode.cn/problems/stone-game/)

下联：[return false](https://leetcode.cn/problems/strictly-palindromic-number/)

横批：脑筋急转弯

---

在题目的条件下，答案一定为 `false`，证明如下：

根据带余除法，$n=qb+r$，其中 $0\le r < b$。

取 $b=n-2$，那么当 $n>4$ 时，上式的 $q=1$，$r=2$，因此当 $n>4$ 时，在 $b=n-2$ 进制下的数值为 $12$，不是回文数。

而对于 $n=4$，在 $b=2$ 进制下的数值为 $100$，也不是回文数。

因此直接返回 `false` 即可。

```py [sol1-Python3]
class Solution:
    def isStrictlyPalindromic(self, n: int) -> bool:
        return False
```

```go [sol1-Go]
func isStrictlyPalindromic(int) bool {
	return false
}
```
