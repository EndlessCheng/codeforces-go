### 方法一：二分答案+滑动窗口

代价越大，越容易做到（存在移除方案）。

二分最大代价 $\textit{mx}$，用滑动窗口来维护，只要窗口尽量大（窗口内 $0$ 的个数不超过 $\textit{mx}$），那么窗口外的 $1$ 的个数就尽量小。

如果某个时刻窗口外的 $1$ 的个数不超过 $\textit{mx}$，那么答案至多为 $\textit{mx}$。根据这个来写 $\text{check}$ 函数。

```go
package main
import("bufio";."fmt";"os";"sort";"strings")

func main() {
	in := bufio.NewReader(os.Stdin)
	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		tot1 := strings.Count(s, "1")
		Println(sort.Search(len(s), func(mx int) bool { // 二分最大代价 mx
			in0 := 0     // 窗口内的 0 的个数
			out1 := tot1 // 窗口外的 1 的个数
			left := 0
			for _, b := range s {
				v := int(b & 1)
				in0 += v ^ 1
				out1 -= v
				for in0 > mx { // 窗口内的 0 太多了
					v = int(s[left] & 1)
					in0 -= v ^ 1
					out1 += v
					left++
				}
				if out1 <= mx { // 答案至多为 mx
					return true
				}
			}
			return false // 答案至少为 mx+1
		}))
	}
}
```

时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是字符串 $s$ 的长度。

### 方法二：利用性质

注意到，如果窗口内的 $0$ 的个数低于窗口外的 $1$ 的个数，那么扩大窗口更好：

- 如果把 $1$ 移入窗口，那么窗口外 $1$ 的个数变小，代价变小。
- 如果把 $0$ 移入窗口，虽然代价不变，但如果后面把 $1$ 移入窗口，也会让代价变小。相当于更有机会让代价变小。

同样，如果窗口内的 $0$ 的个数超过窗口外的 $1$ 的个数，那么缩小窗口更好：

- 如果把 $0$ 移出窗口，那么窗口内 $0$ 的个数变小，代价变小。
- 如果把 $1$ 移出窗口，虽然代价不变，但如果后面把 $0$ 移出窗口，也会让代价变小。

所以无需二分，直接滑窗即可。

```go
package main
import("bufio";. "fmt";"os";"strings")
func min(a, b int) int { if b < a { return b }; return a }

func main() {
	in := bufio.NewReader(os.Stdin)
	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		ans := len(s)
		in0 := 0
		out1 := strings.Count(s, "1")
		left := 0
		for _, b := range s {
			v := int(b & 1)
			in0 += v ^ 1
			out1 -= v
			for in0 > out1 { // 0 多就缩小窗口
				v = int(s[left] & 1)
				in0 -= v ^ 1
				out1 += v
				left++
			}
			ans = min(ans, out1) // in0 <= out1
		}
		Println(ans)
	}
}
```

时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是字符串 $s$ 的长度。
