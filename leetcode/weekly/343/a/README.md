下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

```py [sol1-Python3]
def score(a: List[int]) -> int:
    res = 0
    for i, x in enumerate(a):
        if i and a[i - 1] == 10 or i > 1 and a[i - 2] == 10:
            x *= 2
        res += x
    return res

class Solution:
    def isWinner(self, player1: List[int], player2: List[int]) -> int:
        s1, s2 = score(player1), score(player2)
        return 1 if s1 > s2 else 2 if s1 < s2 else 0
```

```go [sol1-Go]
func score(a []int) (res int) {
	for i, x := range a {
		if i > 0 && a[i-1] == 10 || i > 1 && a[i-2] == 10 {
			x *= 2
		}
		res += x
	}
	return
}

func isWinner(player1, player2 []int) int {
	s1, s2 := score(player1), score(player2)
	if s1 > s2 {
		return 1
	}
	if s1 < s2 {
		return 2
	}
	return 0
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{player}_1$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
