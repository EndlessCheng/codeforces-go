下午 2 点在 B 站直播讲周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

贪心策略：

把 $\textit{pattern}$ 按照 $\texttt{III}\cdots \texttt{IDDD}\cdots \texttt{D}$ 分组，每组前一段是 $\texttt{I}$，后一段是 $\texttt{D}$。

遍历每一段，设当前段的长度为 $x$，我们应该把剩余最小的 $x$ 个数字填到该段上（如果是第一段则填最小的 $x+1$ 个数字），从而保证字典序最小。

假设第一段为 $\texttt{IIIDDD}$，构造方案如下：

- 前 $2$ 个 $\texttt{I}$ 视作长度为 $3$ 的上升段；
- 剩余的 $\texttt{I}$ 和 $\texttt{D}$ 视作长度为 $4$ 的下降段；
- 最小的 $3$ 个数给上升段，然后剩余最小的 $4$ 个数给下降段；
- 构造结果为 $\texttt{1237654}$。

按照该方案分组模拟即可。

```py [sol1-Python3]
class Solution:
    def smallestNumber(self, pattern: str) -> str:
        i, cur, n = 0, 1, len(pattern)
        ans = [''] * (n + 1)
        while i < n:
            if i and pattern[i] == 'I':
                i += 1
            while i < n and pattern[i] == 'I':
                ans[i] = digits[cur]
                cur += 1
                i += 1
            i0 = i
            while i < n and pattern[i] == 'D':
                i += 1
            for j in range(i, i0 - 1, -1):
                ans[j] = digits[cur]
                cur += 1
        return ''.join(ans)
```

```java [sol1-Java]
class Solution {
    public String smallestNumber(String pattern) {
        int i = 0, n = pattern.length();
        var cur = '1';
        var ans = new char[n + 1];
        while (i < n) {
            if (i > 0 && pattern.charAt(i) == 'I') ++i;
            for (; i < n && pattern.charAt(i) == 'I'; ++i) ans[i] = cur++;
            var i0 = i;
            while (i < n && pattern.charAt(i) == 'D') ++i;
            for (var j = i; j >= i0; --j) ans[j] = cur++;
        }
        return new String(ans);
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    string smallestNumber(string pattern) {
        int i = 0, n = pattern.length();
        char cur = '1';
        string ans(n + 1, 0);
        while (i < n) {
            if (i && pattern[i] == 'I') ++i;
            for (; i < n && pattern[i] == 'I'; ++i) ans[i] = cur++;
            int i0 = i;
            while (i < n && pattern[i] == 'D') ++i;
            for (int j = i; j >= i0; --j) ans[j] = cur++;
        }
        return ans;
    }
};
```

```go [sol1-Go]
func smallestNumber(pattern string) string {
	n := len(pattern)
	ans := make([]byte, n+1)
	for i, cur := 0, byte('1'); i < n; {
		if i > 0 && pattern[i] == 'I' {
			i++
		}
		for ; i < n && pattern[i] == 'I'; i++ {
			ans[i] = cur
			cur++
		}
		i0 := i
		for ; i < n && pattern[i] == 'D'; i++ {
		}
		for j := i; j >= i0; j-- {
			ans[j] = cur
			cur++
		}
	}
	return string(ans)
}
```
