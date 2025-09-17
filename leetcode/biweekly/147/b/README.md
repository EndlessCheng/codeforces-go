为了实现 $\texttt{execTop}$，用**最大堆**维护优先级、任务编号、用户编号。

为了实现堆的懒更新和懒删除，用一个哈希表记录每个任务编号对应的**最新的**优先级和用户编号。哈希表的 key 是任务编号，value 是优先级和用户编号组成的二元组。

- $\texttt{edit}$：直接把一个新的优先级、任务编号、用户编号三元组入堆。同时更新哈希表。
- $\texttt{rmv}$：把元素从哈希表中删掉。更简单的写法是，把优先级改成 $-1$。
- $\texttt{execTop}$：不断弹出「货不对板」的堆顶，也就是任务编号和哈希表中记录的数据不一致的堆顶。直到堆为空或者找到一致的数据。

> 注 1：如果你学过 Dijkstra 算法，其中的堆就是懒更新堆。
>
> 注 2：题目保证输入的 $\textit{tasks}$ 数组中没有重复的 $\textit{taskId}$。

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
            if self.mp[-taskId] == (-priority, userId):
                self.rmv(-taskId)
                return userId
            # else 货不对板，堆顶和 mp 中记录的不一样，说明堆顶数据已被修改或删除，不做处理
        return -1
```

```java [sol-Java]
class TaskManager {
    // taskId -> (priority, userId)
    private final Map<Integer, int[]> mp = new HashMap<>();

    // (priority, taskId, userId)
    private final PriorityQueue<int[]> pq =
            new PriorityQueue<>((a, b) -> a[0] != b[0] ? b[0] - a[0] : b[1] - a[1]);

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
        int userId = mp.get(taskId)[1];
        add(userId, taskId, newPriority);
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
            if (p[0] == priority && p[1] == userId) {
                rmv(taskId);
                return userId;
            }
            // else 货不对板，堆顶和 mp 中记录的不一样，说明堆顶数据已被修改或删除，不做处理
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class TaskManager {
    unordered_map<int, pair<int, int>> mp; // taskId -> (priority, userId)
    priority_queue<tuple<int, int, int>> pq; // (priority, taskId, userId)

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
            if (mp[taskId] == pair(priority, userId)) {
                rmv(taskId);
                return userId;
            }
            // else 货不对板，堆顶和 mp 中记录的不一样，说明堆顶数据已被修改或删除，不做处理
        }
        return -1;
    }
};
```

```go [sol-Go]
type pair struct{ priority, userId int }

type TaskManager struct {
	mp map[int]pair // taskId -> (priority, userId)
	h  *hp          // (priority, taskId, userId)
}

func Constructor(tasks [][]int) TaskManager {
	n := len(tasks)
	mp := make(map[int]pair, n) // 预分配空间
	h := make(hp, n)
	for i, t := range tasks {
		userId, taskId, priority := t[0], t[1], t[2]
		mp[taskId] = pair{priority, userId}
		h[i] = tuple{priority, taskId, userId}
	}
	heap.Init(&h)
	return TaskManager{mp, &h}
}

func (t *TaskManager) Add(userId, taskId, priority int) {
	t.mp[taskId] = pair{priority, userId}
	heap.Push(t.h, tuple{priority, taskId, userId})
}

func (t *TaskManager) Edit(taskId, newPriority int) {
	// 懒修改
	t.Add(t.mp[taskId].userId, taskId, newPriority)
}

func (t *TaskManager) Rmv(taskId int) {
	// 懒删除
	t.mp[taskId] = pair{-1, -1}
}

func (t *TaskManager) ExecTop() int {
	for t.h.Len() > 0 {
		top := heap.Pop(t.h).(tuple)
		priority, taskId, userId := top.priority, top.taskId, top.userId
		if t.mp[taskId] == (pair{priority, userId}) {
			t.Rmv(taskId)
			return userId
		}
		// else 货不对板，堆顶和 mp 中记录的不一样，说明堆顶数据已被修改或删除，不做处理
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

```js [sol-JavaScript]
class TaskManager {
    constructor(tasks) {
        // taskId -> [priority, userId]
        this.mp = new Map();

        // 最大堆 [priority, taskId, userId]
        this.pq = new PriorityQueue((a, b) => a[0] !== b[0] ? b[0] - a[0] : b[1] - a[1]);

        for (const [userId, taskId, priority] of tasks) {
            this.add(userId, taskId, priority);
        }
    }

    add(userId, taskId, priority) {
        this.mp.set(taskId, [priority, userId]);
        this.pq.enqueue([priority, taskId, userId]);
    }

    edit(taskId, newPriority) {
        // 懒修改
        const userId = this.mp.get(taskId)[1];
        this.add(userId, taskId, newPriority);
    }

    rmv(taskId) {
        // 懒删除
        this.mp.get(taskId)[0] = -1;
    }

    execTop() {
        while (!this.pq.isEmpty()) {
            const [priority, taskId, userId] = this.pq.dequeue();
            const [p, u] = this.mp.get(taskId);
            if (p === priority && u === userId) {
                this.rmv(taskId);
                return userId;
            }
            // else 货不对板，堆顶和 mp 中记录的不一样，说明堆顶数据已被修改或删除，不做处理
        }
        return -1;
    }
}
```

```rust [sol-Rust]
use std::collections::{BinaryHeap, HashMap};

struct TaskManager {
    mp: HashMap<i32, (i32, i32)>, // taskId -> (priority, userId)
    h: BinaryHeap<(i32, i32, i32)>, // (priority, taskId, userId)
}

impl TaskManager {
    fn new(tasks: Vec<Vec<i32>>) -> Self {
        let mut manager = Self {
            mp: HashMap::new(),
            h: BinaryHeap::new(),
        };
        for task in tasks {
            manager.add(task[0], task[1], task[2]);
        }
        manager
    }

    fn add(&mut self, user_id: i32, task_id: i32, priority: i32) {
        self.mp.insert(task_id, (priority, user_id));
        self.h.push((priority, task_id, user_id));
    }

    fn edit(&mut self, task_id: i32, new_priority: i32) {
        // 懒修改
        let user_id = self.mp.get(&task_id).unwrap().1;
        self.add(user_id, task_id, new_priority);
    }

    fn rmv(&mut self, task_id: i32) {
        // 懒删除
        *self.mp.get_mut(&task_id).unwrap() = (-1, -1);
    }

    fn exec_top(&mut self) -> i32 {
        while let Some((priority, task_id, user_id)) = self.h.pop() {
            let (p, u) = self.mp[&task_id];
            if p == priority && u == user_id {
                self.rmv(task_id);
                return user_id;
            }
            // else 货不对板，堆顶和 mp 中记录的不一样，说明堆顶数据已被修改或删除，不做处理
        }
        -1
    }
}
```

#### 复杂度分析

- 时间复杂度：
  - 初始化：$\mathcal{O}(n)$ 或者 $\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{tasks}$ 的长度。Python 和 Go 使用了堆化，复杂度是 $\mathcal{O}(n)$ 的。
  - $\texttt{add}$ 和 $\texttt{edit}$：$\mathcal{O}(\log (n+q))$，其中 $q$ 是 $\texttt{add}$ 和 $\texttt{edit}$ 的调用次数之和。
  - $\texttt{rmv}$：$\mathcal{O}(1)$。
  - $\texttt{execTop}$：均摊 $\mathcal{O}(\log (n+q))$。每个元素至多入堆出堆各一次。
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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
