[本题视频讲解](https://www.bilibili.com/video/BV1pC4y1j7Pw/)

## 方法一：离线做法+最小堆

> 离线：按照自己定义的某种顺序回答询问（而不是按照输入顺序一个个地回答）。

不妨设 $a_i \le b_i$。

如果 $a_i = b_i$ 或者 $\textit{heights}[a_i] <\textit{heights}[b_i]$，那么 Alice 可以直接跳到 Bob 的位置，即 $\textit{ans}[i] = b_i$。

否则，我们可以在 $b_i$ 处**记录**「左边有个 $a_i$，它属于第 $i$ 个询问」。

然后遍历 $\textit{heights}$，同时用一个**最小堆**维护上面说的「记录」：遍历到 $\textit{heights}[i]$ 时，把在下标 $i$ 处的「记录」全部加到最小堆中。

在加到最小堆之前，我们可以回答左边所有高度小于 $\textit{heights}[i]$ 的「记录」，其答案就是 $i$。

#### 总结

这个算法涉及到三个位置，按照**从左到右**的顺序，它们分别是：

1. $a_i$：回答询问时，用它的高度来和当前高度判断。
2. $b_i$：决定了在什么位置把询问加入堆中。
3. 回答询问的位置。

```py [sol-Python3]
class Solution:
    def leftmostBuildingQueries(self, heights: List[int], queries: List[List[int]]) -> List[int]:
        ans = [-1] * len(queries)
        left = [[] for _ in heights]
        for qi, (i, j) in enumerate(queries):
            if i > j:
                i, j = j, i  # 保证 i <= j
            if i == j or heights[i] < heights[j]:
                ans[qi] = j  # i 直接跳到 j
            else:
                left[j].append((heights[i], qi))  # 离线

        h = []
        for i, x in enumerate(heights):  # 从小到大枚举下标 i
            while h and h[0][0] < x:
                ans[heappop(h)[1]] = i  # 可以跳到 i（此时 i 是最小的）
            for p in left[i]:
                heappush(h, p)  # 后面再回答
        return ans
```

```java [sol-Java]
class Solution {
    public int[] leftmostBuildingQueries(int[] heights, int[][] queries) {
        int[] ans = new int[queries.length];
        Arrays.fill(ans, -1);
        List<int[]>[] left = new ArrayList[heights.length];
        Arrays.setAll(left, e -> new ArrayList<>());
        for (int qi = 0; qi < queries.length; qi++) {
            int i = queries[qi][0], j = queries[qi][1];
            if (i > j) {
                int temp = i;
                i = j;
                j = temp; // 保证 i <= j
            }
            if (i == j || heights[i] < heights[j]) {
                ans[qi] = j; // i 直接跳到 j
            } else {
                left[j].add(new int[]{heights[i], qi}); // 离线
            }
        }

        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]);
        for (int i = 0; i < heights.length; i++) { // 从小到大枚举下标 i
            while (!pq.isEmpty() && pq.peek()[0] < heights[i]) {
                ans[pq.poll()[1]] = i; // 可以跳到 i（此时 i 是最小的）
            }
            for (int[] p : left[i]) {
                pq.offer(p); // 后面再回答
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> leftmostBuildingQueries(vector<int> &heights, vector<vector<int>> &queries) {
        vector<int> ans(queries.size(), -1);
        vector<vector<pair<int, int>>> left(heights.size());
        for (int qi = 0; qi < queries.size(); qi++) {
            int i = queries[qi][0], j = queries[qi][1];
            if (i > j) {
                swap(i, j); // 保证 i <= j
            }
            if (i == j || heights[i] < heights[j]) {
                ans[qi] = j; // i 直接跳到 j
            } else {
                left[j].emplace_back(heights[i], qi); // 离线
            }
        }

        priority_queue<pair<int, int>, vector<pair<int, int>>, greater<>> pq;
        for (int i = 0; i < heights.size(); i++) { // 从小到大枚举下标 i
            while (!pq.empty() && pq.top().first < heights[i]) {
                ans[pq.top().second] = i; // 可以跳到 i（此时 i 是最小的）
                pq.pop();
            }
            for (auto &p: left[i]) {
                pq.emplace(p); // 后面再回答
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for i := range ans {
		ans[i] = -1
	}
	left := make([][]pair, len(heights))
	for qi, q := range queries {
		i, j := q[0], q[1]
		if i > j {
			i, j = j, i // 保证 i <= j
		}
		if i == j || heights[i] < heights[j] {
			ans[qi] = j // i 直接跳到 j
		} else {
			left[j] = append(left[j], pair{heights[i], qi}) // 离线
		}
	}

	h := hp{}
	for i, x := range heights { // 从小到大枚举下标 i
		for h.Len() > 0 && h[0].h < x {
			ans[heap.Pop(&h).(pair).qi] = i // 可以跳到 i（此时 i 是最小的）
		}
		for _, p := range left[i] {
			heap.Push(&h, p) // 后面再回答
		}
	}
	return ans
}

type pair struct{ h, qi int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].h < h[j].h }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q\log q)$，其中 $n$ 为 $\textit{heights}$ 的长度，$q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n + q)$。

## 方法二：在线做法+线段树二分

构建一棵维护区间**最大值** $\textit{mx}$ 的线段树。

方法一中用堆回答的询问，相当于问区间 $[j+1,n-1]$ 中的大于 $v = \textit{heights}[i]$ 的最小下标。

由于代码中线段树的下标是从 $1$ 开始的，所以区间是 $[j+2,n]$。不过为了避免讨论 $j+2>n$ 的情况，代码中用的 $j+1$。

- 如果当前区间 $\textit{mx}\le v$，则整个区间都不存在大于 $v$ 的数，返回 $0$。
- 如果当前区间只包含一个元素，则找到答案，返回该元素的下标。
- 如果左子树包含下标 $j+1$，则递归左子树。
- 如果左子树返回 $0$，则返回递归右子树的结果。

```py [sol-Python3]
class Solution:
    def leftmostBuildingQueries(self, heights: List[int], queries: List[List[int]]) -> List[int]:
        n = len(heights)
        mx = [0] * (n * 4)

        def build(o: int, l: int, r: int) -> None:
            if l == r:
                mx[o] = heights[l - 1]
                return
            m = (l + r) // 2
            build(o * 2, l, m)
            build(o * 2 + 1, m + 1, r)
            mx[o] = max(mx[o * 2], mx[o * 2 + 1])

        # 返回 [L,n] 中 > v 的最小下标（前三个参数表示线段树的节点信息）
        def query(o: int, l: int, r: int, L: int, v: int) -> int:
            if v >= mx[o]:  # 最大值 <= v，没法找到 > v 的数
                return 0
            if l == r:  # 找到了
                return l
            m = (l + r) // 2
            if L <= m:
                pos = query(o * 2, l, m, L, v)  # 递归左子树
                if pos > 0:  # 找到了
                    return pos
            return query(o * 2 + 1, m + 1, r, L, v)  # 递归右子树

        build(1, 1, n)
        ans = []
        for i, j in queries:
            if i > j:
                i, j = j, i
            if i == j or heights[i] < heights[j]:
                ans.append(j)
            else:
                pos = query(1, 1, n, j + 1, heights[i])
                ans.append(pos - 1)  # 不存在时刚好得到 -1
        return ans
```

```java [sol-Java]
public class Solution {
    public int[] leftmostBuildingQueries(int[] heights, int[][] queries) {
        int n = heights.length;
        mx = new int[n * 4];
        build(1, 1, n, heights);

        int[] ans = new int[queries.length];
        for (int qi = 0; qi < queries.length; qi++) {
            int i = queries[qi][0], j = queries[qi][1];
            if (i > j) {
                int temp = i;
                i = j;
                j = temp;
            }
            if (i == j || heights[i] < heights[j]) {
                ans[qi] = j;
            } else {
                int pos = query(1, 1, n, j + 1, heights[i]);
                ans[qi] = pos - 1; // 不存在时刚好得到 -1
            }
        }
        return ans;
    }

    private int[] mx;

    private void build(int o, int l, int r, int[] heights) {
        if (l == r) {
            mx[o] = heights[l - 1];
            return;
        }
        int m = (l + r) / 2;
        build(o * 2, l, m, heights);
        build(o * 2 + 1, m + 1, r, heights);
        mx[o] = Math.max(mx[o * 2], mx[o * 2 + 1]);
    }

    // 返回 [L,n] 中 > v 的最小下标（前三个参数表示线段树的节点信息）
    private int query(int o, int l, int r, int L, int v) {
        if (v >= mx[o]) { // 最大值 <= v，没法找到 > v 的数
            return 0;
        }
        if (l == r) { // 找到了
            return l;
        }
        int m = (l + r) / 2;
        if (L <= m) {
            int pos = query(o * 2, l, m, L, v); // 递归左子树
            if (pos > 0) { // 找到了
                return pos;
            }
        }
        return query(o * 2 + 1, m + 1, r, L, v); // 递归右子树
    }
}
```

```cpp [sol-C++]
class Solution {
private:
    vector<int> mx;

    void build(int o, int l, int r, vector<int> &heights) {
        if (l == r) {
            mx[o] = heights[l - 1];
            return;
        }
        int m = (l + r) / 2;
        build(o * 2, l, m, heights);
        build(o * 2 + 1, m + 1, r, heights);
        mx[o] = max(mx[o * 2], mx[o * 2 + 1]);
    }

    // 返回 [L,n] 中 > v 的最小下标（前三个参数表示线段树的节点信息）
    int query(int o, int l, int r, int L, int v) {
        if (v >= mx[o]) { // 最大值 <= v，没法找到 > v 的数
            return 0;
        }
        if (l == r) { // 找到了
            return l;
        }
        int m = (l + r) / 2;
        if (L <= m) {
            int pos = query(o * 2, l, m, L, v); // 递归左子树
            if (pos > 0) { // 找到了
                return pos;
            }
        }
        return query(o * 2 + 1, m + 1, r, L, v); // 递归右子树
    }

public:
    vector<int> leftmostBuildingQueries(vector<int> &heights, vector<vector<int>> &queries) {
        int n = heights.size();
        mx.resize(n * 4);
        build(1, 1, n, heights);

        vector<int> ans;
        for (auto &q : queries) {
            int i = q[0], j = q[1];
            if (i > j) {
                swap(i, j);
            }
            if (i == j || heights[i] < heights[j]) {
                ans.push_back(j);
            } else {
                int pos = query(1, 1, n, j + 1, heights[i]);
                ans.push_back(pos - 1); // 不存在时刚好得到 -1
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
type seg []int

func (t seg) build(a []int, o, l, r int) {
	if l == r {
		t[o] = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t[o] = max(t[o<<1], t[o<<1|1])
}

// 返回 [L,n] 中 > v 的最小下标（前三个参数表示线段树的节点信息）
func (t seg) query(o, l, r, L, v int) int {
	if v >= t[o] { // 最大值 <= v，没法找到 > v 的数
		return 0
	}
	if l == r { // 找到了
		return l
	}
	m := (l + r) >> 1
	if L <= m {
		pos := t.query(o<<1, l, m, L, v) // 递归左子树
		if pos > 0 { // 找到了
			return pos
		}
	}
	return t.query(o<<1|1, m+1, r, L, v) // 递归右子树
}

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	n := len(heights)
	t := make(seg, n*4)
	t.build(heights, 1, 1, n)
	ans := make([]int, len(queries))
	for qi, q := range queries {
		i, j := q[0], q[1]
		if i > j {
			i, j = j, i
		}
		if i == j || heights[i] < heights[j] {
			ans[qi] = j
		} else {
			pos := t.query(1, 1, n, j+1, heights[i])
			ans[qi] = pos - 1 // 不存在时刚好得到 -1
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q\log n)$，其中 $n$ 为 $\textit{heights}$ 的长度，$q$ 为 $\textit{queries}$ 的长度。对于左子树的递归，时间是 $\mathcal{O}(\log n)$ 的（同单点更新）；对于右子树的递归，由于区间满足 $\textit{max}\le v$ 则不递归，否则只会向下递归，所以这部分的时间也是 $\mathcal{O}(\log n)$ 的，所以线段树二分的时间复杂度为 $\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

#### 相似题目

下面这题强制在线：

- [2286. 以组为单位订音乐会的门票](https://leetcode.cn/problems/booking-concert-tickets-in-groups/)
