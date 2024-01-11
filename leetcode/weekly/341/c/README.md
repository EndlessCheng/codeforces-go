[视频讲解](https://www.bilibili.com/video/BV1ng4y1T7QA/) 第三题。

## 方法一：考虑相邻字母

下文将 $\textit{word}$ 简记为 $s$。

考察两个相邻字母，分别设为 $x=s[i-1]$ 和 $y=s[i]$。

使 $s$ 有效的话，我们需要在 $x$ 和 $y$ 之间插入

$$
y-x-1
$$

个字母。

考虑到这可能是个负数，可以通过如下技巧转换在 $[0,2]$ 内：

$$
(y-x-1+3)\bmod 3
$$

- 例如 $x=\text{`a'},y=\text{`c'}$，则 $(\text{`c'}-\text{`a'}+2)\bmod 3 = 1$，表示需要插入 $1$ 个字母，即字母 $\text{`b'}$。
- 例如 $x=\text{`a'},y=\text{`a'}$，则 $(\text{`a'}-\text{`a'}+2)\bmod 3 = 2$，表示需要插入 $2$ 个字母，即字母 $\text{`b'}$ 和字母 $\text{`c'}$。
- 例如 $x=\text{`c'},y=\text{`a'}$，则 $(\text{`a'}-\text{`c'}+2)\bmod 3 = 0$，表示无需插入字母。

最后，如果 $s[0]$ 不是 $\text{`a'}$，那么需要在 $s$ 前面插入 $s[0]-\text{`a'}$ 个字母，$s[n-1]$ 也同理，需要插入 $\text{`c'}-s[n-1]$ 个字母。这俩可以合并为 $s[0]-s[n-1]+2$。

```py [sol-Python3]
class Solution:
    def addMinimum(self, s: str) -> int:
        ans = ord(s[0]) - ord(s[-1]) + 2
        for x, y in pairwise(map(ord, s)):
            ans += (y - x + 2) % 3
        return ans
```

```java [sol-Java]
class Solution {
    public int addMinimum(String word) {
        char[] s = word.toCharArray();
        int ans = s[0] + 2 - s[s.length - 1];
        for (int i = 1; i < s.length; i++) {
            ans += (s[i] + 2 - s[i - 1]) % 3;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int addMinimum(string s) {
        int ans = s[0] + 2 - s.back();
        for (int i = 1; i < s.length(); i++) {
            ans += (s[i] + 2 - s[i - 1]) % 3;
        }
        return ans;
    }
};
```

```go [sol-Go]
func addMinimum(s string) int {
	ans := int(s[0]) - int(s[len(s)-1]) + 2
	for i := 1; i < len(s); i++ {
		ans += (int(s[i]) - int(s[i-1]) + 2) % 3
	}
	return ans
}
```

```js [sol-JavaScript]
var addMinimum = function(s) {
    let ans = s.charCodeAt(0) + 2 - s.charCodeAt(s.length - 1);
    for (let i = 1; i < s.length; i++) {
        ans += (s.charCodeAt(i) + 2 - s.charCodeAt(i - 1)) % 3;
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn add_minimum(word: String) -> i32 {
        let s = word.as_bytes();
        let mut ans = s[0] as i32 - *s.last().unwrap() as i32 + 2;
        for i in 1..s.len() {
            ans += (s[i] as i32 - s[i - 1] as i32 + 2) % 3;
        }
        ans
    }
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{word}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

## 方法二：考虑 abc 的个数

假设答案由 $t$ 个 $\text{``abc''}$ 组成，那么需要插入的字母个数为 $3t-n$。

对于两个相邻字母 $x$ 和 $y$（$x$ 在 $y$ 左侧）：

- 如果 $x<y$，那么 $x$ 和 $y$ 可以在同一个 $\text{``abc''}$ 内。
- 如果 $x\ge y$，那么 $x$ 和 $y$ 一定不在同一个 $\text{``abc''}$ 内。

例如 $s=\text{``caa''}$ 中的 $s[0]\ge s[1], s[1]\ge s[2]$，所以需要 $t=3$ 个 $\text{``abc''}$，即 $\text{``ab}\textbf{ca}\text{bc}\textbf{a}\text{bc''}$。

所以 $t$ 就是 $x\ge y$ 的次数加一。

```py [sol-Python3]
class Solution:
    def addMinimum(self, s: str) -> int:
        t = 1 + sum(x >= y for x, y in pairwise(s))
        return t * 3 - len(s)
```

```java [sol-Java]
class Solution {
    public int addMinimum(String word) {
        char[] s = word.toCharArray();
        int t = 1;
        for (int i = 1; i < s.length; i++) {
            if (s[i - 1] >= s[i]) { // 必须生成一个新的 abc
                t++;
            }
        }
        return t * 3 - s.length;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int addMinimum(string s) {
        int t = 1;
        for (int i = 1; i < s.length(); i++) {
            t += s[i - 1] >= s[i]; // 必须生成一个新的 abc
        }
        return t * 3 - s.length();
    }
};
```

```go [sol-Go]
func addMinimum(s string) int {
	t := 1
	for i := 1; i < len(s); i++ {
		if s[i-1] >= s[i] { // 必须生成一个新的 abc
			t++
		}
	}
	return t*3 - len(s)
}
```

```js [sol-JavaScript]
var addMinimum = function(s) {
    let t = 1;
    for (let i = 1; i < s.length; i++) {
        t += s[i - 1] >= s[i]; // 必须生成一个新的 abc
    }
    return t * 3 - s.length;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn add_minimum(word: String) -> i32 {
        let s = word.as_bytes();
        let mut t = 1;
        for i in 1..s.len() {
            if s[i - 1] >= s[i] {
                t += 1; // 必须生成一个新的 abc
            }
        }
        return t * 3 - s.len() as i32;
    }
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{word}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

更多精彩题解，请看 [往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
