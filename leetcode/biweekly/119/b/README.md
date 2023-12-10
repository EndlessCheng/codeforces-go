[本题视频讲解](https://www.bilibili.com/video/BV1dC4y1X7PE/)

下文把 $\textit{word}$ 简记为 $s$。

从左到右遍历 $s$，如果发现 $s[i-1]$ 和 $s[i]$ 近似相等，应当改 $s[i-1]$ 还是 $s[i]$？

如果改 $s[i-1]$，那么 $s[i]$ 和 $s[i+1]$ 是可能近似相等的，但如果改 $s[i]$，就可以避免 $s[i]$ 和 $s[i+1]$ 近似相等。

所以每次发现两个相邻字母近似相等，就改右边那个。

```py [sol-Python3]
class Solution:
    def removeAlmostEqualCharacters(self, s: str) -> int:
        ans = 0
        i, n = 1, len(s)
        while i < n:
            if abs(ord(s[i - 1]) - ord(s[i])) <= 1:
                ans += 1
                i += 2
            else:
                i += 1
        return ans
```

```java [sol-Java]
public class Solution {
    public int removeAlmostEqualCharacters(String s) {
        int ans = 0;
        for (int i = 1; i < s.length(); i++) {
            if (Math.abs(s.charAt(i - 1) - s.charAt(i)) <= 1) {
                ans++;
                i++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int removeAlmostEqualCharacters(string s) {
        int ans = 0;
        for (int i = 1; i < s.length(); i++) {
            if (abs(s[i - 1] - s[i]) <= 1) {
                ans++;
                i++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func removeAlmostEqualCharacters(s string) (ans int) {
	for i := 1; i < len(s); i++ {
		if abs(int(s[i-1])-int(s[i])) <= 1 {
			ans++
			i++
		}
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

- [2560. 打家劫舍 IV](https://leetcode.cn/problems/house-robber-iv/)
