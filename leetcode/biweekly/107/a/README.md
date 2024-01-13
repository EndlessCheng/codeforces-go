遍历 $\textit{words}$，对于 $s=\textit{words}[i]$：

- 如果在 $s$ 之前，已经遇到 $s$ 反转后的字符串，那么就找到了一个匹配。由于题目保证 $\textit{words}$ 中的字符串互不相同，这个匹配是唯一的。
- 如果前面没有遇到，那么把 $s$ 加入哈希表（数组）$\textit{seen}$，方便在后续遍历中，快速判断是否有能匹配的字符串。

```py [sol-Python3]
class Solution:
    def maximumNumberOfStringPairs(self, words: List[str]) -> int:
        ans = 0
        seen = set()
        for s in words:
            if s[::-1] in seen:
                ans += 1  # s 和 seen 中的 s[::-1] 匹配
            else:
                seen.add(s)
        return ans
```

```java [sol-Java]
class Solution {
    public int maximumNumberOfStringPairs(String[] words) {
        int ans = 0;
        boolean[][] seen = new boolean[26][26];
        for (String s : words) {
            int x = s.charAt(0) - 'a';
            int y = s.charAt(1) - 'a';
            if (seen[y][x]) {
                ans++; // s 和 seen 中的 y+x 匹配
            } else {
                seen[x][y] = true;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumNumberOfStringPairs(vector<string>& words) {
        int ans = 0;
        bool seen[26][26]{};
        for (auto &s : words) {
            int x = s[0] - 'a';
            int y = s[1] - 'a';
            if (seen[y][x]) {
                ans++; // s 和 seen 中的 y+x 匹配
            } else {
                seen[x][y] = true;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumNumberOfStringPairs(words []string) (ans int) {
	seen := [26][26]bool{}
	for _, s := range words {
		x, y := s[0]-'a', s[1]-'a'
		if seen[y][x] {
			ans++ // s 和 seen 中的 y+x 匹配
		} else {
			seen[x][y] = true
		}
	}
	return
}
```

```js [sol-JavaScript]
var maximumNumberOfStringPairs = function(words) {
    let ans = 0;
    const seen = new Set();
    for (const s of words) {
        if (seen.has(s.split('').reverse().join(''))) {
            ans++; // s 和 seen 中的反转字符串匹配
        } else {
            seen.add(s);
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn maximum_number_of_string_pairs(words: Vec<String>) -> i32 {
        let mut ans = 0;
        let mut seen = [[false; 26]; 26];
        for s in words {
            let s = s.as_bytes();
            let x = (s[0] - b'a') as usize;
            let y = (s[1] - b'a') as usize;
            if seen[y][x] {
                ans += 1; // s 和 seen 中的 y+x 匹配
            } else {
                seen[x][y] = true;
            }
        }
        ans
    }
}
```

#### 复杂度分析（哈希表实现）

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{words}$ 的长度。字符串长度视作 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

如果 $\textit{words}$ 中有相同字符串要怎么做？

见 [视频讲解](https://www.bilibili.com/video/BV1am4y1a7Zi/)
