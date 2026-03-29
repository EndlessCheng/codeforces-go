用一个堆维护所有活跃事件。

但是，如何**修改**堆中元素的优先级？

我们可以用一种「**懒更新**」的想法：

- 对于 $\texttt{updatePriority}$，不做更新，而是把一条新的 $(\textit{newPriority},\textit{eventId})$ 数据**插入**堆。同时，额外用一个哈希表记录每个 $\textit{eventId}$ 对应的**最新**优先级。
- 对于 $\texttt{pollHighest}$，如果堆顶优先级与哈希表中记录的优先级不一致，说明堆顶是个旧数据，弹出堆顶即可。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class EventManager:
    def __init__(self, events: List[List[int]]):
        self.id_to_priority = {}
        self.h = []
        for eid, priority in events:
            self.id_to_priority[eid] = priority
            self.h.append((-priority, eid))
        heapify(self.h)

    def updatePriority(self, eventId: int, newPriority: int) -> None:
        self.id_to_priority[eventId] = newPriority
        heappush(self.h, (-newPriority, eventId))

    def pollHighest(self) -> int:
        while self.h:
            priority, eid = heappop(self.h)
            if self.id_to_priority.get(eid, -1) == -priority:
                del self.id_to_priority[eid]
                return eid
            # else 货不对板，继续找下一个
        return -1
```

```java [sol-Java]
class EventManager {
    private final Map<Integer, Integer> idToPriority = new HashMap<>();
    private final PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) ->
            a[0] != b[0] ? b[0] - a[0] : a[1] - b[1]
    );

    public EventManager(int[][] events) {
        for (int[] e : events) {
            int id = e[0];
            int priority = e[1];
            idToPriority.put(id, priority);
            pq.offer(new int[]{priority, id});
        }
    }

    public void updatePriority(int eventId, int newPriority) {
        idToPriority.put(eventId, newPriority);
        pq.offer(new int[]{newPriority, eventId});
    }

    public int pollHighest() {
        while (!pq.isEmpty()) {
            int[] e = pq.poll();
            int priority = e[0];
            int id = e[1];
            if (idToPriority.getOrDefault(id, -1) == priority) {
                idToPriority.remove(id);
                return id;
            }
            // else 货不对板，继续找下一个
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class EventManager {
    unordered_map<int, int> id_to_priority;
    priority_queue<pair<int, int>> pq;

public:
    EventManager(vector<vector<int>>& events) {
        id_to_priority.reserve(events.size()); // 预分配空间
        for (auto& e : events) {
            int id = e[0], priority = e[1];
            id_to_priority[id] = priority;
            pq.emplace(priority, -id);
        }
    }

    void updatePriority(int eventId, int newPriority) {
        id_to_priority[eventId] = newPriority;
        pq.emplace(newPriority, -eventId);
    }

    int pollHighest() {
        while (!pq.empty()) {
            auto [priority, id] = pq.top();
            pq.pop();
            id = -id;
            if (id_to_priority[id] == priority) {
                id_to_priority[id] = -1;
                return id;
            }
            // else 货不对板，继续找下一个
        }
        return -1;
    }
};
```

```go [sol-Go]
type EventManager struct {
	idToPriority map[int]int
	h            *hp
}

func Constructor(events [][]int) EventManager {
	n := len(events)
	idToPriority := make(map[int]int, n) // 预分配空间
	h := make(hp, n)
	for i, e := range events {
		id, priority := e[0], e[1]
		idToPriority[id] = priority
		h[i] = event{priority, id}
	}
	heap.Init(&h)
	return EventManager{idToPriority, &h}
}

func (m EventManager) UpdatePriority(eventId, newPriority int) {
	m.idToPriority[eventId] = newPriority
	heap.Push(m.h, event{newPriority, eventId})
}

func (m EventManager) PollHighest() int {
	for m.h.Len() > 0 {
		e := heap.Pop(m.h).(event)
		if m.idToPriority[e.id] == e.priority {
			delete(m.idToPriority, e.id)
			return e.id
		}
		// else 货不对板，继续找下一个
	}
	return -1
}

type event struct{ priority, id int }
type hp []event
func (h hp) Len() int      { return len(h) }
func (h hp) Less(i, j int) bool {
	return h[i].priority > h[j].priority || h[i].priority == h[j].priority && h[i].id < h[j].id
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)   { *h = append(*h, v.(event)) }
func (h *hp) Pop() any     { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```

#### 复杂度分析

- 时间复杂度：
   - 初始化：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{events}$ 的长度。
   - $\texttt{updatePriority}$：$\mathcal{O}(\log (n+q))$，其中 $q$ 是 $\textit{updatePriority}$ 的调用次数。
   - $\texttt{pollHighest}$：均摊 $\mathcal{O}(\log (n+q))$。
- 空间复杂度：$\mathcal{O}(n+q)$。

## 专题训练

见下面数据结构题单的「**§5.6 懒删除堆**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)
