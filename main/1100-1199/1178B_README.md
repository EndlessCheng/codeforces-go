### 方法一：前后缀分解（两次遍历）

预处理每个 $\texttt{o}$ 左右的 $\texttt{vv}$ 个数，分别记作 $\textit{pre}$ 和 $\textit{suf}$。

然后枚举 $\texttt{o}$，根据乘法原理，把 $\textit{pre}\cdot \textit{suf}$ 加到答案中。

```go
package main
import("bufio";."fmt";"os")

func main() {
	var s string
	Fscan(bufio.NewReader(os.Stdin), &s)
	var pre, suf, ans int
	n := len(s)
	for i := 1; i < n-1; i++ {
		if s[i] == 'v' && s[i+1] == 'v' {
			suf++
		}
	}
	for i := 1; i < n-2; i++ {
		if s[i] == 'o' {
			ans += pre * suf
		} else {
			if s[i-1] == 'v' {
				pre++
			}
			if s[i+1] == 'v' {
				suf--
			}
		}
	}
	Print(ans)
}
```

### 方法二：状态机 DP（一次遍历）

定义：

- $f_{i,0}$ 表示考虑前 $i$ 个字母，能得到的 $\texttt{w}$ 子序列的个数。
- $f_{i,1}$ 表示考虑前 $i$ 个字母，能得到的 $\texttt{wo}$ 子序列的个数。
- $f_{i,2}$ 表示考虑前 $i$ 个字母，能得到的 $\texttt{wow}$ 子序列的个数。

分类讨论：

- 如果 $s_i=\texttt{o}$ 那么 $f_{i,1} = f_{i-1,1} + f_{i-1,0}$，分别表示不选和选 $s_i$。
- 如果 $s_i=\texttt{v}$ 且 $s_{i-1}=\texttt{v}$，那么 $f_{i,0} = f_{i-1,0} + 1$，$f_{i,2} = f_{i-1,2} + f_{i-1,1}$，理由同上。
- 其余情况 $f_{i,j}=f_{i-1,j}$。

答案为 $f_{n,2}$。

代码实现时，第一个维度可以去掉。

```go
package main
import("bufio";."fmt";"os")

func main() {
	var s string
	Fscan(bufio.NewReader(os.Stdin), &s)
	var f0, f1, f2 int
	for i := 1; i < len(s); i++ {
		if s[i] == 'o' {
			f1 += f0
		} else if s[i-1] == 'v' {
			f2 += f1
			f0++
		}
	}
	Print(f2)
}
```

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
