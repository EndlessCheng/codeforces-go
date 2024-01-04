对于二进制加法，$a+b$ 可以拆分 $a\operatorname{xor} b$ 和两个 $a\operatorname{and} b$，即

$$
a+b = (a\operatorname{xor} b) + 2\cdot (a\operatorname{and} b)
$$

结合题目给的公式

$$
a+b = 2\cdot (a\operatorname{xor} b)
$$

得

$$
a\operatorname{and} b = \dfrac{a\operatorname{xor} b}{2}
$$

如果 $a\operatorname{xor} b$ 是奇数，或者 $(a\operatorname{xor} b) \operatorname{and}\  (a\operatorname{and} b) \ne 0$，那么无解。后者是因为 $a\operatorname{xor} b$ 中的 $1$ 说明在这个比特位上必定一个是 $0$ 另一个是 $1$，$\operatorname{and}$ 结果必定是 $0$，所以 $a\operatorname{xor} b$ 和 $a\operatorname{and} b$ 在同一个比特位上不可能都是 $1$，所以 $(a\operatorname{xor} b) \operatorname{and}\  (a\operatorname{and} b) = 0$ 恒成立。从集合的角度上来说就是 $a\operatorname{xor} b$ 和 $a\operatorname{and} b$ **不相交**。

现在 $a\operatorname{and} b$ 和 $a\operatorname{xor} b$ 都知道了，可以构造出一组解了。

由于 $a\operatorname{and} b$ 中的 $1$ 在 $a$ 和 $b$ 中都有，而 $a\operatorname{xor} b$ 中的 $1$ 要么在 $a$ 中，要么在 $b$ 中，不妨全部给 $a$，得

$$
\begin{cases}
a = (a\operatorname{xor} b) \operatorname{or}\ (a\operatorname{and} b) = (a\operatorname{xor} b) + (a\operatorname{and} b) = \dfrac{3(a\operatorname{xor} b)}{2}\\
b = a\operatorname{and} b = \dfrac{a\operatorname{xor} b}{2}
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
