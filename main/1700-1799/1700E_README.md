定义**好格子**为：要么是 $1$，要么至少有一个小于自己的邻居。若不满足，则为**坏格子**。

**结论**：如果所有格子都是好格子，那么无需交换。

**证明**：反证法。假设从 $1$ 开始，某个时刻遇到了一堵「墙」。设此时没有访问的最小元素为 $x$，那么 $x$ 必然在「墙」的另一侧。由于所有比 $x$ 小的数我们都访问过了，所以 $x$ 的邻居必然都大于 $x$，所以 $x$ 必然是个坏格子，矛盾，故原命题成立。

我们可以通过交换坏格子（或者坏格子的邻居）与其他的格子，来让坏格子变成好格子。

**结论**：如果有 $\ge 4$ 个坏格子，那么至少要操作 $2$ 次。

**证明**：考虑下图，红色为坏格子。唯一的交换方法是找一个外面的格子，和正中间的蓝色格子交换，但是交换后，外面仍然有一圈无法通过的由灰色格子组成的「墙」。

![](https://cdn.luogu.com.cn/upload/image_hosting/869tu0ee.png)

如果 $4$ 个坏格子不是上图这样，那么更加不可能只通过 $1$ 次操作就使其都变成好格子。

如果有 $1$ 到 $3$ 个坏格子，那么可能只需操作 $1$ 次，也可能至少要操作 $2$ 次。

枚举坏格子和其邻居 $(x,y)$（至多 $15$ 个格子），然后枚举所有 $(i,j)$，交换二者。交换后，检查是否有坏格子仍然是坏格子，如果是，那么交换失败。如果不是，那么继续判断 $(x,y),(i,j)$ 及其邻居（至多 $10$ 个格子）是否均为好格子。如果是，那么记录 $((x,y),(i,j))$ 到集合 $\textit{ans}$ 中（目的是去重，注意保证坐标字典序小的在前面）。

如果枚举结束后，$\textit{ans}$ 是空的，说明至少要操作 $2$ 次；否则只需操作 $1$ 次，且方案数为 $\textit{ans}$ 的大小。

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
	type pair struct{ i, j int }
	badPos := []pair{} // 坏格子
	swapPos := map[pair]int{}
	for i := range n {
		for j := range m {
			if ok(i, j) {
				continue
			}
			badPos = append(badPos, pair{i, j}) // 坏格子
			if len(badPos) >= 4 {
				Print(2)
				return
			}
			// 除了交换 (i,j)，也可以通过交换 (i,j) 的邻居，使自己变成一个好格子
			swapPos[pair{i, j}] = 1
			swapPos[pair{i, j - 1}] = 1
			swapPos[pair{i, j + 1}] = 1
			swapPos[pair{i - 1, j}] = 1
			swapPos[pair{i + 1, j}] = 1
		}
	}
	if len(badPos) == 0 {
		Print(0)
		return
	}

	// (i,j)，以及 (i,j) 的邻居，都是好格子
	ok2 := func(i, j int) bool {
		return ok(i, j) &&
			(j == 0 || ok(i, j-1)) &&
			(j+1 == m || ok(i, j+1)) &&
			(i == 0 || ok(i-1, j)) &&
			(i+1 == n || ok(i+1, j))
	}
	ans := map[pair]struct{}{}
	for p := range swapPos {
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
