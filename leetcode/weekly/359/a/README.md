请看 [视频讲解](https://www.bilibili.com/video/BV1Rx4y1f75Y/)。

如果 $\textit{words}$ 长度和 $s$ 长度不同，直接返回 `false`。

如果存在 $\textit{words}[i][0] \ne s[i]$，也返回 `false`。

否则返回 `true`。

```py [sol-Python3]
class Solution:
    def isAcronym(self, words: List[str], s: str) -> bool:
        return len(words) == len(s) and all(w[0] == c for w, c in zip(words, s))
```

```java [sol-Java]
class Solution {
    public boolean isAcronym(List<String> words, String s) {
        if (words.size() != s.length())
            return false;
        for (int i = 0; i < words.size(); i++)
            if (words.get(i).charAt(0) != s.charAt(i))
                return false;
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool isAcronym(vector<string> &words, string s) {
        if (words.size() != s.length())
            return false;
        for (int i = 0; i < words.size(); i++)
            if (words[i][0] != s[i])
                return false;
        return true;
    }
};
```

```go [sol-Go]
func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}
	for i, w := range words {
		if w[0] != s[i] {
			return false
		}
	}
	return true
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
