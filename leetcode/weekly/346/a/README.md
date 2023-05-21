## 本题视频讲解

见[【周赛 346】](https://www.bilibili.com/video/BV1Qm4y1t7cx/)第一题，欢迎点赞投币！

## 算法一：暴力

本题思路和 [20. 有效的括号](https://leetcode.cn/problems/valid-parentheses/) 是一样的。

暴力做法是不断把 AB 和 CD 去掉，直到 $s$ 中没有 AB 和 CD 为止。

```py [sol1-Python3]
class Solution:
    def minLength(self, s: str) -> int:
        while "AB" in s or "CD" in s:
            s = s.replace("AB", "").replace("CD", "")
        return len(s)
```

```java [sol1-Java]
class Solution {
    public int minLength(String s) {
        while (s.contains("AB") || s.contains("CD"))
            s = s.replace("AB", "").replace("CD", "");
        return s.length();
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int minLength(string s) {
        regex ab("AB"), cd("CD");
        while (s.find("AB") != string::npos || s.find("CD") != string::npos) {
            s = regex_replace(s, ab, "");
            s = regex_replace(s, cd, "");
        }
        return s.length();
    }
};
```

```go [sol1-Go]
func minLength(s string) int {
	for strings.Contains(s, "AB") || strings.Contains(s, "CD") {
		s = strings.ReplaceAll(s, "AB", "")
		s = strings.ReplaceAll(s, "CD", "")
	}
	return len(s)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $s$ 的长度。对于 AAA...BBB 这样的字符串，需要循环 $\mathcal{O}(n)$ 次，每次需要 $\mathcal{O}(n)$ 的时间。
- 空间复杂度：$\mathcal{O}(n)$。替换过程中生成的字符串需要 $\mathcal{O}(n)$ 的空间。

## 算法二：栈

用栈记录遍历过的，没有删除的字母。

如果当前字母是 B，且栈顶为 A，那么这两个字母都可以删除。同理，如果当前字母是 D，且栈顶为 C，那么这两个字母都可以删除。

否则，把当前字母入栈。

```py [sol-Python3]
class Solution:
    def minLength(self, s: str) -> int:
        st = []
        for c in s:
            if st and (c == 'B' and st[-1] == 'A' or c == 'D' and st[-1] == 'C'):
                st.pop()
            else:
                st.append(c)
        return len(st)
```

```java [sol-Java]
class Solution {
    public int minLength(String s) {
        var st = new ArrayDeque<Character>();
        for (var c : s.toCharArray()) {
            if (!st.isEmpty() && (c == 'B' && st.peek() == 'A' || c == 'D' && st.peek() == 'C'))
                st.pop();
            else
                st.push(c);
        }
        return st.size();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minLength(string s) {
        stack<char> st;
        for (char c: s) {
            if (!st.empty() && (c == 'B' && st.top() == 'A' || c == 'D' && st.top() == 'C'))
                st.pop();
            else
                st.push(c);
        }
        return st.size();
    }
};
```

```go [sol-Go]
func minLength(s string) int {
	st := []rune{}
	for _, c := range s {
		if len(st) > 0 && (c == 'B' && st[len(st)-1] == 'A' || c == 'D' && st[len(st)-1] == 'C') {
			st = st[:len(st)-1]
		} else {
			st = append(st, c)
		}
	}
	return len(st)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
