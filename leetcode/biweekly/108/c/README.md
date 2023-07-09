下午两点[【b站@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，欢迎关注！

## 前置知识：动态规划入门

详见 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)

## 思路

由于字符串的长度为 $n$，所以二进制的值是低于 $2^n$ 的。那么可以预处理 $2^n$ 以内的 $5$ 的幂的二进制表示，记作 $\textit{pow}_5$。由于测试数据很多，可以用全局变量，预处理 $2^{15}$ 以内的 $5$ 的幂（有 $7$ 个）。

定义 $\textit{dfs}(i)$ 表示划分从 $s[i]$ 开始的后缀，最少要划分成多少段。

枚举 $\textit{pow}_5$ 中的字符串 $t$，设其长度为 $m$，如果从 $s[i]$ 到 $s[i+m-1]$ 与 $t$ 相等，那么有

$$
\textit{dfs}(i) = \textit{dfs}(i+m) + 1
$$

所有情况取最小值。

如果 $s[i]=\texttt{0}$，或者不存在这样的 $t$，那么 $\textit{dfs}(i)=\infty$。

递归边界：$\textit{dfs}(n)=0$。

递归入口：$\textit{dfs}(0)$。

注意在比较字符串时，由于 $\textit{pow}_5$ 中字符串的公共前缀很短，很容易就失配了，所以除了匹配上的那次是 $\mathcal{O}(n)$ 的，其余匹配都可以视作是 $\mathcal{O}(1)$ 的。

```py [sol-Python3]
# 预处理 2**15 以内的 5 的幂
pow5 = [bin(5 ** i)[2:] for i in range(7)]

class Solution:
    def minimumBeautifulSubstrings(self, s: str) -> int:
        n = len(s)
        @cache
        def dfs(i: int) -> int:
            if i == n: return 0
            if s[i] == '0': return inf  # 不能包含前导 0
            res = inf
            for t in pow5:
                if i + len(t) > n:
                    break
                if s[i: i + len(t)] == t:  # 忽略切片的时间，这里的比较视作均摊 O(1)
                    res = min(res, dfs(i + len(t)) + 1)
            return res
        ans = dfs(0)
        return ans if ans < inf else -1
```

按照视频中的做法，1:1 翻译成递推。

倒着遍历的好处是方便判断是否有前导零。

```py [sol-Python3]
# 预处理 2**15 以内的 5 的幂
pow5 = [bin(5 ** i)[2:] for i in range(7)]

class Solution:
    def minimumBeautifulSubstrings(self, s: str) -> int:
        n = len(s)
        f = [inf] * n + [0]
        for i in range(n - 1, -1, -1):
            if s[i] == '0': continue  # 不能包含前导 0
            for t in pow5:
                if i + len(t) > n:
                    break
                if s[i: i + len(t)] == t:  # 忽略切片的时间，这里的比较视作均摊 O(1)
                    f[i] = min(f[i], f[i + len(t)] + 1)
        return f[0] if f[0] < inf else -1
```

```go [sol-Go]
var pow5 []string

func init() {
	// 预处理 2**15 以内的 5 的幂
	for p5 := 1; p5 < 1<<15; p5 *= 5 {
		pow5 = append(pow5, strconv.FormatUint(uint64(p5), 2))
	}
}

func minimumBeautifulSubstrings(s string) int {
	n := len(s)
	f := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		f[i] = n + 1
		if s[i] == '0' { // 不能包含前导 0
			continue
		}
		for _, t := range pow5 {
			if i+len(t) > n {
				break
			}
			if s[i:i+len(t)] == t {
				f[i] = min(f[i], f[i+len(t)]+1)
			}
		}
	}
	if f[0] > n {
		return -1
	}
	return f[0]
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $s$ 的长度。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(n)$，所以动态规划的时间复杂度为 $\mathcal{O}(n^2)$。注意在比较字符串时，由于 $\textit{pow}_5$ 中字符串的公共前缀很短，很容易就失配了，所以除了匹配上的那次是 $\mathcal{O}(n)$ 的，其余匹配都可以视作是 $\mathcal{O}(1)$ 的。
- 空间复杂度：$\mathcal{O}(n)$。
