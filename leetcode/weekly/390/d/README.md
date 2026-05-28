**前置题目**：[208. 实现 Trie (前缀树)](https://leetcode.cn/problems/implement-trie-prefix-tree/)，[我的题解](https://leetcode.cn/problems/implement-trie-prefix-tree/solutions/2993894/cong-er-cha-shu-dao-er-shi-liu-cha-shu-p-xsj4/)。

把每个字符串都反转（或者**从右到左遍历**），本题就变成**前缀**问题了：

- 在 $\textit{wordsContainer}$ 中，找到与 $q = \textit{wordsQuery}[i]$ 有**最长**公共前缀的所有字符串。在这些字符串中，找到长度**最短**的。长度相同时，下标小的更优。

最长公共前缀用字典树解决。此外，**有着相同前缀的字符串，都在同一棵子树中**。为了 $\mathcal{O}(1)$ 找到子树中的最短字符串，我们可以在子树的根节点处，保存子树中的最短字符串长度 $\textit{minLen}$ 及其下标。在插入字符串 $s$ 的过程中，用 $s$ 的长度，更新路径上每个节点的 $\textit{minLen}$ 的最小值。

[本题视频讲解](https://www.bilibili.com/video/BV1wr421h7xY/?t=35m56s)，欢迎点赞关注~

```py [sol-Python3]
class Node:
    __slots__ = 'son', 'min_len', 'best_index'

    def __init__(self):
        self.son = [None] * 26
        self.min_len = inf  # 子树中的最短字符串的长度

class Solution:
    def stringIndices(self, wordsContainer: List[str], wordsQuery: List[str]) -> List[int]:
        ord_a = ord('a')
        root = Node()
        for i, s in enumerate(wordsContainer):
            len_s = len(s)
            if len_s < root.min_len:
                root.min_len = len_s
                root.best_index = i

            # 把 s[::-1] 插入字典树
            cur = root
            for ch in reversed(s):
                c = ord(ch) - ord_a
                if cur.son[c] is None:
                    cur.son[c] = Node()
                cur = cur.son[c]
                # 维护 cur 子树中的最短字符串的长度及其下标
                # 由于我们是按照 i 从小到大的顺序遍历，字符串长度相同时不更新 best_index
                if len_s < cur.min_len:
                    cur.min_len = len_s  
                    cur.best_index = i

        ans = []
        for s in wordsQuery:
            cur = root
            for ch in reversed(s):
                c = ord(ch) - ord_a
                if cur.son[c] is None:
                    break
                cur = cur.son[c]
            # 退出循环时，cur 即最长公共前缀（的对应节点），cur.best_index 是前缀为 cur 的最短字符串的下标
            ans.append(cur.best_index)
        return ans
```

```java [sol-Java]
class Node {
    Node[] son = new Node[26];
    int minLen = Integer.MAX_VALUE; // 子树中的最短字符串的长度
    int bestIndex; // 子树中的最短字符串的下标
}

class Solution {
    public int[] stringIndices(String[] wordsContainer, String[] wordsQuery) {
        Node root = new Node();
        for (int i = 0; i < wordsContainer.length; i++) {
            char[] s = wordsContainer[i].toCharArray();
            if (s.length < root.minLen) {
                root.minLen = s.length;
                root.bestIndex = i;
            }

            // 把 reverse(s) 插入字典树
            Node cur = root;
            for (int j = s.length - 1; j >= 0; j--) {
                int b = s[j] - 'a';
                if (cur.son[b] == null) {
                    cur.son[b] = new Node();
                }
                cur = cur.son[b];
                // 维护 cur 子树中的最短字符串的长度及其下标
                // 由于我们是按照 i 从小到大的顺序遍历，字符串长度相同时不更新 bestIndex
                if (s.length < cur.minLen) {
                    cur.minLen = s.length;
                    cur.bestIndex = i;
                }
            }
        }

        int[] ans = new int[wordsQuery.length];
        for (int i = 0; i < wordsQuery.length; i++) {
            String s = wordsQuery[i]; // 由于下面会中途退出循环，不转成 char[] 更快
            Node cur = root;
            for (int j = s.length() - 1; j >= 0 && cur.son[s.charAt(j) - 'a'] != null; j--) {
                cur = cur.son[s.charAt(j) - 'a'];
            }
            // 退出循环时，cur 即最长公共前缀（的对应节点），cur.bestIndex 是前缀为 cur 的最短字符串的下标
            ans[i] = cur.bestIndex;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    struct Node {
        Node* son[26]{};
        int min_len = INT_MAX; // 子树中的最短字符串的长度
        int best_index; // 子树中的最短字符串的下标

        // 本题需要写析构函数，否则会 MLE
        ~Node() {
            for (Node* s : son) {
                delete s;
            }
        }
    };

public:
    vector<int> stringIndices(vector<string>& wordsContainer, vector<string>& wordsQuery) {
        Node root{};
        for (int i = 0; i < wordsContainer.size(); i++) {
            auto& s = wordsContainer[i];
            int len = s.size();
            if (len < root.min_len) {
                root.min_len = len;
                root.best_index = i;
            }

            // 把 reverse(s) 插入字典树
            Node* cur = &root;
            for (int j = len - 1; j >= 0; j--) {
                int b = s[j] - 'a';
                if (cur->son[b] == nullptr) {
                    cur->son[b] = new Node();
                }
                cur = cur->son[b];
                // 维护 cur 子树中的最短字符串的长度及其下标
                // 由于我们是按照 i 从小到大的顺序遍历，字符串长度相同时不更新 best_index
                if (len < cur->min_len) {
                    cur->min_len = len;
                    cur->best_index = i;
                }
            }
        }

        vector<int> ans;
        ans.reserve(wordsQuery.size());
        for (auto& s : wordsQuery) {
            Node* cur = &root;
            for (int j = s.size() - 1; j >= 0 && cur->son[s[j] - 'a']; j--) {
                cur = cur->son[s[j] - 'a'];
            }
            // 退出循环时，cur 即最长公共前缀（的对应节点），cur->best_index 是前缀为 cur 的最短字符串的下标
            ans.push_back(cur->best_index);
        }
        return ans;
    }
};
```

```go [sol-Go]
func stringIndices(wordsContainer, wordsQuery []string) []int {
	type node struct {
		son       [26]*node
		minLen    int // 子树中的最短字符串的长度
		bestIndex int // 子树中的最短字符串的下标
	}
	root := &node{minLen: math.MaxInt}

	for i, s := range wordsContainer {
		l := len(s)
		if l < root.minLen {
			root.minLen = l
			root.bestIndex = i
		}

		// 把 reverse(s) 插入字典树
		cur := root
		for j := l - 1; j >= 0; j-- {
			b := s[j] - 'a'
			if cur.son[b] == nil {
				cur.son[b] = &node{minLen: math.MaxInt}
			}
			cur = cur.son[b]
			// 维护 cur 子树中的最短字符串的长度及其下标
			// 由于我们是按照 i 从小到大的顺序遍历，字符串长度相同时不更新 bestIndex
			if l < cur.minLen {
				cur.minLen = l
				cur.bestIndex = i
			}
		}
	}

	ans := make([]int, len(wordsQuery))
	for i, s := range wordsQuery {
		cur := root
		for j := len(s) - 1; j >= 0 && cur.son[s[j]-'a'] != nil; j-- {
			cur = cur.son[s[j]-'a']
		}
		// 退出循环时，cur 即最长公共前缀（的对应节点），cur.bestIndex 是前缀为 cur 的最短字符串的下标
		ans[i] = cur.bestIndex
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L_1|\Sigma| + L_2)$，其中 $L_1$ 是 $\textit{wordsContainer}$ 中的所有字符串的长度**之和**，$L_2$ 是 $\textit{wordsQuery}$ 中的所有字符串的长度**之和**，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(L_1|\Sigma|)$。返回值不计入。

## 专题训练

见下面数据结构题单的「**六、字典树**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
