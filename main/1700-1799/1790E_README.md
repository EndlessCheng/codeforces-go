对于二进制加法，$a+b$ 可以拆分 $a\oplus b$ 和两个 $a\&b$，即

$$
a+b = (a\oplus b) + 2\cdot (a\&b)
$$

结合题目给的公式

$$
a+b = 2\cdot (a\oplus b)
$$

得

$$
a\&b = \dfrac{a\oplus b}{2}
$$

如果 $a\oplus b$ 是奇数，或者 $(a\oplus b) \& (a\&b) \ne 0$，那么无解。后者是因为 $a\oplus b$ 中的 $1$ 说明在这个比特位上必定一个是 $0$ 另一个是 $1$，AND 结果必定是 $0$，所以 $a\oplus b$ 和 $a\&b$ 在同一个比特位上不可能都是 $1$，所以 $(a\oplus b) \& (a\&b) = 0$ 恒成立。

否则解一定存在，例如下面是一组解：

$$
\begin{cases}
a = (a\oplus b) | (a\&b)\\
b = a\&b
\end{cases}
$$

```go
package main
import("bufio";."fmt";"os")

func main() {
	in := bufio.NewReader(os.Stdin)
	var T, xor int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &xor)
		and := xor >> 1
		if xor&1 > 0 || xor&and > 0 {
			Println(-1)
		} else {
			Println(xor|and, and)
		}
	}
}
```

时间复杂度：$\mathcal{O}(1)$。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
