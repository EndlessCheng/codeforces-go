定义**好格子**为：要么是 $1$，要么至少有一个小于自己的邻居。若不满足，则为**坏格子**。

**观察**：如果所有格子都是好格子，那么无需交换。

**证明**：反证法。假设从 $1$ 开始，某个时刻怎么走也走不下去了（遇到了一堵墙）。设此时没有被访问的最小元素为 $x$，那么 $x$ 必然在墙的另一侧。由于所有比 $x$ 小的数我们都访问过了，所以 $x$ 的邻居必然都大于 $x$，所以 $x$ 必然是个坏格子，矛盾，故原命题成立。

我们可以通过交换坏格子（或者坏格子的邻居）与其他的格子，来让坏格子变成好格子。

如果存在坏格子，我们只需考察其中一个坏格子和它的 $4$ 个邻居，因为这个坏格子必须变成好格子。

枚举其中一个坏格子及其邻居，记作 $(x,y)$，然后枚举所有 $(i,j)$，交换二者。交换后，检查是否有坏格子仍然是坏格子，如果是，那么交换失败。如果不是，继续判断 $(x,y),(i,j)$ 及其邻居（至多 $10$ 个格子）是否均为好格子。如果是，那么记录 $((x,y),(i,j))$ 到集合 $\textit{ans}$ 中（目的是去重，注意保证坐标字典序小的在前面）。

如果枚举结束后，$\textit{ans}$ 是空的，说明至少要操作 $2$ 次；否则只需操作 $1$ 次，且方案数为 $\textit{ans}$ 的大小。

⚠**注意**：提前判断有 $4$ 个坏格子就输出 $2$ 的做法是错误的，hack 数据如下：

```
6 7
1  12 11 10 38 37 36
14 13  9  3 39 35 34
15  8  2 42  4 40 33
16 17  7  5 41 31 32
18 19 20  6 28 29 30
21 22 23 24 25 26 27
```

因为交换 $1$ 和 $42$ 即可。

AC 代码（Golang）：

```go
package main
import("bufio";."fmt";"os")

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	Fscan(in, &n, &m)
	a := make([][]int32, n)
	for i := range a {
		a[i] = make([]int32, m)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}

	// 好格子：自己是 1，或者存在一个小于自己的邻居
	ok := func(i, j int) bool {
		return a[i][j] == 1 ||
			j > 0 && a[i][j-1] < a[i][j] ||
			j+1 < m && a[i][j+1] < a[i][j] ||
			i > 0 && a[i-1][j] < a[i][j] ||
			i+1 < n && a[i+1][j] < a[i][j]
	}
	// 判断 (i,j)，以及 (i,j) 的邻居，是否都是好格子
	ok2 := func(i, j int) bool {
		return ok(i, j) &&
			(j == 0 || ok(i, j-1)) &&
			(j+1 == m || ok(i, j+1)) &&
			(i == 0 || ok(i-1, j)) &&
			(i+1 == n || ok(i+1, j))
	}

	// 收集坏格子
	type pair struct{ i, j int }
	badPos := []pair{}
	for i := range n {
		for j := range m {
			if !ok(i, j) {
				badPos = append(badPos, pair{i, j})
			}
		}
	}
	if len(badPos) == 0 {
		Print(0)
		return
	}

	ans := map[pair]struct{}{}
	// 除了交换 (bi,bj)，也可以通过交换 (bi,bj) 的邻居，使 (bi,bj) 变成一个好格子
	// 只需检查至多 5 个位置，因为 (bi,bj) 必须变成好格子
	bi, bj := badPos[0].i, badPos[0].j
	for _, p := range []pair{{bi, bj}, {bi, bj - 1}, {bi, bj + 1}, {bi - 1, bj}, {bi + 1, bj}} {
		if p.i < 0 || p.i == n || p.j < 0 || p.j == m {
			continue
		}
		for i := range n {
			for j := range m {
				// 交换其他所有点
				a[p.i][p.j], a[i][j] = a[i][j], a[p.i][p.j]
				// 交换离坏格子很远的点，必然是无效交换，所以先检查是否有坏格子仍然是坏格子
				for _, q := range badPos {
					if !ok(q.i, q.j) {
						goto next
					}
				}
				// 有效交换！进一步检查受到影响的 10 个点是否正常
				if ok2(p.i, p.j) && ok2(i, j) {
					// 注意去重
					ans[pair{min(p.i*m+p.j, i*m+j), max(p.i*m+p.j, i*m+j)}] = struct{}{}
				}
			next:
				a[p.i][p.j], a[i][j] = a[i][j], a[p.i][p.j]
			}
		}
	}

	if len(ans) > 0 {
		Println(1, len(ans))
	} else {
		Print(2)
	}
}
```

**时间复杂度**：$\mathcal{O}(nm)$。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
