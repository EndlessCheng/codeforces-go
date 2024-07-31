## 方法一：转换 + 二分

![](https://pic.leetcode.cn/1695696658-skpYuS-2251.jpg)

把 $\textit{flowers}$ 中的数据分成两个数组：开始时间 $\textit{starts}$ 和结束时间 $\textit{ends}$。

第 $i$ 个人能看到的花的数目，等价于 $\textit{starts}$ 中的小于等于 $\textit{people}[i]$ 的花的数目，减去 $\textit{ends}$ 中的小于 $\textit{people}[i]$ 的花的数目，即**开花数减去凋落数**。例如上图中的 $\textit{people}[2]=7$，有 $3$ 朵花在小于等于 $7$ 的时刻开花，有 $1$ 朵花在小于 $7$ 的时刻凋落，所以这个人能看到的花的数目为 $3-1=2$。

如何快速计算开花数和凋落数？

把 $\textit{starts}$ 和 $\textit{ends}$ 从小到大排序，然后：

- 在 $\textit{starts}$ 中二分查找大于 $\textit{people}[i]$ 的下一个数的下标（若不存在则为 $n$），即为开花数。
- 在 $\textit{ends}$ 中二分查找大于等于 $\textit{people}[i]$ 的下一个数的下标（若不存在则为 $n$），即为凋落数。

关于二分查找的原理，请看视频讲解：[二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)

```Python [sol-Python3]
class Solution:
    def fullBloomFlowers(self, flowers: List[List[int]], people: List[int]) -> List[int]:
        starts = sorted(s for s, _ in flowers)
        ends = sorted(e for _, e in flowers)
        return [bisect_right(starts, p) - bisect_left(ends, p) for p in people]
```

```java [sol-Java]
class Solution {
    public int[] fullBloomFlowers(int[][] flowers, int[] people) {
        int n = flowers.length;
        var starts = new int[n];
        var ends = new int[n];
        for (int i = 0; i < n; i++) {
            starts[i] = flowers[i][0];
            ends[i] = flowers[i][1];
        }
        Arrays.sort(starts);
        Arrays.sort(ends);

        for (int i = 0; i < people.length; i++) {
            people[i] = lowerBound(starts, people[i] + 1) - lowerBound(ends, people[i]);
        }
        return people;
    }

    // 返回 >= x 的第一个数的下标
    // 如果不存在（所有元素都小于 x），则返回 nums.length
    private int lowerBound(int[] nums, int x) {
        int left = -1, right = nums.length; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] < x
            // nums[right] >= x
            int mid = left + (right - left) / 2;
            if (nums[mid] < x) {
                left = mid; // 区间缩小为 (mid, right)
            } else {
                right = mid; // 区间缩小为 (left, mid)
            }
        }
        return right; // 根据循环不变量，此时 right 就是满足 nums[right] >= x 的最小值
    }
}
```

```java [sol-Java Stream]
class Solution {
    public int[] fullBloomFlowers(int[][] flowers, int[] people) {
        int[] starts = Arrays.stream(flowers).mapToInt(f -> f[0]).sorted().toArray();
        int[] ends = Arrays.stream(flowers).mapToInt(f -> f[1]).sorted().toArray();
        return Arrays.stream(people).map(p -> lowerBound(starts, p + 1) - lowerBound(ends, p)).toArray();
    }

    // 返回 >= x 的第一个数的下标
    // 如果不存在（所有元素都小于 x），则返回 nums.length
    private int lowerBound(int[] nums, int x) {
        int left = -1, right = nums.length; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] < x
            // nums[right] >= x
            int mid = left + (right - left) / 2;
            if (nums[mid] < x) {
                left = mid; // 区间缩小为 (mid, right)
            } else {
                right = mid; // 区间缩小为 (left, mid)
            }
        }
        return right; // 根据循环不变量，此时 right 就是满足 nums[right] >= x 的最小值
    }
}
```

```C++ [sol-C++]
class Solution {
public:
    vector<int> fullBloomFlowers(vector<vector<int>> &flowers, vector<int> &people) {
        int n = flowers.size();
        vector<int> starts(n), ends(n);
        for (int i = 0; i < n; i++) {
            starts[i] = flowers[i][0];
            ends[i] = flowers[i][1];
        }
        sort(starts.begin(), starts.end());
        sort(ends.begin(), ends.end());

        for (int &p: people)
            p = (upper_bound(starts.begin(), starts.end(), p) - starts.begin()) -
                (lower_bound(ends.begin(), ends.end(), p) - ends.begin());
        return people;
    }
};
```

```go [sol-Go]
func fullBloomFlowers(flowers [][]int, people []int) []int {
	n := len(flowers)
	starts := make([]int, n)
	ends := make([]int, n)
	for i, f := range flowers {
		starts[i] = f[0]
		ends[i] = f[1]
	}
	sort.Ints(starts)
	sort.Ints(ends)

	for i, p := range people {
		people[i] = sort.SearchInts(starts, p+1) - sort.SearchInts(ends, p)
	}
	return people
}
```

```js [sol-JavaScript]
var fullBloomFlowers = function (flowers, people) {
    const starts = flowers.map(f => f[0]).sort((a, b) => a - b);
    const ends = flowers.map(f => f[1]).sort((a, b) => a - b);
    return people.map(p => lowerBound(starts, p + 1) - lowerBound(ends, p));
};

// 返回 >= x 的第一个数的下标
// 如果不存在（所有元素都小于 x），则返回 nums.length
var lowerBound = function (nums, x) {
    let left = -1, right = nums.length; // 开区间 (left, right)
    while (left + 1 < right) { // 区间不为空
        // 循环不变量：
        // nums[left] < x
        // nums[right] >= x
        const mid = left + ((right - left) >> 1);
        if (nums[mid] < x)
            left = mid; // 区间缩小为 (mid, right)
        else
            right = mid; // 区间缩小为 (left, mid)
    }
    return right; // 根据循环不变量，此时 right 就是满足 nums[right] >= x 的最小值
};
```

```rust [sol-Rust]
impl Solution {
    pub fn full_bloom_flowers(flowers: Vec<Vec<i32>>, people: Vec<i32>) -> Vec<i32> {
        let mut starts: Vec<i32> = flowers.iter().map(|f| f[0]).collect();
        let mut ends: Vec<i32> = flowers.iter().map(|f| f[1]).collect();
        starts.sort();
        ends.sort();
        people.iter().map(|&p| Solution::lower_bound(&starts, p + 1) - Solution::lower_bound(&ends, p)).collect()
    }

    fn lower_bound(nums: &Vec<i32>, x: i32) -> i32 {
        let mut left = 0;
        let mut right = nums.len();
        while left < right {
            let mid = left + (right - left) / 2;
            if nums[mid] < x {
                left = mid + 1;
            } else {
                right = mid;
            }
        }
        left as i32
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+m)\log n)$，其中 $n$ 是 $\textit{flowers}$ 的长度，$m$ 是 $\textit{people}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。不计入返回值的空间。

## 方法二：差分

![](https://pic.leetcode.cn/1695696658-skpYuS-2251.jpg)

把 $\textit{flowers}[i]$ 看成是将区间 $[\textit{start}_i,\textit{end}_i]$ 上的每个时间点都增加一朵花。

例如示例 1，我们可以从一个全为 $0$ 的数组开始，把下标在区间 $[1,6],[3,7],[9,12],[4,13]$ 内的数都加一，如果一个下标出现在多个区间内，就多次加一。

加完后，我们得到的数组就是 $a=[0,1,1,2,3,3,3,2,1,2,2,2,2,1]$。

那么 $\textit{answer}[i]$ 就等于 $\textit{person}[i]$ 时间点上有多少朵花，即 $a[\textit{person}[i]]$。

如何快速地「把一个区间内的数都加一」？请看[【算法小课堂】差分数组](https://leetcode.cn/circle/discuss/FfMCgb/)，并至少完成一道练习题。

对于本题，由于元素值很大，需要把数组改成有序集合，或者用哈希表+排序。

第 $i$ 个人能看到的花的数目，即为不超过 $\textit{person}[i]$ 时间点的差分累加值。为了一边累加一边计算答案，我们需要把 $\textit{person}$ 也排序（实际排序的是 $\textit{person}$ 的下标）。

```Python [sol-Python3]
class Solution:
    def fullBloomFlowers(self, flowers: List[List[int]], people: List[int]) -> List[int]:
        diff = Counter()
        for start, end in flowers:
            diff[start] += 1
            diff[end + 1] -= 1
        times = sorted(diff.keys())

        j = s = 0
        for p, i in sorted(zip(people, range(len(people)))):
            while j < len(times) and times[j] <= p:
                s += diff[times[j]]  # 累加不超过 people[i] 的差分值
                j += 1
            people[i] = s  # 从而得到这个时刻花的数量
        return people
```

```java [sol-Java]
class Solution {
    public int[] fullBloomFlowers(int[][] flowers, int[] people) {
        var diff = new TreeMap<Integer, Integer>();
        for (var f : flowers) {
            diff.merge(f[0], 1, Integer::sum);
            diff.merge(f[1] + 1, -1, Integer::sum);
        }

        int n = people.length;
        var id = new Integer[n];
        for (int i = 0; i < n; i++) {
            id[i] = i;
        }
        Arrays.sort(id, (i, j) -> people[i] - people[j]);

        // Java 迭代器没有 peek，可以自己实现一个带 peek 的迭代器
        var it = new PeekingIterator<>(diff.entrySet().iterator());
        int sum = 0;
        for (int i : id) {
            while (it.hasNext() && it.peek().getKey() <= people[i]) {
                sum += it.next().getValue(); // 累加不超过 people[i] 的差分值
            }
            people[i] = sum; // 从而得到这个时刻花的数量
        }
        return people;
    }

    // 284. 顶端迭代器 https://leetcode.cn/problems/peeking-iterator/
    private static class PeekingIterator<E> implements Iterator<E> {
        private final Iterator<E> iterator;
        private E nextElement;

        public PeekingIterator(Iterator<E> iterator) {
            this.iterator = iterator;
            if (iterator.hasNext()) {
                nextElement = iterator.next();
            }
        }

        public E peek() {
            return nextElement;
        }

        @Override
        public E next() {
            E currentElement = nextElement;
            nextElement = iterator.hasNext() ? iterator.next() : null;
            return currentElement;
        }

        @Override
        public boolean hasNext() {
            return nextElement != null;
        }
    }
}
```

```C++ [sol-C++]
class Solution {
public:
    vector<int> fullBloomFlowers(vector<vector<int>> &flowers, vector<int> &people) {
        map<int, int> diff;
        for (auto &f : flowers) {
            diff[f[0]]++;
            diff[f[1] + 1]--;
        }

        int n = people.size();
        vector<int> id(n);
        iota(id.begin(), id.end(), 0); // id[i] = i
        sort(id.begin(), id.end(), [&](int i, int j) { return people[i] < people[j]; });

        auto it = diff.begin();
        int sum = 0;
        for (int i : id) {
            while (it != diff.end() && it->first <= people[i])
                sum += it++->second; // 累加不超过 people[i] 的差分值
            people[i] = sum; // 从而得到这个时刻花的数量
        }
        return people;
    }
};
```

```go [sol-Go]
func fullBloomFlowers(flowers [][]int, people []int) []int {
	diff := map[int]int{}
	for _, f := range flowers {
		diff[f[0]]++
		diff[f[1]+1]--
	}

	n := len(diff)
	times := make([]int, 0, n)
	for t := range diff {
		times = append(times, t)
	}
	sort.Ints(times)

	id := make([]int, len(people))
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return people[id[i]] < people[id[j]] })

	j, sum := 0, 0
	for _, i := range id {
		for ; j < n && times[j] <= people[i]; j++ {
			sum += diff[times[j]] // 累加不超过 people[i] 的差分值
		}
		people[i] = sum // 从而得到这个时刻花的数量
	}
	return people
}
```

```js [sol-JavaScript]
var fullBloomFlowers = function (flowers, people) {
    const diff = new Map();
    for (const [start, end] of flowers) {
        diff.set(start, (diff.get(start) ?? 0) + 1);
        diff.set(end + 1, (diff.get(end + 1) ?? 0) - 1);
    }
    const times = [...diff.keys()].sort((a, b) => a - b);

    const id = [...people.keys()].sort((i, j) => people[i] - people[j]);
    let j = 0, sum = 0;
    for (const i of id) {
        while (j < times.length && times[j] <= people[i]) {
            sum += diff.get(times[j++]); // 累加不超过 people[i] 的差分值
        }
        people[i] = sum; // 从而得到这个时刻花的数量
    }
    return people;
};
```

```rust [sol-Rust]
use std::collections::BTreeMap;

impl Solution {
    pub fn full_bloom_flowers(flowers: Vec<Vec<i32>>, people: Vec<i32>) -> Vec<i32> {
        let mut diff = BTreeMap::new();
        for f in &flowers {
            *diff.entry(f[0]).or_insert(0) += 1;
            *diff.entry(f[1] + 1).or_insert(0) -= 1;
        }

        let n = people.len();
        let mut id: Vec<usize> = (0..n).collect();
        id.sort_by(|&i, &j| people[i].cmp(&people[j]));

        let mut ans = vec![0; n];
        let mut it = diff.iter().peekable();
        let mut sum = 0;
        for &i in &id {
            while let Some((&t, &d)) = it.peek() {
                if t > people[i] {
                    break;
                }
                sum += d; // 累加不超过 people[i] 的差分值
                it.next();
            }
            ans[i] = sum; // 从而得到这个时刻花的数量
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + m\log m)$，其中 $n$ 是 $\textit{flowers}$ 的长度，$m$ 是 $\textit{people}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
