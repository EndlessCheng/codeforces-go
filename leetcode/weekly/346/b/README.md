## 本题视频讲解

见[【周赛 346】](https://www.bilibili.com/video/BV1Qm4y1t7cx/)第二题，欢迎点赞投币！

## 思路

对于两个中心对称的字母 $x=s[i]$ 和 $y=s[n-1-i]$，如果 $x\ne y$，那么只需要修改一次，就可以让这两个字母相同：把 $x$ 改成 $y$ 或者把 $y$ 改成 $x$。

- 如果 $x>y$，那么把 $x$ 修改成 $y$ 更好，这样字典序更小。
- 如果 $x<y$，那么把 $y$ 修改成 $x$ 更好，这样字典序更小。

代码实现时可以把 $x=y$ 的情况合并到 $x<y$ 中，从而少写一个 `else if` 的判断逻辑。

```py [sol-Python3]
class Solution:
    def makeSmallestPalindrome(self, s: str) -> str:
        s = list(s)
        for i in range(len(s) // 2):
            x, y = s[i], s[-1 - i]
            if x > y: s[i] = y
            else: s[-1 - i] = x
        return ''.join(s)
```

```java [sol-Java]
class Solution {
    public String makeSmallestPalindrome(String S) {
        var s = S.toCharArray();
        for (int i = 0, n = s.length; i < n / 2; i++) {
            char x = s[i], y = s[n - 1 - i];
            if (x > y) s[i] = y;
            else s[n - 1 - i] = x;
        }
        return new String(s);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string makeSmallestPalindrome(string s) {
        for (int i = 0, n = s.length(); i < n / 2; i++) {
            char x = s[i], y = s[n - 1 - i];
            if (x > y) s[i] = y;
            else s[n - 1 - i] = x;
        }
        return s;
    }
};
```

```go [sol-Go]
func makeSmallestPalindrome(S string) string {
	s := []byte(S)
	for i, n := 0, len(s); i < n/2; i++ {
		x, y := s[i], s[n-1-i]
		if x > y {
			s[i] = y
		} else {
			s[n-1-i] = x
		}
	}
	return string(s)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。取决于能否直接修改 $s$。
