思路和 [70. 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/) 类似。

把相同字符分为一组，考虑如下 DP：

对于字符不为 $\texttt{7}$ 或 $\texttt{9}$ 的情况，定义 $f[i]$ 表示长为 $i$ 的只有一种字符的字符串对应的文字信息种类数，我们可以将末尾的 $1$ 个、$2$ 个或 $3$ 个字符单独视作一个字母，那么有转移方程

$$
f[i] = f[i-1]+f[i-2]+f[i-3]
$$

对于字符为 $\texttt{7}$ 或 $\texttt{9}$ 的情况，定义 $g[i]$ 表示长为 $i$ 的只有一种字符的字符串对应的文字信息种类数，可以得到类似的转移方程

$$
g[i] = g[i-1]+g[i-2]+g[i-3]+g[i-4]
$$

这样能算出每组字符串的文字信息种类数。

由于不同组之间互不影响，根据乘法原理，把不同组的文字信息种类数相乘，得到答案。

```go
const mod, mx int = 1e9 + 7, 1e5

var f = [mx + 1]int{1, 1, 2, 4}
var g = f

func init() {
	for i := 4; i <= mx; i++ { // 预处理所有长度的结果
		f[i] = (f[i-1] + f[i-2] + f[i-3]) % mod
		g[i] = (g[i-1] + g[i-2] + g[i-3] + g[i-4]) % mod
	}
}

func countTexts(s string) int {
	ans, cnt := 1, 0
	for i, c := range s {
		cnt++
		if i == len(s)-1 || s[i+1] != byte(c) { // 找到一个完整的组
			if c != '7' && c != '9' {
				ans = ans * f[cnt] % mod
			} else {
				ans = ans * g[cnt] % mod
			}
			cnt = 0
		}
	}
	return ans
}
```
