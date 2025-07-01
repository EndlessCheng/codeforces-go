## 方法一：哈希集合

按题意模拟即可。

为了快速判断当前字符串 $t$ 是否在答案中，用一个哈希集合保存在答案中的字符串。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1j6gZzqEdc/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def partitionString(self, s: str) -> List[str]:
        ans = []
        vis = set()
        t = ''
        for c in s:
            t += c
            if t not in vis:
                vis.add(t)
                ans.append(t)
                t = ''
        return ans
```

```java [sol-Java]
class Solution {
    public List<String> partitionString(String s) {
        List<String> ans = new ArrayList<>();
        Set<String> vis = new HashSet<>();
        String t = "";
        for (char c : s.toCharArray()) {
            t += c;
            if (vis.add(t)) { // t 不在 vis 中
                ans.add(t);
                t = "";
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<string> partitionString(string s) {
        vector<string> ans;
        unordered_set<string> vis;
        string t;
        for (char c : s) {
            t += c;
            if (vis.insert(t).second) { // t 不在 vis 中
                ans.push_back(t);
                t.clear();
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func partitionString(s string) (ans []string) {
	vis := map[string]bool{}
	left := 0
	for i := range s {
		t := s[left : i+1]
		if !vis[t] {
			vis[t] = true
			ans = append(ans, t)
			left = i + 1
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt n)$，其中 $n$ 是 $s$ 的长度。最坏情况下 $n$ 个一样的字母，加到答案中的字符串长度为 $1,2,3,\ldots,k$，解不等式 $1+2+3+\cdots+k = \dfrac{k(k+1)}{2}\le n$，得 $k = \mathcal{O}(\sqrt n)$。每次判断一个长为 $\mathcal{O}(\sqrt n)$ 的字符串是否在哈希集合中，需要 $\mathcal{O}(\sqrt n)$ 的时间，一共判断 $n$ 次。所以时间复杂度为 $\mathcal{O}(n\sqrt n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：字典树

原理见 [208. 实现 Trie (前缀树)](https://leetcode.cn/problems/implement-trie-prefix-tree/)，[我的题解](https://leetcode.cn/problems/implement-trie-prefix-tree/solutions/2993894/cong-er-cha-shu-dao-er-shi-liu-cha-shu-p-xsj4/)。

```py [sol-Python3]
class Solution:
    def partitionString(self, s: str) -> List[str]:
        ans = []
        cur = root = {}
        left = 0
        for i, c in enumerate(s):
            if c not in cur:  # 无路可走？
                cur[c] = {}  # 那就造路！
                ans.append(s[left: i + 1])
                left = i + 1
                cur = root  # 重置
            else:
                cur = cur[c]
        return ans
```

```java [sol-Java]
class Solution {
    public List<String> partitionString(String s) {
        record Node(Node[] son) {
            Node() { this(new Node[26]); }
        }

        List<String> ans = new ArrayList<>();
        Node root = new Node();
        Node cur = root;
        int left = 0;
        for (int i = 0; i < s.length(); i++) {
            int c = s.charAt(i) - 'a';
            if (cur.son[c] == null) { // 无路可走？
                cur.son[c] = new Node(); // 那就造路！
                ans.add(s.substring(left, i + 1));
                left = i + 1;
                cur = root; // 重置
            } else {
                cur = cur.son[c];
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
struct Node {
    Node* son[26]{};
};

class Solution {
public:
    vector<string> partitionString(string s) {
        vector<string> ans;
        Node* root = new Node();
        Node* cur = root;
        int left = 0;
        for (int i = 0; i < s.size(); i++) {
            int c = s[i] - 'a';
            if (cur->son[c] == nullptr) { // 无路可走？
                cur->son[c] = new Node(); // 那就造路！
                ans.push_back(s.substr(left, i - left + 1));
                left = i + 1;
                cur = root; // 重置
            } else {
                cur = cur->son[c];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func partitionString(s string) (ans []string) {
	type node struct{ son [26]*node }
	root := &node{}
	cur := root
	left := 0
	for i, c := range s {
		c -= 'a'
		if cur.son[c] == nil { // 无路可走？
			cur.son[c] = &node{} // 那就造路！
			ans = append(ans, s[left:i+1])
			left = i + 1
			cur = root // 重置
		} else {
			cur = cur.son[c]
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。用哈希表实现可以做到 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n|\Sigma|)$。用哈希表实现可以做到 $\mathcal{O}(n)$。

## 专题训练

见下面数据结构题单的「**六、字典树（trie）**」。

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
