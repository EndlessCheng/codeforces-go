双倍经验：[CF798D](https://www.luogu.com.cn/problem/CF798D)

按照 $a_i$ 从小到大排序。

**结论**：如下两个方案，必定有一个是符合要求的：

- 方案一：选择装有 $a_1,a_3,a_5,\cdots,a_{2n-5},a_{2n-3},a_{2n-1}$ 苹果的 $n$ 个盒子。
- 方案二：选择装有 $a_2,a_4,a_6,\cdots,a_{2n-4},a_{2n-2},a_{2n-1}$ 苹果的 $n$ 个盒子。

**证明**：

对于方案一，除去 $a_1$ 的其余每个数，都大于等于它左边的数，所以这些苹果的个数必然大于等于没有选的苹果个数，也就是说，这些苹果的数量至少是全部苹果的一半。

对于方案二，除去 $a_{2n-1}$ 的其余每个数，都大于等于它左边的数，所以这些苹果的个数必然大于等于没有选的苹果个数，也就是说，这些苹果的数量至少是全部苹果的一半。

所以只需要证明，至少有一种方案中的橘子个数是符合要求的。

设 $s_1 = o_1 + o_3 + o_5 + \cdots + o_{2n-5} + o_{2n-3}$。

设 $s_2 = o_2 + o_4 + o_6 + \cdots + o_{2n-4} + o_{2n-2}$。

那么所有橘子的个数为 $s = s_1 + s_2 + o_{2n-1}$。

对于方案一，我们选的橘子个数为 $s_1 + o_{2n-1}$。

对于方案二，我们选的橘子个数为 $s_2 + o_{2n-1}$。

利用反证法，只要证明如下两个不等式不可能同时成立，就能证明结论。

$$
\begin{aligned}
&s_1 + o_{2n-1} < s/2\\
&s_2 + o_{2n-1} < s/2
\end{aligned}
$$

把这两个不等式相加，得

$$
s_1 + s_2 + 2\cdot o_{2n-1} < s = s_1 + s_2 + o_{2n-1}
$$

化简得

$$
o_{2n-1} < 0
$$

矛盾，所以结论成立。

```go
package main
import("bufio";."fmt";"os";"sort")

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		m := n*2 - 1
		a := make([]struct{ x, y, i int }, m)
		s := 0
		for i := range a {
			Fscan(in, &a[i].x, &a[i].y)
			s += a[i].y
			a[i].i = i + 1
		}

		sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })
		se := 0
		for i := 0; i < m; i += 2 {
			se += a[i].y
		}

		Fprintln(out, "YES")
		if se*2 >= s { // 方案一
			for i := 0; i < m; i += 2 {
				Fprint(out, a[i].i, " ")
			}
			Fprintln(out)
		} else { // 方案二
			for i := 1; i < m; i += 2 {
				Fprint(out, a[i].i, " ")
			}
			Fprintln(out, a[m-1].i)
		}
	}
}
```

时间复杂度：$\mathcal{O}(n\log n)$。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
