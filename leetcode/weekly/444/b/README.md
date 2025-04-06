题目要求 FIFO（先进先出），这可以用**队列**实现。

为了判重，可以用哈希表记录数据包。

对于 $\texttt{getCount}$，需要按照 $\textit{destination}$ 分组，所以要用哈希表套队列。

具体来说，创建三个数据结构：

1. $\textit{packetQ}$：存储数据包的队列。
2. $\textit{packetSet}$：存储所有未转发的数据包，方便判重。
3. $\textit{destToTimestamps}$：哈希表套队列，key 是 $\textit{destination}$，value 是对应的由 $\textit{timestamp}$ 组成的队列。

> 注：可以只把 $\textit{destination}$ 保存到 $\textit{packetQ}$ 中，$\textit{destToTimestamps}$ 中额外保存 $\textit{source}$。但考虑到二分是本题的瓶颈，所以 $\textit{destToTimestamps}$ 只保存 $\textit{timestamp}$ 更好。

其他就是模拟了，具体见代码。

$\texttt{getCount}$ 可以用 [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/) 的做法，二分查找。

为了方便二分，可以用列表模拟队列。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1ezRvYiE27/)，欢迎点赞关注~

```py [sol-Python3]
class Router:
    def __init__(self, memoryLimit: int):
        self.memory_limit = memoryLimit
        self.packet_q = deque()  # packet 队列
        self.packet_set = set()  # packet 集合
        self.dest_to_timestamps = defaultdict(deque)  # destination -> [timestamp]

    def addPacket(self, source: int, destination: int, timestamp: int) -> bool:
        packet = (source, destination, timestamp)
        if packet in self.packet_set:
            return False
        self.packet_set.add(packet)
        if len(self.packet_q) == self.memory_limit:  # 太多了
            self.forwardPacket()
        self.packet_q.append(packet)  # 入队
        self.dest_to_timestamps[destination].append(timestamp)
        return True

    def forwardPacket(self) -> List[int]:
        if not self.packet_q:
            return []
        packet = self.packet_q.popleft()  # 出队
        self.packet_set.remove(packet)
        self.dest_to_timestamps[packet[1]].popleft()
        return packet  # list(packet)

    def getCount(self, destination: int, startTime: int, endTime: int) -> int:
        timestamps = self.dest_to_timestamps[destination]
        left = bisect_left(timestamps, startTime)  # deque 访问不是 O(1) 的，可以看另一份代码
        right = bisect_right(timestamps, endTime)
        return right - left
```

```py [sol-Python3 list]
class Router:
    def __init__(self, memoryLimit: int):
        self.memory_limit = memoryLimit
        self.packet_q = deque()  # packet 队列
        self.packet_set = set()  # packet 集合
        self.dest_to_timestamps = defaultdict(lambda: [[], 0])  # destination -> [[timestamp], head]

    def addPacket(self, source: int, destination: int, timestamp: int) -> bool:
        packet = (source, destination, timestamp)
        if packet in self.packet_set:
            return False
        self.packet_set.add(packet)
        if len(self.packet_q) == self.memory_limit:  # 太多了
            self.forwardPacket()
        self.packet_q.append(packet)  # 入队
        self.dest_to_timestamps[destination][0].append(timestamp)
        return True

    def forwardPacket(self) -> List[int]:
        if not self.packet_q:
            return []
        packet = self.packet_q.popleft()  # 出队
        self.packet_set.remove(packet)
        self.dest_to_timestamps[packet[1]][1] += 1  # 队首下标加一，模拟出队
        return packet  # list(packet)

    def getCount(self, destination: int, startTime: int, endTime: int) -> int:
        timestamps, head = self.dest_to_timestamps[destination]
        left = bisect_left(timestamps, startTime, head)
        right = bisect_right(timestamps, endTime, head)
        return right - left
```

```java [sol-Java]
class Router {
    private record Packet(int source, int destination, int timestamp) {
    }

    private record Queue(List<Integer> timestamps, int head) {
    }

    private final int memoryLimit;
    private final Deque<Packet> packetQ = new ArrayDeque<>(); // Packet 队列
    private final Set<Packet> packetSet = new HashSet<>(); // Packet 集合
    private final Map<Integer, Queue> destToTimestamps = new HashMap<>(); // destination -> [[timestamp], head]

    public Router(int memoryLimit) {
        this.memoryLimit = memoryLimit;
    }

    public boolean addPacket(int source, int destination, int timestamp) {
        Packet packet = new Packet(source, destination, timestamp);
        if (!packetSet.add(packet)) {
            return false;
        }
        if (packetQ.size() == memoryLimit) { // 太多了
            forwardPacket();
        }
        packetQ.offer(packet); // 入队
        destToTimestamps.computeIfAbsent(destination, k -> new Queue(new ArrayList<>(), 0)).timestamps.add(timestamp);
        return true;
    }

    public int[] forwardPacket() {
        if (packetQ.isEmpty()) {
            return new int[]{};
        }
        Packet packet = packetQ.poll(); // 出队
        packetSet.remove(packet);
        destToTimestamps.compute(packet.destination, (k, q) -> new Queue(q.timestamps, q.head + 1)); // 队首下标加一，模拟出队
        return new int[]{packet.source, packet.destination, packet.timestamp};
    }

    public int getCount(int destination, int startTime, int endTime) {
        Queue q = destToTimestamps.get(destination);
        if (q == null) {
            return 0;
        }
        int left = lowerBound(q.timestamps, startTime, q.head - 1);
        int right = lowerBound(q.timestamps, endTime + 1, q.head - 1);
        return right - left;
    }

    private int lowerBound(List<Integer> nums, int target, int left) {
        int right = nums.size();
        while (left + 1 < right) {
            int mid = (left + right) >>> 1;
            if (nums.get(mid) >= target) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Router {
    int memory_limit;
    queue<tuple<int, int, int>> packet_q; // packet 队列
    set<tuple<int, int, int>> packet_set; // 本来要用 unordered_set，但 tuple 需要手写哈希，为了方便直接用 set
    unordered_map<int, pair<vector<int>, int>> dest_to_timestamps; // destination -> [[timestamp], head]

public:
    Router(int memoryLimit) {
        memory_limit = memoryLimit;
    }

    bool addPacket(int source, int destination, int timestamp) {
        auto packet = make_tuple(source, destination, timestamp);
        if (!packet_set.insert(packet).second) {
            return false;
        }
        if (packet_q.size() == memory_limit) {  // 太多了
            forwardPacket();
        }
        packet_q.emplace(packet); // 入队
        dest_to_timestamps[destination].first.push_back(timestamp);
        return true;
    }

    vector<int> forwardPacket() {
        if (packet_q.empty()) {
            return {};
        }
        auto packet = packet_q.front(); // 出队
        packet_q.pop();
        packet_set.erase(packet);
        auto [source, destination, timestamp] = packet;
        dest_to_timestamps[destination].second++; // 队首下标加一，模拟出队
        return {source, destination, timestamp};
    }

    int getCount(int destination, int startTime, int endTime) {
        auto& [timestamps, head] = dest_to_timestamps[destination];
        auto left = ranges::lower_bound(timestamps.begin() + head, timestamps.end(), startTime);
        auto right = ranges::upper_bound(timestamps.begin() + head, timestamps.end(), endTime);
        return right - left;
    }
};
```

```go [sol-Go]
type packet struct {
	source, destination, timestamp int
}

type Router struct {
	memoryLimit      int
	packetQ          []packet            // packet 队列
	packetSet        map[packet]struct{} // packet 集合
	destToTimestamps map[int][]int       // destination -> [timestamp]
}

func Constructor(memoryLimit int) Router {
	return Router{
		memoryLimit:      memoryLimit,
		packetSet:        map[packet]struct{}{},
		destToTimestamps: map[int][]int{},
	}
}

func (r *Router) AddPacket(source, destination, timestamp int) bool {
	pkt := packet{source, destination, timestamp}
	if _, ok := r.packetSet[pkt]; ok {
		return false
	}
	r.packetSet[pkt] = struct{}{}
	if len(r.packetQ) == r.memoryLimit { // 太多了
		r.ForwardPacket()
	}
	r.packetQ = append(r.packetQ, pkt) // 入队
	r.destToTimestamps[destination] = append(r.destToTimestamps[destination], timestamp)
	return true
}

func (r *Router) ForwardPacket() []int {
	if len(r.packetQ) == 0 {
		return nil
	}
	pkt := r.packetQ[0]
	r.packetQ = r.packetQ[1:] // 出队
	r.destToTimestamps[pkt.destination] = r.destToTimestamps[pkt.destination][1:]
	delete(r.packetSet, pkt)
	return []int{pkt.source, pkt.destination, pkt.timestamp}
}

func (r *Router) GetCount(destination, startTime, endTime int) int {
	timestamps := r.destToTimestamps[destination]
	return sort.SearchInts(timestamps, endTime+1) - sort.SearchInts(timestamps, startTime)
}
```

#### 复杂度分析

- 时间复杂度：$\texttt{GetCount}$ 是 $\mathcal{O}(\log \min(q, \textit{memoryLimit}))$，其中 $q$ 是 $\texttt{addPacket}$ 的调用次数。其余操作为 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(\min(q, \textit{memoryLimit}))$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
