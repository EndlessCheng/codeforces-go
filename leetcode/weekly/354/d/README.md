[视频讲解](https://www.bilibili.com/video/BV1DM4y1x7bR/) 第四题。

## 提示 1

$\textit{forbidden}[i]$ 的长度至多为 $10$。

## 提示 2

[滑动窗口](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

## 提示 3

初始化子串左端点 $\textit{left}=0$，枚举子串右端点 $\textit{right}$。

对于示例 2，只要 $\textit{right}\ge 1$，那么合法子串是不能包含 $\texttt{le}$ 的，所以左端点 $\textit{left}$ 必须向右移，不可能再回到 $0$（否则就包含 $\texttt{le}$ 了）。因为左端点只会向右移动，不会向左移动，这样的**单调性**保证了算法的效率。

当 $\textit{right}$ 右移到一个新的字母时，**枚举**以该字母为右端点的 $\textit{forbidden}[i]$ 的最短长度。如果发现子串 $\textit{word}[i]$ 到 $\textit{word}[\textit{right}]$ 在 $\textit{forbidden}$ 中（用哈希表实现），那么更新 $\textit{left}=i+1$ 并结束枚举，从而避免合法子串包含 $\textit{forbidden}$ 中的字符串。枚举结束后，更新答案为合法子串长度 $\textit{right}-\textit{left}+1$ 的最大值。

```py [sol-Python3]
class Solution:
    def longestValidSubstring(self, word: str, forbidden: List[str]) -> int:
        fb = set(forbidden)
        ans = left = 0
        for right in range(len(word)):
            for i in range(right, max(right - 10, left - 1), -1):
                if word[i: right + 1] in fb:
                    left = i + 1  # 当子串右端点 >= right 时，合法子串一定不能包含 word[i]
                    break
            ans = max(ans, right - left + 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int longestValidSubstring(String word, List<String> forbidden) {
        var fb = new HashSet<String>();
        fb.addAll(forbidden);
        int ans = 0, left = 0, n = word.length();
        for (int right = 0; right < n; right++) {
            for (int i = right; i >= left && i > right - 10; i--) {
                if (fb.contains(word.substring(i, right + 1))) {
                    left = i + 1; // 当子串右端点 >= right 时，合法子串一定不能包含 word[i]
                    break;
                }
            }
            ans = Math.max(ans, right - left + 1);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestValidSubstring(string word, vector<string> &forbidden) {
        unordered_set<string> fb{forbidden.begin(), forbidden.end()};
        int ans = 0, left = 0, n = word.length();
        for (int right = 0; right < n; right++) {
            for (int i = right; i >= left && i > right - 10; i--) {
                if (fb.count(word.substr(i, right - i + 1))) {
                    left = i + 1; // 当子串右端点 >= right 时，合法子串一定不能包含 word[i]
                    break;
                }
            }
            ans = max(ans, right - left + 1);
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestValidSubstring(word string, forbidden []string) (ans int) {
	has := make(map[string]bool, len(forbidden))
	for _, s := range forbidden {
		has[s] = true
	}

	left := 0
	for right := range word {
		for i := right; i >= left && i > right-10; i-- {
			if has[word[i:right+1]] {
				left = i + 1 // 当子串右端点 >= right 时，合法子串一定不能包含 word[i]
				break
			}
		}
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

```js [sol-JavaScript]
var longestValidSubstring = function (word, forbidden) {
    let fb = new Set();
    for (const f of forbidden) fb.add(f);
    const n = word.length;
    let ans = 0, left = 0;
    for (let right = 0; right < n; right++) {
        for (let i = right; i >= left && i > right - 10; i--) {
            if (fb.has(word.substring(i, right + 1))) {
                left = i + 1; // 当子串右端点 >= right 时，合法子串一定不能包含 word[i]
                break;
            }
        }
        ans = Math.max(ans, right - left + 1);
    }
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L+nM^2)$，其中 $L$ 为所有 $\textit{forbidden}[i]$ 的长度之和，$n$ 为 $\textit{word}$ 的长度，$M=10$ 表示 $\textit{forbidden}[i]$ 的最长长度。请注意，在哈希表中查询一个长为 $M$ 的字符串的时间是 $\mathcal{O}(M)$，每次移动右指针会执行至多 $M$ 次这样的查询。
- 空间复杂度：$\mathcal{O}(L)$。

更多滑窗题目，请看[【题单】滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
