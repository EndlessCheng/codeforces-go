[视频讲解](https://www.bilibili.com/video/BV15V4y1m7Sb/) 第二题。

根据题意，把 $\texttt{a}$ 替换成 $\texttt{z}$ 会让字典序变大，所以子串里面是不能包含 $\texttt{a}$ 的。

替换其它字符可以让字典序变小。

那么从左到右找到第一个不等于 $\texttt{a}$ 的字符 $s[i]$，继续向后遍历，每个字符都减一，直到 $s$ 末尾或者遇到了 $\texttt{a}$。

特别地，如果 $s$ 全为 $\texttt{a}$，由于题目要求选择的子串是非空的，且必须操作一次，那么就把最后一个 $\texttt{a}$ 改成 $\texttt{z}$。

```py [sol-Python3]
class Solution:
    def smallestString(self, s: str) -> str:
        t = list(s)
        for i, c in enumerate(t):
            if c != 'a':
                for j in range(i, len(t)):
                    if t[j] == 'a': break
                    t[j] = chr(ord(t[j]) - 1)
                return ''.join(t)
        t[-1] = 'z'
        return ''.join(t)
```

```java [sol-Java]
class Solution {
    public String smallestString(String S) {
        var s = S.toCharArray();
        int n = s.length;
        for (int i = 0; i < n; i++) {
            if (s[i] > 'a') {
                for (; i < n && s[i] > 'a'; i++)
                    s[i]--;
                return new String(s);
            }
        }
        s[n - 1] = 'z';
        return new String(s);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string smallestString(string s) {
        int n = s.length();
        for (int i = 0; i < n; i++) {
            if (s[i] > 'a') {
                for (; i < n && s[i] > 'a'; i++)
                    s[i]--;
                return s;
            }
        }
        s.back() = 'z';
        return s;
    }
};
```

```go [sol-Go]
func smallestString(s string) (ans string) {
	t := []byte(s)
	for i, c := range t {
		if c > 'a' {
			for ; i < len(t) && t[i] > 'a'; i++ {
				t[i]--
			}
			return string(t)
		}
	}
	t[len(t)-1] = 'z'
	return string(t)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$，如果可以直接修改 $s$ 则为 $\mathcal{O}(1)$ 空间（C++）。
