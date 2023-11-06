遍历下标在 $[\textit{left},\textit{right}]$ 内的字符串 $s=\textit{words}[i]$，如果 $s$ 的第一个字母和最后一个字母都是元音，把答案加一。

```py [sol-Python3]
class Solution:
    def vowelStrings(self, words: List[str], left: int, right: int) -> int:
        return sum(s[0] in "aeiou" and s[-1] in "aeiou" for s in words[left:right+1])
```

```java [sol-Java]
class Solution {
    private static final String VOWEL = "aeiou";

    public int vowelStrings(String[] words, int left, int right) {
        int ans = 0;
        for (int i = left; i <= right; i++) {
            String s = words[i];
            if (VOWEL.indexOf(s.charAt(0)) != -1 &&
                VOWEL.indexOf(s.charAt(s.length() - 1)) != -1) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int vowelStrings(vector<string> &words, int left, int right) {
        const string vowel = "aeiou";
        int ans = 0;
        for (int i = left; i <= right; i++) {
            string &s = words[i];
            ans += vowel.find(s[0]) != string::npos &&
                   vowel.find(s.back()) != string::npos;
        }
        return ans;
    }
};
```

```go [sol-Go]
func vowelStrings(words []string, left, right int) (ans int) {
	for _, s := range words[left : right+1] {
		if strings.Contains("aeiou", s[:1]) && strings.Contains("aeiou", s[len(s)-1:]) {
			ans++
		}
	}
	return
}
```

```js [sol-JavaScript]
var vowelStrings = function(words, left, right) {
    let ans = 0;
    for (let i = left; i <= right; i++) {
        const s = words[i];
        if ("aeiou".includes(s[0]) && "aeiou".includes(s[s.length - 1])) {
            ans++;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn vowel_strings(words: Vec<String>, left: i32, right: i32) -> i32 {
        let mut ans = 0;
        for i in left..=right {
            let s = &words[i as usize];
            if "aeiou".contains(s.chars().next().unwrap()) &&
               "aeiou".contains(s.chars().last().unwrap()) {
                ans += 1;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\textit{right}-\textit{left})$。注意每个字符串只取第一个和最后一个字母，不会遍历整个字符串，所以处理每个字符串的时间是 $\mathcal{O}(1)$ 的。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
