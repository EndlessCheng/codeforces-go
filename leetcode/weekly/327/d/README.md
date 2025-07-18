复杂模拟题，推荐观看 [视频讲解](https://www.bilibili.com/video/BV1KG4y1j73o/?t=16m25s)，详细介绍了本题的思考过程。

---

建立 $4$ 个堆，每个堆都记录工人下标 $i$ 和完成时间 $t$（到达桥的时间），这 $4$ 个堆**从左到右**分别表示:

1. $\textit{workL}$：新仓库正在放箱的工人；
2. $\textit{waitL}$：左边等待过桥的工人；
3. $\textit{waitR}$：右边等待过桥的工人；
4. $\textit{workR}$：旧仓库正在搬箱的工人。

记录当前时间 $\textit{cur}$，不断循环直到所有箱子被搬完，每次循环：

1. 把完成时间不超过 $\textit{cur}$ 的 $\textit{workL}$ 弹出，放入 $\textit{waitL}$ 中；
2. 把完成时间不超过 $\textit{cur}$ 的 $\textit{workR}$ 弹出，放入 $\textit{waitR}$ 中；
3. 如果 $\textit{waitR}$ 不为空，出堆，过桥，更新 $\textit{cur}$ 为过完桥的时间，然后把这个工人放入 $\textit{workL}$ 中（记录完成时间）；
4. 否则如果 $\textit{waitL}$ 不为空，出堆，过桥，更新 $\textit{cur}$ 为过完桥的时间，然后把这个工人放入 $\textit{workR}$ 中（记录完成时间），同时把 $n$ 减一；
5. 否则说明 $\textit{cur}$ 过小，找个最小的放箱/搬箱完成时间来更新 $\textit{cur}$。

循环结束后，不断弹出 $\textit{workR}$，过桥，最后一个工人过完桥的时间即为答案。

代码实现时，可以先把 $\textit{time}$ 从小到大**稳定排序**，这样下标越大的工人效率越低，只看下标就能比较工人的效率。

```py [sol-Python3]
class Solution:
    def findCrossingTime(self, n: int, k: int, time: List[List[int]]) -> int:
        time.sort(key=lambda t: t[0] + t[2])  # 稳定排序
        cur = 0
        workL, waitL, waitR, workR = [], [[-i, 0] for i in range(k - 1, -1, -1)], [], []  # 下标越大效率越低
        while n:
            while workL and workL[0][0] <= cur:
                p = heappop(workL)
                p[0], p[1] = p[1], p[0]
                heappush(waitL, p)  # 左边完成放箱
            while workR and workR[0][0] <= cur:
                p = heappop(workR)
                p[0], p[1] = p[1], p[0]
                heappush(waitR, p)  # 右边完成搬箱
            if waitR:  # 右边过桥，注意加到 waitR 中的都是 <= cur 的（下同）
                p = heappop(waitR)
                cur += time[-p[0]][2]
                p[1] = p[0]
                p[0] = cur + time[-p[0]][3]
                heappush(workL, p)  # 放箱
            elif waitL:  # 左边过桥
                p = heappop(waitL)
                cur += time[-p[0]][0]
                p[1] = p[0]
                p[0] = cur + time[-p[0]][1]
                heappush(workR, p)  # 搬箱
                n -= 1
            elif len(workL) == 0: cur = workR[0][0]  # cur 过小，找个最小的放箱/搬箱完成时间来更新 cur
            elif len(workR) == 0: cur = workL[0][0]
            else: cur = min(workL[0][0], workR[0][0])
        while workR:
            t, i = heappop(workR)  # 右边完成搬箱
            # 如果没有排队，直接过桥；否则由于无论谁先过桥，最终完成时间都一样，所以也可以直接计算
            cur = max(t, cur) + time[-i][2]
        return cur  # 最后一个过桥的时间
```

```java [sol-Java]
class Solution {
    public int findCrossingTime(int n, int k, int[][] time) {
        Arrays.sort(time, (a, b) -> a[0] + a[2] - b[0] - b[2]); // 稳定排序
        
        var workL = new PriorityQueue<int[]>((a, b) -> a[1] - b[1]);
        var workR = new PriorityQueue<int[]>(workL.comparator());
        var waitL = new PriorityQueue<int[]>((a, b) -> b[0] - a[0]); // 下标越大效率越低
        var waitR = new PriorityQueue<int[]>(waitL.comparator());
        for (int i = k - 1; i >= 0; --i) 
            waitL.add(new int[]{i, 0});
        
        int cur = 0;
        while (n > 0) {
            while (!workL.isEmpty() && workL.peek()[1] <= cur) waitL.add(workL.poll()); // 左边完成放箱
            while (!workR.isEmpty() && workR.peek()[1] <= cur) waitR.add(workR.poll()); // 右边完成搬箱
            if (!waitR.isEmpty()) { // 右边过桥，注意加到 waitR 中的都是 <= cur 的（下同）
                var p = waitR.poll();
                cur += time[p[0]][2];
                p[1] = cur + time[p[0]][3];
                workL.add(p); // 放箱
            } else if (!waitL.isEmpty()) { // 左边过桥
                var p = waitL.poll();
                cur += time[p[0]][0];
                p[1] = cur + time[p[0]][1];
                workR.add(p); // 搬箱
                --n;
            } else if (workL.isEmpty()) cur = workR.peek()[1]; // cur 过小，找个最小的放箱/搬箱完成时间来更新 cur
            else if (workR.isEmpty()) cur = workL.peek()[1];
            else cur = Math.min(workL.peek()[1], workR.peek()[1]);
        }
        while (!workR.isEmpty()) {
            var p = workR.poll(); // 右边完成搬箱
            // 如果没有排队，直接过桥；否则由于无论谁先过桥，最终完成时间都一样，所以也可以直接计算
            cur = Math.max(p[1], cur) + time[p[0]][2];
        }
        return cur; // 最后一个过桥的时间
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findCrossingTime(int n, int k, vector<vector<int>> &time) {
        stable_sort(time.begin(), time.end(), [](auto &a, auto &b) {
            return a[0] + a[2] < b[0] + b[2];
        });
        
        priority_queue<pair<int, int>> waitL, waitR;
        priority_queue<pair<int, int>, vector<pair<int, int>>, greater<>> workL, workR;
        for (int i = k - 1; i >= 0; --i) 
            waitL.emplace(i, 0); // 下标越大效率越低
        
        int cur = 0;
        while (n) {
            while (!workL.empty() && workL.top().first <= cur) {
                auto[t, i] = workL.top();
                workL.pop();
                waitL.emplace(i, t); // 左边完成放箱
            }
            while (!workR.empty() && workR.top().first <= cur) {
                auto[t, i] = workR.top();
                workR.pop();
                waitR.emplace(i, t); // 右边完成搬箱
            }
            if (!waitR.empty()) { // 右边过桥，注意加到 waitR 中的都是 <= cur 的（下同）
                auto[i, t] = waitR.top();
                waitR.pop();
                cur += time[i][2];
                workL.emplace(cur + time[i][3], i); // 放箱
            } else if (!waitL.empty()) { // 左边过桥
                auto[i, t] = waitL.top();
                waitL.pop();
                cur += time[i][0];
                workR.emplace(cur + time[i][1], i); // 搬箱
                --n;
            } else if (workL.empty()) cur = workR.top().first; // cur 过小，找个最小的放箱/搬箱完成时间来更新 cur
            else if (workR.empty()) cur = workL.top().first;
            else cur = min(workL.top().first, workR.top().first);
        }
        while (!workR.empty()) {
            auto[t, i] = workR.top(); // 右边完成搬箱
            workR.pop();
            // 如果没有排队，直接过桥；否则由于无论谁先过桥，最终完成时间都一样，所以也可以直接计算
            cur = max(t, cur) + time[i][2];
        }
        return cur; // 最后一个过桥的时间
    }
};
```

```go [sol-Go]
func findCrossingTime(n, k int, time [][]int) (cur int) {
	sort.SliceStable(time, func(i, j int) bool {
		a, b := time[i], time[j]
		return a[0]+a[2] < b[0]+b[2]
	})

	waitL, waitR := make(hp, k), hp{}
	for i := range waitL {
		waitL[i].i = k - 1 - i // 下标越大效率越低
	}
	workL, workR := hp2{}, hp2{}

	for n > 0 {
		for len(workL) > 0 && workL[0].t <= cur {
			heap.Push(&waitL, heap.Pop(&workL)) // 左边完成放箱
		}
		for len(workR) > 0 && workR[0].t <= cur {
			heap.Push(&waitR, heap.Pop(&workR)) // 右边完成搬箱
		}
		if len(waitR) > 0 { // 右边过桥，注意加到 waitR 中的都是 <= cur 的（下同）
			p := heap.Pop(&waitR).(pair)
			cur += time[p.i][2]
			heap.Push(&workL, pair{p.i, cur + time[p.i][3]}) // 放箱，记录完成时间
		} else if len(waitL) > 0 { // 左边过桥
			p := heap.Pop(&waitL).(pair)
			cur += time[p.i][0]
			heap.Push(&workR, pair{p.i, cur + time[p.i][1]}) // 搬箱，记录完成时间
			n--
		} else if len(workL) == 0 { // cur 过小，找个最小的放箱/搬箱完成时间来更新 cur
			cur = workR[0].t
		} else if len(workR) == 0 {
			cur = workL[0].t
		} else {
			cur = min(workL[0].t, workR[0].t)
		}
	}
	for len(workR) > 0 {
		p := heap.Pop(&workR).(pair) // 右边完成搬箱
		// 如果没有排队，直接过桥；否则由于无论谁先过桥，最终完成时间都一样，所以也可以直接计算
		cur = max(p.t, cur) + time[p.i][2]
	}
	return cur // 最后一个过桥的时间
}

type pair struct{ i, t int }
type hp []pair
func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].i > h[j].i }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)         { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any           { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

type hp2 []pair
func (h hp2) Len() int            { return len(h) }
func (h hp2) Less(i, j int) bool  { return h[i].t < h[j].t }
func (h hp2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v any)         { *h = append(*h, v.(pair)) }
func (h *hp2) Pop() any           { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```

```js [sol-JavaScript]
var findCrossingTime = function(n, k, time) {
    time.sort((a, b) => (a[0] + a[2]) - (b[0] + b[2])); // 稳定排序

    const workL = new MinPriorityQueue(e => e[1]);
    const workR = new MinPriorityQueue(e => e[1]);
    const waitL = new MaxPriorityQueue(e => e[0]);
    const waitR = new MaxPriorityQueue(e => e[0]);
    for (let i = k - 1; i >= 0; i--) {
        waitL.enqueue([i, 0]); // 下标越大效率越低
    }

    let cur = 0;
    while (n) {
        while (!workL.isEmpty() && workL.front()[1] <= cur) {
            waitL.enqueue(workL.dequeue()); // 左边完成放箱
        }
        while (!workR.isEmpty() && workR.front()[1] <= cur) {
            waitR.enqueue(workR.dequeue()); // 右边完成搬箱
        }
        if (!waitR.isEmpty()) { // 右边过桥，注意加到 waitR 中的都是 <= cur 的（下同）
            const p = waitR.dequeue();
            cur += time[p[0]][2];
            p[1] = cur + time[p[0]][3];
            workL.enqueue(p); // 放箱
        } else if (!waitL.isEmpty()) { // 左边过桥
            const p = waitL.dequeue();
            cur += time[p[0]][0];
            p[1] = cur + time[p[0]][1];
            workR.enqueue(p); // 搬箱
            n--;
        } else if (workL.isEmpty()) { // cur 过小，下面找个最小的放箱/搬箱完成时间来更新 cur
            cur = workR.front()[1];
        } else if (workR.isEmpty()) {
            cur = workL.front()[1];
        } else {
            cur = Math.min(workL.front()[1], workR.front()[1]);
        }
    }

    while (!workR.isEmpty()) {
        const [i, t] = workR.dequeue(); // 右边完成搬箱
        // 如果没有排队，直接过桥；否则由于无论谁先过桥，最终完成时间都一样，所以也可以直接计算
        cur = Math.max(t, cur) + time[i][2];
    }
    return cur; // 最后一个过桥的时间
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log k)$。
- 空间复杂度：$\mathcal{O}(k)$。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
