[视频讲解](https://www.bilibili.com/video/BV1yu4y1z7sE/) 第四题。

![w364D-2-c.png](https://pic.leetcode.cn/1708991317-yqozUn-w364D-2-c.png)

代码实现时，为避免反复 DFS 同一个非质数连通块，可以把每个非质数所处的连通块的大小记录下来（类似记忆化搜索）。如果之前计算过，就无需再次 DFS。

```py [sol-Python3]
# 标记 10**5 以内的质数
MX = 10 ** 5 + 1
is_prime = [True] * MX
is_prime[1] = False
for i in range(2, isqrt(MX) + 1):
    if is_prime[i]:
        for j in range(i * i, MX, i):
            is_prime[j] = False

class Solution:
    def countPaths(self, n: int, edges: List[List[int]]) -> int:
        g = [[] for _ in range(n + 1)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        nodes = []
        def dfs(x: int, fa: int) -> None:
            nodes.append(x)
            for y in g[x]:
                if y != fa and not is_prime[y]:
                    dfs(y, x)

        ans = 0
        size = [0] * (n + 1)
        for x in range(1, n + 1):
            if not is_prime[x]:  # 跳过非质数
                continue
            s = 0
            for y in g[x]:  # 质数 x 把这棵树分成了若干个连通块
                if is_prime[y]:
                    continue
                if size[y] == 0:  # 尚未计算过
                    nodes.clear()
                    dfs(y, -1)  # 遍历 y 所在连通块，在不经过质数的前提下，统计有多少个非质数
                    for z in nodes:
                        size[z] = len(nodes)
                # 这 size[y] 个非质数与之前遍历到的 s 个非质数，两两之间的路径只包含质数 x
                ans += size[y] * s
                s += size[y]
            ans += s  # 从 x 出发的路径
        return ans
```

```java [sol-Java]
class Solution {
    private final static int MX = (int) 1e5;
    private final static boolean[] np = new boolean[MX + 1]; // 质数=false 非质数=true

    static {
        np[1] = true;
        for (int i = 2; i * i <= MX; i++) {
            if (!np[i]) {
                for (int j = i * i; j <= MX; j += i) {
                    np[j] = true;
                }
            }
        }
    }

    public long countPaths(int n, int[][] edges) {
        List<Integer>[] g = new ArrayList[n + 1];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (var e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        long ans = 0;
        int[] size = new int[n + 1];
        var nodes = new ArrayList<Integer>();
        for (int x = 1; x <= n; x++) {
            if (np[x]) { // 跳过非质数
                continue;
            }
            int sum = 0;
            for (int y : g[x]) { // 质数 x 把这棵树分成了若干个连通块
                if (!np[y]) {
                    continue;
                }
                if (size[y] == 0) { // 尚未计算过
                    nodes.clear();
                    dfs(y, -1, g, nodes); // 遍历 y 所在连通块，在不经过质数的前提下，统计有多少个非质数
                    for (int z : nodes) {
                        size[z] = nodes.size();
                    }
                }
                // 这 size[y] 个非质数与之前遍历到的 sum 个非质数，两两之间的路径只包含质数 x
                ans += (long) size[y] * sum;
                sum += size[y];
            }
            ans += sum; // 从 x 出发的路径
        }
        return ans;
    }

    private void dfs(int x, int fa, List<Integer>[] g, List<Integer> nodes) {
        nodes.add(x);
        for (int y : g[x]) {
            if (y != fa && np[y]) {
                dfs(y, x, g, nodes);
            }
        }
    }
}
```

```cpp [sol-C++]
const int MX = 1e5;
bool np[MX + 1]; // 质数=false 非质数=true
int init = []() {
    np[1] = true;
    for (int i = 2; i * i <= MX; i++) {
        if (!np[i]) {
            for (int j = i * i; j <= MX; j += i) {
                np[j] = true;
            }
        }
    }
    return 0;
}();

class Solution {
public:
    long long countPaths(int n, vector<vector<int>> &edges) {
        vector<vector<int>> g(n + 1);
        for (auto &e: edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        vector<int> size(n + 1);
        vector<int> nodes;
        function<void(int, int)> dfs = [&](int x, int fa) {
            nodes.push_back(x);
            for (int y: g[x]) {
                if (y != fa && np[y]) {
                    dfs(y, x);
                }
            }
        };

        long long ans = 0;
        for (int x = 1; x <= n; x++) {
            if (np[x]) continue; // 跳过非质数
            int sum = 0;
            for (int y: g[x]) { // 质数 x 把这棵树分成了若干个连通块
                if (!np[y]) continue;
                if (size[y] == 0) { // 尚未计算过
                    nodes.clear();
                    dfs(y, -1); // 遍历 y 所在连通块，在不经过质数的前提下，统计有多少个非质数
                    for (int z: nodes) {
                        size[z] = nodes.size();
                    }
                }
                // 这 size[y] 个非质数与之前遍历到的 sum 个非质数，两两之间的路径只包含质数 x
                ans += (long long) size[y] * sum;
                sum += size[y];
            }
            ans += sum; // 从 x 出发的路径
        }
        return ans;
    }
};
```

```go [sol-Go]
const mx int = 1e5 + 1
var np = [mx]bool{1: true}
func init() { // 质数=false 非质数=true
	for i := 2; i*i < mx; i++ {
		if !np[i] {
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func countPaths(n int, edges [][]int) (ans int64) {
	g := make([][]int, n+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	size := make([]int, n+1)
	nodes := []int{}
	var dfs func(int, int)
	dfs = func(x, fa int) {
		nodes = append(nodes, x)
		for _, y := range g[x] {
			if y != fa && np[y] {
				dfs(y, x)
			}
		}
	}
	for x := 1; x <= n; x++ {
		if np[x] { // 跳过非质数
			continue
		}
		sum := 0
		for _, y := range g[x] { // 质数 x 把这棵树分成了若干个连通块
			if !np[y] {
				continue
			}
			if size[y] == 0 { // 尚未计算过
				nodes = nodes[:0]
				dfs(y, -1) // 遍历 y 所在连通块，在不经过质数的前提下，统计有多少个非质数
				for _, z := range nodes {
					size[z] = len(nodes)
				}
			}
			// 这 size[y] 个非质数与之前遍历到的 sum 个非质数，两两之间的路径只包含质数 x
			ans += int64(size[y]) * int64(sum)
			sum += size[y]
		}
		ans += int64(sum) // 从 x 出发的路径
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。预处理的时间不计入。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

这题本质上是求一种类似于「非质数-质数-非质数」的路径个数。

这让我想到了另外一道题目 [2242. 节点序列的最大得分](https://leetcode.cn/problems/maximum-score-of-a-node-sequence/)。

这两题的共同点在于「枚举中间」，请读者细细品味。
