## 方法一：哈希表 + 有序集合

为了实现 $\texttt{find}$，我们需要对每个 $\textit{number}$ 创建一个有序集合，维护这个 $\textit{number}$ 对应的所有下标。用有序集合可以快速地获取最小下标。

对于 $\texttt{change}$，如果 $\textit{index}$ 处有数字，我们需要先删除旧的数字，所以还需要知道每个 $\textit{index}$ 对应的 $\textit{number}$ 是多少，这可以用一个哈希表记录。

具体来说，创建一个哈希表 $\textit{indexToNumber}$，以及一个哈希表套有序集合 $\textit{numberToIndices}$。

对于 $\texttt{change}$：

- 如果 $\textit{index}$ 处有数字 $x$，那么从 $\textit{numberToIndices}[x]$ 中删除 $\textit{index}$（删除旧的数据）。
- 然后，更新（或者插入）$\textit{indexToNumber}[\textit{index}] = \textit{number}$，往 $\textit{numberToIndices}[\textit{number}]$ 中添加 $\textit{index}$。

对于 $\texttt{find}$，获取 $\textit{numberToIndices}[\textit{number}]$ 中的最小元素即可。

```py [sol-Python3]
class NumberContainers:
    def __init__(self):
        self.index_to_number = {}
        # from sortedcontainers import SortedSet
        self.number_to_indices = defaultdict(SortedSet)

    def change(self, index: int, number: int) -> None:
        # 移除旧数据
        old_number = self.index_to_number.get(index, None)
        if old_number is not None:
            self.number_to_indices[old_number].discard(index)

        # 添加新数据
        self.index_to_number[index] = number
        self.number_to_indices[number].add(index)

    def find(self, number: int) -> int:
        indices = self.number_to_indices[number]
        return indices[0] if indices else -1
```

```java [sol-Java]
class NumberContainers {
    private final Map<Integer, Integer> indexToNumber = new HashMap<>();
    private final Map<Integer, TreeSet<Integer>> numberToIndices = new HashMap<>();

    public void change(int index, int number) {
        // 移除旧数据
        Integer oldNumber = indexToNumber.get(index);
        if (oldNumber != null) {
            numberToIndices.get(oldNumber).remove(index);
        }

        // 添加新数据
        indexToNumber.put(index, number);
        numberToIndices.computeIfAbsent(number, _ -> new TreeSet<>()).add(index);
    }

    public int find(int number) {
        TreeSet<Integer> indices = numberToIndices.get(number);
        return indices == null || indices.isEmpty() ? -1 : indices.first();
    }
}
```

```cpp [sol-C++]
class NumberContainers {
    unordered_map<int, int> index_to_number;
    unordered_map<int, set<int>> number_to_indices;

public:
    void change(int index, int number) {
        // 移除旧数据
        auto it = index_to_number.find(index);
        if (it != index_to_number.end()) {
            number_to_indices[it->second].erase(index);
        }

        // 添加新数据
        index_to_number[index] = number;
        number_to_indices[number].insert(index);
    }

    int find(int number) {
        auto it = number_to_indices.find(number);
        return it == number_to_indices.end() || it->second.empty() ? -1 : *it->second.begin();
    }
};
```

```go [sol-Go]
// import "github.com/emirpasic/gods/v2/trees/redblacktree"
type NumberContainers struct {
	indexToNumber   map[int]int
	numberToIndices map[int]*redblacktree.Tree[int, struct{}]
}

func Constructor() NumberContainers {
	return NumberContainers{map[int]int{}, map[int]*redblacktree.Tree[int, struct{}]{}}
}

func (n NumberContainers) Change(index, number int) {
	// 移除旧数据
	if oldNumber, ok := n.indexToNumber[index]; ok {
		n.numberToIndices[oldNumber].Remove(index)
	}

	// 添加新数据
	n.indexToNumber[index] = number
	if n.numberToIndices[number] == nil {
		n.numberToIndices[number] = redblacktree.New[int, struct{}]()
	}
	n.numberToIndices[number].Put(index, struct{}{})
}

func (n NumberContainers) Find(number int) int {
	indices, ok := n.numberToIndices[number]
	if !ok || indices.Empty() {
		return -1
	}
	return indices.Left().Key
}
```

#### 复杂度分析

- 时间复杂度：
   - 初始化 $\mathcal{O}(1)$。
   - $\texttt{change}$：$\mathcal{O}(\log q)$，其中 $q$ 是 $\texttt{change}$ 的调用次数。
   - $\texttt{find}$：$\mathcal{O}(\log q)$ 或者 $\mathcal{O}(1)$，取决于有序集合是否额外维护最小值。
- 空间复杂度：$\mathcal{O}(q)$。

## 方法二：哈希表 + 懒删除堆

$\textit{numberToIndices}$ 改成哈希表套最小堆。

对于 $\texttt{change}$，不删除旧数据。

对于 $\texttt{find}$，查看堆顶是否等于 $\textit{number}$，若不相同，则意味着堆顶是之前没有删除的旧数据，弹出堆顶；否则堆顶就是答案。

```py [sol-Python3]
class NumberContainers:
    def __init__(self):
        self.index_to_number = {}
        self.number_to_indices = defaultdict(list)

    def change(self, index: int, number: int) -> None:
        # 添加新数据
        self.index_to_number[index] = number
        heappush(self.number_to_indices[number], index)

    def find(self, number: int) -> int:
        indices = self.number_to_indices[number]
        while indices and self.index_to_number[indices[0]] != number:
            heappop(indices)  # 堆顶货不对板，说明是旧数据，删除
        return indices[0] if indices else -1
```

```java [sol-Java]
class NumberContainers {
    private final Map<Integer, Integer> indexToNumber = new HashMap<>();
    private final Map<Integer, PriorityQueue<Integer>> numberToIndices = new HashMap<>();

    public void change(int index, int number) {
        // 添加新数据
        indexToNumber.put(index, number);
        numberToIndices.computeIfAbsent(number, _ -> new PriorityQueue<>()).offer(index);
    }

    public int find(int number) {
        PriorityQueue<Integer> indices = numberToIndices.get(number);
        if (indices == null) {
            return -1;
        }
        while (!indices.isEmpty() && indexToNumber.get(indices.peek()) != number) {
            indices.poll(); // 堆顶货不对板，说明是旧数据，删除
        }
        return indices.isEmpty() ? -1 : indices.peek();
    }
}
```

```cpp [sol-C++]
class NumberContainers {
    unordered_map<int, int> index_to_number;
    unordered_map<int, priority_queue<int, vector<int>, greater<>>> number_to_indices;

public:
    void change(int index, int number) {
        // 添加新数据
        index_to_number[index] = number;
        number_to_indices[number].push(index);
    }

    int find(int number) {
        auto& indices = number_to_indices[number];
        while (!indices.empty() && index_to_number[indices.top()] != number) {
            indices.pop(); // 堆顶货不对板，说明是旧数据，删除
        }
        return indices.empty() ? -1 : indices.top();
    }
};
```

```go [sol-Go]
type NumberContainers struct {
	indexToNumber   map[int]int
	numberToIndices map[int]*hp
}

func Constructor() NumberContainers {
	return NumberContainers{map[int]int{}, map[int]*hp{}}
}

func (n NumberContainers) Change(index, number int) {
	// 添加新数据
	n.indexToNumber[index] = number
	if _, ok := n.numberToIndices[number]; !ok {
		n.numberToIndices[number] = &hp{}
	}
	heap.Push(n.numberToIndices[number], index)
}

func (n NumberContainers) Find(number int) int {
	indices, ok := n.numberToIndices[number]
	if !ok {
		return -1
	}
	for indices.Len() > 0 && n.indexToNumber[indices.IntSlice[0]] != number {
		heap.Pop(indices) // 堆顶货不对板，说明是旧数据，删除
	}
	if indices.Len() == 0 {
		return -1
	}
	return indices.IntSlice[0]
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
```

#### 复杂度分析

- 时间复杂度：
  - 初始化 $\mathcal{O}(1)$。
  - $\texttt{change}$：$\mathcal{O}(\log q)$，其中 $q$ 是 $\texttt{change}$ 的调用次数。
  - $\texttt{find}$：均摊 $\mathcal{O}(\log q)$。
- 空间复杂度：$\mathcal{O}(q)$。

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
