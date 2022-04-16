#### 提示 1

设序列为 $a-x-y-b$（$-$ 表示边），枚举 $\textit{edges}$ 中的每条边，作为序列中的 $x-y$。

#### 提示 2

根据提示 1，我们需要把与 $x$ 相邻的点中，分数最大且不同于 $y$ 和 $b$ 的点作为 $a$；把与 $y$ 相邻的点中，分数最大且不同于 $x$ 和 $a$ 的点作为 $b$。

#### 提示 3

与 $x$ 相邻的点中，由于只需要与 $y$ 和 $b$ 不一样，我们仅需要保留分数最大的三个点，$a$ 必定在这三个点中。

#### 提示 4

剩下要做的，就是在枚举 $\textit{edges}$ 前，预处理出这三个点。

代码实现时，可以用排序、堆或者分治（`nth_element`）求前三大。

```Python [sol1-Python3]
class Solution:
    def maximumScore(self, scores: List[int], edges: List[List[int]]) -> int:
        g = [[] for _ in range(len(scores))]
        for x, y in edges:
            g[x].append((scores[y], y))
            g[y].append((scores[x], x))
        for i, vs in enumerate(g):
            g[i] = nlargest(3, vs)

        # 下面这一段可以简写成一行，为了可读性这里就不写了
        ans = -1
        for x, y in edges:
            for (score_a, a), (score_b, b) in product(g[x], g[y]):
                if y != a != b != x:
                    ans = max(ans, score_a + scores[x] + scores[y] + score_b)
        return ans
```

```go [sol1-Go]
func maximumScore(scores []int, edges [][]int) int {
	type nb struct{ to, s int }
	g := make([][]nb, len(scores))
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], nb{y, scores[y]})
		g[y] = append(g[y], nb{x, scores[x]})
	}
	for i, vs := range g {
		sort.Slice(vs, func(i, j int) bool { return vs[i].s > vs[j].s })
		if len(vs) > 3 {
			vs = vs[:3]
		}
		g[i] = vs
	}

	ans := -1
	for _, e := range edges {
		x, y := e[0], e[1]
		for _, p := range g[x] {
			for _, q := range g[y] {
				if p.to != y && q.to != x && p.to != q.to {
					ans = max(ans, p.s+scores[x]+scores[y]+q.s)
				}
			}
		}
	}
	return ans
}

func max(a, b int) int { if b > a { return b }; return a }
```

```C++ [sol1-C++]
class Solution {
public:
    int maximumScore(vector<int> &scores, vector<vector<int>> &edges) {
        int n = scores.size();
        vector<vector<pair<int, int>>> g(n);
        for (auto &e : edges) {
            int x = e[0], y = e[1];
            g[x].emplace_back(-scores[y], y);
            g[y].emplace_back(-scores[x], x);
        }
        for (auto &vs : g)
            if (vs.size() > 3) {
                nth_element(vs.begin(), vs.begin() + 3, vs.end());
                vs.resize(3);
            }

        int ans = -1;
        for (auto &e : edges) {
            int x = e[0], y = e[1];
            for (auto &[score_a, a] : g[x])
                for (auto &[score_b, b] : g[y])
                    if (a != y && b != x && a != b)
                        ans = max(ans, -score_a + scores[x] + scores[y] - score_b);
        }
        return ans;
    }
};
```

```java [sol1-Java]
class Solution {
    public int maximumScore(int[] scores, int[][] edges) {
        var n = scores.length;
        List<int[]>[] g = new ArrayList[n];
        for (var i = 0; i < n; i++)
            g[i] = new ArrayList<>();
        for (var e : edges) {
            int x = e[0], y = e[1];
            g[x].add(new int[]{scores[y], y});
            g[y].add(new int[]{scores[x], x});
        }
        for (var i = 0; i < n; i++)
            if (g[i].size() > 3) {
                Collections.sort(g[i], (a, b) -> (b[0] - a[0]));
                g[i] = new ArrayList<>(g[i].subList(0, 3));
            }

        var ans = -1;
        for (var e : edges) {
            int x = e[0], y = e[1];
            for (var p : g[x]) {
                var a = p[1];
                for (var q : g[y]) {
                    var b = q[1];
                    if (a != y && b != x && a != b)
                        ans = Math.max(ans, p[0] + scores[x] + scores[y] + q[0]);
                }
            }
        }
        return ans;
    }
}
```
