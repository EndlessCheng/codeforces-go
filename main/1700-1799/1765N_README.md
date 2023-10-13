本题属于【字典序最小子序列问题】，可以用单调栈解决，时间复杂度严格 $\mathcal{O}(n)$。

先看看**允许有前导零**要怎么做：直接遍历字符串，用单调栈维护，只要当前字符小于栈顶，就弹出栈顶。最后栈底到栈顶是单调非降的，把栈顶的 $k$ 个元素弹出。

回到原题，**不允许有前导零**，我们可以在前 $k$ 个字符中找到第一个最小的非 $0$ 字符，设其为下标为 $j$，那么肯定要把下标小于 $j$ 的字符都删掉。然后从 $j+1$ 开始按照上面的做法跑单调栈即可。

```go
package main
import("bufio";."fmt";"os")

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T, k int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &k)
		j := 0
		for i, b := range s[:k+1] {
			if b > '0' && b < s[j] { // 最小非 0 字符（且下标最小）
				j = i
			}
		}
		k -= j // 把 j 左边的都删掉
		st := s[j : j+1] // 保留 s[j]
		for _, b := range s[j+1:] { // j 右边的正常跑单调栈即可
			for len(st) > 1 && k > 0 && b < st[len(st)-1] {
				st = st[:len(st)-1]
				k--
			}
			st = append(st, b)
		}
		st = st[:len(st)-k] // 还剩下操作次数，就去掉后面的（因为栈顶元素大）
		Fprintf(out, "%s\n", st)
	}
}
```
