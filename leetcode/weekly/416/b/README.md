## 方法一：最小堆模拟

循环 $\textit{mountainHeight}$ 次，每次选一个「工作后总用时」最短的工人，把山的高度降低 $1$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1WRtDejEjD/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minNumberOfSeconds(self, mountainHeight: int, workerTimes: List[int]) -> int:
        h = [(t, t, t) for t in workerTimes]
        heapify(h)
        for _ in range(mountainHeight):
            # 工作后总用时，当前工作（山高度降低 1）用时，workerTimes[i]
            nxt, delta, base = h[0]
            heapreplace(h, (nxt + delta + base, delta + base, base))
        return nxt  # 最后一个出堆的 nxt 即为答案
```

```java [sol-Java]
class Solution {
    public long minNumberOfSeconds(int mountainHeight, int[] workerTimes) {
        PriorityQueue<long[]> pq = new PriorityQueue<>((a, b) -> Long.compare(a[0], b[0]));
        for (int t : workerTimes) {
            pq.offer(new long[]{t, t, t});
        }
        long ans = 0;
        while (mountainHeight-- > 0) {
            // 工作后总用时，当前工作（山高度降低 1）用时，workerTimes[i]
            long[] w = pq.poll();
            long nxt = w[0], delta = w[1], base = w[2];
            ans = nxt; // 最后一个出堆的 nxt 即为答案
            pq.offer(new long[]{nxt + delta + base, delta + base, base});
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minNumberOfSeconds(int mountainHeight, vector<int>& workerTimes) {
        priority_queue<tuple<long long, long long, int>, vector<tuple<long long, long long, int>>, greater<>> pq;
        for (int t : workerTimes) {
            pq.emplace(t, t, t);
        }
        long long ans = 0;
        while (mountainHeight--) {
            // 工作后总用时，当前工作（山高度降低 1）用时，workerTimes[i]
            auto [nxt, delta, base] = pq.top(); pq.pop();
            ans = nxt; // 最后一个出堆的 nxt 即为答案
            pq.emplace(nxt + delta + base, delta + base, base);
        }
        return ans;
    }
};
```

```go [sol-Go]
func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	h := make(hp, len(workerTimes))
	for i, t := range workerTimes {
		h[i] = worker{t, t, t}
	}
	heap.Init(&h)

	ans := 0
	for ; mountainHeight > 0; mountainHeight-- {
		ans = h[0].nxt // 最后一个出堆的 nxt 即为答案
		h[0].delta += h[0].base
		h[0].nxt += h[0].delta
		heap.Fix(&h, 0)
	}
	return int64(ans)
}

// 工作后总用时，当前工作（山高度降低 1）用时，workerTimes[i]
type worker struct{ nxt, delta, base int }
type hp []worker
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].nxt < h[j].nxt }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\textit{mountainHeight}\log n)$，其中 $n$ 是 $\textit{workerTimes}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：二分答案

由于花的时间越多，能够降低的高度也越多，所以有**单调性**，可以二分答案。

问题变成：

- 每个工人**至多**花费 $m$ 秒，总共降低的高度是多少？能否大于等于 $\textit{mountainHeight}$？

遍历 $\textit{workerTimes}$，设 $t=\textit{workerTimes}[i]$，那么有

$$
t + 2t+ \cdots + xt = t\cdot \dfrac{x(x+1)}{2} \le m
$$

即

$$
\dfrac{x(x+1)}{2} \le \left\lfloor\dfrac{m}{t}\right\rfloor = k
$$

解得

$$
x \le \dfrac{-1 + \sqrt{1 + 8k}}{2}
$$

所以第 $i$ 名工人可以把山的高度降低

$$
\left\lfloor \dfrac{-1 + \sqrt{1 + 8k}}{2} \right\rfloor = \left\lfloor \dfrac{-1 + \lfloor\sqrt{1 + 8k}\rfloor}{2} \right\rfloor
$$

累加上式，如果和 $\ge \textit{mountainHeight}$，则说明答案 $\le m$，否则说明答案 $> m$。

最后，讨论二分的上下界。这里用开区间二分，其他二分写法也是可以的。

- 开区间二分下界：$0$，无法把山的高度降低到 $0$。
- 开区间二分上界：设 $\textit{maxT}$ 为 $\textit{workerTimes}$ 的最大值，假设每个工人都是最慢的 $\textit{maxT}$，那么单个工人要把山降低 $h=\left\lceil\dfrac{mountainHeight}{n}\right\rceil$，耗时 $\textit{maxT}\cdot(1+2+\cdots+h)=\textit{maxT}\cdot\dfrac{h(h+1)}{2}$，将其作为开区间的二分上界，一定可以把山的高度降低到 $\le 0$。

代码实现时，可以用等式

$$
\left\lceil\dfrac{a}{b}\right\rceil = \left\lfloor\dfrac{a-1}{b}\right\rfloor + 1
$$

计算上取整。讨论 $a$ 被 $b$ 整除，和不被 $b$ 整除两种情况，可以证明上式的正确性。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1WRtDejEjD/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minNumberOfSeconds(self, mountainHeight: int, workerTimes: List[int]) -> int:
        def check(m: int) -> bool:
            left_h = mountainHeight
            for t in workerTimes:
                left_h -= (isqrt(m // t * 8 + 1) - 1) // 2
                if left_h <= 0:
                    return True
            return False

        max_t = max(workerTimes)
        h = (mountainHeight - 1) // len(workerTimes) + 1
        return bisect_left(range(max_t * h * (h + 1) // 2), True, 1, key=check)
```

```py [sol-Python3 写法二]
class Solution:
    def minNumberOfSeconds(self, mountainHeight: int, workerTimes: List[int]) -> int:
        f = lambda m: sum((isqrt(m // t * 8 + 1) - 1) // 2 for t in workerTimes)
        max_t = max(workerTimes)
        h = (mountainHeight - 1) // len(workerTimes) + 1
        return bisect_left(range(max_t * h * (h + 1) // 2), mountainHeight, 1, key=f)
```

```java [sol-Java]
class Solution {
    public long minNumberOfSeconds(int mountainHeight, int[] workerTimes) {
        int maxT = 0;
        for (int t : workerTimes) {
            maxT = Math.max(maxT, t);
        }
        int h = (mountainHeight - 1) / workerTimes.length + 1;
        long left = 0;
        long right = (long) maxT * h * (h + 1) / 2;
        while (left + 1 < right) {
            long mid = (left + right) / 2;
            if (check(mid, mountainHeight, workerTimes)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    private boolean check(long m, int leftH, int[] workerTimes) {
        for (int t : workerTimes) {
            leftH -= ((int) Math.sqrt(m / t * 8 + 1) - 1) / 2;
            if (leftH <= 0) {
                return true;
            }
        }
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minNumberOfSeconds(int mountainHeight, vector<int>& workerTimes) {
        auto check = [&](long long m) {
            int left_h = mountainHeight;
            for (int t : workerTimes) {
                left_h -= ((int) sqrt(m / t * 8 + 1) - 1) / 2;
                if (left_h <= 0) {
                    return true;
                }
            }
            return false;
        };

        int max_t = ranges::max(workerTimes);
        int h = (mountainHeight - 1) / workerTimes.size() + 1;
        long long left = 0, right = (long long) max_t * h * (h + 1) / 2;
        while (left + 1 < right) {
            long long mid = (left + right) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol-Go]
func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	maxT := slices.Max(workerTimes)
	h := (mountainHeight-1)/len(workerTimes) + 1
	ans := 1 + sort.Search(maxT*h*(h+1)/2-1, func(m int) bool {
		m++
		leftH := mountainHeight
		for _, t := range workerTimes {
			leftH -= (int(math.Sqrt(float64(m/t*8+1))) - 1) / 2
			if leftH <= 0 {
				return true
			}
		}
		return false
	})
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{workerTimes}$ 的长度，$U\le 5\cdot 10^{10}(10^5+1)$ 是二分上界。二分 $\mathcal{O}(\log U)$ 次，每次 $\mathcal{O}(n)$ 时间。开平方有专门的 CPU 指令，可以视作 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

更多相似题目，见下面二分题单中的「**二分答案：求最小**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
