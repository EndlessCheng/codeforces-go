## 方法一：Z 函数 + 字典树

对于两个字符串 $s$ 和 $t$，设 $s$ 的长度为 $m$。如果 $s$ 既是 $t$ 的前缀又是 $t$ 的后缀，那么对于 $t$ 来说，它的长为 $m$ 的前后缀必须相等。

怎么快速判断 $t$ 的某个长度的前后缀相等？我们可以用 Z 函数解决。

由于 $z[i]$ 的定义是后缀 $t[i:]$ 与 $t$ 的最长公共前缀的长度，所以只要 $z[i] = n-1-i$，那么 $t[i:]$ 和与其等长的 $t$ 的前缀是相等的。

如果 $z[i] = n-1-i$ 成立，我们只需要判断 $s$ 是否为 $t$ 的前缀。

枚举 $t=\textit{words}[j]$，怎么统计有多少个 $s=\textit{words}[i]$ 是 $t$ 的前缀？

这可以用**字典树**解决，在遍历 $\textit{words}$ 的同时，维护每个字符串的出现次数。当我们遍历 $t$ 时，同时遍历字典树上的对应节点，并把 $t$ 插入字典树。具体请看代码。

```py [sol-Python3]
class Node:
    __slots__ = 'son', 'cnt'

    def __init__(self):
        self.son = dict()
        self.cnt = 0

class Solution:
    def countPrefixSuffixPairs(self, words: List[str]) -> int:
        ans = 0
        root = Node()
        for t in words:
            z = self.calc_z(t)
            cur = root
            for i, c in enumerate(t):
                if c not in cur.son:
                    cur.son[c] = Node()
                cur = cur.son[c]
                if z[-1 - i] == i + 1:  # t[-1-i:] == t[:i+1]
                    ans += cur.cnt
            cur.cnt += 1
        return ans

    def calc_z(self, s: str) -> List[int]:
        n = len(s)
        z = [0] * n
        l, r = 0, 0
        for i in range(1, n):
            if i <= r:
                z[i] = min(z[i - l], r - i + 1)
            while i + z[i] < n and s[z[i]] == s[i + z[i]]:
                l, r = i, i + z[i]
                z[i] += 1
        z[0] = n
        return z
```

```java [sol-Java]
class Node {
    Node[] son = new Node[26];
    int cnt;
}

class Solution {
    public long countPrefixSuffixPairs(String[] words) {
        long ans = 0;
        Node root = new Node();
        for (String T : words) {
            char[] t = T.toCharArray();
            int n = t.length;
            int[] z = new int[n];
            int l = 0, r = 0;
            for (int i = 1; i < n; i++) {
                if (i <= r) {
                    z[i] = Math.min(z[i - l], r - i + 1);
                }
                while (i + z[i] < n && t[z[i]] == t[i + z[i]]) {
                    l = i;
                    r = i + z[i];
                    z[i]++;
                }
            }
            z[0] = n;

            Node cur = root;
            for (int i = 0; i < n; i++) {
                int c = t[i] - 'a';
                if (cur.son[c] == null) {
                    cur.son[c] = new Node();
                }
                cur = cur.son[c];
                if (z[n - 1 - i] == i + 1) { // t 的长为 i+1 的前后缀相同
                    ans += cur.cnt;
                }
            }
            cur.cnt++;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
struct Node {
    Node *son[26]{};
    int cnt = 0;
};

class Solution {
public:
    long long countPrefixSuffixPairs(vector<string> &words) {
        long long ans = 0;
        Node *root = new Node();
        for (string &t: words) {
            int n = t.length();
            vector<int> z(n);
            int l = 0, r = 0;
            for (int i = 1; i < n; i++) {
                if (i <= r) {
                    z[i] = min(z[i - l], r - i + 1);
                }
                while (i + z[i] < n && t[z[i]] == t[i + z[i]]) {
                    l = i;
                    r = i + z[i];
                    z[i]++;
                }
            }
            z[0] = n;

            auto cur = root;
            for (int i = 0; i < n; i++) {
                int c = t[i] - 'a';
                if (cur->son[c] == nullptr) {
                    cur->son[c] = new Node();
                }
                cur = cur->son[c];
                if (z[n - 1 - i] == i + 1) { // t 的长为 i+1 的前后缀相同
                    ans += cur->cnt;
                }
            }
            cur->cnt++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func calcZ(s string) []int {
	n := len(s)
	z := make([]int, n)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
	}
	z[0] = n
	return z
}

func countPrefixSuffixPairs(words []string) (ans int64) {
	type node struct {
		son [26]*node
		cnt int
	}
	root := &node{}
	for _, t := range words {
		z := calcZ(t)
		cur := root
		for i, c := range t {
			c -= 'a'
			if cur.son[c] == nil {
				cur.son[c] = &node{}
			}
			cur = cur.son[c]
			if z[len(t)-1-i] == i+1 { // t[:i+1] == t[len(t)-1-i:]
				ans += int64(cur.cnt)
			}
		}
		cur.cnt++
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L)$，其中 $L$ 为所有 $\textit{words}[i]$ 的长度之和。
- 空间复杂度：$\mathcal{O}(L)$。

## 方法二：只用字典树

把每个字符串 $s$ 视作一个 pair 列表：

$$
[(s[0],s[n-1]), (s[1],s[n-2]), (s[2],s[n-3]), \cdots, (s[n-1], s[0])]
$$

只要这个 pair 列表是另一个字符串 $t$ 的 pair 列表的前缀，那么 $s$ 就是 $t$ 的前后缀。

具体例子请看视频：[字典树的巧妙运用](https://www.bilibili.com/video/BV1jZ42127Yf/)。

```py [sol-Python3]
class Node:
    __slots__ = 'son', 'cnt'

    def __init__(self):
        self.son = dict()
        self.cnt = 0

class Solution:
    def countPrefixSuffixPairs(self, words: List[str]) -> int:
        ans = 0
        root = Node()
        for s in words:
            cur = root
            for p in zip(s, reversed(s)):
                if p not in cur.son:
                    cur.son[p] = Node()
                cur = cur.son[p]
                ans += cur.cnt
            cur.cnt += 1
        return ans
```

```java [sol-Java]
class Node {
    Map<Integer, Node> son = new HashMap<>();
    int cnt;
}

class Solution {
    public long countPrefixSuffixPairs(String[] words) {
        long ans = 0;
        Node root = new Node();
        for (String S : words) {
            char[] s = S.toCharArray();
            int n = s.length;
            Node cur = root;
            for (int i = 0; i < n; i++) {
                int p = (s[i] - 'a') << 5 | (s[n - 1 - i] - 'a');
                cur = cur.son.computeIfAbsent(p, k -> new Node());
                ans += cur.cnt;
            }
            cur.cnt++;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
struct Node {
    unordered_map<int, Node*> son;
    int cnt = 0;
};

class Solution {
public:
    long long countPrefixSuffixPairs(vector<string> &words) {
        long long ans = 0;
        Node *root = new Node();
        for (string &s: words) {
            int n = s.length();
            auto cur = root;
            for (int i = 0; i < n; i++) {
                int p = (int) (s[i] - 'a') << 5 | (s[n - 1 - i] - 'a');
                if (cur->son[p] == nullptr) {
                    cur->son[p] = new Node();
                }
                cur = cur->son[p];
                ans += cur->cnt;
            }
            cur->cnt++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countPrefixSuffixPairs(words []string) (ans int64) {
	type pair struct{ x, y byte }
	type node struct {
		son map[pair]*node
		cnt int
	}
	root := &node{son: map[pair]*node{}}
	for _, s := range words {
		cur := root
		for i := range s {
			p := pair{s[i], s[len(s)-1-i]}
			if cur.son[p] == nil {
				cur.son[p] = &node{son: map[pair]*node{}}
			}
			cur = cur.son[p]
			ans += int64(cur.cnt)
		}
		cur.cnt++
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L)$，其中 $L$ 为所有 $\textit{words}[i]$ 的长度之和。
- 空间复杂度：$\mathcal{O}(L)$。

## 练习：Z 函数

- [3036. 匹配模式数组的子数组数目 II](https://leetcode.cn/problems/number-of-subarrays-that-match-a-pattern-ii/)
- [3008. 找出数组中的美丽下标 II](https://leetcode.cn/problems/find-beautiful-indices-in-the-given-array-ii/) 2016
- [2223. 构造字符串的总得分和](https://leetcode.cn/problems/sum-of-scores-of-built-strings/) 2220
- [3031. 将单词恢复初始状态所需的最短时间 II](https://leetcode.cn/problems/minimum-time-to-revert-word-to-initial-state-ii/) 2278

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
