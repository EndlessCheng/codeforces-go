[双周赛视频讲解](https://www.bilibili.com/video/BV18u411Y7Gt/)

## 前置知识：滑动窗口

请看[【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)

## 思路

移动右指针 $\textit{right}$，并统计相邻相同的情况出现了多少次，记作 $\textit{same}$

如果 $\textit{same}>1$，则不断移动左指针 $\textit{left}$ 直到 $s[\textit{left}]=s[\textit{left}-1]$，此时将一对相同的字符移到窗口之外。然后将 $\textit{same}$ 置为 $1$。

然后统计子串长度 $\textit{right}-\textit{left}+1$ 的最大值。

```py [sol-Python3]
class Solution:
    def longestSemiRepetitiveSubstring(self, s: str) -> int:
        ans, left, same = 1, 0, 0
        for right in range(1, len(s)):
            same += s[right] == s[right - 1]
            if same > 1:
                left += 1
                while s[left] != s[left - 1]:
                    left += 1
                same = 1
            ans = max(ans, right - left + 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int longestSemiRepetitiveSubstring(String S) {
        var s = S.toCharArray();
        int ans = 1, left = 0, same = 0, n = s.length;
        for (int right = 1; right < n; right++) {
            if (s[right] == s[right - 1] && ++same > 1) {
                for (left++; s[left] != s[left - 1]; left++);
                same = 1;
            }
            ans = Math.max(ans, right - left + 1);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestSemiRepetitiveSubstring(string s) {
        int ans = 1, left = 0, same = 0, n = s.length();
        for (int right = 1; right < n; right++) {
            if (s[right] == s[right - 1] && ++same > 1) {
                for (left++; s[left] != s[left - 1]; left++);
                same = 1;
            }
            ans = max(ans, right - left + 1);
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestSemiRepetitiveSubstring(s string) int {
	ans, left, same := 1, 0, 0
	for right := 1; right < len(s); right++ {
		if s[right] == s[right-1] {
			same++
			if same > 1 {
				left++
				for s[left] != s[left-1] {
					left++
				}
				same = 1
			}
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。注意 $\textit{left}$ 只会增加不会减少，所以二重循环的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
