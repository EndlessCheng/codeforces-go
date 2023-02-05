下午两点【bilibili@灵茶山艾府】直播讲题，记得关注哦~

---

### 算法小课堂：前缀和

定义前缀和 $s[0]=0$，$s[i+1] = \sum\limits_{j=0}^{i}p[j]$。

例如 $p=[2,3,3,5]$，对应的前缀和数组为 $s=[0, 2, 5, 8, 13]$。

通过前缀和，我们可以把**子数组的和转换成两个前缀和的差**，即

$$
\sum_{j=\textit{left}}^{\textit{right}}p[j] = \sum\limits_{j=0}^{\textit{right}}p[j] - \sum\limits_{j=0}^{\textit{left}-1}p[j] = s[\textit{right}+1] - s[\textit{left}]
$$

例如 $p$ 的子数组 $[3,3,5]$ 的和就可以用 $s[4]-s[1]=13-2=11$ 算出来。

> 注：为方便计算，常用左闭右开区间 $[\textit{left},\textit{right})$ 来表示子数组，此时子数组的和为 $s[\textit{right}]-s[\textit{left}]$。

### 思路

如果 $\textit{words}[i]$ 符合要求，就视作 $1$，否则视作 $0$。

这样求前缀和后，就可以 $O(1)$ 回答每个询问了。

```py [sol1-Python3]
class Solution:
    def vowelStrings(self, words: List[str], queries: List[List[int]]) -> List[int]:
        s = list(accumulate((w[0] in "aeiou" and w[-1] in "aeiou" for w in words), initial=0))
        return [s[r + 1] - s[l] for l, r in queries]
```

```go [sol1-Go]
func vowelStrings(words []string, queries [][]int) []int {
	sum := make([]int, len(words)+1)
	for i, w := range words {
		sum[i+1] = sum[i]
		if strings.Contains("aeiou", w[:1]) && strings.Contains("aeiou", w[len(w)-1:]) {
			sum[i+1]++
		}
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = sum[q[1]+1] - sum[q[0]]
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$O(n+q)$，其中 $n$ 为 $\textit{words}$ 的长度, $q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$O(n)$。
