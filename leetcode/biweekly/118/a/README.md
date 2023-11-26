遍历 $\textit{words}$，判断 $x$ 是否在 $\textit{words}[i]$ 中，如果是则把 $i$ 加入答案。

```py [sol-Python3]
class Solution:
    def findWordsContaining(self, words: List[str], x: str) -> List[int]:
        return [i for i, s in enumerate(words) if x in s]
```

```java [sol-Java]
class Solution {
    public List<Integer> findWordsContaining(String[] words, char x) {
        List<Integer> ans = new ArrayList<>();
        for (int i = 0; i < words.length; i++) {
            if (words[i].indexOf(x) != -1) {
                ans.add(i);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findWordsContaining(vector<string> &words, char x) {
        vector<int> ans;
        for (int i = 0; i < words.size(); i++) {
            auto &s = words[i];
            if (find(s.begin(), s.end(), x) != s.end()) {
                ans.push_back(i);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func findWordsContaining(words []string, x byte) (ans []int) {
	for i, s := range words {
		if strings.IndexByte(s, x) >= 0 {
			ans = append(ans, i)
		}
	}
	return
}
```

```js [sol-JavaScript]
var findWordsContaining = function(words, x) {
    const ans = [];
    for (let i = 0; i < words.length; i++) {
        if (words[i].includes(x)) {
            ans.push(i);
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_words_containing(words: Vec<String>, x: char) -> Vec<i32> {
        let mut ans = Vec::new();
        for (i, s) in words.iter().enumerate() {
            if s.chars().any(|c| c == x) {
                ans.push(i as i32);
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L)$，其中 $L$ 为所有字符串的长度之和。
- 空间复杂度：$\mathcal{O}(1)$。
