首先，两个数组中都有的数（交集），不需要参与交换，都去掉。去掉后，每个剩余数字在两个数组中的总出现次数必须为偶数，这样才能均分。

在示例 1 中，两个数组的交集是 $[4,2]$，去掉后，两个数组分别为 $[2,2]$ 和 $[1,1]$。选择 $2$ 和 $1$ 交换，就能让两个数组的数字及其出现次数相同。所以只需交换一次，交换成本为 $\min(2,1) = 1$，答案是 $1$。

对于数组中的相同元素，我们会选择其中一半元素，交换到另一个数组；而另一半保留，不参与交换。**下文只关注需要交换的那一半元素**。在示例 1 中，两个数组的交集是 $[4,2]$，去掉后，两个数组分别为 $[2,2]$ 和 $[1,1]$。只看需要交换的元素，两个数组分别为 $[2]$ 和 $[1]$。

把去掉交集，再去掉相同元素的一半后的两个数组分别称作 $a$ 和 $b$。这两个数组中的每个数都需要交换。

如何让交换成本尽量小？考虑 $a$ 中最小的数，找 $b$ 中的哪个数交换，是最合适的？

交换成本是 $\min(a_i,b_j)$，假如 $\min(a_i,b_j) = a_i$，那么看上去，$b_j$ 越大越好？所以应该用小的数和大的数交换？

比如 $a=[1,4]$，$b=[3,2]$，把 $1$ 和 $3$ 交换，$4$ 和 $2$ 交换，交换成本之和为 $\min(1,3) + \min(4,2) = 1+2=3$。如果把 $1$ 和 $2$ 交换，$4$ 和 $3$ 交换，交换成本之和为 $\min(1,2) + \min(4,3) = 1+3=4$，不是最优的。

**猜想**：在每个数至多交换一次的前提下，最优交换方式是把 $a$ **从小到大**排序，$b$ **从大到小**排序，然后交换 $a_i$ 和 $b_i$。

**证明**：考察 $a$ 中的两个数 $p,q$ 和 $b$ 中的两个数 $s,t$，满足 $p\le q$ 且 $s\ge t$。

方案一：$p$ 与 $s$ 交换，$q$ 与 $t$ 交换，交换成本之和为 $S_1 = \min(p,s) + \min(q,t)$。

方案二：$p$ 与 $t$ 交换，$q$ 与 $s$ 交换，交换成本之和为 $S_2 = \min(p,t) + \min(q,s)$。

分类讨论：

- 如果 $p\le t$，那么 $p\le t\le s$，所以 $S_1 = p + \min(q,t)$，$S_2 = p + \min(q,s)$。由于 $t\le s$，所以 $\min(q,t)\le \min(q,s)$，所以 $S_1\le S_2$。
- 如果 $p > t$，那么 $t < p \le q$，所以 $S_1 = \min(p,s) + t$，$S_2 = t + \min(q,s)$。由于 $p\le q$，所以 $\min(p,s)\le \min(q,s)$，所以 $S_1\le S_2$。

所以 $S_1\le S_2$ 恒成立，方案一更好。

根据交换论证法（见 [贪心题单](https://leetcode.cn/circle/discuss/g6KTKL/) §1.7 节的介绍），猜想的交换方式是最优的。 $\square$

但是，题目没说每个数至多交换一次，我们还有一种方案。

设原数组 $\textit{basket}_1$ 和 $\textit{basket}_2$ 的最小值为 $\textit{mn}$。把 $\textit{mn}$ 当作中介，我们可以把 $a_i$ 和 $b_i$ 分别与 $\textit{mn}$ 交换一次，也能完成 $a_i$ 和 $b_i$ 的交换。注意 $\textit{mn}$ 一共交换了两次，在两个数组之间来回跳了一次，所以交换前后，$\textit{mn}$ 所在数组是不变的。

所以，交换 $a_i$ 和 $b_i$ 的实际最小代价为

$$
\min(a_i, b_i, 2\cdot \textit{mn})
$$

累加上式，即为答案。

## 答疑

**问**：为什么 $a$ 和 $b$ 的长度一定相同？

**答**：看代码，最上面的循环会执行 $n$ 次加一和 $n$ 次减一，所以 $\textit{cnt}$ 中的正数 $c$ 之和（除以 $2$ 后是 $a$ 的长度），等于负数 $c$ 之和的绝对值（除以 $2$ 后是 $b$ 的长度）。所以 $a$ 和 $b$ 的长度一定相同。

## 优化前

```py [sol-Python3]
class Solution:
    def minCost(self, basket1: List[int], basket2: List[int]) -> int:
        cnt = defaultdict(int)
        for x, y in zip(basket1, basket2):
            cnt[x] += 1
            cnt[y] -= 1  # 交集元素互相抵消

        a, b = [], []
        for x, c in cnt.items():
            if c % 2:  # 奇数无法均分
                return -1
            # 剩余元素的一半放入 a 或者 b
            if c > 0:
                a.extend([x] * (c // 2))
            else:
                b.extend([x] * (-c // 2))

        a.sort()
        b.sort(reverse=True)
        mn = min(cnt)  # 中介

        return sum(min(x, y, mn * 2) for x, y in zip(a, b))
```

```java [sol-Java]
class Solution {
    public long minCost(int[] basket1, int[] basket2) {
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int i = 0; i < basket1.length; i++) {
            cnt.merge(basket1[i], 1, Integer::sum);  // cnt[basket1[i]]++
            cnt.merge(basket2[i], -1, Integer::sum); // cnt[basket2[i]]--
            // 交集元素在这个过程中互相抵消
        }

        List<Integer> a = new ArrayList<>();
        List<Integer> b = new ArrayList<>();
        int mn = Integer.MAX_VALUE; // 中介
        for (Map.Entry<Integer, Integer> e : cnt.entrySet()) {
            int x = e.getKey();
            int c = e.getValue();
            if (c % 2 != 0) { // 奇数无法均分
                return -1;
            }
            mn = Math.min(mn, x);
            // 剩余元素的一半放入 a 或者 b
            if (c > 0) {
                for (int i = 0; i < c / 2; i++) {
                    a.add(x);
                }
            } else {
                for (int i = 0; i < -c / 2; i++) {
                    b.add(x);
                }
            }
        }

        Collections.sort(a);
        b.sort(Collections.reverseOrder());

        long ans = 0;
        for (int i = 0; i < a.size(); i++) {
            ans += Math.min(Math.min(a.get(i), b.get(i)), mn * 2); // 累加最小交换代价
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minCost(vector<int>& basket1, vector<int>& basket2) {
        unordered_map<int, int> cnt;
        for (int i = 0; i < basket1.size(); i++) {
            cnt[basket1[i]]++;
            cnt[basket2[i]]--; // 交集元素互相抵消
        }

        vector<int> a, b;
        int mn = INT_MAX;
        for (auto [x, c] : cnt) {
            if (c % 2 != 0) { // 奇数无法均分
                return -1;
            }
            mn = min(mn, x);
            // 剩余元素的一半放入 a 或者 b
            if (c > 0) {
                for (int i = 0; i < c / 2; i++) {
                    a.push_back(x);
                }
            } else {
                for (int i = 0; i < -c / 2; i++) {
                    b.push_back(x);
                }
            }
        }

        ranges::sort(a);
        ranges::sort(b, greater());

        long long ans = 0;
        for (int i = 0; i < a.size(); i++) {
            ans += min({a[i], b[i], mn * 2}); // 累加最小交换代价
        }
        return ans;
    }
};
```

```go [sol-Go]
func minCost(basket1, basket2 []int) (ans int64) {
	cnt := map[int]int{}
	for i, x := range basket1 {
		cnt[x]++
		cnt[basket2[i]]-- // 交集元素互相抵消
	}

	var a, b []int
	mn := math.MaxInt // 中介
	for x, c := range cnt {
		if c%2 != 0 { // 奇数无法均分
			return -1
		}
		mn = min(mn, x)
		// 剩余元素的一半放入 a 或者 b
		if c > 0 {
			for range c / 2 {
				a = append(a, x)
			}
		} else {
			for range -c / 2 {
				b = append(b, x)
			}
		}
	}

	slices.Sort(a)
	slices.SortFunc(b, func(a, b int) int { return b - a })

	for i, x := range a {
		ans += int64(min(x, b[i], mn*2)) // 累加最小交换代价
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

```js [sol-JavaScript]
var minCost = function(basket1, basket2) {
    const cnt = new Map();
    for (let i = 0; i < basket1.length; i++) {
        cnt.set(basket1[i], (cnt.get(basket1[i]) ?? 0) + 1);
        cnt.set(basket2[i], (cnt.get(basket2[i]) ?? 0) - 1); // 交集元素互相抵消
    }

    const a = [], b = [];
    let mn = Infinity; // 中介
    for (const [x, c] of cnt.entries()) {
        if (c % 2 !== 0) { // 奇数无法均分
            return -1;
        }
        mn = Math.min(mn, x);
        // 剩余元素的一半放入 a 或者 b
        if (c > 0) {
            for (let i = 0; i < c / 2; i++) {
                a.push(x);
            }
        } else {
            for (let i = 0; i < -c / 2; i++) {
                b.push(x);
            }
        }
    }

    a.sort((x, y) => x - y);
    b.sort((x, y) => y - x);

    let ans = 0;
    for (let i = 0; i < a.length; i++) {
        ans += Math.min(a[i], b[i], mn * 2); // 累加最小交换代价
    }
    return ans;
};
```

```rust [sol-Rust]
use std::collections::HashMap;

impl Solution {
    pub fn min_cost(basket1: Vec<i32>, basket2: Vec<i32>) -> i64 {
        let mut cnt = HashMap::new();
        for (x, y) in basket1.into_iter().zip(basket2) {
            *cnt.entry(x).or_insert(0) += 1;
            *cnt.entry(y).or_insert(0) -= 1; // 交集元素互相抵消
        }

        let mut a = vec![];
        let mut b = vec![];
        let mut mn = i32::MAX; // 中介
        for (x, c) in cnt {
            if c % 2 != 0 { // 奇数无法均分
                return -1;
            }
            mn = mn.min(x);
            // 剩余元素的一半放入 a 或者 b
            if c > 0 {
                for _ in 0..(c / 2) {
                    a.push(x);
                }
            } else {
                for _ in 0..(-c / 2) {
                    b.push(x);
                }
            }
        }

        a.sort_unstable();
        b.sort_unstable_by_key(|x| -x);

        // 累加最小交换代价
        a.into_iter().zip(b).map(|(x, y)| x.min(y).min(mn * 2) as i64).sum()
    }
}
```

## 优化

考察 $\min(a_i,b_i)$ 的值来自 $a$ 和 $b$ 中的哪些数。

由于 $a$ 是递增的，$b$ 是递减的，所以随着 $i$ 的变大，$\min(a_i,b_i)$ 的值会从 $a_i$ 变成 $b_i$。

换句话说，$\min(a_i,b_i)$ 的值来自 $a$ 的前缀与 $b$ 的后缀。且前缀与后缀的长度之和恰好等于 $a$ 的长度 $m$。

设这 $m$ 个数为 $A=[a_0,a_1,\ldots,a_i,b_{i+1},b_{i+2},\ldots,b_{m-1}]$。

**定理**：$A$ 中的这 $m$ 个数，也是 $a+b$ 中最小的 $m$ 个数。

**证明**：由于 $a_{i+1}\ge b_{i+1}\ge b_{i+2}\ge \cdots\ge b_{m-1}$ 且 $a_{i+1} \ge a_{i}\ge a_{i-1}\ge \cdots \ge a_0$，所以 $a_{i+1}$ 大于等于 $A$ 中每个数，又由于 $a_{i+2},a_{i+3},\ldots,a_{m-1}$ 都大于等于 $a_{i+1}$，所以 $a_{i+2},a_{i+3},\ldots,a_{m-1}$ 也都大于等于 $A$ 中每个数。对于 $b_0,b_1,\ldots,b_i$ 同理。所以 $A$ 中的这 $m$ 个数，也是 $a+b$ 中最小的 $m$ 个数。 

因此，直接把 $a$ 和 $b$ 拼起来，从小到大排序后，只需遍历前一半的数，计算最小代价。甚至可以用快速选择算法做到 $\mathcal{O}(n)$，见 C++ 代码。

```py [sol-Python3]
class Solution:
    def minCost(self, basket1: List[int], basket2: List[int]) -> int:
        cnt = defaultdict(int)
        for x, y in zip(basket1, basket2):
            cnt[x] += 1
            cnt[y] -= 1

        a = []
        for x, c in cnt.items():
            if c % 2:
                return -1
            a.extend([x] * (abs(c) // 2))  # 剩余元素全部加到 a 中

        a.sort()
        mn = min(cnt)

        return sum(min(x, mn * 2) for x in a[:len(a) // 2])
```

```java [sol-Java]
class Solution {
    public long minCost(int[] basket1, int[] basket2) {
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int i = 0; i < basket1.length; i++) {
            cnt.merge(basket1[i], 1, Integer::sum);  // cnt[basket1[i]]++
            cnt.merge(basket2[i], -1, Integer::sum); // cnt[basket2[i]]--
        }

        List<Integer> a = new ArrayList<>();
        int mn = Integer.MAX_VALUE;
        for (Map.Entry<Integer, Integer> e : cnt.entrySet()) {
            int x = e.getKey();
            int c = e.getValue();
            if (c % 2 != 0) {
                return -1;
            }
            mn = Math.min(mn, x);
            for (c = Math.abs(c) / 2; c > 0; c--) {
                a.add(x);
            }
        }

        Collections.sort(a);

        long ans = 0;
        for (int i = 0; i < a.size() / 2; i++) {
            ans += Math.min(a.get(i), mn * 2);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minCost(vector<int>& basket1, vector<int>& basket2) {
        unordered_map<int, int> cnt;
        for (int i = 0; i < basket1.size(); i++) {
            cnt[basket1[i]]++;
            cnt[basket2[i]]--;
        }

        vector<int> a;
        int mn = INT_MAX;
        for (auto [x, c] : cnt) {
            if (c % 2) {
                return -1;
            }
            mn = min(mn, x);
            for (c = abs(c) / 2; c > 0; c--) {
                a.push_back(x);
            }
        }

        ranges::nth_element(a, a.begin() + a.size() / 2); // 快速选择

        long long ans = 0;
        for (int i = 0; i < a.size() / 2; i++) {
            ans += min(a[i], mn * 2);
        }
        return ans;
    }
};
```

```go [sol-Go]
func minCost(basket1, basket2 []int) (ans int64) {
	cnt := map[int]int{}
	for i, x := range basket1 {
		cnt[x]++
		cnt[basket2[i]]--
	}

	a := []int{}
	mn := math.MaxInt
	for x, c := range cnt {
		if c%2 != 0 {
			return -1
		}
		mn = min(mn, x)
		for range abs(c) / 2 {
			a = append(a, x)
		}
	}

	slices.Sort(a)

	for _, x := range a[:len(a)/2] {
		ans += int64(min(x, mn*2))
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

```js [sol-JavaScript]
var minCost = function(basket1, basket2) {
    const cnt = new Map();
    for (let i = 0; i < basket1.length; i++) {
        cnt.set(basket1[i], (cnt.get(basket1[i]) ?? 0) + 1);
        cnt.set(basket2[i], (cnt.get(basket2[i]) ?? 0) - 1);
    }

    const a = [];
    let mn = Infinity;
    for (const [x, c] of cnt.entries()) {
        if (c % 2 !== 0) {
            return -1;
        }
        mn = Math.min(mn, x);
        for (let k = Math.abs(c) / 2; k > 0; k--) {
            a.push(x);
        }
    }

    a.sort((x, y) => x - y);

    let ans = 0;
    for (let i = 0; i < a.length / 2; i++) {
        ans += Math.min(a[i], mn * 2);
    }
    return ans;
};
```

```rust [sol-Rust]
use std::collections::HashMap;

impl Solution {
    pub fn min_cost(basket1: Vec<i32>, basket2: Vec<i32>) -> i64 {
        let mut cnt: HashMap<i32, i32> = HashMap::new();
        for (x, y) in basket1.into_iter().zip(basket2) {
            *cnt.entry(x).or_insert(0) += 1;
            *cnt.entry(y).or_insert(0) -= 1;
        }

        let mut a = vec![];
        let mut mn = i32::MAX;
        for (x, c) in cnt {
            if c % 2 != 0 {
                return -1;
            }
            mn = mn.min(x);
            for _ in 0..c.abs() / 2 {
                a.push(x);
            }
        }

        if a.is_empty() {
            return 0;
        }

        let k = a.len() / 2;
        a.select_nth_unstable(k); // 快速选择

        a[..k].iter().map(|&x| x.min(mn * 2) as i64).sum()
    }
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$ 或 $\mathcal{O}(n)$，其中 $n$ 为 $\textit{basket}_1$ 的长度。瓶颈在排序上。用快速选择算法可以做到 $\mathcal{O}(n)$，见 C++ 代码。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

下面贪心题单的「**§1.3 双序列配对**」和「**§1.7 交换论证法**」。

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
