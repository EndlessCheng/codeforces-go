## 需要注意的细节

1. 坐船回来的那个人，不一定是刚才过河的人，也可能是之前（比如上上次）过河的人。
2. 不能干等，船必须时时刻刻都在移动。
3. 可以用一个人来来回回过河。这看似毫无意义，但我们可以用这个操作来近似地实现干等，从而调整到一个倍率很小的阶段上，在这个时候过河，是更优的。

## 思路

由于存在来来回回过河的情况，计算过程中可能会形成环，所以 DP（记忆化搜索）不太合适。

改成建图，跑 Dijkstra 最短路。

把 $(\textit{stage},S)$ 当作节点，其中 $\textit{stage}$ 表示当前阶段，$S$ 表示剩余为过河的人的下标集合。

起点为 $(0,U)$，其中全集 $U=\{0,1,2,\ldots, n-1\}$。

终点为 $S=\varnothing$ 的节点。

对于节点 $(\textit{stage},S)$ 来说，我们枚举 $S$ 的大小 $\le k$ 的非空子集 $T$，作为这次过河的人群。

如果 $T\ne S$，也就是还有没过河的人，那么需要从已过河的人 $\complement_U (S\setminus T)$ 中枚举一个人回来。

代码实现时，用二进制表示集合，用位运算操作集合，具体见 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

## 答疑

**问**：为什么回来的人不能贪心地选 $\textit{time}$ 最小的？

**答**：这只是局部最优，不是全局最优。可能存在这样一种情况，选一个 $\textit{time}$ 大的人回来，刚好到了一个 $\textit{mul}$ 很小的阶段，后面过河消耗的时间更少。

**问**：为什么要枚举大小小于 $k$ 的子集？直接坐满 $k$ 人，不是更好吗？

**答**：同上，这只是局部最优，不是全局最优。可能存在这样一种情况，先运一小部分人过河（这些人的 $\textit{time}$ 的最大值比较小），再回来一个人，此时刚好到了一个 $\textit{mul}$ 很小的阶段，适合运 $\textit{time}$ 比较大的人过河，消耗的时间更少。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1GCNRzgEYp/?t=32m18s)，欢迎点赞关注~

## 写法一

```py [sol-Python3]
class Solution:
    def minTime(self, n: int, k: int, m: int, time: List[int], mul: List[float]) -> float:
        u = 1 << n
        # 预处理每个 time 子集的最大值
        max_time = [0] * u
        for i, t in enumerate(time):
            high_bit = 1 << i
            for mask in range(high_bit):
                max_time[high_bit | mask] = max(max_time[mask], t)

        # 预处理每个集合的大小 <= k 的非空子集
        sub_masks = [[] for _ in range(u)]
        for i in range(u):
            sub = i
            while sub:
                if sub.bit_count() <= k:
                    sub_masks[i].append(sub)
                sub = (sub - 1) & i

        dis = [[inf] * u for _ in range(m)]
        h = []

        def push(d: float, stage: int, mask: int) -> None:
            if d < dis[stage][mask]:
                dis[stage][mask] = d
                heappush(h, (d, stage, mask))

        push(0, 0, u - 1)  # 起点

        while h:
            d, stage, left = heappop(h)  # left 是剩余没有过河的人
            if left == 0:  # 所有人都过河了
                return d
            if d > dis[stage][left]:
                continue
            # 枚举 sub 这群人坐一艘船过河
            for sub in sub_masks[left]:
                cost = max_time[sub] * mul[stage]
                cur_stage = (stage + floor(cost)) % m  # 过河后的阶段
                if sub == left:  # 所有人都过河了
                    push(d + cost, cur_stage, 0)
                    continue
                # 枚举回来的人（可以是之前过河的人）
                s = (u - 1) ^ left ^ sub
                while s:
                    lb = s & -s
                    return_time = max_time[lb] * mul[cur_stage]
                    push(d + cost + return_time, (cur_stage + floor(return_time)) % m, left ^ sub ^ lb)
                    s ^= lb

        return -1
```

```java [sol-Java]
class Solution {
    private record Tuple(double dis, int stage, int mask) {
    }

    public double minTime(int n, int k, int m, int[] time, double[] mul) {
        int u = 1 << n;
        // 计算每个 time 子集的最大值
        int[] maxTime = new int[u];
        for (int i = 0; i < n; i++) {
            int highBit = 1 << i;
            for (int mask = 0; mask < highBit; mask++) {
                maxTime[highBit | mask] = Math.max(maxTime[mask], time[i]);
            }
        }
        // 把 maxTime 中的大小大于 k 的集合改为 inf
        for (int i = 0; i < u; i++) {
            if (Integer.bitCount(i) > k) {
                maxTime[i] = Integer.MAX_VALUE;
            }
        }

        double[][] dis = new double[m][u];
        for (double[] row : dis) {
            Arrays.fill(row, Double.MAX_VALUE);
        }

        PriorityQueue<Tuple> h = new PriorityQueue<>(Comparator.comparingDouble(t -> t.dis));
        push(0, 0, u - 1, dis, h); // 起点

        while (!h.isEmpty()) {
            Tuple top = h.poll();
            double d = top.dis;
            int stage = top.stage;
            int left = top.mask; // 剩余没有过河的人
            if (left == 0) { // 所有人都过河了
                return d;
            }
            if (d > dis[stage][left]) {
                continue;
            }
            // 枚举 sub 这群人坐一艘船
            for (int sub = left; sub > 0; sub = (sub - 1) & left) {
                if (maxTime[sub] == Integer.MAX_VALUE) {
                    continue;
                }
                // sub 过河
                double cost = maxTime[sub] * mul[stage];
                int curStage = (stage + (int) cost) % m; // 过河后的阶段
                // 所有人都过河了
                if (sub == left) {
                    push(d + cost, curStage, 0, dis, h);
                    continue;
                }
                // 枚举回来的人（可以是之前过河的人）
                for (int s = (u - 1) ^ left ^ sub, lb = 0; s > 0; s ^= lb) {
                    lb = s & -s;
                    double returnTime = maxTime[lb] * mul[curStage];
                    push(d + cost + returnTime, (curStage + (int) returnTime) % m, left ^ sub ^ lb, dis, h);
                }
            }
        }

        return -1;
    }

    private void push(double d, int stage, int mask, double[][] dis, PriorityQueue<Tuple> pq) {
        if (d < dis[stage][mask]) {
            dis[stage][mask] = d;
            pq.add(new Tuple(d, stage, mask));
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    double minTime(int n, int k, int m, vector<int>& time, vector<double>& mul) {
        int u = 1 << n;
        // 计算每个 time 子集的最大值
        vector<int> max_time(u);
        for (int i = 0; i < n; i++) {
            int t = time[i];
            int high_bit = 1 << i;
            for (int mask = 0; mask < high_bit; mask++) {
                max_time[high_bit | mask] = max({max_time[high_bit | mask], max_time[mask], t});
            }
        }
        // 把 max_time 中的大小大于 k 的集合改为 inf
        for (uint32_t i = 0; i < u; i++) {
            if (popcount(i) > k) {
                max_time[i] = INT_MAX;
            }
        }

        vector dis(m, vector<double>(u, DBL_MAX));
        using T = tuple<double, int, int>;
        priority_queue<T, vector<T>, greater<>> pq;

        auto push = [&](double d, int stage, int mask) {
            if (d < dis[stage][mask]) {
                dis[stage][mask] = d;
                pq.emplace(d, stage, mask);
            }
        };

        push(0, 0, u - 1); // 起点

        while (!pq.empty()) {
            auto [d, stage, left] = pq.top();
            pq.pop();
            if (left == 0) { // 所有人都过河了
                return d;
            }
            if (d > dis[stage][left]) {
                continue;
            }
            // 枚举 sub 这群人坐一艘船
            for (int sub = left; sub > 0; sub = (sub - 1) & left) {
                if (max_time[sub] == INT_MAX) {
                    continue;
                }
                // sub 过河
                double cost = max_time[sub] * mul[stage];
                int cur_stage = (stage + int(cost)) % m; // 过河后的阶段
                // 所有人都过河了
                if (sub == left) {
                    push(d + cost, cur_stage, 0);
                    continue;
                }
                // 枚举回来的人（可以是之前过河的人）
                for (int s = (u - 1) ^ left ^ sub, lb; s > 0; s ^= lb) {
                    lb = s & -s;
                    double return_time = max_time[lb] * mul[cur_stage];
                    push(d + cost + return_time, (cur_stage + int(return_time)) % m, left ^ sub ^ lb);
                }
            }
        }
        return -1;
    }
};
```

```go [sol-Go]
func minTime(n, k, m int, time []int, mul []float64) float64 {
	u := 1 << n
	// 计算每个 time 子集的最大值
	maxTime := make([]int, u)
	for i, t := range time {
		highBit := 1 << i
		for mask, mx := range maxTime[:highBit] {
			maxTime[highBit|mask] = max(mx, t)
		}
	}
	// 把 maxTime 中的大小大于 k 的集合改为 inf
	for i := range maxTime {
		if bits.OnesCount(uint(i)) > k {
			maxTime[i] = math.MaxInt
		}
	}

	dis := make([][]float64, m)
	for i := range dis {
		dis[i] = make([]float64, u)
		for j := range dis[i] {
			dis[i][j] = math.MaxFloat64
		}
	}
	h := hp{}
	push := func(d float64, stage, mask int) {
		if d < dis[stage][mask] {
			dis[stage][mask] = d
			heap.Push(&h, tuple{d, stage, mask})
		}
	}

	push(0, 0, u-1) // 起点

	for len(h) > 0 {
		top := heap.Pop(&h).(tuple)
		d := top.dis
		stage := top.stage
		left := top.mask // 剩余没有过河的人
		if left == 0 {   // 所有人都过河了
			return d
		}
		if d > dis[stage][left] {
			continue
		}
		// 枚举 sub 这群人坐一艘船
		for sub := left; sub > 0; sub = (sub - 1) & left {
			if maxTime[sub] == math.MaxInt {
				continue
			}
			// sub 过河
			cost := float64(maxTime[sub]) * mul[stage]
			curStage := (stage + int(cost)) % m // 过河后的阶段
			// 所有人都过河了
			if sub == left {
				push(d+cost, curStage, 0)
				continue
			}
			// 枚举回来的人（可以是之前过河的人）
			for s, lb := u-1^left^sub, 0; s > 0; s ^= lb {
				lb = s & -s
				returnTime := float64(maxTime[lb]) * mul[curStage]
				push(d+cost+returnTime, (curStage+int(returnTime))%m, left^sub^lb)
			}
		}
	}
	return -1
}

type tuple struct {
	dis         float64
	stage, mask int
}
type hp []tuple
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
```

#### 复杂度分析

对于大小为 $n$ 的集合，它的大小为 $m$ 的子集有 $\binom n m$ 个，每个子集又有 $2^m$ 个子集。根据二项式定理，$\sum\limits_{m=0}^n \binom n m 2^m = (2+1)^n = 3^n$，所以「枚举子集的子集」的时间复杂度为 $\mathcal{O}(3^n)$。

- 时间复杂度：$\mathcal{O}(M\log M)$，其中 $M=nm3^n$ 是图中的边数上限。
- 空间复杂度：$\mathcal{O}(M)$。

## 写法二：状态机

写法一的优化点是，不同的集合 $A$ 在去掉坐船过河的子集后，得到的集合 $B$ 可以是相同的，这些集合 $B$ 在「枚举坐船回来的人」时，枚举的内容是完全一样的，这些重复的枚举可以优化。

添加一个维度 $\textit{direction}$，把 $(\textit{stage},S,\textit{direction})$ 当作节点，其中：

- $\textit{direction}=0$ 表示一群人要过河的状态。
- $\textit{direction}=1$ 表示一个人要回来的状态。

起点为 $(0,U,0)$，其中全集 $U=\{0,1,2,\ldots, n-1\}$。

终点为 $S=\varnothing$ 的节点。

```py [sol-Python3]
class Solution:
    def minTime(self, n: int, k: int, m: int, time: List[int], mul: List[float]) -> float:
        u = 1 << n
        # 预处理每个 time 子集的最大值
        max_time = [0] * u
        for i, t in enumerate(time):
            high_bit = 1 << i
            for mask in range(high_bit):
                max_time[high_bit | mask] = max(max_time[mask], t)

        # 预处理每个集合的大小 <= k 的非空子集
        sub_masks = [[] for _ in range(u)]
        for i in range(u):
            sub = i
            while sub:
                if sub.bit_count() <= k:
                    sub_masks[i].append(sub)
                sub = (sub - 1) & i

        dis = [[[inf, inf] for _ in range(u)] for _ in range(m)]
        h = []

        def push(d: float, stage: int, mask: int, direction: int) -> None:
            if d < dis[stage][mask][direction]:
                dis[stage][mask][direction] = d
                heappush(h, (d, stage, mask, direction))

        push(0, 0, u - 1, 0)  # 起点

        while h:
            d, stage, left, direction = heappop(h)  # left 是剩余没有过河的人
            if left == 0:  # 所有人都过河了
                return d
            if d > dis[stage][left][direction]:
                continue
            if direction == 0:
                # 枚举 sub 这群人坐一艘船过河
                for sub in sub_masks[left]:
                    cost = max_time[sub] * mul[stage]
                    push(d + cost, (stage + floor(cost)) % m, left ^ sub, 1)
            else:
                # 枚举回来的人
                s = (u - 1) ^ left
                while s:
                    lb = s & -s
                    cost = max_time[lb] * mul[stage]
                    push(d + cost, (stage + floor(cost)) % m, left ^ lb, 0)
                    s ^= lb

        return -1
```

```java [sol-Java]
class Solution {
    private record Tuple(double dis, int stage, int mask, int direction) {
    }

    public double minTime(int n, int k, int m, int[] time, double[] mul) {
        int u = 1 << n;
        // 计算每个 time 子集的最大值
        int[] maxTime = new int[u];
        for (int i = 0; i < n; i++) {
            int highBit = 1 << i;
            for (int mask = 0; mask < highBit; mask++) {
                maxTime[highBit | mask] = Math.max(maxTime[mask], time[i]);
            }
        }
        // 把 maxTime 中的大小大于 k 的集合改为 inf
        for (int i = 0; i < u; i++) {
            if (Integer.bitCount(i) > k) {
                maxTime[i] = Integer.MAX_VALUE;
            }
        }

        double[][][] dis = new double[m][u][2];
        for (double[][] mat : dis) {
            for (double[] row : mat) {
                Arrays.fill(row, Double.MAX_VALUE);
            }
        }

        PriorityQueue<Tuple> h = new PriorityQueue<>(Comparator.comparingDouble(t -> t.dis));
        push(0, 0, u - 1, 0, dis, h); // 起点

        while (!h.isEmpty()) {
            Tuple top = h.poll();
            double d = top.dis;
            int stage = top.stage;
            int left = top.mask; // 剩余没有过河的人
            int direction = top.direction;
            if (left == 0) { // 所有人都过河了
                return d;
            }
            if (d > dis[stage][left][direction]) {
                continue;
            }
            if (direction == 0) {
                // 枚举 sub 这群人坐一艘船过河
                for (int sub = left; sub > 0; sub = (sub - 1) & left) {
                    if (maxTime[sub] != Integer.MAX_VALUE) {
                        double cost = maxTime[sub] * mul[stage];
                        push(d + cost, (stage + (int) cost) % m, left ^ sub, 1, dis, h);
                    }
                }
            } else {
                // 枚举回来的人
                for (int s = (u - 1) ^ left, lb = 0; s > 0; s ^= lb) {
                    lb = s & -s;
                    double cost = maxTime[lb] * mul[stage];
                    push(d + cost, (stage + (int) cost) % m, left ^ lb, 0, dis, h);
                }
            }
        }

        return -1;
    }

    private void push(double d, int stage, int mask, int direction, double[][][] dis, PriorityQueue<Tuple> pq) {
        if (d < dis[stage][mask][direction]) {
            dis[stage][mask][direction] = d;
            pq.add(new Tuple(d, stage, mask, direction));
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    double minTime(int n, int k, int m, vector<int>& time, vector<double>& mul) {
        int u = 1 << n;
        // 计算每个 time 子集的最大值
        vector<int> max_time(u);
        for (int i = 0; i < n; i++) {
            int t = time[i];
            int high_bit = 1 << i;
            for (int mask = 0; mask < high_bit; mask++) {
                max_time[high_bit | mask] = max({max_time[high_bit | mask], max_time[mask], t});
            }
        }
        // 把 max_time 中的大小大于 k 的集合改为 inf
        for (uint32_t i = 0; i < u; i++) {
            if (popcount(i) > k) {
                max_time[i] = INT_MAX;
            }
        }

        vector dis(m, vector<array<double, 2>>(u, {DBL_MAX, DBL_MAX}));
        using T = tuple<double, int, int, uint8_t>;
        priority_queue<T, vector<T>, greater<>> pq;

        auto push = [&](double d, int stage, int mask, uint8_t direction) {
            if (d < dis[stage][mask][direction]) {
                dis[stage][mask][direction] = d;
                pq.emplace(d, stage, mask, direction);
            }
        };

        push(0, 0, u - 1, 0); // 起点

        while (!pq.empty()) {
            auto [d, stage, left, direction] = pq.top();
            pq.pop();
            if (left == 0) { // 所有人都过河了
                return d;
            }
            if (d > dis[stage][left][direction]) {
                continue;
            }
            if (direction == 0) {
                // 枚举 sub 这群人坐一艘船过河
                for (int sub = left; sub > 0; sub = (sub - 1) & left) {
                    if (max_time[sub] != INT_MAX) {
                        double cost = max_time[sub] * mul[stage];
                        push(d + cost, (stage + int(cost)) % m, left ^ sub, 1);
                    }
                }
            } else {
                // 枚举回来的人
                for (int s = (u - 1) ^ left, lb; s > 0; s ^= lb) {
                    lb = s & -s;
                    double cost = max_time[lb] * mul[stage];
                    push(d + cost, (stage + int(cost)) % m, left ^ lb, 0);
                }
            }
        }
        return -1;
    }
};
```

```go [sol-Go]
func minTime(n, k, m int, time []int, mul []float64) float64 {
	u := 1 << n
	// 计算每个 time 子集的最大值
	maxTime := make([]int, u)
	for i, t := range time {
		highBit := 1 << i
		for mask, mx := range maxTime[:highBit] {
			maxTime[highBit|mask] = max(mx, t)
		}
	}
	// 把 maxTime 中的大小大于 k 的集合改为 inf
	for i := range maxTime {
		if bits.OnesCount(uint(i)) > k {
			maxTime[i] = math.MaxInt
		}
	}

	dis := make([][][2]float64, m)
	for i := range dis {
		dis[i] = make([][2]float64, u)
		for j := range dis[i] {
			dis[i][j] = [2]float64{math.MaxFloat64, math.MaxFloat64}
		}
	}
	h := hp{}
	push := func(d float64, stage, mask int, direction uint8) {
		if d < dis[stage][mask][direction] {
			dis[stage][mask][direction] = d
			heap.Push(&h, tuple{d, stage, mask, direction})
		}
	}

	push(0, 0, u-1, 0) // 起点

	for len(h) > 0 {
		top := heap.Pop(&h).(tuple)
		d := top.dis
		stage := top.stage
		left := top.mask // 剩余没有过河的人
		direction := top.direction
		if left == 0 { // 所有人都过河了
			return d
		}
		if d > dis[stage][left][direction] {
			continue
		}
		if direction == 0 {
			// 枚举 sub 这群人坐一艘船
			for sub := left; sub > 0; sub = (sub - 1) & left {
				if maxTime[sub] != math.MaxInt {
					cost := float64(maxTime[sub]) * mul[stage]
					push(d+cost, (stage+int(cost))%m, left^sub, 1)
				}
			}
		} else {
			// 枚举回来的人
			for s, lb := u-1^left, 0; s > 0; s ^= lb {
				lb = s & -s
				cost := float64(maxTime[lb]) * mul[stage]
				push(d+cost, (stage+int(cost))%m, left^lb, 0)
			}
		}
	}
	return -1
}

type tuple struct {
	dis         float64
	stage, mask int
	direction   uint8 // 状态机：0 未过河，1 已过河
}
type hp []tuple
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(M\log M)$，其中 $M=m3^n$ 是图中的边数上限。
- 空间复杂度：$\mathcal{O}(M)$。

## 相似题目

1. 图论题单的「**§3.1 单源最短路：Dijkstra 算法**」。
2. 动态规划题单的「**§9.4 子集状压 DP**」。

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
