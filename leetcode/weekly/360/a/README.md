请看 [视频讲解](https://www.bilibili.com/video/BV1Em4y1T7Bq/)。

根据题意，向左走就是 $-1$，向右走就是 $1$。

最后的位置就等于一堆 $-1$ 和一堆 $1$ 相加。

由于加法满足**交换律**，所以我们可以先只考虑 L 和 R，然后考虑下划线。

用 R 的个数减去 L 的个数，得到 $x$。

- 如果 $x>0$，那么所有的下划线都应该变成 R；
- 如果 $x<0$，那么所有的下划线都应该变成 L；
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
