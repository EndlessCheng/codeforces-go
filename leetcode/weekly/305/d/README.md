下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

定义 $f[i][c]$ 表示 $s$ 的前 $i$ 个字母中的以 $c$ 结尾的理想字符串的最长长度。

考虑 $s[i]$ 选或不选，根据题意：

- 选，需要从 $f[i-1]$ 中的 $[s[i]-k,s[i]+k]$ 范围内的字符转移过来，即

  $$
  f[i][s[i]] = 1 + \max_{c=\max(s[i]-k,0)}^{\min(s[i]+k,25)} f[i-1][c]
  $$

- 不选，则 $f[i][c] = f[i-1][c]$。

答案为 $\max(f[n-1])$。

代码实现时第一维可以压缩掉。

#### 复杂度分析

- 时间复杂度：$O(nk)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$O(|\Sigma|)$，其中 $|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。

```py [sol1-Python3]
class Solution:
    def longestIdealString(self, s: str, k: int) -> int:
        f = [0] * 26
        for c in s:
            c = ord(c) - ord('a')
            f[c] = 1 + max(f[max(c - k, 0): c + k + 1])
        return max(f)
```

```java [sol1-Java]
class Solution {
    public int longestIdealString(String s, int k) {
        var f = new int[26];
        for (var i = 0; i < s.length(); i++) {
            var c = s.charAt(i) - 'a';
            for (var j = Math.max(c - k, 0); j <= Math.min(c + k, 25); j++)
                f[c] = Math.max(f[c], f[j]);
            f[c]++;
        }
        return Arrays.stream(f).max().getAsInt();
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int longestIdealString(string &s, int k) {
        int f[26] = {};
        for (char c : s) {
            c -= 'a';
            f[c] = 1 + *max_element(f + max(c - k, 0), f + min(c + k + 1, 26));
        }
        return *max_element(f, f + 26);
    }
};
```

```go [sol1-Go]
func longestIdealString(s string, k int) (ans int) {
	f := [26]int{}
	for _, c := range s {
		c := int(c - 'a')
		for j := max(c-k, 0); j <= min(c+k, 25); j++ {
			f[c] = max(f[c], f[j])
		}
		f[c]++
	}
	for _, v := range f {
		ans = max(ans, v)
	}
	return
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
```
