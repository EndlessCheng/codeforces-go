注意 $\textit{groups}$ 中只有 $0$ 和 $1$。

例如示例 2 的 $1011$ 有 $3$ 个连续相同段 $1,0,11$，对应的最长子序列的长度为 $3$。每一段选一个下标，从 $\textit{words}$ 中得到对应的字符串，加入答案。

一般地，形如 $000111000111\cdots$ 中的每一个连续相同段只能选一个下标。为了让子序列尽量长，每个连续相同段都必须选一个下标。

```py [sol-Python3]
class Solution:
    def getWordsInLongestSubsequence(self, n: int, words: List[str], groups: List[int]) -> List[str]:
        return [w for (x, y), w in zip(pairwise(groups), words) if x != y] + [words[-1]]
```

```java [sol-Java]
class Solution {
    public List<String> getWordsInLongestSubsequence(int n, String[] words, int[] groups) {
        List<String> ans = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            if (i == n - 1 || groups[i] != groups[i + 1]) {
                ans.add(words[i]);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<string> getWordsInLongestSubsequence(int n, vector<string>& words, vector<int>& groups) {
        vector<string> ans;
        for (int i = 0; i < n; i++) {
            if (i == n - 1 || groups[i] != groups[i + 1]) {
                ans.push_back(words[i]);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func getWordsInLongestSubsequence(n int, words []string, groups []int) (ans []string) {
	for i, x := range groups {
		if i == n-1 || x != groups[i+1] {
			ans = append(ans, words[i])
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。
