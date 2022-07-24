下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

## 方法一：平衡树

由于数据范围很大，我们可以用一个**哈希表** $m$ 记录每个下标对应的元素，另一个**哈希表套平衡树** $\textit{ms}$ 记录每个元素对应的下标集合。

对于 `change` 操作，如果 $\textit{index}$ 处已有数字，则先从 $\textit{ms}[m[\textit{index}]]$ 中删掉 $\textit{index}$。然后将 $\textit{index}$ 和 $\textit{number}$ 记录到 $m$ 和 $\textit{ms}$ 中。

```py [sol1-Python3]
from sortedcontainers import SortedSet

class NumberContainers:
    def __init__(self):
        self.m = {}
        self.ms = defaultdict(SortedSet)

    def change(self, index: int, number: int) -> None:
        if index in self.m:
            self.ms[self.m[index]].remove(index)
        self.m[index] = number
        self.ms[number].add(index)

    def find(self, number: int) -> int:
        s = self.ms[number]
        return s[0] if s else -1
```

```java [sol1-Java]
class NumberContainers {
    Map<Integer, Integer> m = new HashMap<>();
    Map<Integer, TreeSet<Integer>> ms = new HashMap<>();

    public void change(int index, int number) {
        var old = m.get(index);
        if (old != null) ms.get(old).remove(index);
        m.put(index, number);
        ms.computeIfAbsent(number, k -> new TreeSet<>()).add(index);
    }

    public int find(int number) {
        var s = ms.get(number);
        return s == null || s.isEmpty() ? -1 : s.first();
    }
}
```

```cpp [sol1-C++]
class NumberContainers {
    map<int, int> m;
    map<int, set<int>> ms;

public:
    void change(int index, int number) {
        auto it = m.find(index);
        if (it != m.end()) {
            ms[it->second].erase(index);
            it->second = number; // 优化：直接在迭代器上赋值
        } else m[index] = number;
        ms[number].insert(index);
    }

    int find(int number) {
        auto it = ms.find(number);
        return it == ms.end() || it->second.empty() ? -1 : *it->second.begin();
    }
};
```

```go [sol1-Go]
type NumberContainers struct {
	m  map[int]int
	ms map[int]*redblacktree.Tree
}

func Constructor() NumberContainers {
	return NumberContainers{map[int]int{}, map[int]*redblacktree.Tree{}}
}

func (n NumberContainers) Change(index int, number int) {
	if num, ok := n.m[index]; ok {
		n.ms[num].Remove(index)
	}
	n.m[index] = number
	if n.ms[number] == nil {
		n.ms[number] = redblacktree.NewWithIntComparator()
	}
	n.ms[number].Put(index, nil)
}

func (n NumberContainers) Find(number int) int {
	s, ok := n.ms[number]
	if !ok || s.Size() == 0 {
		return -1
	}
	return s.Left().Key.(int)
}
```

## 方法二：堆

另一种做法是用堆：

- 对于 `change` 操作，直接往 $\textit{ms}$ 中记录，不做任何删除操作；
- 对于 `find` 操作，查看堆顶下标对应的元素是否和 $\textit{number}$ 相同，若不相同则意味着对应的元素已被替换成了其他值，直接弹出堆顶；否则堆顶就是答案。

```py [sol2-Python3]
class NumberContainers:
    def __init__(self):
        self.m = {}
        self.ms = defaultdict(list)

    def change(self, index: int, number: int) -> None:
        self.m[index] = number
        heappush(self.ms[number], index)

    def find(self, number: int) -> int:
        h = self.ms[number]
        while h and self.m[h[0]] != number:  # 意味着 h[0] 处的元素已被替换成了其他值
            heappop(h)
        return h[0] if h else -1
```

```java [sol2-Java]
class NumberContainers {
    Map<Integer, Integer> m = new HashMap<>();
    Map<Integer, Queue<Integer>> ms = new HashMap<>();

    public void change(int index, int number) {
        m.put(index, number);
        ms.computeIfAbsent(number, k -> new PriorityQueue<>()).offer(index);
    }

    public int find(int number) {
        var q = ms.get(number);
        if (q == null) return -1;
        while (!q.isEmpty() && m.get(q.peek()) != number) q.poll();
        return q.isEmpty() ? -1 : q.peek();
    }
}
```

```cpp [sol2-C++]
class NumberContainers {
    map<int, int> m;
    map<int, priority_queue<int, vector<int>, greater<>>> ms;

public:
    void change(int index, int number) {
        m[index] = number;
        ms[number].push(index);
    }

    int find(int number) {
        auto it = ms.find(number);
        if (it == ms.end()) return -1;
        auto &q = it->second;
        while (!q.empty() && m[q.top()] != number) q.pop();
        return q.empty() ? -1 : q.top();
    }
};
```

```go [sol2-Go]
type NumberContainers struct {
	m  map[int]int
	ms map[int]*hp
}

func Constructor() NumberContainers {
	return NumberContainers{map[int]int{}, map[int]*hp{}}
}

func (n NumberContainers) Change(index int, number int) {
	n.m[index] = number
	if n.ms[number] == nil {
		n.ms[number] = &hp{}
	}
	heap.Push(n.ms[number], index)
}

func (n NumberContainers) Find(number int) int {
	h, ok := n.ms[number]
	if !ok {
		return -1
	}
	for h.Len() > 0 && n.m[h.IntSlice[0]] != number {
		heap.Pop(h)
	}
	if h.Len() == 0 {
		return -1
	}
	return h.IntSlice[0]
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
```
