### 提示 1

设拼接后的字符串为 $S$。如果 $S$ 能满足后两个要求，那么 $S$ 的长度一定是奇数。所以只需关注后面两个要求。

### 提示 2

题目要求 $S$ 恰好有 $25$ 种不同字母。我们可以枚举不包含的那个字母，记作 $k$。

### 提示 3

下面的讨论，假设字符串均不包含 $k$。

题目要求每种字母的出现次数均为奇数，这可以用二进制表示，二进制从低到高的第 $i$ 位是 $1$，就表示第 $i$ 个字母出现了奇数次。

拼接成 $S$ 的这两个字符串的二进制，分别记作 $m_i$ 和 $m_j$。

问题相当于求满足

$$
m_i\operatorname{xor} m_j = 2^{26}-1-2^k
$$

的 $(m_i,m_j)$ 数对个数。

套路和 [1. 两数之和](https://leetcode.cn/problems/two-sum/) 是一样的，用哈希表或者数组解决。

```go
package main
import("bufio";."fmt";"os")

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, ans int
	var s string
	Fscan(in, &n)
	a := make([]struct{ m, all uint32 }, n)
	for i := range a {
		Fscan(in, &s)
		for _, c := range s {
			b := uint32(1) << (c - 'a')
			a[i].m ^= b
			a[i].all |= b
		}
	}
	for k := 0; k < 26; k++ {
		cnt := map[uint32]int{}
		for _, p := range a {
			if p.all>>k&1 == 0 { // 只统计不包含 k 的字符串
				ans += cnt[1<<26-1^1<<k^p.m]
				cnt[p.m]++
			}
		}
	}
	Print(ans)
}
```

时间复杂度：$\mathcal{O}(L+|\Sigma|n)$。其中 $L$ 为字符串长度之和，$|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
