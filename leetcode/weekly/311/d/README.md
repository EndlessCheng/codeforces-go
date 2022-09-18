下午 2 点在 B 站直播讲周赛和双周赛的题目，**包括本题字典树（trie）的原理**，[欢迎关注](https://space.bilibili.com/206214/dynamic)~

---

根据题意，我们可以用字典树（trie）存储所有字符串，由于每个节点都是其子树节点的前缀，题干中的**分数**就是对应节点的子树大小。

然后 DFS 这颗字典树，累加路径上的分数，就可以得到每个字符串的所有非空前缀的分数总和了。

代码实现时，由于可能有相同字符串，每个字符串对应的节点需要用一个列表存储该字符串在 $\textit{words}$ 中的所有下标。

```py [sol1-Python3]
class Node:
    __slots__ = 'son', 'ids', 'score'

    def __init__(self):
        self.son = defaultdict(Node)
        self.ids = []
        self.score = 0

class Solution:
    def sumPrefixScores(self, words: List[str]) -> List[int]:
        root = Node()
        for i, word in enumerate(words):
            cur = root
            for c in word:
                cur = cur.son[c]
                cur.score += 1  # 更新所有前缀的分数
            cur.ids.append(i)

        ans = [0] * len(words)
        def dfs(node: Node, sum: int) -> None:
            sum += node.score  # 累加分数，即可得到答案
            for i in node.ids:
                ans[i] = sum
            for child in node.son.values():
                if child:
                    dfs(child, sum)
        dfs(root, 0)
        return ans
```

```java [sol1-Java]
class Node {
    Node[] son = new Node[26];
    List<Integer> ids = new ArrayList<>();
    int score;
}

class Solution {
    int[] ans;

    private void dfs(Node node, int sum) {
        sum += node.score; // 累加分数，即可得到答案
        for (var i : node.ids)
            ans[i] += sum;
        for (var child : node.son)
            if (child != null)
                dfs(child, sum);
    }

    public int[] sumPrefixScores(String[] words) {
        var n = words.length;
        var root = new Node();
        for (int i = 0; i < n; i++) {
            var cur = root;
            for (var c : words[i].toCharArray()) {
                c -= 'a';
                if (cur.son[c] == null) cur.son[c] = new Node();
                cur = cur.son[c];
                ++cur.score; // 更新所有前缀的分数
            }
            cur.ids.add(i);
        }
        ans = new int[n];
        dfs(root, 0);
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
    struct Node {
        Node *son[26]{};
        vector<int> ids;
        int score = 0;
    };

    vector<int> ans;

    void dfs(Node *node, int sum) {
        sum += node->score; // 累加分数，即可得到答案
        for (int i : node->ids)
            ans[i] += sum;
        for (auto child : node->son)
            if (child != nullptr)
                dfs(child, sum);
    }

public:
    vector<int> sumPrefixScores(vector<string> &words) {
        int n = words.size();
        Node *root = new Node();
        for (int i = 0; i < n; ++i) {
            auto cur = root;
            for (char c : words[i]) {
                c -= 'a';
                if (cur->son[c] == nullptr) cur->son[c] = new Node();
                cur = cur->son[c];
                ++cur->score; // 更新所有前缀的分数
            }
            cur->ids.push_back(i);
        }
        ans.resize(n);
        dfs(root, 0);
        return ans;
    }
};
```

```go [sol1-Go]
func sumPrefixScores(words []string) []int {
	type node struct {
		son   [26]*node
		ids   []int
		score int
	}
	root := &node{}
	for i, word := range words {
		cur := root
		for _, c := range word {
			c -= 'a'
			if cur.son[c] == nil {
				cur.son[c] = &node{}
			}
			cur = cur.son[c]
			cur.score++ // 更新所有前缀的分数
		}
		cur.ids = append(cur.ids, i)
	}

	ans := make([]int, len(words))
	var dfs func(*node, int)
	dfs = func(node *node, sum int) {
		sum += node.score // 累加分数，即可得到答案
		for _, i := range node.ids {
			ans[i] = sum
		}
		for _, child := range node.son {
			if child != nil {
				dfs(child, sum)
			}
		}
	}
	dfs(root, 0)
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$O(L)$，其中 $L$ 为 $\textit{word}_i$ 的长度之和。
- 空间复杂度：$O(L)$。
