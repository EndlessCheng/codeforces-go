对每个非质数 $x$，在 $x$ 和 $x$ 操作一次得到的数 $y$（非质数）之间连一条有向边，边权为 $y$。

答案就是从 $n$ 到 $m$ 的最短路长度，加上 $n$。

这可以用 **Dijkstra 算法**解决。[Dijkstra 算法介绍](https://leetcode.cn/problems/network-delay-time/solution/liang-chong-dijkstra-xie-fa-fu-ti-dan-py-ooe8/)。

代码实现时，不需要连边，而是在 Dijkstra 算法的过程中，计算出 $y$ 是多少。

此外，预处理 $[1,9999]$ 中的每个数是否为质数，这可以用**筛法**。下面用的是埃氏筛。

```py [sol-Python3]
# 埃氏筛，标记每个数是否为质数
MX = 10000
is_prime = [True] * MX
is_prime[1] = False
for i in range(2, MX):
    if is_prime[i]:
        for j in range(i * i, MX, i):
            is_prime[j] = False

class Solution:
    def minOperations(self, n: int, m: int) -> int:
        if is_prime[n] or is_prime[m]:
            return -1
        len_n = len(str(n))
        dis = [inf] * (10 ** len_n)
        dis[n] = n
        h = [(n, n)]
        while h:
            dis_x, x = heappop(h)
            if x == m:
                return dis_x
            if dis_x > dis[x]:
                continue
            v = x
            pow10 = 1
            while v:
                v, d = divmod(v, 10)
                if d > 0:  # 减少
                    y = x - pow10
                    if not is_prime[y] and (new_d := dis_x + y) < dis[y]:
                        dis[y] = new_d
                        heappush(h, (new_d, y))
                if d < 9:  # 增加
                    y = x + pow10
                    if not is_prime[y] and (new_d := dis_x + y) < dis[y]:
                        dis[y] = new_d
                        heappush(h, (new_d, y))
                pow10 *= 10
        return -1
```

```java [sol-Java]
class Solution {
    private static final int MX = 10000;
    private static final boolean[] np = new boolean[MX];

    static {
        np[1] = true;
        // 埃氏筛，标记每个数是否为合数（或者 1）
        for (int i = 2; i < MX; i++) {
            if (!np[i]) {
                for (int j = i * i; j < MX; j += i) {
                    np[j] = true; // 合数
                }
            }
        }
    }

    public int minOperations(int n, int m) {
        if (!np[n] || !np[m]) {
            return -1;
        }
        int lenN = Integer.toString(n).length();
        int[] dis = new int[(int) Math.pow(10, lenN)];
        Arrays.fill(dis, Integer.MAX_VALUE);
        dis[n] = n;
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> (a[0] - b[0]));
        pq.offer(new int[]{n, n});
        while (!pq.isEmpty()) {
            int[] top = pq.poll();
            int disX = top[0];
            int x = top[1];
            if (x == m) {
                return disX;
            }
            if (disX > dis[x]) {
                continue;
            }
            int pow10 = 1;
            for (int v = x; v > 0; v /= 10) {
                int d = v % 10;
                if (d > 0) { // 减少
                    int y = x - pow10;
                    int newD = disX + y;
                    if (np[y] && newD < dis[y]) {
                        dis[y] = newD;
                        pq.offer(new int[]{newD, y});
                    }
                }
                if (d < 9) { // 增加
                    int y = x + pow10;
                    int newD = disX + y;
                    if (np[y] && newD < dis[y]) {
                        dis[y] = newD;
                        pq.offer(new int[]{newD, y});
                    }
                }
                pow10 *= 10;
            }
        }
        return -1;
    }
}
```

```cpp [sol-C++]
const int MX = 10000;
bool np[MX];

int init = [] {
    np[1] = true;
    // 埃氏筛，标记每个数是否为合数（或者 1）
    for (int i = 2; i < MX; i++) {
        if (!np[i]) {
            for (int j = i * i; j < MX; j += i) {
                np[j] = true; // 合数
            }
        }
    }
    return 0;
}();

class Solution {
public:
    int minOperations(int n, int m) {
        if (!np[n] || !np[m]) {
            return -1;
        }
        int len_n = to_string(n).length();
        vector<int> dis(pow(10, len_n), INT_MAX);
        dis[n] = n;
        priority_queue<pair<int, int>, vector<pair<int, int>>, greater<>> pq;
        pq.emplace(n, n);
        while (!pq.empty()) {
            auto [dix_x, x] = pq.top();
            pq.pop();
            if (x == m) {
                return dix_x;
            }
            if (dix_x > dis[x]) {
                continue;
            }
            int pow10 = 1;
            for (int v = x; v; v /= 10) {
                int d = v % 10;
                if (d > 0) { // 减少
                    int y = x - pow10;
                    int new_d = dix_x + y;
                    if (np[y] && new_d < dis[y]) {
                        dis[y] = new_d;
                        pq.emplace(new_d, y);
                    }
                }
                if (d < 9) { // 增加
                    int y = x + pow10;
                    int new_d = dix_x + y;
                    if (np[y] && new_d < dis[y]) {
                        dis[y] = new_d;
                        pq.emplace(new_d, y);
                    }
                }
                pow10 *= 10;
            }
        }
        return -1;
    }
};
```

```go [sol-Go]
const mx = 10000
var np = [mx]bool{1: true}

func init() {
	// 埃氏筛，标记每个数是否为合数（或者 1）
	for i := 2; i < mx; i++ {
		if !np[i] {
			for j := i * i; j < mx; j += i {
				np[j] = true // 合数
			}
		}
	}
}

func minOperations(n, m int) int {
	if !np[n] || !np[m] {
		return -1
	}
	lenN := len(strconv.Itoa(n))
	dis := make([]int, int(math.Pow10(lenN)))
	for i := range dis {
		dis[i] = math.MaxInt
	}
	dis[n] = n
	h := hp{{n, n}}
	for len(h) > 0 {
		top := heap.Pop(&h).(pair)
		x := top.x
		if x == m {
			return top.dis
		}
		disX := top.dis
		if disX > dis[x] {
			continue
		}
		pow10 := 1
		for v := x; v > 0; v /= 10 {
			d := v % 10
			if d > 0 { // 减少
				y := x - pow10
				newD := disX + y
				if np[y] && newD < dis[y] {
					dis[y] = newD
					heap.Push(&h, pair{newD, y})
				}
			}
			if d < 9 { // 增加
				y := x + pow10
				newD := disX + y
				if np[y] && newD < dis[y] {
					dis[y] = newD
					heap.Push(&h, pair{newD, y})
				}
			}
			pow10 *= 10
		}
	}
	return -1
}

type pair struct{ dis, x int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
```

#### 复杂度分析

忽略预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(M\log M)$，其中 $M=n\log n$。图中有 $\mathcal{O}(n)$ 个节点，$\mathcal{O}(M)$ 条边，每条边需要 $\mathcal{O}(\log M)$ 的堆操作。
- 空间复杂度：$\mathcal{O}(M)$。堆中有 $\mathcal{O}(M)$ 个元素。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. 【本题相关】[图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. 【本题相关】[数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
