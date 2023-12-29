## 本题视频讲解

请看[【周赛 339】](https://www.bilibili.com/video/BV1va4y1M7Fr/)

## 思路

记录上一段连续相同字符个数 $\textit{pre}$，以及当前连续相同字符个数 $\textit{cur}$。

如果当前字符是 $1$，那么上一段的字符是 $0$，这两段可以组成一个 $01$ 串，由于 $0$ 和 $1$ 的个数需要相等，那么当前这个 $01$ 串的最大长度就是 

$$
2\cdot \min(\textit{pre}, \textit{cur})
$$

取所有长度的最大值，即为答案。

```py [sol-Python3]
class Solution:
    def findTheLongestBalancedSubstring(self, s: str) -> int:
        ans = pre = cur = 0
        for i, c in enumerate(s):
            cur += 1
            if i == len(s) - 1 or c != s[i + 1]:  # i 是连续相同段的末尾
                if c == '1':
                    ans = max(ans, min(pre, cur) * 2)
                pre = cur
                cur = 0
        return ans
```

```java [sol-Java]
class Solution {
    public int findTheLongestBalancedSubstring(String S) {
        char[] s = S.toCharArray(); // 不想产生额外空间的话下面可以用 charAt
        int ans = 0, pre = 0, cur = 0, n = s.length;
        for (int i = 0; i < n; i++) {
            cur++;
            if (i == s.length - 1 || s[i] != s[i + 1]) { // i 是连续相同段的末尾
                if (s[i] == '1') {
                    ans = Math.max(ans, Math.min(pre, cur) * 2);
                }
                pre = cur;
                cur = 0;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findTheLongestBalancedSubstring(string s) {
        int ans = 0, pre = 0, cur = 0, n = s.length();
        for (int i = 0; i < n; i++) {
            cur++;
            if (i == s.length() - 1 || s[i] != s[i + 1]) { // i 是连续相同段的末尾
                if (s[i] == '1') {
                    ans = max(ans, min(pre, cur) * 2);
                }
                pre = cur;
                cur = 0;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func findTheLongestBalancedSubstring(s string) (ans int) {
	pre, cur := 0, 0
	for i, c := range s {
		cur++
		if i == len(s)-1 || byte(c) != s[i+1] { // i 是连续相同段的末尾
			if c == '1' {
				ans = max(ans, min(pre, cur)*2)
			}
			pre = cur
			cur = 0
		}
	}
	return
}
```

```js [sol-JavaScript]
var findTheLongestBalancedSubstring = function(s) {
    let ans = 0, pre = 0, cur = 0;
    for (let i = 0; i < s.length; i++) {
        cur++;
        if (i === s.length - 1 || s[i] !== s[i + 1]) { // i 是连续相同段的末尾
            if (s[i] === '1') {
                ans = Math.max(ans, Math.min(pre, cur) * 2);
            }
            pre = cur;
            cur = 0;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_the_longest_balanced_substring(s: String) -> i32 {
        let s: Vec<u8> = s.bytes().collect();
        let mut ans = 0;
        let mut pre = 0;
        let mut cur = 0;
        for (i, &c) in s.iter().enumerate() {
            cur += 1;
            if i == s.len() - 1 || c != s[i + 1] { // i 是连续相同段的末尾
                if c == '1' as u8 {
                    ans = ans.max(pre.min(cur) * 2);
                }
                pre = cur;
                cur = 0;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

更多精彩题解，请看 [往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
