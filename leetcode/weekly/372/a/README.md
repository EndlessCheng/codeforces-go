[本题视频讲解](https://www.bilibili.com/video/BV1pC4y1j7Pw/)

设 $\textit{lcp}$ 为三个字符串的最长公共前缀的长度。

如果 $\textit{lcp}=0$，无法操作成一样的，返回 $-1$。

否则返回三个字符串的长度之和，减去剩下的长度 $3\cdot \textit{lcp}$。

```py [sol-Python3]
class Solution:
    def findMinimumOperations(self, s1: str, s2: str, s3: str) -> int:
        lcp = 0
        for x, y, z in zip(s1, s2, s3):
            if x != y or x != z:
                break
            lcp += 1
        return -1 if lcp == 0 else len(s1) + len(s2) + len(s3) - lcp * 3
```

```java [sol-Java]
class Solution {
    public int findMinimumOperations(String s1, String s2, String s3) {
        int n = Math.min(Math.min(s1.length(), s2.length()), s3.length());
        int i = 0;
        while (i < n && s2.charAt(i) == s1.charAt(i) && s3.charAt(i) == s1.charAt(i)) {
            i++;
        }
        return i == 0 ? -1 : s1.length() + s2.length() + s3.length() - i * 3;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findMinimumOperations(string s1, string s2, string s3) {
        int n = min({s1.length(), s2.length(), s3.length()});
        int i = 0;
        while (i < n && s2[i] == s1[i] && s3[i] == s1[i]) {
            i++;
        }
        return i == 0 ? -1 : s1.length() + s2.length() + s3.length() - i * 3;
    }
};
```

```go [sol-Go]
func findMinimumOperations(s1, s2, s3 string) int {
	n := min(len(s1), len(s2), len(s3))
	i := 0
	for i < n && s2[i] == s1[i] && s3[i] == s1[i] {
		i++
	}
	if i == 0 {
		return -1
	}
	return len(s1) + len(s2) + len(s3) - i*3
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为三个字符串中的最短字符串的长度。
- 空间复杂度：$\mathcal{O}(1)$。
