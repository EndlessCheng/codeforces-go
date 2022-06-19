本题 [视频讲解](https://www.bilibili.com/video/BV1CW4y1k7B3) 已出炉，欢迎点赞三连~

---

#### 提示 1

子序列的长度**至少**可以是多少？

#### 提示 2

前导零不会改变二进制数的大小，因此要尽可能地往子序列前面添加前导零。

#### 提示 3

我们应该在 $s$ 中的一个靠后的位置找值不超过 $k$ 的子序列，越靠后越好，因为这样前面能添加的前导零也就越多。

#### 提示 4

找 $s$ 中值不超过 $k$ 的最长后缀，往这个后缀前面添加前导零。

#### 提示 5

设 $s$ 的长度为 $n$，$k$ 的二进制长度为 $m$。

分类讨论：

- 如果 $n<m$，由于 $k$ 最高位为 $1$，整个 $s$ 对应的数字必然小于 $k$；
- 如果 $s$ 长为 $m$ 的后缀 $\textit{suf}$ 对应的数字不超过 $k$，那么我们可以从 $s$ 的其余部分找尽可能多的 $0$，拼在 $\textit{suf}$ 前面；
- 如果 $s$ 长为 $m$ 的后缀对应的数字超过 $k$，那么我们可以取 $s$ 长为 $m-1$ 的后缀 $\textit{suf}$，这样对应的数字小于 $k$，然后同上，从其余部分找尽可能多的 $0$，拼在 $\textit{suf}$ 前面。

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干变量。（忽略子串的空间开销）

```Python [sol1-Python3]
class Solution:
    def longestSubsequence(self, s: str, k: int) -> int:
        n, m = len(s), k.bit_length()
        if n < m: return n
        ans = m if int(s[n - m:], 2) <= k else m - 1
        return ans + s.count('0', 0, n - m)
```

```java [sol1-Java]
class Solution {
    public int longestSubsequence(String s, int k) {
        int n = s.length(), m = 32 - Integer.numberOfLeadingZeros(k);
        if (n < m) return n;
        var ans = Integer.parseInt(s.substring(n - m), 2) <= k ? m : m - 1;
        return ans + (int) s.substring(0, n - m).chars().filter(c -> c == '0').count();
    }
}
```

```C++ [sol1-C++]
class Solution {
public:
    int longestSubsequence(string s, int k) {
        int n = s.length(), m = 32 - __builtin_clz(k);
        if (n < m) return n;
        int ans = stoi(s.substr(n - m), nullptr, 2) <= k ? m : m - 1;
        return ans + count(s.begin(), s.end() - m, '0');
    }
};
```

```go [sol1-Go]
func longestSubsequence(s string, k int) int {
	n, m := len(s), bits.Len(uint(k))
	if n < m {
		return n
	}
	ans := m
	v, _ := strconv.ParseInt(s[n-m:], 2, 64)
	if int(v) > k {
		ans--
	}
	return ans + strings.Count(s[:n-m], "0")
}
```
