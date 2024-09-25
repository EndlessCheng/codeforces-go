## 方法一：枚举因子

### 分析

为方便描述，把 $\textit{nums}_1$ 和 $\textit{nums}_2$ 记作 $a$ 和 $b$。

$a[i]$ 能被 $b[j]\cdot k$ 整除，等价于 $a[i]$ 是 $k$ 的倍数且 $\dfrac{a[i]}{k}$ 能被 $b[j]$ 整除。

也就是说，$\dfrac{a[i]}{k}$ 有一个因子 $d$ 等于 $b[j]$。

### 算法

1. 遍历 $a$，枚举 $\dfrac{a[i]}{k}$ 的所有因子，统计到哈希表 $\textit{cnt}$ 中。比如遍历完后 $\textit{cnt}[3] = 5$，说明有 $5$ 个 $\dfrac{a[i]}{k}$ 可以被 $3$ 整除，等价于有 $5$ 个 $a[i]$ 可以被 $3\cdot k$ 整除。
2. 遍历 $b$，把 $\textit{cnt}[b[j]]$ 加入答案。例如 $b[j]=3$，那么就找到了 $\textit{cnt}[3]$ 个优质数对。

枚举因子的技巧请看 [视频讲解](https://www.bilibili.com/video/BV17t421N7L6/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def numberOfPairs(self, nums1: List[int], nums2: List[int], k: int) -> int:
        cnt = defaultdict(int)
        for x in nums1:
            if x % k:
                continue
            x //= k
            for d in range(1, isqrt(x) + 1):  # 枚举因子
                if x % d:
                    continue
                cnt[d] += 1  # 统计因子
                if d * d < x:
                    cnt[x // d] += 1  # 因子总是成对出现
        return sum(cnt[x] for x in nums2)
```

```java [sol-Java]
class Solution {
    public long numberOfPairs(int[] nums1, int[] nums2, int k) {
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int x : nums1) {
            if (x % k != 0) {
                continue;
            }
            x /= k;
            for (int d = 1; d * d <= x; d++) { // 枚举因子
                if (x % d > 0) {
                    continue;
                }
                cnt.merge(d, 1, Integer::sum); // cnt[d]++
                if (d * d < x) {
                    cnt.merge(x / d, 1, Integer::sum); // cnt[x/d]++
                }
            }
        }

        long ans = 0;
        for (int x : nums2) {
            ans += cnt.getOrDefault(x, 0);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long numberOfPairs(vector<int>& nums1, vector<int>& nums2, int k) {
        unordered_map<int, int> cnt;
        for (int x : nums1) {
            if (x % k) {
                continue;
            }
            x /= k;
            for (int d = 1; d * d <= x; d++) { // 枚举因子
                if (x % d) {
                    continue;
                }
                cnt[d]++; // 统计因子
                if (d * d < x) {
                    cnt[x / d]++; // 因子总是成对出现
                }
            }
        }

        long long ans = 0;
        for (int x : nums2) {
            ans += cnt.contains(x) ? cnt[x] : 0;
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfPairs(nums1, nums2 []int, k int) (ans int64) {
    cnt := map[int]int{}
    for _, x := range nums1 {
        if x%k > 0 {
            continue
        }
        x /= k
        for d := 1; d*d <= x; d++ { // 枚举因子
            if x%d == 0 {
                cnt[d]++ // 统计因子
                if d*d < x {
                    cnt[x/d]++ // 因子总是成对出现
                }
            }
        }
    }

    for _, x := range nums2 {
        ans += int64(cnt[x])
    }
    return
}
```

```js [sol-JavaScript]
var numberOfPairs = function(nums1, nums2, k) {
    const cnt = new Map();
    for (let x of nums1) {
        if (x % k) {
            continue;
        }
        x /= k;
        for (let d = 1; d * d <= x; d++) { // 枚举因子
            if (x % d) {
                continue;
            }
            cnt.set(d, (cnt.get(d) || 0) + 1); // 统计因子
            if (d * d < x) {
                cnt.set(x / d, (cnt.get(x / d) ?? 0) + 1); // 因子总是成对出现
            }
        }
    }

    let ans = 0;
    for (const x of nums2) {
        ans += cnt.get(x) ?? 0;
    }
    return ans;
};
```

```rust [sol-Rust]
use std::collections::HashMap;

impl Solution {
    pub fn number_of_pairs(nums1: Vec<i32>, nums2: Vec<i32>, k: i32) -> i64 {
        let mut cnt = HashMap::new();
        for mut x in nums1 {
            if x % k != 0 {
                continue;
            }
            x /= k;
            let mut d = 1;
            while d * d <= x {
                if x % d == 0 {
                    *cnt.entry(d).or_insert(0) += 1;
                    if d * d < x {
                        *cnt.entry(x / d).or_insert(0) += 1;
                    }
                }
                d += 1;
            }
        }

        nums2.iter().map(|x| *cnt.get(x).unwrap_or(&0) as i64).sum()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt{U/k} + m)$，其中 $n$ 是 $\textit{nums}_1$ 的长度，$m$ 是 $\textit{nums}_2$ 的长度，$U=\max(\textit{nums}_1)$。
- 空间复杂度：$\mathcal{O}(U/k)$。不同因子个数不会超过 $U/k$。

## 方法二：枚举倍数

### 分析

横看成岭侧成峰，我们还可以枚举 $b[j]$ 的倍数。

例如 $b[j]=3$，枚举 $3,6,9,12,\cdots$，统计 $a$ 中有多少个 $\dfrac{a[i]}{k}$ 等于 $3,6,9,12,\cdots$

### 算法

1. 统计 $\dfrac{a[i]}{k}$ 的出现次数，保存到哈希表 $\textit{cnt}_1$ 中。
2. 统计 $b[j]$ 的出现次数（相同 $b[j]$ 无需重复计算），保存到哈希表 $\textit{cnt}_2$ 中。
3. 设 $\textit{cnt}_1$ 中的最大 key 为 $u$。
4. 枚举 $\textit{cnt}_2$ 中的元素 $x$，然后枚举 $x$ 的倍数 $y=x,2x,3x,\cdots$（不超过 $u$），累加 $\textit{cnt}_1[y]$，再乘上 $\textit{cnt}_2[x]$，加入答案。

```py [sol-Python3]
class Solution:
    def numberOfPairs(self, nums1: List[int], nums2: List[int], k: int) -> int:
        cnt1 = Counter(x // k for x in nums1 if x % k == 0)
        if not cnt1:
            return 0

        ans = 0
        u = max(cnt1)
        for x, cnt in Counter(nums2).items():
            s = sum(cnt1[y] for y in range(x, u + 1, x))  # 枚举 x 的倍数
            ans += s * cnt
        return ans
```

```java [sol-Java]
class Solution {
    public long numberOfPairs(int[] nums1, int[] nums2, int k) {
        Map<Integer, Integer> cnt1 = new HashMap<>();
        for (int x : nums1) {
            if (x % k == 0) {
                cnt1.merge(x / k, 1, Integer::sum); // cnt1[x/k]++
            }
        }
        if (cnt1.isEmpty()) {
            return 0;
        }

        Map<Integer, Integer> cnt2 = new HashMap<>();
        for (int x : nums2) {
            cnt2.merge(x, 1, Integer::sum); // cnt2[x]++
        }

        long ans = 0;
        int u = Collections.max(cnt1.keySet());
        for (Map.Entry<Integer, Integer> e : cnt2.entrySet()) {
            int x = e.getKey();
            int cnt = e.getValue();
            int s = 0;
            for (int y = x; y <= u; y += x) { // 枚举 x 的倍数
                if (cnt1.containsKey(y)) {
                    s += cnt1.get(y);
                }
            }
            ans += (long) s * cnt;
        }
        return ans;
    }
}
```

```java [sol-Java 数组]
class Solution {
    public long numberOfPairs(int[] nums1, int[] nums2, int k) {
        int mx1 = 0;
        for (int x : nums1) {
            if (x % k == 0) {
                mx1 = Math.max(mx1, x / k);
            }
        }
        if (mx1 == 0) {
            return 0;
        }

        int[] cnt1 = new int[mx1 + 1];
        for (int x : nums1) {
            if (x % k == 0) {
                cnt1[x / k]++;
            }
        }

        int mx2 = 0;
        for (int x : nums2) {
            mx2 = Math.max(mx2, x);
        }
        int[] cnt2 = new int[mx2 + 1];
        for (int x : nums2) {
            cnt2[x]++;
        }

        long ans = 0;
        for (int x = 1; x <= mx2; x++) {
            if (cnt2[x] == 0) {
                continue;
            }
            int s = 0;
            for (int y = x; y <= mx1; y += x) { // 枚举 x 的倍数
                s += cnt1[y];
            }
            ans += (long) s * cnt2[x];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long numberOfPairs(vector<int>& nums1, vector<int>& nums2, int k) {
        unordered_map<int, int> cnt1;
        for (int x : nums1) {
            if (x % k == 0) {
                cnt1[x / k]++;
            }
        }
        if (cnt1.empty()) {
            return 0;
        }

        unordered_map<int, int> cnt2;
        for (int x : nums2) {
            cnt2[x]++;
        }

        long long ans = 0;
        int u = ranges::max_element(cnt1)->first;
        for (auto& [x, cnt] : cnt2) {
            int s = 0;
            for (int y = x; y <= u; y += x) { // 枚举 x 的倍数
                s += cnt1.contains(y) ? cnt1[y] : 0;
            }
            ans += (long long) s * cnt;
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfPairs(nums1, nums2 []int, k int) (ans int64) {
    cnt1 := map[int]int{}
    u := 0
    for _, x := range nums1 {
        if x%k == 0 {
            u = max(u, x/k)
            cnt1[x/k]++
        }
    }
    if u == 0 {
        return
    }

    cnt2 := map[int]int{}
    for _, x := range nums2 {
        cnt2[x]++
    }

    for x, cnt := range cnt2 {
        s := 0
        for y := x; y <= u; y += x { // 枚举 x 的倍数
            s += cnt1[y]
        }
        ans += int64(s * cnt)
    }
    return
}
```

```js [sol-JavaScript]
const numberOfPairs = function(nums1, nums2, k) {
    const cnt1 = new Map();
    let u = 0;
    for (const x of nums1) {
        if (x % k === 0) {
            u = Math.max(u, x / k);
            cnt1.set(x / k, (cnt1.get(x / k) ?? 0) + 1);
        }
    }
    if (u === 0) {
        return 0;
    }

    const cnt2 = new Map();
    for (const x of nums2) {
        cnt2.set(x, (cnt2.get(x) ?? 0) + 1);
    }

    let ans = 0;
    for (const [x, cnt] of cnt2.entries()) {
        let s = 0;
        for (let y = x; y <= u; y += x) { // 枚举 x 的倍数
            s += cnt1.get(y) ?? 0;
        }
        ans += s * cnt;
    }
    return ans;
}
```

```rust [sol-Rust]
use std::collections::HashMap;

impl Solution {
    pub fn number_of_pairs(nums1: Vec<i32>, nums2: Vec<i32>, k: i32) -> i64 {
        let mut cnt1 = HashMap::new();
        for x in nums1 {
            if x % k == 0 {
                *cnt1.entry(x / k).or_insert(0) += 1;
            }
        }
        if cnt1.is_empty() {
            return 0;
        }

        let mut cnt2 = HashMap::new();
        for x in nums2 {
            *cnt2.entry(x).or_insert(0) += 1;
        }

        let mut ans = 0i64;
        let u = *cnt1.keys().max().unwrap();
        for (x, cnt) in cnt2 {
            let mut s = 0;
            for y in (x..=u).step_by(x as usize) { // 枚举 x 的倍数
                if let Some(&c) = cnt1.get(&y) {
                    s += c;
                }
            }
            ans += s as i64 * cnt as i64;
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m + (U/k)\log m)$，其中 $n$ 是 $\textit{nums}_1$ 的长度，$m$ 是 $\textit{nums}_2$ 的长度，$U=\max(\textit{nums}_1)$。复杂度根据调和级数可得。详细解释请看 [视频讲解](https://www.bilibili.com/video/BV17t421N7L6/)。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
