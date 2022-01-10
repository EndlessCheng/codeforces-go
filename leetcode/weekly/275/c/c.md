逆向思维

考虑 $\textit{targetWords}[i]$ 能否还原回 $\textit{startWords}$ 中的某个字符串。

由于可以任意重排，我们可以对 $\textit{startWords}$ 和 $\textit{targetWords}$ 的每个字符串都排序。

由于需要追加字符，且题目保证所有字符串都没有重复字符，因此我们可以枚举排序后的 $\textit{targetWords}[i]$ 的所有字符，将其去掉后去看看是否在每个字符串都排序后的 $\textit{startWords}$ 中存在。这可以用哈希表实现。

代码实现时，我们并不需要排序每个字符串，而是记录每个字符是否出现过，这可以用位运算实现。

```go [sol1-Go]
func wordCount(startWords, targetWords []string) (ans int) {
	has := map[int]bool{}
	for _, word := range startWords {
		mask := 0
		for _, ch := range word {
			mask |= 1 << (ch - 'a')
		}
		has[mask] = true
	}
	for _, word := range targetWords {
		mask := 0
		for _, ch := range word {
			mask |= 1 << (ch - 'a')
		}
		for i := 0; i < 26; i++ {
			if mask&(1<<i) > 0 && has[mask^(1<<i)] { // 去掉这个字符
				ans++
				break
			}
		}
	}
	return
}
```

```C++ [sol1-C++]
class Solution {
public:
    int wordCount(vector<string> &startWords, vector<string> &targetWords) {
        unordered_set<int> s;
        for (string &word : startWords) {
            int mask = 0;
            for (char ch : word) {
                mask |= 1 << (ch - 'a');
            }
            s.insert(mask);
        }
        int ans = 0;
        for (string &word : targetWords) {
            int mask = 0;
            for (char ch : word) {
                mask |= 1 << (ch - 'a');
            }
            for (int i = 0; i < 26; ++i) {
                if (mask & (1 << i) && s.count(mask ^ (1 << i))) { // 去掉这个字符
                    ++ans;
                    break;
                }
            }
        }
        return ans;
    }
};
```

```Python [sol1-Python3]
class Solution:
    def wordCount(self, startWords: List[str], targetWords: List[str]) -> int:
        s = set()
        for word in startWords:
            mask = 0
            for ch in word:
                mask |= 1 << (ord(ch) - ord('a'))
            s.add(mask)
        ans = 0
        for word in targetWords:
            mask = 0
            for ch in word:
                mask |= 1 << (ord(ch) - ord('a'))
            for i in range(26):
                if mask & (1 << i) and mask ^ (1 << i) in s:  # 去掉这个字符
                    ans += 1
                    break
        return ans
```

