#### 提示 1

枚举出现次数最多的字符（记作 $a$）与出现次数最少的字符（记作 $b$），答案一定在其中。

#### 提示 2

由于题目求的是这两个字符出现次数的差，我们可以把 $a$ 视作 $1$，$b$ 视作 $-1$，其余字符视作 $0$。则本题转换成了一个类似 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/) 的问题。

#### 提示 3

注意 $a$ 和 $b$ 必须都出现在子串中，不能把只有 $a$ 的子串作为答案。

我们可以用变量 $\textit{diff}$ 维护 $a$ 和 $b$ 的出现次数之差，初始值为 $0$。

同时用另一个变量 $\textit{diffWithB}$ 维护在包含 $b$ 时的 $a$ 和 $b$ 的出现次数之差，初始为 $-\infty$，因为还没有遇到 $b$。

遍历字符串 $s$：

- 当遇到 $a$ 时，$\textit{diff}$ 和 $\textit{diffWithB}$ 均加一。
- 当遇到 $b$ 时，$\textit{diff}$ 减一，$\textit{diffWithB}$ 记录此时的 $\textit{diff}$ 值。若 $\textit{diff}$ 为负则将其置为 $0$（根据提示 1，$\textit{diff}$ 不能为负）。

统计所有 $\textit{diffWithB}$ 的最大值，即为答案。若 $s$ 只有一个字符则答案为 $0$。

```Python [sol1-Python3]
class Solution:
    def largestVariance(self, s: str) -> int:
        if s.count(s[0]) == len(s):
            return 0
        ans = 0
        for a in ascii_lowercase:
            for b in ascii_lowercase:
                if b == a:
                    continue
                diff, diff_with_b = 0, -inf
                for ch in s:
                    if ch == a:
                        diff += 1
                        diff_with_b += 1
                    elif ch == b:
                        diff -= 1
                        diff_with_b = diff  # 记录包含 b 时的 diff
                        if diff < 0:
                            diff = 0
                    if diff_with_b > ans:
                        ans = diff_with_b
        return ans
```

```java [sol1-Java]
class Solution {
    public int largestVariance(String s) {
        var ans = 0;
        for (var a = 'a'; a <= 'z'; ++a)
            for (var b = 'a'; b <= 'z'; ++b) {
                if (a == b) continue;
                var diff = 0;
                var diffWithB = -s.length();
                for (var i = 0; i < s.length(); i++) {
                    if (s.charAt(i) == a) {
                        ++diff;
                        ++diffWithB;
                    } else if (s.charAt(i) == b) {
                        diffWithB = --diff;
                        diff = Math.max(diff, 0);
                    }
                    ans = Math.max(ans, diffWithB);
                }
            }
        return ans;
    }
}
```

```C++ [sol1-C++]
class Solution {
public:
    int largestVariance(string &s) {
        int ans = 0;
        for (char a = 'a'; a <= 'z'; ++a)
            for (char b = 'a'; b <= 'z'; ++b) {
                if (a == b) continue;
                int diff = 0, diff_with_b = -s.length();
                for (char ch : s) {
                    if (ch == a) {
                        ++diff;
                        ++diff_with_b;
                    } else if (ch == b) {
                        diff_with_b = --diff;
                        diff = max(diff, 0);
                    }
                    ans = max(ans, diff_with_b);
                }
            }
        return ans;
    }
};
```

```go [sol1-Go]
func largestVariance2(s string) (ans int) {
	for a := 'a'; a <= 'z'; a++ {
		for b := 'a'; b <= 'z'; b++ {
			if b == a {
				continue
			}
			diff, diffWithB := 0, -len(s)
			for _, ch := range s {
				if ch == a {
					diff++
					diffWithB++
				} else if ch == b {
					diff--
					diffWithB = diff // 记录包含 b 时的 diff
					diff = max(diff, 0)
				}
				ans = max(ans, diffWithB)
			}
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

我们可以在遍历 $s$ 的过程中将 $s[i]$ 作为 $a$ 或 $b$，从而优化时间复杂度。

```Python [sol1-Python3]
class Solution:
    def largestVariance(self, s: str) -> int:
        if s.count(s[0]) == len(s):
            return 0
        ans = 0
        diff = [[0] * 26 for _ in range(26)]
        diff_with_b = [[-inf] * 26 for _ in range(26)]
        for ch in s:
            ch = ord(ch) - ord('a')
            for i in range(26):
                if i == ch:
                    continue
                diff[ch][i] += 1  # a=ch, b=i
                diff_with_b[ch][i] += 1
                diff[i][ch] -= 1  # a=i, b=ch
                diff_with_b[i][ch] = diff[i][ch]
                if diff[i][ch] < 0:
                    diff[i][ch] = 0
                ans = max(ans, diff_with_b[ch][i], diff_with_b[i][ch])
        return ans
```

```java [sol1-Java]
class Solution {
    public int largestVariance(String s) {
        var ans = 0;
        var diff = new int[26][26];
        var diffWithB = new int[26][26];
        for (var i = 0; i < 26; i++) Arrays.fill(diffWithB[i], -s.length());
        for (var k = 0; k < s.length(); k++) {
            var ch = s.charAt(k) - 'a';
            for (var i = 0; i < 26; ++i) {
                if (i == ch) continue;
                ++diff[ch][i]; // a=ch, b=i
                ++diffWithB[ch][i];
                diffWithB[i][ch] = --diff[i][ch]; // a=i, b=ch
                diff[i][ch] = Math.max(diff[i][ch], 0);
                ans = Math.max(ans, Math.max(diffWithB[ch][i], diffWithB[i][ch]));
            }
        }
        return ans;
    }
}
```

```C++ [sol1-C++]
class Solution {
public:
    int largestVariance(string &s) {
        int ans = 0;
        int diff[26][26] = {}, diff_with_b[26][26];
        memset(diff_with_b, 0x80, sizeof(diff_with_b));
        for (char ch : s) {
            ch -= 'a';
            for (char i = 0; i < 26; ++i) {
                if (i == ch) continue;
                ++diff[ch][i]; // a=ch, b=i
                ++diff_with_b[ch][i];
                diff_with_b[i][ch] = --diff[i][ch]; // a=i, b=ch
                diff[i][ch] = max(diff[i][ch], 0);
                ans = max(ans, max(diff_with_b[ch][i], diff_with_b[i][ch]));
            }
        }
        return ans;
    }
};
```

```go [sol1-Go]
func largestVariance(s string) (ans int) {
	var diff, diffWithB [26][26]int
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			diffWithB[i][j] = -len(s)
		}
	}
	for _, ch := range s {
		ch -= 'a'
		for i := rune(0); i < 26; i++ {
			if i == ch {
				continue
			}
			diff[ch][i]++ // a=ch, b=i
			diffWithB[ch][i]++
			diff[i][ch]-- // a=i, b=ch
			diffWithB[i][ch] = diff[i][ch]
			diff[i][ch] = max(diff[i][ch], 0)
			ans = max(ans, max(diffWithB[ch][i], diffWithB[i][ch]))
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

