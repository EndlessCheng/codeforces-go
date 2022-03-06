#### 提示 1

考虑 $s[0]$，如果存在另一个和 $s[0]$ 相同的字符，要将它们交换到关于 $s$ 中心对称的位置上，至少需要交换多少次？

#### 提示 2

去掉这两个字符后，重复上述过程。（注意可能会有一个只出现一次的字符，如何处理这种情况？）

---

设 $s$ 的长度为 $n$。若能从 $s$ 中选两个相同的字符 $s[0]$ 和 $s[i]$，将它们交换到关于中心对称的位置上，那么最小的交换次数为 $n-1-i$。我们可以通过固定 $s[0]$，将 $s[i]$ 交换到字符串末尾来做到。

去掉这两个字符后，重复上述过程。

由于交换后去掉字符，得到的字符串和交换前去掉字符是相同的，所以实际上可以不用模拟交换操作，直接去掉字符。

若有一个只出现一次的字符，由于题目保证一定能变成回文串，那么其只能是回文串的中心。

为方便删除操作，下面代码每次固定的是 $s[n-1]$。

```go [sol1-Go]
func minMovesToMakePalindrome(s string) (ans int) {
	for s != "" {
		i := strings.IndexByte(s, s[len(s)-1])
		if i == len(s)-1 { // 只出现一次的字符
			ans += i / 2 // 交换到回文中心上
		} else {
			ans += i // 交换到字符串开头
			s = s[:i] + s[i+1:] // 移除 s[i]
		}
		s = s[:len(s)-1] // 移除最后一个字符
	}
	return
}
```

```C++ [sol1-C++]
class Solution {
public:
    int minMovesToMakePalindrome(string s) {
        int ans = 0;
        while (!s.empty()) {
            int i = s.find(s.back());
            if (i == s.length() - 1) { // 只出现一次的字符
                ans += i / 2; // 交换到回文中心上
            } else {
                ans += i; // 交换到字符串开头
                s.erase(i, 1); // 移除 s[i]
            }
            s.pop_back(); // 移除最后一个字符
        }
        return ans;
    }
};
```

```Python [sol1-Python3]
class Solution:
    def minMovesToMakePalindrome(self, s: str) -> int:
        ans = 0
        while s:
            i = s.index(s[-1])
            if i == len(s) - 1:  # 只出现一次的字符
                ans += i // 2  # 交换到回文中心上
            else:
                ans += i  # 交换到字符串开头
                s = s[:i] + s[i + 1:]  # 移除 s[i]
            s = s[:-1]  # 移除最后一个字符
        return ans
```

```java [sol1-Java]
class Solution {
    public int minMovesToMakePalindrome(String S) {
        var ans = 0;
        var s = new ArrayList<>(S.chars().boxed().toList());
        while (s.size() > 0) {
            var i = s.indexOf(s.get(s.size() - 1));
            if (i == s.size() - 1) { // 只出现一次的字符
                ans += i / 2; // 交换到回文中心上
            } else {
                ans += i; // 交换到字符串开头
                s.remove(i); // 移除 s[i]
            }
            s.remove(s.size() - 1); // 移除最后一个字符
        }
        return ans;
    }
}
```
