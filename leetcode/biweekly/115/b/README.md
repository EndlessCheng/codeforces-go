**题意**：选一个 $\textit{words}$ 的子序列，要求相邻字符串对应的 $\textit{groups}[i]$ 不同。

为方便描述，把 $\textit{groups}$ 当作一个 $01$ 字符串。

比如 $\textit{groups}=0001100$，分成三组 $000,11,00$。根据题意，每一组内**只能选一个** $\textit{words}[i]$，所以一共选 $3$ 个字符串。如果选超过 $3$ 个字符串，根据鸽巢原理，一定有两个字符串在同一组，这违背了「相邻字符串对应的 $\textit{groups}[i]$ 不同」的要求。

所以遍历 $\textit{groups}$ 的每一个连续相同段，每一段中选一个相应的 $\textit{words}[i]$。

```py [sol-Python3]
class Solution:
    def getLongestSubsequence(self, words: List[str], groups: List[int]) -> List[str]:
        n = len(groups)
        ans = []
        for i, g in enumerate(groups):
            if i == n - 1 or g != groups[i + 1]:  # i 是连续相同段的末尾
                ans.append(words[i])
        return ans
```

```py [sol-Python3 一行]
class Solution:
    def getLongestSubsequence(self, words: List[str], groups: List[int]) -> List[str]:
        return [w for (x, y), w in zip(pairwise(groups), words) if x != y] + [words[-1]]
```

```java [sol-Java]
class Solution {
    public List<String> getLongestSubsequence(String[] words, int[] groups) {
        List<String> ans = new ArrayList<>();
        int n = groups.length;
        for (int i = 0; i < n; i++) {
            if (i == n - 1 || groups[i] != groups[i + 1]) { // i 是连续相同段的末尾
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
    vector<string> getLongestSubsequence(vector<string>& words, vector<int>& groups) {
        vector<string> ans;
        int n = groups.size();
        for (int i = 0; i < n; i++) {
            if (i == n - 1 || groups[i] != groups[i + 1]) { // i 是连续相同段的末尾
                ans.push_back(words[i]);
            }
        }
        return ans;
    }
};
```

```c [sol-C]
char** getLongestSubsequence(char** words, int wordsSize, int* groups, int groupsSize, int* returnSize) {
    char** ans = malloc(sizeof(char*) * groupsSize);
    int idx = 0;
    for (int i = 0; i < groupsSize; i++) {
        if (i == groupsSize - 1 || groups[i] != groups[i + 1]) { // i 是连续相同段的末尾
            ans[idx++] = words[i];
        }
    }
    *returnSize = idx;
    return ans;
}
```

```go [sol-Go]
func getLongestSubsequence(words []string, groups []int) (ans []string) {
	n := len(groups)
	for i, g := range groups {
		if i == n-1 || g != groups[i+1] { // i 是连续相同段的末尾
			ans = append(ans, words[i])
		}
	}
	return
}
```

```js [sol-JavaScript]
var getLongestSubsequence = function(words, groups) {
    const n = groups.length;
    const ans = [];
    for (let i = 0; i < n; i++) {
        if (i === n - 1 || groups[i] !== groups[i + 1]) { // i 是连续相同段的末尾
            ans.push(words[i]);
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn get_longest_subsequence(words: Vec<String>, groups: Vec<i32>) -> Vec<String> {
        let n = groups.len();
        let mut ans = vec![];
        for (i, word) in words.into_iter().enumerate() {
            if i == n - 1 || groups[i] != groups[i + 1] { // i 是连续相同段的末尾
                ans.push(word);
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{words}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
