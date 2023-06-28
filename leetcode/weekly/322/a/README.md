转换成判断：

1. 首尾字符是否相同。
2. 每个空格左右的字符是否相同。

```py [sol-Python3]
class Solution:
    def isCircularSentence(self, s: str) -> bool:
        return s[0] == s[-1] and all(c != ' ' or s[i-1] == s[i+1] for i, c in enumerate(s))
```

```java [sol-Java]
class Solution {
    public boolean isCircularSentence(String sentence) {
        var s = sentence.toCharArray();
        int n = s.length;
        if (s[0] != s[n - 1])
            return false;
        for (int i = 1; i < n - 1; i++)
            if (s[i] == ' ' && s[i - 1] != s[i + 1])
                return false;
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool isCircularSentence(string s) {
        if (s[0] != s.back())
            return false;
        for (int i = 1, n = s.length(); i < n - 1; i++)
            if (s[i] == ' ' && s[i - 1] != s[i + 1])
                return false;
        return true;
    }
};
```

```go [sol-Go]
func isCircularSentence(s string) bool {
	if s[0] != s[len(s)-1] {
		return false
	}
	for i, c := range s {
		if c == ' ' && s[i-1] != s[i+1] {
			return false
		}
	}
	return true
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{sentence}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干额外变量。

---

[往期每日一题题解（按 tag 分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注[ biIibiIi@灵茶山艾府](https://space.bilibili.com/206214)，高质量算法教学，持续输出中~
