既然是先来先匹配，那就用两个队列，分别保存乘客编号和司机编号。

此外，我们还需要知道在队列中的乘客是否取消了请求，可以用一个哈希集合（或者布尔数组）保存**在队列中且没有取消**的乘客编号。对于 $\texttt{cancelRider}$，把乘客编号从哈希集合中删除即可。

> 吐槽一下，没有叫车的乘客也能发送取消请求，这项目好烂……

[本题视频讲解](https://www.bilibili.com/video/BV1hd64BcEBQ/?t=15m48s)，欢迎点赞关注~

```py [sol-Python3]
class RideSharingSystem:
    def __init__(self):
        self.riders = deque()
        self.drivers = deque()
        self.waiting_riders = set()

    def addRider(self, riderId: int) -> None:
        self.riders.append(riderId)
        self.waiting_riders.add(riderId)

    def addDriver(self, driverId: int) -> None:
        self.drivers.append(driverId)

    def matchDriverWithRider(self) -> List[int]:
        # 弹出队列中的已取消乘客
        while self.riders and self.riders[0] not in self.waiting_riders:
            self.riders.popleft()
        # 没有乘客或者司机
        if not self.riders or not self.drivers:
            return [-1, -1]
        # 配对（这里没有删除 waiting_riders 中的乘客编号，面试的话建议写上删除的逻辑）
        return [self.drivers.popleft(), self.riders.popleft()]

    def cancelRider(self, riderId: int) -> None:
        self.waiting_riders.discard(riderId)
```

```java [sol-Java]
class RideSharingSystem {
    private final Deque<Integer> riders = new ArrayDeque<>();
    private final Deque<Integer> drivers = new ArrayDeque<>();
    private final Set<Integer> waitingEiders = new HashSet<>();

    public void addRider(int riderId) {
        riders.addLast(riderId);
        waitingEiders.add(riderId);
    }

    public void addDriver(int driverId) {
        drivers.addLast(driverId);
    }

    public int[] matchDriverWithRider() {
        // 弹出队列中的已取消乘客
        while (!riders.isEmpty() && !waitingEiders.contains(riders.peekFirst())) {
            riders.pollFirst();
        }
        // 没有乘客或者司机
        if (riders.isEmpty() || drivers.isEmpty()) {
            return new int[]{-1, -1};
        }
        // 配对（这里没有删除 waitingEiders 中的乘客编号，面试的话建议写上删除的逻辑）
        return new int[]{drivers.pollFirst(), riders.pollFirst()};
    }

    public void cancelRider(int riderId) {
        waitingEiders.remove(riderId);
    }
}
```

```cpp [sol-C++]
class RideSharingSystem {
    deque<int> riders;
    deque<int> drivers;
    unordered_set<int> waiting_riders;

public:
    void addRider(int riderId) {
        riders.push_back(riderId);
        waiting_riders.insert(riderId);
    }

    void addDriver(int driverId) {
        drivers.push_back(driverId);
    }

    vector<int> matchDriverWithRider() {
        // 弹出队列中的已取消乘客
        while (!riders.empty() && !waiting_riders.contains(riders.front())) {
            riders.pop_front();
        }
        // 没有乘客或者司机
        if (riders.empty() || drivers.empty()) {
            return {-1, -1};
        }
        // 配对（这里没有删除 waiting_riders 中的乘客编号，面试的话建议写上删除的逻辑）
        int driver = drivers.front(); drivers.pop_front();
        int rider = riders.front(); riders.pop_front();
        return {driver, rider};
    }

    void cancelRider(int riderId) {
        waiting_riders.erase(riderId);
    }
};
```

```go [sol-Go]
type RideSharingSystem struct {
	riders        []int
	drivers       []int
	waitingRiders map[int]bool
}

func Constructor() RideSharingSystem {
	return RideSharingSystem{
		waitingRiders: map[int]bool{},
	}
}

func (r *RideSharingSystem) AddRider(riderId int) {
	r.riders = append(r.riders, riderId)
	r.waitingRiders[riderId] = true
}

func (r *RideSharingSystem) AddDriver(driverId int) {
	r.drivers = append(r.drivers, driverId)
}

func (r *RideSharingSystem) MatchDriverWithRider() []int {
	// 弹出队列中的已取消乘客
	for len(r.riders) > 0 && !r.waitingRiders[r.riders[0]] {
		r.riders = r.riders[1:]
	}
	// 没有乘客或者司机
	if len(r.riders) == 0 || len(r.drivers) == 0 {
		return []int{-1, -1}
	}
	// 配对（这里没有删除 waitingRiders 中的乘客编号，面试的话建议写上删除的逻辑）
	ans := []int{r.drivers[0], r.riders[0]}
	r.riders = r.riders[1:]
	r.drivers = r.drivers[1:]
	return ans
}

func (r *RideSharingSystem) CancelRider(riderId int) {
	delete(r.waitingRiders, riderId)
}
```

#### 复杂度分析

- 时间复杂度：$\texttt{matchDriverWithRider}$ 均摊 $\mathcal{O}(1)$，其余为 $\mathcal{O}(1)$。每个乘客只会入队出队各一次。
- 空间复杂度：$\mathcal{O}(q)$。其中 $q$ 是 $\texttt{addRider}$ 和 $\texttt{addDriver}$ 的调用次数之和。

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
