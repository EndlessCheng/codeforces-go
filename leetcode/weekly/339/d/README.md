### 本题视频讲解

见[【周赛 339】](https://www.bilibili.com/video/BV1va4y1M7Fr/)。

### 提示 1

对于子数组 $[L,R]$ 中的任意下标 $i$，翻转后的下标是 $L+R-i$（中心对称翻转，两个下标相加恒等于 $L+R$）。

那么：

- 当子数组向右滑动时，$L$ 和 $R$ 都增加 $1$，所以翻转后的下标会增加 $2$
- 当子数组向左滑动时，$L$ 和 $R$ 都减少 $1$，所以翻转后的下标会减少 $2$

因此，**$i$ 翻转后的所有位置组成了一个公差为 $2$ 的等差数列**（不考虑 $\textit{banned}$）。

如何求出这些位置的范围呢？注意当 $i$ 在数组边界 $0$ 或 $n-1$ 附近时，有些位置是无法翻转到的。

### 提示 2

- 如果不考虑数组的边界，那么范围是 $[i-k+1, i+k-1]$。
- 如果 $i$ 在数组左边界 $0$ 附近，那么翻转时会受到数组左边界的约束，当子数组在最左边时，$L=0,R=k-1$，$i$ 翻转后是 $0+(k-1)-i=k-i-1$，所以小于 $k-i-1$ 的点是无法翻转到的；
- 如果 $i$ 在数组右边界 $n-1$ 附近，那么翻转时会受到数组右边界的约束，当子数组在最右边时，$L=n-k,R=n-1$，$i$ 翻转后是 $(n-k)+(n-1) - i=2n - k - i - 1$，所以大于 $2n - k - i - 1$ 的点是无法翻转到的。

所以实际范围为

$$
[\max(i-k+1,k-i-1), \min(i+k-1,2n - k - i - 1)]
$$

### 提示 3

用两棵平衡树分别维护不等于 $p$ 也不在 $\textit{banned}$ 中的偶数下标和奇数下标。

然后用 BFS **模拟**。

在对应的平衡树上，一边遍历翻转后的所有位置，一边把平衡树上的下标删除，加到队列中。这样可以避免重复访问已经访问过的下标，加快执行时间。

```cpp [sol1-C++]
class Solution {
public:
    vector<int> minReverseOperations(int n, int p, vector<int> &banned, int k) {
        unordered_set<int> ban{banned.begin(), banned.end()};
        set<int> sets[2];
        for (int i = 0; i < n; ++i)
            if (i != p && !ban.count(i))
                sets[i % 2].insert(i);
        sets[0].insert(n);
        sets[1].insert(n); // 哨兵

        vector<int> ans(n, -1);
        vector<int> q = {p};
        for (int step = 0; !q.empty(); ++step) {
            vector<int> nq;
            for (int i: q) {
                ans[i] = step;
                // s 中的从 mn 到 mx 的所有位置都可以翻转到
                int mn = max(i - k + 1, k - i - 1);
                int mx = min(i + k - 1, n * 2 - k - i - 1);
                auto &s = sets[mn % 2];
                for (auto it = s.lower_bound(mn); *it <= mx; it = s.erase(it))
                    nq.push_back(*it);
            }
            q = move(nq);
        }
        return ans;
    }
};
```

```java [sol1-Java]
class Solution {
    public int[] minReverseOperations(int n, int p, int[] banned, int k) {
        var ban = new boolean[n];
        ban[p] = true;
        for (int i : banned) ban[i] = true;
        TreeSet<Integer>[] sets = new TreeSet[2];
        sets[0] = new TreeSet<>();
        sets[1] = new TreeSet<>();
        for (int i = 0; i < n; i++)
            if (!ban[i])
                sets[i % 2].add(i);
        sets[0].add(n);
        sets[1].add(n); // 哨兵

        var ans = new int[n];
        Arrays.fill(ans, -1);
        var q = new ArrayList<Integer>();
        q.add(p);
        for (int step = 0; !q.isEmpty(); ++step) {
            var tmp = q;
            q = new ArrayList<>();
            for (int i : tmp) {
                ans[i] = step;
                // s 中的从 mn 到 mx 的所有位置都可以翻转到
                int mn = Math.max(i - k + 1, k - i - 1);
                int mx = Math.min(i + k - 1, n * 2 - k - i - 1);
                var s = sets[mn % 2];
                for (var j = s.ceiling(mn); j <= mx; j = s.ceiling(mn)) {
                    q.add(j);
                    s.remove(j);
                }
            }
        }
        return ans;
    }
}
```

并查集的思路是，如果要删除一个元素，那么把它的下标 $j$ 和 $j+1$ 合并，这样后面删除的时候就会自动跳过已删除的元素。

```py [sol2-Python3]
class Solution:
    def minReverseOperations(self, n: int, p: int, banned: List[int], k: int) -> List[int]:
        s = set(banned) | {p}
        not_banned = [[], []]
        for i in range(n):
            if i not in s:
                not_banned[i % 2].append(i)
        not_banned[0].append(n)
        not_banned[1].append(n)  # 哨兵

        fa = [list(range(len(not_banned[0]))), list(range(len(not_banned[1])))]

        def find(i: int, x: int) -> int:
            f = fa[i]
            if f[x] != x:
                f[x] = find(i, f[x])
            return f[x]

        def merge(i: int, from_: int, to: int) -> None:
            x, y = find(i, from_), find(i, to)
            fa[i][x] = y

        ans = [-1] * n
        q = [p]
        step = 0
        while q:
            tmp = q
            q = []
            for i in tmp:
                ans[i] = step
                # a 中的从 mn 到 mx 的所有位置都可以翻转到
                mn = max(i - k + 1, k - i - 1)
                mx = min(i + k - 1, n * 2 - k - i - 1)
                a = not_banned[mn % 2]
                j = find(mn % 2, bisect_left(a, mn))
                while a[j] <= mx:
                    q.append(a[j])
                    merge(mn % 2, j, j + 1)  # 删除 j
                    j = find(mn % 2, j + 1)
            step += 1
        return ans
```

```go [sol2-Go]
type uf struct {
	fa []int
}

func newUnionFind(n int) uf {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return uf{fa}
}

func (u *uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *uf) merge(from, to int) {
	x, y := u.find(from), u.find(to)
	u.fa[x] = y
}

func minReverseOperations(n, p int, banned []int, k int) []int {
	ban := map[int]bool{p: true}
	for _, v := range banned {
		ban[v] = true
	}
	notBanned := [2][]int{}
	for i := 0; i < n; i++ {
		if !ban[i] {
			notBanned[i%2] = append(notBanned[i%2], i)
		}
	}
	notBanned[0] = append(notBanned[0], n)
	notBanned[1] = append(notBanned[1], n) // 哨兵
	ufs := [2]uf{newUnionFind(len(notBanned[0])), newUnionFind(len(notBanned[1]))}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	q := []int{p}
	for step := 0; len(q) > 0; step++ {
		tmp := q
		q = nil
		for _, i := range tmp {
			ans[i] = step
			// a 中的从 mn 到 mx 的所有位置都可以翻转到
			mn := max(i-k+1, k-i-1)
			mx := min(i+k-1, n*2-k-i-1)
			a, u := notBanned[mn%2], ufs[mn%2]
			for j := u.find(sort.SearchInts(a, mn)); a[j] <= mx; j = u.find(j + 1) {
				q = append(q, a[j])
				u.merge(j, j+1) // 删除 j
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$。
- 空间复杂度：$O(n)$。

### 相似题目（并查集维护添加或者删除）

- [1851. 包含每个查询的最小区间](https://leetcode.cn/problems/minimum-interval-to-include-each-query/)
- [2382. 删除操作后的最大子段和](https://leetcode.cn/problems/maximum-segment-sum-after-removals/)
- [2334. 元素值大于变化阈值的子数组](https://leetcode.cn/problems/subarray-with-elements-greater-than-varying-threshold/)
