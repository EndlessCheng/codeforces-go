下午两点在B站讲这场双周赛的题目，[欢迎关注](https://space.bilibili.com/206214)~

---

做法同 [1930. 长度为 3 的不同回文子序列](https://leetcode.cn/problems/unique-length-3-palindromic-subsequences/)。

首先倒着遍历 $s$，统计每个字符的出现次数 $\textit{suf}[a]$ 和两个字符的组合个数 $\textit{suf}_2[a][b]$。

例如遍历到 $a=s[i]$ 时，就更新 $\textit{suf}_2[a][0..9] \text{+=} \textit{suf}[0..9]$，然后 $\textit{suf}[a]$ 加一。

然后正着遍历 $s$，统计每个字符的出现次数 $\textit{pre}[a]$ 和两个字符的组合个数 $\textit{pre}_2[a][b]$，更新方法同上。

枚举每个 $s[i]$，当作回文子序列的回文中心，此时的子序列个数就是 $s[0..i-1]$ 中的 $\textit{pre}_2[a][b]$ 与 $s[i+1..n-1]$ 中的 $\textit{suf}_2[a][b]$ 的组合，枚举所有的 $a$ 和 $b$，个数相乘再相加，即为以 $s[i]$ 为回文子序列的回文中心时的答案。

代码实现时，可以一遍遍历一遍计算 $\textit{pre}_2[a][b]$，同时撤销 $\textit{suf}_2[a][b]$ 的统计结果。

注意在最坏情况下（所有字符都相同），答案为

$$
C(n,5) < \dfrac{n^5}{120} < 10^{18}
$$

所以可以最后返回的时候再取模。

```py [sol1-Python3]
class Solution:
    def countPalindromes(self, s: str) -> int:
        suf = [0] * 10
        suf2 = [0] * 100
        for d in map(int, reversed(s)):
            for j, c in enumerate(suf):
                suf2[d * 10 + j] += c
            suf[d] += 1

        ans = 0
        pre = [0] * 10
        pre2 = [0] * 100
        for d in map(int, s):
            suf[d] -= 1
            for j, c in enumerate(suf):
                suf2[d * 10 + j] -= c  # 撤销
            ans += sum(c1 * c2 for c1, c2 in zip(pre2, suf2))  # 枚举所有字符组合
            for j, c in enumerate(pre):
                pre2[d * 10 + j] += c
            pre[d] += 1
        return ans % (10 ** 9 + 7)
```

```java [sol1-Java]
class Solution {
    private static final long MOD = (long) 1e9 + 7;

    public int countPalindromes(String S) {
        var s = S.toCharArray();
        int[] pre = new int[10], suf = new int[10];
        int[][] pre2 = new int[10][10], suf2 = new int[10][10];
        for (var i = s.length - 1; i >= 0; --i) {
            var d = s[i] - '0';
            for (var j = 0; j < 10; ++j)
                suf2[d][j] += suf[j];
            ++suf[d];
        }

        var ans = 0L;
        for (var d : s) {
            d -= '0';
            --suf[d];
            for (var j = 0; j < 10; ++j)
                suf2[d][j] -= suf[j]; // 撤销
            for (var j = 0; j < 10; ++j)
                for (var k = 0; k < 10; ++k)
                    ans += (long) pre2[j][k] * suf2[j][k]; // 枚举所有字符组合
            for (var j = 0; j < 10; ++j)
                pre2[d][j] += pre[j];
            ++pre[d];
        }
        return (int) (ans % MOD);
    }
}
```

```cpp [sol1-C++]
class Solution {
    const long MOD = 1e9 + 7;
public:
    int countPalindromes(string s) {
        int suf[10]{}, suf2[10][10]{}, pre[10]{}, pre2[10][10]{};
        for (int i = s.length() - 1; i >= 0; --i) {
            char d = s[i] - '0';
            for (int j = 0; j < 10; ++j)
                suf2[d][j] += suf[j];
            ++suf[d];
        }

        long ans = 0L;
        for (char d : s) {
            d -= '0';
            --suf[d];
            for (int j = 0; j < 10; ++j)
                suf2[d][j] -= suf[j]; // 撤销
            for (int j = 0; j < 10; ++j)
                for (int k = 0; k < 10; ++k)
                    ans += (long) pre2[j][k] * suf2[j][k]; // 枚举所有字符组合
            for (int j = 0; j < 10; ++j)
                pre2[d][j] += pre[j];
            ++pre[d];
        }
        return ans % MOD;
    }
};
```

```go [sol1-Go]
func countPalindromes(s string) (ans int) {
	const mod int = 1e9 + 7
	n := len(s)
	suf := [10]int{}
	suf2 := [10][10]int{}
	for i := n - 1; i >= 0; i-- {
		d := s[i] - '0'
		for j, c := range suf {
			suf2[d][j] += c
		}
		suf[d]++
	}

	pre := [10]int{}
	pre2 := [10][10]int{}
	for _, d := range s[:n-1] {
		d -= '0'
		suf[d]--
		for j, c := range suf {
			suf2[d][j] -= c // 撤销
		}
		for j, sf := range suf2 {
			for k, c := range sf {
				ans += pre2[j][k] * c // 枚举所有字符组合
			}
		}
		for j, c := range pre {
			pre2[d][j] += c
		}
		pre[d]++
	}
	return ans % mod
}
```

#### 复杂度分析

- 时间复杂度：$O(n|\Sigma|^2)$，其中 $n$ 为 $s$ 的长度，$|\Sigma|$ 为字符集合的大小，本题中字符均为数字，所以 $|\Sigma|=10$。
- 空间复杂度：$O(|\Sigma|^2)$。

#### 相似题目

- [1930. 长度为 3 的不同回文子序列](https://leetcode.cn/problems/unique-length-3-palindromic-subsequences/)
