既然同一个包裹中的苹果可以分装到**不同**的箱子中，那就先把所有苹果堆在一起，然后一个个地装箱。

为了少用箱子，要先装大箱子，再装小箱子。

> 注：题目保证可以将所有苹果重新分装到箱子中。

```py [sol-Python3]
class Solution:
    def minimumBoxes(self, apple: List[int], capacity: List[int]) -> int:
        s = sum(apple)  # 把所有苹果堆在一起
        capacity.sort(reverse=True)  # 先装大箱子，再装小箱子
        for i, x in enumerate(capacity):
            s -= x
            if s <= 0:  # 所有苹果都装入了箱子
                return i + 1  # 从 0 到 i 有 i+1 个箱子
```

```java [sol-Java]
class Solution {
    public int minimumBoxes(int[] apple, int[] capacity) {
        int s = 0;
        for (int x : apple) {
            s += x; // 把所有苹果堆在一起
        }

        Arrays.sort(capacity);

        int m = capacity.length;
        int i = m - 1; // 先装大箱子，再装小箱子
        while (s > 0) { // 还有剩余苹果
            s -= capacity[i--]; // 使用一个箱子
        }
        return m - 1 - i;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumBoxes(vector<int>& apple, vector<int>& capacity) {
        int s = reduce(apple.begin(), apple.end()); // 把所有苹果堆在一起
        ranges::sort(capacity, greater()); // 先装大箱子，再装小箱子
        int i = 0;
        while (s > 0) { // 还有剩余苹果
            s -= capacity[i++]; // 使用一个箱子
        }
        return i;
    }
};
```

```c [sol-C]
int cmp(const void* a, const void* b) {
    return *(int*)b - *(int*)a;
}

int minimumBoxes(int* apple, int appleSize, int* capacity, int capacitySize) {
    int s = 0;
    for (int i = 0; i < appleSize; i++) {
        s += apple[i]; // 把所有苹果堆在一起
    }

    qsort(capacity, capacitySize, sizeof(int), cmp); // 先装大箱子，再装小箱子

    int i = 0;
    while (s > 0) { // 还有剩余苹果
        s -= capacity[i++]; // 使用一个箱子
    }
    return i;
}
```

```go [sol-Go]
func minimumBoxes(apple, capacity []int) int {
	s := 0
	for _, x := range apple {
		s += x // 把所有苹果堆在一起
	}
	slices.SortFunc(capacity, func(a, b int) int { return b - a }) // 先装大箱子，再装小箱子
	for i, c := range capacity {
		s -= c
		if s <= 0 { // 所有苹果都装入了箱子
			return i + 1 // 从 0 到 i 有 i+1 个箱子
		}
	}
	return -1
}
```

```js [sol-JavaScript]
var minimumBoxes = function(apple, capacity) {
    let s = _.sum(apple); // 把所有苹果堆在一起
    capacity.sort((a, b) => b - a); // 先装大箱子，再装小箱子
    let i = 0;
    while (s > 0) { // 还有剩余苹果
        s -= capacity[i++]; // 使用一个箱子
    }
    return i;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn minimum_boxes(apple: Vec<i32>, mut capacity: Vec<i32>) -> i32 {
        let mut s = apple.iter().sum::<i32>(); // 把所有苹果堆在一起
        capacity.sort_unstable_by_key(|a| -a); // 先装大箱子，再装小箱子
        for (i, x) in capacity.iter().enumerate() {
            s -= x;
            if s <= 0 { // 所有苹果都装入了箱子
                return (i + 1) as _; // 从 0 到 i 有 i+1 个箱子
            }
        }
        unreachable!()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m\log m)$，其中 $n$ 为 $\textit{apple}$ 的长度，$m$ 为 $\textit{capacity}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 专题训练

见贪心题单的「**§1.1 从最小/最大开始贪心**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
