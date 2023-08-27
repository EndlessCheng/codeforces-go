下午两点[【b站@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，欢迎关注！

---

先只考虑 L 和 R，也就是用 R 的个数减去 L 的个数，得到当前位置 $x$。

- 如果 $x>0$，那么所有的下划线都应该向右走；
- 如果 $x<0$，那么所有的下划线都应该向左走；
- 如果 $x=0$，向左向右都可以。

```py [sol-Python3]
class Solution:
    def furthestDistanceFromOrigin(self, moves: str) -> int:
        return abs(moves.count('R') - moves.count('L')) + moves.count('_')
```

```go [sol-Go]
func furthestDistanceFromOrigin(moves string) int {
	return abs(strings.Count(moves, "R")-strings.Count(moves, "L")) + strings.Count(moves, "_")
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{moves}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
