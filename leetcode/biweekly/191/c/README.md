## 方法一：完全背包

本题和 [279. 完全平方数](https://leetcode.cn/problems/perfect-squares/) 几乎一样，推荐先把那题做了，并阅读 [我的题解](https://leetcode.cn/problems/perfect-squares/solutions/2830762/dong-tai-gui-hua-cong-ji-yi-hua-sou-suo-3kz1g/)。

本题与 279 的区别：

1. 完全平方数 $x^2$ 改成了三角形数 $1+2+\cdots+x = \dfrac{x(x+1)}{2}$。
2. 每使用一个三角形数 $\dfrac{x(x+1)}{2}$，要花费 $x+1$ 天。特殊情况：使用第一个三角形数花费 $x$ 天。

代码实现时，我们可以初始化 $f[0] = -1$，这样可以统一为使用一个三角形数花费 $x+1$ 天，无论是否为第一个三角形数。

```py [sol-Python3]
MX = 100_001
f = [inf] * MX
f[0] = -1

s = i = 1
while s < MX:
    for j in range(s, MX):
        f[j] = min(f[j], f[j - s] + i + 1)
    i += 1
    s += i

class Solution:
    def minDays(self, n: int) -> int:
        return f[n]
```

```java [sol-Java]
class Solution {
    private static final int MX = 100_001;
    private static final int[] f = new int[MX];
    private static boolean initialized = false;

    // 这样写比 static block 快
    public Solution() {
        if (initialized) {
            return;
        }
        initialized = true;

        Arrays.fill(f, Integer.MAX_VALUE);
        f[0] = -1;

        int sum = 1;
        int i = 1;
        while (sum < MX) {
            for (int j = sum; j < MX; j++) {
                f[j] = Math.min(f[j], f[j - sum] + i + 1);
            }
            i++;
            sum += i;
        }
    }

    public int minDays(int n) {
        return f[n];
    }
}
```

```cpp [sol-C++]
constexpr int MX = 100'001;
int f[MX];

auto init = [] {
    ranges::fill(f, INT_MAX);
    f[0] = -1;

    int sum = 1, i = 1;
    while (sum < MX) {
        for (int j = sum; j < MX; j++) {
            f[j] = min(f[j], f[j - sum] + i + 1);
        }
        i++;
        sum += i;
    }
    return 0;
}();

class Solution {
public:
    int minDays(int n) {
        return f[n];
    }
};
```

```go [sol-Go]
const mx = 100_001
var f [mx]int

func init() {
	f[0] = -1
	for i := 1; i < mx; i++ {
		f[i] = math.MaxInt
	}

	sum, i := 1, 1
	for sum < mx {
		for j := sum; j < mx; j++ {
			f[j] = min(f[j], f[j-sum]+i+1)
		}
		i++
		sum += i
	}
}

func minDays(n int) int {
	return f[n]
}
```

#### 复杂度分析

- 预处理的时间复杂度：$\mathcal{O}(N\sqrt{N})$，其中 $N = 10^5$。
- 预处理的空间复杂度：$\mathcal{O}(N)$。

## 方法二：最短路

> 求最小的完全背包可以转化成计算有向图的最短路长度。

把元素和 $s$ 看成节点，对于每个三角形数 $t = \dfrac{x(x+1)}{2}$，添加一条从 $s$ 到 $s+t$ 的有向边，边权为 $x+1$。

答案为 $0$ 到 $n$ 的最短路长度。

这可以用 **Dijkstra 算法**解决。[Dijkstra 算法介绍](https://leetcode.cn/problems/network-delay-time/solution/liang-chong-dijkstra-xie-fa-fu-ti-dan-py-ooe8/)。

```py [sol-Python3]
MX = 100_001
dis = [inf] * MX
dis[0] = -1
h = [(-1, 0)]
while h:
    d, s = heappop(h)
    if d > dis[s]:
        continue
    t, i = 1, 1
    while s + t < MX:
        w = s + t
        new_d = d + i + 1
        if new_d < dis[w]:
            dis[w] = new_d
            heappush(h, (new_d, w))
        i += 1
        t += i


class Solution:
    def minDays(self, n: int) -> int:
        return dis[n]
```

```java [sol-Java]
class Solution {
    private static final int MX = 100_001;
    private static final int[] dis = new int[MX];
    private static boolean initialized = false;

    // 这样写比 static block 快
    public Solution() {
        if (initialized) {
            return;
        }
        initialized = true;

        Arrays.fill(dis, Integer.MAX_VALUE);
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]);
        dis[0] = -1;
        pq.offer(new int[]{-1, 0});
        while (!pq.isEmpty()) {
            int[] top = pq.poll();
            int d = top[0];
            int s = top[1];
            if (d > dis[s]) {
                continue;
            }
            int t = 1;
            int i = 1;
            while (s + t < MX) {
                int w = s + t;
                int newD = d + i + 1;
                if (newD < dis[w]) {
                    dis[w] = newD;
                    pq.offer(new int[]{newD, w});
                }
                i++;
                t += i;
            }
        }
    }

    public int minDays(int n) {
        return dis[n];
    }
}
```

```cpp [sol-C++]
constexpr int MX = 100'001;
int dis[MX];

auto init = [] {
    ranges::fill(dis, INT_MAX);
    priority_queue<pair<int, int>, vector<pair<int, int>>, greater<>> pq;
    dis[0] = -1;
    pq.emplace(-1, 0);
    while (!pq.empty()) {
        auto [d, s] = pq.top();
        pq.pop();
        if (d > dis[s]) {
            continue;
        }
        int t = 1, i = 1;
        while (s + t < MX) {
            int w = s + t;
            int new_d = d + i + 1;
            if (new_d < dis[w]) {
                dis[w] = new_d;
                pq.emplace(new_d, w);
            }
            i++;
            t += i;
        }
    }
    return 0;
}();

class Solution {
public:
    int minDays(int n) {
        return dis[n];
    }
};
```

```go [sol-Go]
const mx = 100_001
var dis [mx]int

func init() {
	for i := range dis {
		dis[i] = math.MaxInt
	}
	dis[0] = -1
	h := &hp{{-1, 0}}
	for h.Len() > 0 {
		top := heap.Pop(h).(pair)
		d, s := top.dis, top.v
		if d > dis[s] {
			continue
		}
		t, i := 1, 1
		for s+t < mx {
			w := s + t
			newD := d + i + 1
			if newD < dis[w] {
				dis[w] = newD
				heap.Push(h, pair{newD, w})
			}
			i++
			t += i
		}
	}
}

func minDays(n int) int {
	return dis[n]
}

type pair struct{ dis, v int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
```

#### 复杂度分析

- 预处理的时间复杂度：$\mathcal{O}(N\sqrt{N}\log N)$，其中 $N = 10^5$。
- 预处理的空间复杂度：$\mathcal{O}(N\sqrt{N})$。

## 专题训练

1. 动态规划题单的「**§3.2 完全背包**」。
2. 图论题单的「**§3.1 单源最短路：Dijkstra 算法**」。

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
