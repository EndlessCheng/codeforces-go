下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

定义 $f(x, \textit{left})$ 表示当前在 $x$，还剩 $\textit{left}$ 步时，走到终点的方案数。

枚举下一步往左或者往右，累加即为答案。

注意在递归过程中如果 $|x-\textit{endPos}| > \textit{left}$，可以直接返回。

```py [sol1-Python3]
class Solution:
    def numberOfWays(self, startPos: int, endPos: int, k: int) -> int:
        MOD = 10 ** 9 + 7
        @cache
        def f(x: int, left: int) -> int:
            if abs(x - endPos) > left: return 0
            if left == 0: return 1
            return (f(x - 1, left - 1) + f(x + 1, left - 1)) % MOD
        return f(startPos, k)
```

```go [sol1-Go]
func numberOfWays(startPos, endPos, k int) int {
	type pair struct{ x, y int }
	dp := map[pair]int{}
	var f func(int, int) int
	f = func(x, left int) int {
		if abs(x-endPos) > left {
			return 0
		}
		if left == 0 {
			return 1
		}
		p := pair{x, left}
		if v, ok := dp[p]; ok {
			return v
		}
		res := (f(x-1, left-1) + f(x+1, left-1)) % (1e9 + 7)
		dp[p] = res
		return res
	}
	return f(startPos, k)
}
func abs(x int) int { if x < 0 { return -x }; return x }
```
