计算每个子串是否回文，可以用**中心扩展法**，思路参考 [647. 回文子串](https://leetcode.cn/problems/palindromic-substrings/) 的 [官方题解](https://leetcode.cn/problems/palindromic-substrings/solutions/379987/hui-wen-zi-chuan-by-leetcode-solution/)，下面只讲解 DP 部分。

定义 $f[i]$ 表示 $s[0..i-1]$ 中的不重叠回文子字符串的最大数目。

特别地，定义 $f[0] = 0$，方便我们表示空字符串。

如果 $s[i]$ 不在回文子串中，那么有 $f[i+1] = f[i]$。

采用中心扩展法，如果 $s[l..r]$ 是回文子串，且 $r-l+1\ge k$，那么有状态转移方程

$$
f[r+1] = \max(f[r+1], f[l]+1)
$$

最后答案为 $f[n]$，这里 $n$ 为 $s$ 的长度。

```py [sol1-Python3]
class Solution:
    def maxPalindromes(self, s: str, k: int) -> int:
        n = len(s)
        f = [0] * (n + 1)
        for i in range(2 * n - 1):
            l, r = i // 2, i // 2 + i % 2  # 中心扩展法
            f[l + 1] = max(f[l + 1], f[l])
            while l >= 0 and r < n and s[l] == s[r]:
                if r - l + 1 >= k:
                    f[r + 1] = max(f[r + 1], f[l] + 1)
                l -= 1
                r += 1
        return f[n]
```

```java [sol1-Java]
class Solution {
    public int maxPalindromes(String S, int k) {
        var s = S.toCharArray();
        var n = s.length;
        var f = new int[n + 1];
        for (var i = 0; i < 2 * n - 1; ++i) {
            int l = i / 2, r = l + i % 2; // 中心扩展法
            f[l + 1] = Math.max(f[l + 1], f[l]);
            for (; l >= 0 && r < n && s[l] == s[r]; --l, ++r)
                if (r - l + 1 >= k)
                    f[r + 1] = Math.max(f[r + 1], f[l] + 1);
        }
        return f[n];
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int maxPalindromes(string s, int k) {
        int n = s.length(), f[n + 1];
        memset(f, 0, sizeof(f));
        for (int i = 0; i < 2 * n - 1; ++i) {
            int l = i / 2, r = l + i % 2; // 中心扩展法
            f[l + 1] = max(f[l + 1], f[l]);
            for (; l >= 0 && r < n && s[l] == s[r]; --l, ++r)
                if (r - l + 1 >= k)
                    f[r + 1] = max(f[r + 1], f[l] + 1);
        }
        return f[n];
    }
};
```

```go [sol1-Go]
func maxPalindromes(s string, k int) int {
	n := len(s)
	f := make([]int, n+1)
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2 // 中心扩展法
		f[l+1] = max(f[l+1], f[l])
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 >= k {
				f[r+1] = max(f[r+1], f[l]+1)
			}
			l--
			r++
		}
	}
	return f[n]
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n^2)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$O(n)$。
