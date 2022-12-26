[视频讲解](https://www.bilibili.com/video/BV1FV4y1F7v7/) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

设从左侧取到第 $i$ 个字符，从右侧取到第 $j$ 个字符。

由于随着 $i$ 的变大，$j$ 也会单调变大，因此可以用双指针，一边从小到大枚举 $i$，一边维护 $j$ 的最大位置（$j$ 尽量向右移）。

对于左侧没有取字符的情况需要单独计算。

```py [sol1-Python3]
class Solution:
    def takeCharacters(self, s: str, k: int) -> int:
        j = n = len(s)
        c = Counter()
        while c['a'] < k or c['b'] < k or c['c'] < k:
            if j == 0: return -1  # 所有字母都取也无法满足要求
            j -= 1
            c[s[j]] += 1
        ans = n - j  # 左侧没有取字符
        for i, ch in enumerate(s):
            c[ch] += 1
            while j < n and c[s[j]] > k:  # 维护 j 的最大下标
                c[s[j]] -= 1
                j += 1
            ans = min(ans, i + 1 + n - j)
            if j == n: break
        return ans
```

```go [sol1-Go]
func takeCharacters(s string, k int) int {
	n := len(s)
	c, j := [3]int{}, n
	for c[0] < k || c[1] < k || c[2] < k {
		if j == 0 {
			return -1 // 所有字母都取也无法满足要求
		}
		j--
		c[s[j]-'a']++
	}
	ans := n - j // 左侧没有取字符
	for i := 0; i < n && j < n; i++ { // 双指针
		c[s[i]-'a']++
		for j < n && c[s[j]-'a'] > k { // 维护 j 的最大下标
			c[s[j]-'a']--
			j++
		}
		ans = min(ans, i+1+n-j)
	}
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $s$ 的长度。注意二重循环中，`j++` 的次数不会超过 $O(n)$，所以二重循环的时间复杂度为 $O(n)$。
- 空间复杂度：$O(|\Sigma|)$，其中 $|\Sigma|$ 为字符集合的大小，本题中 $|\Sigma|=3$。
