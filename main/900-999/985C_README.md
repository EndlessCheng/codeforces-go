以下讨论，下标从 $0$ 开始。

把 $a$ 排序，找到大于 $a_0+L$ 的最小的数，设其下标为 $i$。

木桶的体积只能在 $a_0,a_1,\cdots,a_{i-1}$ 这前 $i$ 条木板中。如果 $i < n$ 则输出 $0$。

贪心地想，**最小木板的长度越大越好**，但木桶必须恰好由 $k$ 条木板组成，所以前面的 $x$ 个木桶只能选 $a_0,a_k,a_{2k},\cdots,a_{(x-1)k}$ 这些木板。

剩下 $n-x$ 个木桶可以取前 $i$ 条木板中最大的 $n-x$ 个，这要求

$$
i-(n-x) \le xk
$$

解得

$$
x\ge \left\lceil\dfrac{i-n}{k-1}\right\rceil
$$

所以 $x$ 最小是

$$
\left\lceil\dfrac{i-n}{k-1}\right\rceil = \left\lfloor\dfrac{i-n+k-2}{k-1}\right\rfloor
$$

如果 $k=1$ 则 $x$ 最小是 $0$。

```go
package main
import("bufio";."fmt";"os";"slices";"sort")

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, k, l, x, ans int
	Fscan(in, &n, &k, &l)
	a := make([]int, n*k)
	for i := range a {
		Fscan(in, &a[i])
	}

	slices.Sort(a)
	i := sort.SearchInts(a, a[0]+l+1)
	if i < n {
		Print(0)
		return
	}

	if k > 1 {
		x = (i - n + k - 2) / (k - 1)
	}
	for j := 0; j <= (x-1)*k; j += k {
		ans += a[j]
	}
	for _, v := range a[i-n+x : i] {
		ans += v
	}
	Print(ans)
}
```

时间复杂度：$\mathcal{O}(nk\log(nk))$。瓶颈在排序上。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
