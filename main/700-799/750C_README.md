设第一场比赛前的初始 rating 为 $r$。

设从开始到第 $i$ 场比赛之前（注意是比赛前），rating 的累计变化量为 $s$。

- 如果第 $i$ 场比赛是 div.1，那么有 $r+s \ge 1900$，即 $r \ge 1900-s$。
- 如果第 $i$ 场比赛是 div.2，那么有 $r+s \le 1899$，即 $r \le 1899-s$。

这样可以得到 $n$ 个不等式。

设 $r$ 的最小值为 $\textit{minR}$（初始值为 $-\infty$），最大值为 $\textit{maxR}$（初始值为 $\infty$）。

遍历这 $n$ 个不等式，同时维护 $\textit{minR}$ 的最大值，和 $\textit{maxR}$ 的最小值。

遍历结束后：

- 如果 $\textit{minR} > \textit{maxR}$，矛盾，输出 $\texttt{Impossible}$。
- 否则如果 $\textit{maxR}=\infty$，输出 $\texttt{Infinity}$。
- 否则输出 $\textit{maxR}+s$，表示初始 rating 的最大值，加上 $n$ 场比赛的 rating 累计变化量，就是最终 rating 的最大值。

```go
package main
import("bufio";."fmt";"os")
func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, c, d, s int
	minR, maxR := int(-1e9), int(1e9)
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &c, &d)
		if d == 1 {
			minR = max(minR, 1900-s)
		} else {
			maxR = min(maxR, 1899-s)
		}
		s += c
	}
	if minR > maxR {
		Print("Impossible")
	} else if maxR == 1e9 {
		Print("Infinity")
	} else {
		Print(maxR + s)
	}
}
```

- 时间复杂度：$\mathcal{O}(n)$。
