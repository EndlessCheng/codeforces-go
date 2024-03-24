**前置知识**：[208. 实现 Trie (前缀树)](https://leetcode.cn/problems/implement-trie-prefix-tree/)。

从左到右遍历 $\textit{wordsContainer}$，设 $s=\textit{wordsContainer}[i]$。

倒着遍历 $s$，插入字典树。插入时，**对于每个经过的节点**，更新节点对应的最小字符串长度及其下标。

对于查询 $s=\textit{wordsQuery}[i]$，仍然倒着遍历 $s$。在字典树上**找到最后一个匹配的节点**，那么该节点保存的下标就是答案。

具体例子请看 [视频讲解](https://www.bilibili.com/video/BV1wr421h7xY/) 第四题。

```py [sol-Python3]
class Node:
    __slots__ = 'son', 'min_l', 'i'

    def __init__(self):
        self.son = [None] * 26
        self.min_l = inf

class Solution:
    def stringIndices(self, wordsContainer: List[str], wordsQuery: List[str]) -> List[int]:
        ord_a = ord('a')
        root = Node()
        for idx, s in enumerate(wordsContainer):
            l = len(s)
            cur = root
            if l < cur.min_l:
                cur.min_l, cur.i = l, idx
            for c in map(ord, reversed(s)):
                c -= ord_a
                if cur.son[c] is None:
                    cur.son[c] = Node()
                cur = cur.son[c]
                if l < cur.min_l:
                    cur.min_l, cur.i = l, idx

        ans = []
        for s in wordsQuery:
            cur = root
            for c in map(ord, reversed(s)):
                c -= ord_a
                if cur.son[c] is None:
                    break
                cur = cur.son[c]
            ans.append(cur.i)
        return ans
```

```java [sol-Java]
class Node {
    Node[] son = new Node[26];
    int minL = Integer.MAX_VALUE;
    int i;
}

class Solution {
    public int[] stringIndices(String[] wordsContainer, String[] wordsQuery) {
        Node root = new Node();
        for (int idx = 0; idx < wordsContainer.length; ++idx) {
            char[] s = wordsContainer[idx].toCharArray();
            int l = s.length;
            Node cur = root;
            if (l < cur.minL) {
                cur.minL = l;
                cur.i = idx;
            }
            for (int i = s.length - 1; i >= 0; i--) {
                int b = s[i] - 'a';
                if (cur.son[b] == null) {
                    cur.son[b] = new Node();
                }
                cur = cur.son[b];
                if (l < cur.minL) {
                    cur.minL = l;
                    cur.i = idx;
                }
            }
        }

        int[] ans = new int[wordsQuery.length];
        for (int idx = 0; idx < wordsQuery.length; idx++) {
            char[] s = wordsQuery[idx].toCharArray();
            Node cur = root;
            for (int i = s.length - 1; i >= 0 && cur.son[s[i] - 'a'] != null; i--) {
                cur = cur.son[s[i] - 'a'];
            }
            ans[idx] = cur.i;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
struct Node {
    Node *son[26]{};
    int min_l = INT_MAX, i;
};

class Solution {
public:
    vector<int> stringIndices(vector<string> &wordsContainer, vector<string> &wordsQuery) {
        Node *root = new Node();
        for (int idx = 0; idx < wordsContainer.size(); ++idx) {
            auto &s = wordsContainer[idx];
            int l = s.length();
            auto cur = root;
            if (l < cur->min_l) {
                cur->min_l = l;
                cur->i = idx;
            }
            for (int i = s.length() - 1; i >= 0; i--) {
                int b = s[i] - 'a';
                if (cur->son[b] == nullptr) {
                    cur->son[b] = new Node();
                }
                cur = cur->son[b];
                if (l < cur->min_l) {
                    cur->min_l = l;
                    cur->i = idx;
                }
            }
        }

        vector<int> ans;
        ans.reserve(wordsQuery.size());
        for (auto &s: wordsQuery) {
            auto cur = root;
            for (int i = s.length() - 1; i >= 0 && cur->son[s[i] - 'a']; i--) {
                cur = cur->son[s[i] - 'a'];
            }
            ans.push_back(cur->i);
        }
        return ans;
    }
};
```

```go [sol-Go]
func stringIndices(wordsContainer, wordsQuery []string) []int {
	type node struct {
		son     [26]*node
		minL, i int
	}
	root := &node{minL: math.MaxInt}

	for idx, s := range wordsContainer {
		l := len(s)
		cur := root
		if l < cur.minL {
			cur.minL, cur.i = l, idx
		}
		for i := len(s) - 1; i >= 0; i-- {
			b := s[i] - 'a'
			if cur.son[b] == nil {
				cur.son[b] = &node{minL: math.MaxInt}
			}
			cur = cur.son[b]
			if l < cur.minL {
				cur.minL, cur.i = l, idx
			}
		}
	}

	ans := make([]int, len(wordsQuery))
	for idx, s := range wordsQuery {
		cur := root
		for i := len(s) - 1; i >= 0 && cur.son[s[i]-'a'] != nil; i-- {
			cur = cur.son[s[i]-'a']
		}
		ans[idx] = cur.i
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L_1|\Sigma| + L_2)$，其中 $L_1$ 为 $\textit{wordsContainer}$ 中的所有字符串的长度**之和**，$L_2$ 为 $\textit{wordsQuery}$ 中的所有字符串的长度**之和**，$|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(L_1|\Sigma|)$。返回值不计入。

#### 相似题目

- [208. 实现 Trie (前缀树)](https://leetcode.cn/problems/implement-trie-prefix-tree/)
- [2416. 字符串的前缀分数和](https://leetcode.cn/problems/sum-of-prefix-scores-of-strings/) 1725
- [336. 回文对](https://leetcode.cn/problems/palindrome-pairs/)
- [745. 前缀和后缀搜索](https://leetcode.cn/problems/prefix-and-suffix-search/)
- [3045. 统计前后缀下标对 II](https://leetcode.cn/problems/count-prefix-and-suffix-pairs-ii/) 2328
- [527. 单词缩写](https://leetcode.cn/problems/word-abbreviation/)（会员题）
- [1804. 实现 Trie （前缀树） II](https://leetcode.cn/problems/implement-trie-ii-prefix-tree/)（会员题）

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。
