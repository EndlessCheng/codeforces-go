**题意**：数轴上有 $n$ 个点，要求任意长为 $m-1$ 的闭区间不能包含 $\ge k$ 个点，至少要去掉多少个点？

排序后，用**不定长滑窗**做。

维护窗口内的点的个数 $\textit{cnt}$。枚举 $a_i$ 作为窗口右端点位置，如果窗口长度 $\ge m$ 则移动窗口左端点 $\textit{left}$。如果发现 $\textit{cnt}=k$，则优先删除窗口最右边的点，这样更右边的窗口可以包含的点更少。

代码实现时，把 $a_i$ 置为 $0$，表示删除。移动窗口左端点时，如果 $a_{\textit{left}}>0$ 则把 $\textit{cnt}$ 减一，否则不变。

Go 语言代码：

```go
package main
import("bufio";."fmt";"os";"slices")

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m, k, cnt, left, ans int
	Fscan(in, &n, &m, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)
	for i, x := range a {
		cnt++ // 移入窗口
		for x-a[left] >= m {
			if a[left] > 0 {
				cnt-- // 移出窗口
			}
			left++
		}
		if cnt == k {
			a[i] = 0 // 删除
			cnt--
			ans++
		}
	}
	Print(ans)
}
```

时间复杂度：$\mathcal{O}(n\log n)$。瓶颈在排序上。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
