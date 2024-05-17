用哈希表 $\textit{cnt}$ 统计每个数的出现次数。

用哈希表 $\textit{freq}$ 统计出现次数的出现次数，从而可以 $\mathcal{O}(1)$ 回答 $\texttt{hasFrequency}$。

增删元素的时候，除了修改 $\textit{cnt}[\textit{number}]$，还需要根据 $\textit{cnt}[\textit{number}]$ 的变化来修改 $\textit{freq}$，具体见代码。

```py [sol-Python3]
class FrequencyTracker:
    def __init__(self):
        self.cnt = defaultdict(int)  # number 的出现次数
        self.freq = defaultdict(int)  # number 的出现次数的出现次数

    def add(self, number: int, delta=1) -> None:
        self.freq[self.cnt[number]] -= 1  # 去掉一个旧的 cnt[number]
        self.cnt[number] += delta
        self.freq[self.cnt[number]] += 1  # 添加一个新的 cnt[number]

    def deleteOne(self, number: int) -> None:
        if self.cnt[number]:
            self.add(number, -1)

    def hasFrequency(self, frequency: int) -> bool:
        return self.freq[frequency] > 0  # 至少有一个 number 的出现次数恰好为 frequency
```

```java [sol-Java]
class FrequencyTracker {
    private final Map<Integer, Integer> cnt = new HashMap<>(); // number 的出现次数
    private final Map<Integer, Integer> freq = new HashMap<>(); // number 的出现次数的出现次数

    public FrequencyTracker() {}

    private void update(int number, int delta) {
        int c = cnt.merge(number, delta, Integer::sum);
        freq.merge(c - delta, -1, Integer::sum); // 去掉一个旧的 cnt[number]
        freq.merge(c, 1, Integer::sum); // 添加一个新的 cnt[number]
    }

    public void add(int number) {
        update(number, 1);
    }

    public void deleteOne(int number) {
        if (cnt.getOrDefault(number, 0) > 0) {
            update(number, -1);
        }
    }

    public boolean hasFrequency(int frequency) {
        return freq.getOrDefault(frequency, 0) > 0; // 至少有一个 number 的出现次数恰好为 frequency
    }
}
```

```cpp [sol-C++]
class FrequencyTracker {
    unordered_map<int, int> cnt; // number 的出现次数
    unordered_map<int, int> freq; // number 的出现次数的出现次数
public:
    FrequencyTracker() {}

    void add(int number) {
        --freq[cnt[number]]; // 去掉一个旧的 cnt[number]
        ++freq[++cnt[number]]; // 添加一个新的 cnt[number]
    }

    void deleteOne(int number) {
        if (!cnt[number]) return; // 不删除任何内容
        --freq[cnt[number]]; // 去掉一个旧的 cnt[number]
        ++freq[--cnt[number]]; // 添加一个新的 cnt[number]
    }

    bool hasFrequency(int frequency) {
        return freq[frequency]; // 至少有一个 number 的出现次数恰好为 frequency
    }
};
```

```go [sol-Go]
type FrequencyTracker struct {
	cnt  map[int]int // number 的出现次数
	freq map[int]int // number 的出现次数的出现次数
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{map[int]int{}, map[int]int{}}
}

func (f FrequencyTracker) update(number, delta int) {
	f.freq[f.cnt[number]]-- // 去掉一个旧的 cnt[number]
	f.cnt[number] += delta
	f.freq[f.cnt[number]]++ // 添加一个新的 cnt[number]
}

func (f FrequencyTracker) Add(number int) {
	f.update(number, 1)
}

func (f FrequencyTracker) DeleteOne(number int) {
	if f.cnt[number] > 0 {
		f.update(number, -1)
	}
}

func (f FrequencyTracker) HasFrequency(frequency int) bool {
	return f.freq[frequency] > 0 // 至少有一个 number 的出现次数恰好为 frequency
}
```

```js [sol-JavaScript]
class FrequencyTracker {
    constructor() {
        this.cnt = new Map(); // number 的出现次数
        this.freq = new Map(); // number 的出现次数的出现次数
    }

    add(number, delta = 1) {
        let c = this.cnt.get(number) ?? 0;
        this.freq.set(c, (this.freq.get(c) ?? 0) - 1); // 去掉一个旧的 cnt[number]
        c += delta;
        this.cnt.set(number, c);
        this.freq.set(c, (this.freq.get(c) ?? 0) + 1); // 添加一个新的 cnt[number]
    }

    deleteOne(number) {
        if ((this.cnt.get(number) ?? 0) > 0) {
            this.add(number, -1);
        }
    }

    hasFrequency(frequency) {
        return (this.freq.get(frequency) ?? 0) > 0; // 至少有一个 number 的出现次数恰好为 frequency
    }
}
```

```rust [sol-Rust]
use std::collections::HashMap;

struct FrequencyTracker {
    cnt: HashMap<i32, i32>, // number 的出现次数
    freq: HashMap<i32, i32>, // number 的出现次数的出现次数
}

impl FrequencyTracker {
    fn new() -> Self {
        Self { cnt: HashMap::new(), freq: HashMap::new() }
    }

    fn update(&mut self, number: i32, delta: i32) {
        let c = self.cnt.entry(number).or_insert(0);
        *self.freq.entry(*c).or_insert(0) -= 1; // 去掉一个旧的 cnt[number]
        *c += delta;
        *self.freq.entry(*c).or_insert(0) += 1; // 添加一个新的 cnt[number]
    }

    fn add(&mut self, number: i32) {
        self.update(number, 1);
    }

    fn delete_one(&mut self, number: i32) {
        if *self.cnt.get(&number).unwrap_or(&0) > 0 {
            self.update(number, -1);
        }
    }

    fn has_frequency(&self, frequency: i32) -> bool {
        *self.freq.get(&frequency).unwrap_or(&0) > 0 // 至少有一个 number 的出现次数恰好为 frequency
    }
}
```

#### 复杂度分析

- 时间复杂度：所有操作均为 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(q)$。其中 $q$ 为操作次数。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
