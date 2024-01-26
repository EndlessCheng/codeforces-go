[视频讲解](https://www.bilibili.com/video/BV1CN4y1V7uE) 第四题。

看到**子序列**和**相邻**就可以往 DP 上想（回顾一下经典题 [300. 最长递增子序列](https://leetcode.cn/problems/longest-increasing-subsequence/)，它也是子序列和相邻）。

字符串题目套路：枚举字符。定义 $f[i][c]$ 表示 $s$ 的前 $i$ 个字母中的以 $c$ 结尾的理想字符串的最长长度。

根据题意：

- 选 $s[i]$ 作为理想字符串中的字符，需要从 $f[i-1]$ 中的 $[s[i]-k,s[i]+k]$ 范围内的字符转移过来，即

  $$
  f[i][s[i]] = 1 + \max_{c=\max(s[i]-k,0)}^{\min(s[i]+k,25)} f[i-1][c]
  $$

- 其余情况，$f[i][c] = f[i-1][c]$。

答案为 $\max(f[n-1])$。

代码实现时第一维可以压缩掉。

```py [sol-Python3]
class Solution:
    def longestIdealString(self, s: str, k: int) -> int:
        f = [0] * 26
        for c in s:
            c = ord(c) - ord('a')
            f[c] = 1 + max(f[max(c - k, 0): c + k + 1])
        return max(f)
```

```java [sol-Java]
class Solution {
    public int longestIdealString(String s, int k) {
        int[] f = new int[26];
        for (char c : s.toCharArray()) {
            c -= 'a';
            for (int j = Math.max(c - k, 0); j <= Math.min(c + k, 25); j++) {
                f[c] = Math.max(f[c], f[j]);
            }
            f[c]++;
        }
        return Arrays.stream(f).max().getAsInt();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestIdealString(string &s, int k) {
        int f[26]{};
        for (char c : s) {
            c -= 'a';
            f[c] = 1 + *max_element(f + max(c - k, 0), f + min(c + k + 1, 26));
        }
        return ranges::max(f);
    }
};
```

```go [sol-Go]
func longestIdealString(s string, k int) int {
	f := [26]int{}
	for _, c := range s {
		c := int(c - 'a')
		f[c] = 1 + slices.Max(f[max(c-k, 0):min(c+k+1, 26)])
	}
	return slices.Max(f[:])
}
```

#### 复杂度分析

- 时间复杂度：$O(nk)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$O(|\Sigma|)$，其中 $|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。
