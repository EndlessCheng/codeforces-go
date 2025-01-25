问题相当于把 $a$ 划分成若干个集合。

用回溯实现，对于 $a_i$：

- 单独形成一个集合。
- 加到前面的某个集合中。

用一个数组 $b$ 存储每个集合的元素和。

AC 代码（Golang）：

```go
package main
import ."fmt"

func main() {
	var n int
	Scan(&n)
	a := make([]int, n)
	for i := range a {
		Scan(&a[i])
	}

	ans := map[int]bool{}
	b := []int{}
	var dfs func(int, int)
	dfs = func(i, xor int) {
		if i == n {
			ans[xor] = true
			return
		}
		v := a[i]
		// a[i] 单独组成一个集合
		b = append(b, v)
		dfs(i+1, xor^v)
		b = b[:len(b)-1]
		// a[i] 加到前面的集合中
		for j := range b {
			b[j] += v
			dfs(i+1, xor^(b[j]-v)^b[j])
			b[j] -= v
		}
	}
	dfs(0, 0)
	Print(len(ans))
}
```
