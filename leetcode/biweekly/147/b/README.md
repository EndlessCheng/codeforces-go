用**最大堆**维护优先级、任务编号、用户编号。

同时额外用一个哈希表记录每个任务编号对应的**最新的**优先级和用户编号。

- 对于 $\texttt{edit}$，直接把一个新的优先级、任务编号、用户编号三元组入堆。同时更新哈希表。
- 对于 $\texttt{rmv}$，直接把元素从哈希表删掉（或者优先级改成 $-1$）。
- 对于 $\texttt{execTop}$，不断弹出「货不对板」的堆顶，也就是任务编号和哈希表中记录的数据不一致的堆顶。直到找到一致的数据或者堆为空。

> 注 1：如果你学过 Dijkstra 算法，其中的堆就是懒更新堆。
>
> 注 2：题目保证输入的 $\textit{tasks}$ 数组中没有重复的 $\textit{taskId}$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1SzrAYMESJ/?t=2m46s)，欢迎点赞关注~

```py [sol-Python3]
class TaskManager:
    def __init__(self, tasks: List[List[int]]):
        self.mp = {taskId: (priority, userId) for userId, taskId, priority in tasks}
        self.h = [(-priority, -taskId, userId) for userId, taskId, priority in tasks]  # 取相反数，变成最大堆
        heapify(self.h)

    def add(self, userId: int, taskId: int, priority: int) -> None:
        self.mp[taskId] = (priority, userId)
        heappush(self.h, (-priority, -taskId, userId))

    def edit(self, taskId: int, newPriority: int) -> None:
        # 懒修改
        self.add(self.mp[taskId][1], taskId, newPriority)

    def rmv(self, taskId: int) -> None:
        # 懒删除
        self.mp[taskId] = (-1, -1)

    def execTop(self) -> int:
        while self.h:
            priority, taskId, userId = heappop(self.h)
            # 如果货不对板，堆顶和 mp 中记录的不一样，说明这个数据已被修改/删除，不做处理
            if self.mp[-taskId] == (-priority, userId):
                self.rmv(-taskId)
                return userId
        return -1
```

```java [sol-Java]
class TaskManager {
    private final Map<Integer, int[]> mp = new HashMap<>(); // taskId -> (priority, userId)
    private final PriorityQueue<int[]> pq =
            new PriorityQueue<>((a, b) -> a[0] != b[0] ? b[0] - a[0] : b[1] - a[1]); // (priority, taskId, userId)

    public TaskManager(List<List<Integer>> tasks) {
        for (List<Integer> task : tasks) {
            add(task.get(0), task.get(1), task.get(2));
        }
    }

    public void add(int userId, int taskId, int priority) {
        mp.put(taskId, new int[]{priority, userId});
        pq.offer(new int[]{priority, taskId, userId});
    }

    public void edit(int taskId, int newPriority) {
        // 懒修改
        add(mp.get(taskId)[1], taskId, newPriority);
    }

    public void rmv(int taskId) {
        // 懒删除
        mp.get(taskId)[0] = -1;
    }

    public int execTop() {
        while (!pq.isEmpty()) {
            int[] top = pq.poll();
            int priority = top[0], taskId = top[1], userId = top[2];
            int[] p = mp.get(taskId);
            // 如果货不对板，堆顶和 mp 中记录的不一样，说明这个数据已被修改/删除，不做处理
            if (p[0] == priority && p[1] == userId) {
                rmv(taskId);
                return userId;
            }
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class TaskManager {
    priority_queue<tuple<int, int, int>> pq; // (priority, taskId, userId)
    unordered_map<int, pair<int, int>> mp; // taskId -> (priority, userId)

public:
    TaskManager(vector<vector<int>>& tasks) {
        for (auto& task : tasks) {
            add(task[0], task[1], task[2]);
        }
    }

    void add(int userId, int taskId, int priority) {
        mp[taskId] = {priority, userId};
        pq.emplace(priority, taskId, userId);
    }

    void edit(int taskId, int newPriority) {
        // 懒修改
        add(mp[taskId].second, taskId, newPriority);
    }

    void rmv(int taskId) {
        // 懒删除
        mp[taskId].first = -1;
    }

    int execTop() {
        while (!pq.empty()) {
            auto [priority, taskId, userId] = pq.top();
            pq.pop();
            // 如果货不对板，堆顶和 mp 中记录的不一样，说明这个数据已被修改/删除，不做处理
            if (mp[taskId] == pair(priority, userId)) {
                rmv(taskId);
                return userId;
            }
        }
        return -1;
    }
};
```

```go [sol-Go]
type pair struct{ priority, userId int }

type TaskManager struct {
	h  *hp          // (priority, taskId, userId)
	mp map[int]pair // taskId -> (priority, userId)
}

func Constructor(tasks [][]int) TaskManager {
	h := hp{}
	mp := map[int]pair{}
	for _, task := range tasks {
		userId, taskId, priority := task[0], task[1], task[2]
		mp[taskId] = pair{priority, userId}
		h = append(h, tuple{priority, taskId, userId})
	}
	heap.Init(&h)
	return TaskManager{&h, mp}
}

func (tm *TaskManager) Add(userId, taskId, priority int) {
	tm.mp[taskId] = pair{priority, userId}
	heap.Push(tm.h, tuple{priority, taskId, userId})
}

func (tm *TaskManager) Edit(taskId, newPriority int) {
	// 懒修改
	tm.Add(tm.mp[taskId].userId, taskId, newPriority)
}

func (tm *TaskManager) Rmv(taskId int) {
	// 懒删除
	tm.mp[taskId] = pair{-1, -1}
}

func (tm *TaskManager) ExecTop() int {
	for tm.h.Len() > 0 {
		top := heap.Pop(tm.h).(tuple)
		priority, taskId, userId := top.priority, top.taskId, top.userId
		// 如果货不对板，堆顶和 mp 中记录的不一样，说明这个数据已被修改/删除，不做处理
		if tm.mp[taskId] == (pair{priority, userId}) {
			tm.Rmv(taskId)
			return userId
		}
	}
	return -1
}

type tuple struct{ priority, taskId, userId int }
type hp []tuple
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return cmp.Or(h[i].priority-h[j].priority, h[i].taskId-h[j].taskId) > 0 }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```

#### 复杂度分析

- 时间复杂度：
  - 初始化：$\mathcal{O}(n)$ 或者 $\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{tasks}$ 的长度。Python 和 Go 使用了堆化，复杂度是 $\mathcal{O}(n)$ 的。
  - $\texttt{add}$ 和 $\texttt{edit}$：$\mathcal{O}(\log (n+q))$，其中 $q$ 是 $\texttt{add}$ 和 $\texttt{edit}$ 的调用次数之和。
  - $\texttt{rmv}$：$\mathcal{O}(1)$。
  - $\texttt{execTop}$：均摊 $\mathcal{O}(\log (n+q))$。每个元素至多入堆出堆各一次。
- 空间复杂度：$\mathcal{O}(n+q)$。

更多相似题目，见下面数据结构题单中的「**§5.6 懒删除堆**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. 【本题相关】[常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
