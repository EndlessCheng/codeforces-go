既然是先来先匹配，那就用两个队列，分别保存乘客编号和司机编号。

此外，还需要知道**在队列中的**乘客是否取消了请求，这可以用一个哈希集合（或者布尔数组）保存已取消的乘客编号。

如果乘客根本就不在队列中，我们不能将其标记为已取消。所以还需要再用一个哈希集合（或者布尔数组）保存**在队列中的乘客编号**。注意题目保证乘客编号唯一且只会被添加一次。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class RideSharingSystem:
    def __init__(self):
        self.riders = deque()
        self.drivers = deque()
        self.seen_riders = set()
        self.canceled_riders = set()

    def addRider(self, riderId: int) -> None:
        self.riders.append(riderId)
        self.seen_riders.add(riderId)

    def addDriver(self, driverId: int) -> None:
        self.drivers.append(driverId)

    def matchDriverWithRider(self) -> List[int]:
        # 弹出队列中的已取消乘客
        while self.riders and self.riders[0] in self.canceled_riders:
            self.riders.popleft()
        # 没有乘客或者司机
        if not self.riders or not self.drivers:
            return [-1, -1]
        # 配对
        return [self.drivers.popleft(), self.riders.popleft()]

    def cancelRider(self, riderId: int) -> None:
        # 对于不存在的乘客，不能标记为取消
        if riderId in self.seen_riders:
            self.canceled_riders.add(riderId)
```

```java [sol-Java]
class RideSharingSystem {
    private final Deque<Integer> riders = new ArrayDeque<>();
    private final Deque<Integer> drivers = new ArrayDeque<>();
    private final Set<Integer> seenRiders = new HashSet<>();
    private final Set<Integer> canceledRiders = new HashSet<>();

    public void addRider(int riderId) {
        riders.addLast(riderId);
        seenRiders.add(riderId);
    }

    public void addDriver(int driverId) {
        drivers.addLast(driverId);
    }

    public int[] matchDriverWithRider() {
        // 弹出队列中的已取消乘客
        while (!riders.isEmpty() && canceledRiders.contains(riders.peekFirst())) {
            riders.pollFirst();
        }
        // 没有乘客或者司机
        if (riders.isEmpty() || drivers.isEmpty()) {
            return new int[]{-1, -1};
        }
        // 配对
        return new int[]{drivers.pollFirst(), riders.pollFirst()};
    }

    public void cancelRider(int riderId) {
        // 对于不存在的乘客，不能标记为取消
        if (seenRiders.contains(riderId)) {
            canceledRiders.add(riderId);
        }
    }
}
```

```cpp [sol-C++]
class RideSharingSystem {
    deque<int> riders;
    deque<int> drivers;
    unordered_set<int> seen_riders;
    unordered_set<int> canceled_riders;

public:
    void addRider(int riderId) {
        riders.push_back(riderId);
        seen_riders.insert(riderId);
    }

    void addDriver(int driverId) {
        drivers.push_back(driverId);
    }

    vector<int> matchDriverWithRider() {
        // 弹出队列中的已取消乘客
        while (!riders.empty() && canceled_riders.contains(riders.front())) {
            riders.pop_front();
        }
        // 没有乘客或者司机
        if (riders.empty() || drivers.empty()) {
            return {-1, -1};
        }
        // 配对
        int driver = drivers.front(); drivers.pop_front();
        int rider = riders.front(); riders.pop_front();
        return {driver, rider};
    }

    void cancelRider(int riderId) {
        // 对于不存在的乘客，不能标记为取消
        if (seen_riders.contains(riderId)) {
            canceled_riders.insert(riderId);
        }
    }
};
```

```go [sol-Go]
type RideSharingSystem struct {
	riders         []int
	drivers        []int
	seenRiders     map[int]bool
	canceledRiders map[int]bool
}

func Constructor() RideSharingSystem {
	return RideSharingSystem{
		seenRiders:     map[int]bool{},
		canceledRiders: map[int]bool{},
	}
}

func (r *RideSharingSystem) AddRider(riderId int) {
	r.riders = append(r.riders, riderId)
	r.seenRiders[riderId] = true
}

func (r *RideSharingSystem) AddDriver(driverId int) {
	r.drivers = append(r.drivers, driverId)
}

func (r *RideSharingSystem) MatchDriverWithRider() []int {
	// 弹出队列中的已取消乘客
	for len(r.riders) > 0 && r.canceledRiders[r.riders[0]] {
		r.riders = r.riders[1:]
	}
	// 没有乘客或者司机
	if len(r.riders) == 0 || len(r.drivers) == 0 {
		return []int{-1, -1}
	}
	// 配对
	ans := []int{r.drivers[0], r.riders[0]}
	r.riders = r.riders[1:]
	r.drivers = r.drivers[1:]
	return ans
}

func (r *RideSharingSystem) CancelRider(riderId int) {
	// 对于不存在的乘客，不能标记为取消
	if r.seenRiders[riderId] {
		r.canceledRiders[riderId] = true
	}
}
```

#### 复杂度分析

- 时间复杂度：$\texttt{MatchDriverWithRider}$ 均摊 $\mathcal{O}(1)$，其余为 $\mathcal{O}(1)$。每个乘客只会入队出队各一次。
- 空间复杂度：$\mathcal{O}(q)$。其中 $q$ 是 $\texttt{AddRider}$ 和 $\texttt{AddDriver}$ 的调用次数之和。

## 专题训练

见下面数据结构题单的「**四、队列**」。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
