以下描述，下标从 $0$ 开始。

### 提示 1

设答案为 $s$。

假设 $x$ 是个有序数组（递增），那么绝对值符号可以去掉，从而化简公式：

$$
\begin{aligned}
s =\ &(n-1)(x[n-1]-x[0]) + (n-3)(x[n-2]-x[1]) + \cdots\\
=\ &\sum_{i=0}^{\lfloor n/2 \rfloor-1}  (n-1-2i)(x[n-1-i]-x[i])
\end{aligned}
$$

### 提示 2

既然 $x$ 有序很有用，那么怎样可以「把 $x$ 排序」呢？

或者想一想，$x[0]$ 在哪个区间？$x[n-1]$ 在哪个区间？

如果把这些闭区间粗略地排序，那么 $R[i]$ 越小的区间应该越靠左排，$L[i]$ 越大的区间应该越靠右排。

### 提示 3

为了让 $s$ 尽量小，$x[0]$ 要尽量大，$x[n-1]$ 要尽量小。

$x[0]$ 最大是多少？$x[n-1]$ 最小是多少？

最小的 $R[i]$ 即为最大的 $x[0]$。

最大的 $L[i]$ 即为最小的 $x[n-1]$。

如果 $x[0] < x[n-1]$，那么去掉这两个数（区间），问题变成一个规模更小（$n-2$）的子问题。

否则，这些区间的交集不为空，所有 $x[i]$ 都可以取到同一个数。

```go
package main
import("bufio";."fmt";"os";"sort")

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, ans int
	Fscan(in, &n)
	l := make([]int, n)
	r := make([]int, n)
	for i := range l {
		Fscan(in, &l[i], &r[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(l))) // 从大到小排
	sort.Ints(r) // 从小到大排
	for i := 0; l[i] > r[i]; i++ {
		ans += (n - 1 - i*2) * (l[i] - r[i])
	}
	Print(ans)
}
```

时间复杂度：$\mathcal{O}(n\log n)$。瓶颈在排序上。
