### 本题视频讲解

见[【力扣杯2023春·战队赛】](https://www.bilibili.com/video/TODO/)第四题。

### 思路

1. 要使整棵树对应的字符串的字典序最小，对于根节点下面的子树，对应的字符串的字典序也要最小，这意味着可以递归处理。
2. 如何得到最小字典序？把所有子树的结果排序，然后拼接起来，在开头加上 $0$，末尾加上 $1$，然后返回。
3. 题目说「节点指针最终可以停在任何节点上，不一定要回到根节点」。我们可以把结果的末尾的 $1$ 都去掉来实现这一点（开头的一个 $0$ 也要去掉）。但这样的话，我们对字符串排序的做法是否会影响答案的正确性呢？
4. 假设有两棵子树 $A$ 和 $B$，先 $A$ 后 $B$ 的结果是 $000\cdots 11$，先 $B$ 后 $A$ 的结果（字典序更大）是 $001\cdots 0\cdots 11$，注意中间有个 $0$，因为无论按照何种顺序遍历 $A$ 和 $B$，$0$ 和 $1$ 的总数是不变的。所以原来字典序最小的，去掉末尾的 $1$，字典序仍然是最小的。
5. 因此按照上面第 2 点说的，递归实现就好了。

```py [sol1-Python3]
class Solution:
    def evolutionaryRecord(self, parents: List[int]) -> str:
        n = len(parents)
        g = [[] for _ in range(n)]
        for i in range(1, n):
            g[parents[i]].append(i)  # 建树

        def dfs(x: int) -> str:
            a = sorted(dfs(y) for y in g[x])
            return "0" + ''.join(a) + "1"
        return dfs(0)[1:].rstrip('1')  # 去掉根节点以及返回根节点的路径
```

```go [sol1-Go]
func evolutionaryRecord(parents []int) string {
	n := len(parents)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		p := parents[w]
		g[p] = append(g[p], w) // 建树
	}

	var dfs func(int) string
	dfs = func(v int) string {
		a := make([]string, len(g[v]))
		for i, w := range g[v] {
			a[i] = dfs(w)
		}
		sort.Strings(a)
		return "0" + strings.Join(a, "") + "1"
	}
	return strings.TrimRight(dfs(0)[1:], "1") // 去掉根节点以及返回根节点的路径
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{parents}$ 的长度。瓶颈在拼接字符串上（可以用链表等结构优化），对于排序的时间复杂度，感性理解，如果树高比较小，那么比较的字符串就比较短，如果树高比较大，那么参与比较的字符串就比较少。[具体分析](https://leetcode.cn/problems/special-binary-string/solution/on-log-n-by-hqztrue-nrmw/)。
- 空间复杂度：$\mathcal{O}(n)$。
