## 题意解读

构建一棵 $\textit{words}$ 的字典树（Trie）。按照某种顺序 DFS 这棵树，使得访问到的 $\textit{words}[i]$ 的下标顺序尽量与 $\textit{target}$ 保持一致。

什么是「尽量保持一致」？设 $\textit{pos}[i]$ 是下标 $i$ 在 $\textit{target}$ 中的位置。按照 DFS 顺序收集访问到的 $\textit{pos}[i]$，得到一个数组 $a$。最小化 $a$ 的逆序对个数。

## 前置知识

1. **字典树（Trie）**。见 [208. 实现 Trie (前缀树)](https://leetcode.cn/problems/implement-trie-prefix-tree/)，[讲解（含模板代码）](https://leetcode.cn/problems/implement-trie-prefix-tree/solutions/2993894/cong-er-cha-shu-dao-er-shi-liu-cha-shu-p-xsj4/)。
2. **相邻无关排列型状压 DP**。见 [526. 优美的排列](https://leetcode.cn/problems/beautiful-arrangement/)，[教你一步步思考状压 DP：从记忆化搜索到递推](https://leetcode.cn/problems/beautiful-arrangement/solution/jiao-ni-yi-bu-bu-si-kao-zhuang-ya-dpcong-c6kd/)。

## 思路

逆序对来自两种情况：

1. 根与子树。在任意 DFS 序中，都必须先访问子树的根节点，再访问子树中的其他节点。如果子树的根是某个 $\textit{words}[i]$ 的最后一个字母，那么 $\textit{pos}[i]$ 一定在该子树其余节点的 $\textit{pos}$ 值的前面。计算子树中的小于 $\textit{pos}[i]$ 的 $\textit{pos}$ 值个数。这部分是固定的，与 DFS 序无关。
2. 子树与子树。对于一个节点的儿子子树，这些子树的访问顺序可以由我们决定，哪一种顺序（子树的排列）是最优的呢？这是本题的核心问题，下面细讲。

设节点 $\textit{node}$ 有 $m$ 个儿子，把第 $i$ 棵（$i$ 从 $0$ 开始）儿子子树的所有 $\textit{pos}$ 值记在数组 $a[i]$ 中。

于是问题变成：

- 给你 $m$ 个数组 $a[0],a[1],\ldots,a[m-1]$。请你重排这 $m$ 个数组，然后把这 $m$ 个数组合并成一个大数组 $A$，最小化 $A$ 的逆序对个数。

这是一个标准的「**相邻无关排列型状压 DP**」问题。关于如何得到状态定义和状态转移方程，套路和 526 题一样（见上面的链接），本文不再赘述。

定义 $\textit{inv}[i][j]$ 表示数组 $a[i]$ 在数组 $a[j]$ 的左边时，这两个数组之间的逆序对个数。这可以把数组都排序后，用**双指针**计算。

设全集 $U=\{0,1,\ldots,m-1\}$，设 $S$ 是 $U$ 的一个子集。定义状态 $f[S]$ 表示子集 $S$ 的最小逆序对个数。

枚举 $j$，在 $S$ 的右边添加数组 $a[j]$，那么新产生的逆序对个数为

$$
\sum_{i\in S} \textit{inv}[i][j]
$$

用 $f[S] + \sum\limits_{i\in S} \textit{inv}[i][j]$ 更新 $f[S\cup \{j\}]$ 的最小值（刷表法）。

初始值：$f[\varnothing] = 0$。

$A$ 的逆序对个数的最小值为 $f[U]$，加到答案中。

[本题视频讲解](https://www.bilibili.com/video/BV1rHpj6kE8u/?t=8m9s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Node:
    def __init__(self):
        self.son = {}
        self.pos = -1

# 计算 a 在 b 左侧时，a 和 b 之间的逆序对个数
# a 和 b 已经升序排序
def calc_inversions(a: list[int], b: list[int]) -> int:
    n, m = len(a), len(b)
    i = j = inv = 0
    while i < n and j < m:
        if a[i] > b[j]:
            # a 的子数组 [i, n-1] 中的数都大于 b[j]
            inv += n - i
            j += 1
        else:
            i += 1
    return inv

class Solution:
    def minInversions(self, words: list[str], target: list[int]) -> int:
        n = len(target)
        pos = [0] * n
        for i, x in enumerate(target):
            pos[x] = i

        # 创建前缀树（字典树）
        root = Node()
        for i, word in enumerate(words):
            cur = root
            for ch in word:
                if ch not in cur.son:
                    cur.son[ch] = Node()
                cur = cur.son[ch]
            cur.pos = pos[i]

        # 返回 node 子树的 pos 列表（已从小到大排序）
        def dfs(node: Node) -> list[int]:
            nonlocal ans
            a = [dfs(son) for son in node.son.values()]

            # 预处理逆序对个数
            inv = [[calc_inversions(p, q) for q in a] for p in a]

            m = len(a)
            f = [inf] * (1 << m)
            f[0] = 0
            for s, fs in enumerate(f):  # 之前选的下标集合是 s
                # 在 s 的右边添加 j
                for j in range(m):
                    if s >> j & 1:
                        continue
                    sum_inv = fs
                    # 枚举 s 中的下标 i，计算新产生的逆序对个数
                    for i in range(m):
                        if s >> i & 1:
                            sum_inv += inv[i][j]
                    f[s | 1 << j] = min(f[s | 1 << j], sum_inv)
            ans += f[-1]

            res = []
            cur_pos = node.pos
            if cur_pos >= 0:
                res.append(cur_pos)
            for sub in a:
                if cur_pos >= 0:
                    # cur_pos 与子树之间的逆序对个数
                    ans += bisect_left(sub, cur_pos)
                res += sub

            res.sort()
            return res

        ans = 0
        dfs(root)
        return ans
```

```java [sol-Java]
class Solution {
    private static class Node {
        Node[] son = new Node[10];
        int pos = -1;
    }

    private int ans = 0;

    public int minInversions(String[] words, int[] target) {
        int n = target.length;
        int[] pos = new int[n];
        for (int i = 0; i < n; i++) {
            pos[target[i]] = i;
        }

        // 创建前缀树（字典树）
        Node root = new Node();
        for (int i = 0; i < words.length; i++) {
            Node cur = root;
            for (char ch : words[i].toCharArray()) {
                int c = ch - 'a';
                if (cur.son[c] == null) {
                    cur.son[c] = new Node();
                }
                cur = cur.son[c];
            }
            cur.pos = pos[i];
        }

        dfs(root);
        return ans;
    }

    // 返回 node 子树的 pos 列表（已从小到大排序）
    private List<Integer> dfs(Node node) {
        List<List<Integer>> a = new ArrayList<>();
        for (Node son : node.son) {
            if (son != null) {
                a.add(dfs(son));
            }
        }

        // 预处理逆序对个数
        int m = a.size();
        int[][] inv = new int[m][m];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < m; j++) {
                if (j != i) {
                    inv[i][j] = calcInversions(a.get(i), a.get(j));
                }
            }
        }

        int u = 1 << m;
        int[] f = new int[u];
        Arrays.fill(f, Integer.MAX_VALUE);
        f[0] = 0;

        for (int s = 0; s < u; s++) { // 之前选的下标集合是 s
            // 在 s 的右边添加 j
            for (int j = 0; j < m; j++) {
                if ((s >> j & 1) > 0) {
                    continue;
                }
                int sum = f[s];
                // 枚举 s 中的下标 i，计算新产生的逆序对个数
                for (int i = 0; i < m; i++) {
                    if ((s >> i & 1) > 0) {
                        sum += inv[i][j];
                    }
                }
                f[s | 1 << j] = Math.min(f[s | 1 << j], sum);
            }
        }
        ans += f[u - 1];

        List<Integer> res = new ArrayList<>();
        int curPos = node.pos;
        if (curPos >= 0) {
            res.add(curPos);
        }
        for (List<Integer> sub : a) {
            if (curPos >= 0) {
                // curPos 与子树之间的逆序对个数
                ans += ~Collections.binarySearch(sub, curPos);
            }
            res.addAll(sub);
        }

        Collections.sort(res);
        return res;
    }

    // 计算 a 在 b 左侧时，a 和 b 之间的逆序对个数
    // a 和 b 已经升序排序
    private int calcInversions(List<Integer> a, List<Integer> b) {
        int n = a.size(), m = b.size();
        int i = 0, j = 0;
        int inv = 0;
        while (i < n && j < m) {
            if (a.get(i) > b.get(j)) {
                // a 的子数组 [i, n-1] 中的数都大于 b[j]
                inv += n - i;
                j++;
            } else {
                i++;
            }
        }
        return inv;
    }
}
```

```cpp [sol-C++]
struct Node {
    Node* son[10]{};
    int pos = -1;
};

class Solution {
    // 计算 a 在 b 左侧时，a 和 b 之间的逆序对个数
    // a 和 b 已经升序排序
    int calc_inversions(vector<int>& a, vector<int>& b) {
        int n = a.size(), m = b.size();
        int i = 0, j = 0;
        int inv = 0;
        while (i < n && j < m) {
            if (a[i] > b[j]) {
                // a 的子数组 [i, n-1] 中的数都大于 b[j]
                inv += n - i;
                j++;
            } else {
                i++;
            }
        }
        return inv;
    }

public:
    int minInversions(vector<string>& words, vector<int>& target) {
        int n = target.size();
        vector<int> pos(n);
        for (int i = 0; i < n; i++) {
            pos[target[i]] = i;
        }

        // 创建前缀树（字典树）
        Node* root = new Node();
        for (int i = 0; i < words.size(); i++) {
            Node* cur = root;
            for (char ch : words[i]) {
                ch -= 'a';
                if (cur->son[ch] == nullptr) {
                    cur->son[ch] = new Node();
                }
                cur = cur->son[ch];
            }
            cur->pos = pos[i];
        }

        int ans = 0;

        // 返回 node 子树的 pos 列表（已从小到大排序）
        auto dfs = [&](this auto&& dfs, Node* node) -> vector<int> {
            vector<vector<int>> a;
            for (Node* son : node->son) {
                if (son != nullptr) {
                    a.emplace_back(move(dfs(son)));
                }
            }

            // 预处理逆序对个数
            int m = a.size();
            vector inv(m, vector<int>(m));
            for (int i = 0; i < m; i++) {
                for (int j = 0; j < m; j++) {
                    if (j != i) {
                        inv[i][j] = calc_inversions(a[i], a[j]);
                    }
                }
            }

            int u = 1 << m;
            vector<int> f(u, INT_MAX);
            f[0] = 0;
            for (int s = 0; s < u; s++) { // 之前选的下标集合是 s
                // 在 s 的右边添加 j
                for (int j = 0; j < m; j++) {
                    if (s >> j & 1) {
                        continue;
                    }
                    int sum = f[s];
                    // 枚举 s 中的下标 i，计算新产生的逆序对个数
                    for (int i = 0; i < m; i++) {
                        if (s >> i & 1) {
                            sum += inv[i][j];
                        }
                    }
                    f[s | 1 << j] = min(f[s | 1 << j], sum);
                }
            }
            ans += f[u - 1];

            vector<int> res;
            int cur_pos = node->pos;
            if (cur_pos >= 0) {
                res.push_back(cur_pos);
            }
            for (auto& sub : a) {
                if (cur_pos >= 0) {
                    // cur_pos 与子树之间的逆序对个数
                    ans += ranges::lower_bound(sub, cur_pos) - sub.begin();
                }
                res.insert(res.end(), sub.begin(), sub.end());
            }

            ranges::sort(res);
            return res;
        };

        dfs(root);
        return ans;
    }
};
```

```go [sol-Go]
// 计算 a 在 b 左侧时，a 和 b 之间的逆序对个数
// a 和 b 已经升序排序
func calcInversions(a, b []int) (inv int) {
	n, m := len(a), len(b)
	i, j := 0, 0
	for i < n && j < m {
		if a[i] > b[j] {
			// a 的子数组 [i, n-1] 中的数都大于 b[j] 
			inv += n - i
			j++
		} else {
			i++
		}
	}
	return
}

func minInversions(words []string, target []int) (ans int) {
	n := len(target)
	pos := make([]int, n)
	for i, x := range target {
		pos[x] = i
	}

	// 创建前缀树（字典树）
	type node struct {
		son [10]*node
		pos int
	}
	root := &node{}
	for i, word := range words {
		cur := root
		for _, ch := range word {
			ch -= 'a'
			if cur.son[ch] == nil {
				cur.son[ch] = &node{}
			}
			cur = cur.son[ch]
		}
		// 对于不是 word 的节点，pos 是 0
		cur.pos = pos[i] + 1 // +1，与 pos = 0 的节点区分开
	}

	// 返回 node 子树的 pos 列表（已从小到大排序）
	var dfs func(*node) []int
	dfs = func(node *node) []int {
		a := [][]int{}
		for _, son := range node.son {
			if son != nil {
				a = append(a, dfs(son))
			}
		}

		// 预处理逆序对个数
		m := len(a)
		inv := make([][]int, m)
		for i, p := range a {
			inv[i] = make([]int, m)
			for j, q := range a {
				if j != i {
					inv[i][j] = calcInversions(p, q)
				}
			}
		}

		u := 1 << m
		f := make([]int, u)
		for i := 1; i < u; i++ {
			f[i] = math.MaxInt
		}
		for s, fs := range f { // 之前选的下标集合是 s
			// 在 s 的右边添加 j
			for j := range m {
				if s>>j&1 > 0 {
					continue
				}
				sum := fs
				// 枚举 s 中的下标 i，计算新产生的逆序对个数
				for i := range m {
					if s>>i&1 > 0 {
						sum += inv[i][j]
					}
				}
				f[s|1<<j] = min(f[s|1<<j], sum)
			}
		}
		ans += f[u-1]

		res := []int{}
		curPos := node.pos
		if curPos > 0 {
			res = append(res, curPos)
		}
		for _, sub := range a {
			if curPos > 0 {
				// curPos 与子树之间的逆序对个数
				ans += sort.SearchInts(sub, curPos)
			}
			res = append(res, sub...)
		}

		slices.Sort(res)
		return res
	}
	dfs(root)

	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nL|\Sigma|^2 2^{|\Sigma|} + nL\log n)$，其中 $n$ 是 $\textit{words}$ 的长度，$L\le 20$ 是单个字符串的最大长度，$|\Sigma|=10$ 是字符集合的大小。字典树有 $\mathcal{O}(nL)$ 个节点，每个节点计算一次 $\mathcal{O}(|\Sigma|^2 2^{|\Sigma|})$ 的状压 DP。计算逆序对的总循环次数为 $\mathcal{O}(nL|\Sigma|)$，理由见下文。排序的时间复杂度之和为 $\mathcal{O}(nL\log n)$，理由见下文。
- 空间复杂度：$\mathcal{O}(nL|\Sigma| + 2^{|\Sigma|})$。字典树需要 $\mathcal{O}(nL|\Sigma|)$ 的空间，状压 DP 需要 $\mathcal{O}(|\Sigma|^2 + 2^{|\Sigma|}) = \mathcal{O}(2^{|\Sigma|})$ 的空间。

对于每个 $\textit{words}[i]$，从字典树的根到 $\textit{words}[i]$ 的最后一个字母对应的节点，有 $k = |\textit{words}[i]|+1$ 个节点。$\textit{words}[i]$ 会给这 $k$ 个节点都贡献一个 $\textit{pos}$ 值。因此，所有节点的 $\textit{pos}$ 值的个数之和不超过 $n(L+1) = \mathcal{O}(nL)$。

对于单个节点，双指针计算逆序对的总循环次数为

$$
\sum_i \sum_j (|a_i|+|a_j|)
$$

每个 $|a_k|$ 在上式中至多出现 $|\Sigma|=10$ 次。

所以

$$
\sum_i \sum_j (|a_i|+|a_j|) \le |\Sigma|\sum_k |a_k|
$$

对所有节点求和，得到计算逆序对的总循环次数为

$$
\mathcal{O}(nL|\Sigma|)
$$

同理，排序的时间复杂度之和为 

$$
\mathcal{O}(nL\log n)
$$

## 优化

对于每个 $j$，预处理 $\textit{sumInv}_j[S]$ 表示在集合 $S$ 的右边添加 $a[j]$ 后，新增的逆序对个数。

这也是一个 DP。在 $S$ 中添加 $a[i]$ 后，只需增加 $\textit{inv}[i][j]$，即可得到在集合 $S\cup \{i\}$ 的右边添加 $a[j]$ 后，新增的逆序对个数，即

$$
\textit{sumInv}_j[S\cup \{i\}] = \textit{sumInv}_j[S] + \textit{inv}[i][j]
$$

初始值：$\textit{sumInv}_j[\varnothing] = 0$。

这样可以把状压 DP 的时间复杂度从 $\mathcal{O}(|\Sigma|^2 2^{|\Sigma|})$ 优化至 $\mathcal{O}(|\Sigma| 2^{|\Sigma|})$。

```py [sol-Python3]
class Node:
    def __init__(self):
        self.son = {}
        self.pos = -1

# 计算 a 在 b 左侧时，a 和 b 之间的逆序对个数
# a 和 b 已经升序排序
def calc_inversions(a: list[int], b: list[int]) -> int:
    n, m = len(a), len(b)
    i = j = inv = 0
    while i < n and j < m:
        if a[i] > b[j]:
            # a 的子数组 [i, n-1] 中的数都大于 b[j]
            inv += n - i
            j += 1
        else:
            i += 1
    return inv

class Solution:
    def minInversions(self, words: list[str], target: list[int]) -> int:
        n = len(target)
        pos = [0] * n
        for i, x in enumerate(target):
            pos[x] = i

        # 创建前缀树（字典树）
        root = Node()
        for i, word in enumerate(words):
            cur = root
            for ch in word:
                if ch not in cur.son:
                    cur.son[ch] = Node()
                cur = cur.son[ch]
            cur.pos = pos[i]

        # 返回 node 子树的 pos 列表（已从小到大排序）
        def dfs(node: Node) -> list[int]:
            nonlocal ans
            a = [dfs(son) for son in node.son.values()]

            # 预处理逆序对个数
            inv = [[calc_inversions(p, q) for q in a] for p in a]

            m = len(a)
            u = 1 << m
            sum_inv = [None] * m
            for j in range(m):
                s_inv = [0] * u
                for i in range(m):
                    high_bit = 1 << i
                    for mask in range(high_bit):
                        s_inv[high_bit | mask] = s_inv[mask] + inv[i][j]
                sum_inv[j] = s_inv

            f = [inf] * u
            f[0] = 0
            for s, fs in enumerate(f):  # 之前选的下标集合是 s
                # 在 s 的右边添加 j
                for j in range(m):
                    if s >> j & 1 == 0:
                        t = s | 1 << j
                        f[t] = min(f[t], fs + sum_inv[j][s])
            ans += f[-1]

            res = []
            cur_pos = node.pos
            if cur_pos >= 0:
                res.append(cur_pos)
            for sub in a:
                if cur_pos >= 0:
                    # cur_pos 与子树之间的逆序对个数
                    ans += bisect_left(sub, cur_pos)
                res += sub

            res.sort()
            return res

        ans = 0
        dfs(root)
        return ans
```

```java [sol-Java]
class Solution {
    private static class Node {
        Node[] son = new Node[10];
        int pos = -1;
    }

    private int ans = 0;

    public int minInversions(String[] words, int[] target) {
        int n = target.length;
        int[] pos = new int[n];
        for (int i = 0; i < n; i++) {
            pos[target[i]] = i;
        }

        // 创建前缀树（字典树）
        Node root = new Node();
        for (int i = 0; i < words.length; i++) {
            Node cur = root;
            for (char ch : words[i].toCharArray()) {
                int c = ch - 'a';
                if (cur.son[c] == null) {
                    cur.son[c] = new Node();
                }
                cur = cur.son[c];
            }
            cur.pos = pos[i];
        }

        dfs(root);
        return ans;
    }

    // 返回 node 子树的 pos 列表（已从小到大排序）
    private List<Integer> dfs(Node node) {
        List<List<Integer>> a = new ArrayList<>();
        for (Node son : node.son) {
            if (son != null) {
                a.add(dfs(son));
            }
        }

        // 预处理逆序对个数
        int m = a.size();
        int[][] inv = new int[m][m];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < m; j++) {
                if (j != i) {
                    inv[i][j] = calcInversions(a.get(i), a.get(j));
                }
            }
        }

        int u = 1 << m;
        int[][] sumInv = new int[m][u];
        for (int j = 0; j < m; j++) {
            int[] sum = sumInv[j];
            for (int i = 0; i < m; i++) {
                int highBit = 1 << i;
                for (int mask = 0; mask < highBit; mask++) {
                    int s = sum[mask];
                    sum[highBit | mask] = s + inv[i][j];
                }
            }
        }

        int[] f = new int[u];
        Arrays.fill(f, Integer.MAX_VALUE);
        f[0] = 0;

        for (int s = 0; s < u; s++) { // 之前选的下标集合是 s
            // 在 s 的右边添加 j
            for (int j = 0; j < m; j++) {
                if ((s >> j & 1) == 0) {
                    f[s | 1 << j] = Math.min(f[s | 1 << j], f[s] + sumInv[j][s]);
                }
            }
        }
        ans += f[u - 1];

        List<Integer> res = new ArrayList<>();
        int curPos = node.pos;
        if (curPos >= 0) {
            res.add(curPos);
        }
        for (List<Integer> sub : a) {
            if (curPos >= 0) {
                // curPos 与子树之间的逆序对个数
                ans += ~Collections.binarySearch(sub, curPos);
            }
            res.addAll(sub);
        }

        Collections.sort(res);
        return res;
    }

    // 计算 a 在 b 左侧时，a 和 b 之间的逆序对个数
    // a 和 b 已经升序排序
    private int calcInversions(List<Integer> a, List<Integer> b) {
        int n = a.size(), m = b.size();
        int i = 0, j = 0;
        int inv = 0;
        while (i < n && j < m) {
            if (a.get(i) > b.get(j)) {
                // a 的子数组 [i, n-1] 中的数都大于 b[j]
                inv += n - i;
                j++;
            } else {
                i++;
            }
        }
        return inv;
    }
}
```

```cpp [sol-C++]
struct Node {
    Node* son[10]{};
    int pos = -1;
};

class Solution {
    // 计算 a 在 b 左侧时，a 和 b 之间的逆序对个数
    // a 和 b 已经升序排序
    int calc_inversions(vector<int>& a, vector<int>& b) {
        int n = a.size(), m = b.size();
        int i = 0, j = 0;
        int inv = 0;
        while (i < n && j < m) {
            if (a[i] > b[j]) {
                // a 的子数组 [i, n-1] 中的数都大于 b[j]
                inv += n - i;
                j++;
            } else {
                i++;
            }
        }
        return inv;
    }

public:
    int minInversions(vector<string>& words, vector<int>& target) {
        int n = target.size();
        vector<int> pos(n);
        for (int i = 0; i < n; i++) {
            pos[target[i]] = i;
        }

        // 创建前缀树（字典树）
        Node* root = new Node();
        for (int i = 0; i < words.size(); i++) {
            Node* cur = root;
            for (char ch : words[i]) {
                ch -= 'a';
                if (cur->son[ch] == nullptr) {
                    cur->son[ch] = new Node();
                }
                cur = cur->son[ch];
            }
            cur->pos = pos[i];
        }

        int ans = 0;
        int sum_inv[10][1 << 10]; // 避免在 dfs 中反复申请内存
        int f[1 << 10];

        // 返回 node 子树的 pos 列表（已从小到大排序）
        auto dfs = [&](this auto&& dfs, Node* node) -> vector<int> {
            vector<vector<int>> a;
            for (Node* son : node->son) {
                if (son != nullptr) {
                    a.emplace_back(move(dfs(son)));
                }
            }

            // 预处理逆序对个数
            int m = a.size();
            vector inv(m, vector<int>(m));
            for (int i = 0; i < m; i++) {
                for (int j = 0; j < m; j++) {
                    if (j != i) {
                        inv[i][j] = calc_inversions(a[i], a[j]);
                    }
                }
            }

            int u = 1 << m;
            for (int j = 0; j < m; j++) {
                auto& sum = sum_inv[j];
                sum[0] = 0;
                for (int i = 0; i < m; i++) {
                    int high_bit = 1 << i;
                    for (int mask = 0; mask < high_bit; mask++) {
                        int s = sum[mask];
                        sum[high_bit | mask] = s + inv[i][j];
                    }
                }
            }

            fill(f, f + u, INT_MAX);
            f[0] = 0;
            for (int s = 0; s < u; s++) { // 之前选的下标集合是 s
                // 在 s 的右边添加 j
                for (int j = 0; j < m; j++) {
                    if ((s >> j & 1) == 0) {
                        f[s | 1 << j] = min(f[s | 1 << j], f[s] + sum_inv[j][s]);
                    }
                }
            }
            ans += f[u - 1];

            vector<int> res;
            int cur_pos = node->pos;
            if (cur_pos >= 0) {
                res.push_back(cur_pos);
            }
            for (auto& sub : a) {
                if (cur_pos >= 0) {
                    // cur_pos 与子树之间的逆序对个数
                    ans += ranges::lower_bound(sub, cur_pos) - sub.begin();
                }
                res.insert(res.end(), sub.begin(), sub.end());
            }

            ranges::sort(res);
            return res;
        };

        dfs(root);
        return ans;
    }
};
```

```go [sol-Go]
// 计算 a 在 b 左侧时，a 和 b 之间的逆序对个数
// a 和 b 已经升序排序
func calcInversions(a, b []int) (inv int) {
	n, m := len(a), len(b)
	i, j := 0, 0
	for i < n && j < m {
		if a[i] > b[j] {
			// a 的子数组 [i, n-1] 中的数都大于 b[j] 
			inv += n - i
			j++
		} else {
			i++
		}
	}
	return
}

func minInversions(words []string, target []int) (ans int) {
	n := len(target)
	pos := make([]int, n)
	for i, x := range target {
		pos[x] = i
	}

	// 创建前缀树（字典树）
	type node struct {
		son [10]*node
		pos int
	}
	root := &node{}
	for i, word := range words {
		cur := root
		for _, ch := range word {
			ch -= 'a'
			if cur.son[ch] == nil {
				cur.son[ch] = &node{}
			}
			cur = cur.son[ch]
		}
		// 对于不是 word 的节点，pos 是 0
		cur.pos = pos[i] + 1 // +1，与 pos = 0 的节点区分开
	}

	// 返回 node 子树的 pos 列表（已从小到大排序）
	var dfs func(*node) []int
	dfs = func(node *node) []int {
		a := [][]int{}
		for _, son := range node.son {
			if son != nil {
				a = append(a, dfs(son))
			}
		}

		// 预处理逆序对个数
		m := len(a)
		inv := make([][]int, m)
		for i, p := range a {
			inv[i] = make([]int, m)
			for j, q := range a {
				if j != i {
					inv[i][j] = calcInversions(p, q)
				}
			}
		}

		u := 1 << m
		sumInv := make([][]int, m)
		for j := range sumInv {
			sum := make([]int, u)
			for i := range m {
				highBit := 1 << i
				for mask, s := range sum[:highBit] {
					sum[highBit|mask] = s + inv[i][j]
				}
			}
			sumInv[j] = sum
		}

		f := make([]int, u)
		for i := 1; i < u; i++ {
			f[i] = math.MaxInt
		}
		for s, fs := range f { // 之前选的下标集合是 s
			// 在 s 的右边添加 j
			for j := range m {
				if s>>j&1 == 0 {
					f[s|1<<j] = min(f[s|1<<j], fs+sumInv[j][s])
				}
			}
		}
		ans += f[u-1]

		res := []int{}
		curPos := node.pos
		if curPos > 0 {
			res = append(res, curPos)
		}
		for _, sub := range a {
			if curPos > 0 {
				// curPos 与子树之间的逆序对个数
				ans += sort.SearchInts(sub, curPos)
			}
			res = append(res, sub...)
		}

		slices.Sort(res)
		return res
	}
	dfs(root)

	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nL|\Sigma| 2^{|\Sigma|} + nL\log n)$，其中 $n$ 是 $\textit{words}$ 的长度，$L\le 20$ 是单个字符串的最大长度，$|\Sigma|=10$ 是字符集合的大小。详细分析过程见优化前的复杂度分析。
- 空间复杂度：$\mathcal{O}(nL|\Sigma| + |\Sigma|2^{|\Sigma|})$。详细分析过程见优化前的复杂度分析。

## 专题训练

1. 数据结构题单的「**六、字典树（trie）**」。
2. 动态规划题单的「**§9.1 排列型状压 DP ① 相邻无关**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
