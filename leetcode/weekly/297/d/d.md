## 分析

什么样的一对字符串无法交换首字母？

示例 1 中的 $\texttt{coffee}$ 和 $\texttt{time}$，虽然这样两个字符串完全不一样，但如果交换了 $\texttt{coffee}$ 和 $\texttt{time}$ 的首字母，会得到字符串 $\texttt{toffee}$，它在数组 $\textit{ideas}$ 中，不符合题目要求。

又例如 $\textit{ideas}=[\texttt{aa},\texttt{ab},\texttt{ac},\texttt{bc},\texttt{bd},\texttt{be}]$，将其分成两组：

- 第一组：$\texttt{aa},\texttt{ab},\texttt{ac}$。
- 第二组：$\texttt{bc},\texttt{bd},\texttt{be}$。

其中第一组内的字符串是不能交换首字母的，因为交换后字符串不变，必然在 $\textit{ideas}$ 中。第二组也同理。

考虑交换第一组的字符串和第二组的字符串，哪些是可以交换首字母的，哪些是不能交换首字母的？

- 第一组的 $\texttt{ac}$ 无法和第二组的任何字符串交换，因为交换后会得到 $\texttt{bc}$，它在 $\textit{ideas}$ 中。
- 第二组的 $\texttt{bc}$ 无法和第一组的任何字符串交换，因为交换后会得到 $\texttt{ac}$，它在 $\textit{ideas}$ 中。
- 其余字符串对可以交换首字母。

上面的分析立刻引出了如下方法。

## 方法一：按照首字母分组

按照首字母，把 $\textit{ideas}$ 分成（至多）$26$ 组字符串。

例如 $\textit{ideas}=[\texttt{aa},\texttt{ab},\texttt{ac},\texttt{bc},\texttt{bd},\texttt{be}]$ 分成如下两组（**只记录去掉首字母后的字符串**）：

- $A$ 组（集合）：$\{\texttt{a},\texttt{b},\texttt{c}\}$。
- $B$ 组（集合）：$\{\texttt{c},\texttt{d},\texttt{e}\}$。

分组后：

1. 从 $A$ 中选一个不等于 $\texttt{c}$ 的字符串，这有 $2$ 种选法。
2. 从 $B$ 中选一个不等于 $\texttt{c}$ 的字符串，这有 $2$ 种选法。
3. 考虑两个字符串的先后顺序（谁在左谁在右），有 $2$ 种方法。

根据**乘法原理**，有 $2\times 2\times 2=8$ 对符合要求的字符串。

由于无法选交集中的字符串，一般地，从 $A$ 和 $B$ 中可以选出

$$
2\cdot(|A|-|A\cap B|)\cdot(|B|-|A\cap B|)
$$

对符合要求的字符串。其中 $|S|$ 表示集合 $S$ 的大小。

枚举所有组对，计算上式，累加到答案中。

```py [sol-Python3]
class Solution:
    def distinctNames(self, ideas: List[str]) -> int:
        groups = defaultdict(set)
        for s in ideas:
            groups[s[0]].add(s[1:])  # 按照首字母分组

        ans = 0
        for a, b in combinations(groups.values(), 2):  # 枚举所有组对
            m = len(a & b)  # 交集的大小
            ans += (len(a) - m) * (len(b) - m)
        return ans * 2  # 乘 2 放到最后
```

```java [sol-Java]
class Solution {
    public long distinctNames(String[] ideas) {
        Set<String>[] groups = new HashSet[26];
        Arrays.setAll(groups, i -> new HashSet<>());
        for (String s : ideas) {
            groups[s.charAt(0) - 'a'].add(s.substring(1)); // 按照首字母分组
        }

        long ans = 0;
        for (int a = 1; a < 26; a++) { // 枚举所有组对
            for (int b = 0; b < a; b++) {
                int m = 0; // 交集的大小
                for (String s : groups[a]) {
                    if (groups[b].contains(s)) {
                        m++;
                    }
                }
                ans += (long) (groups[a].size() - m) * (groups[b].size() - m);
            }
        }
        return ans * 2; // 乘 2 放到最后
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long distinctNames(vector<string>& ideas) {
        unordered_set<string> groups[26];
        for (auto& s : ideas) {
            groups[s[0] - 'a'].insert(s.substr(1)); // 按照首字母分组
        }

        long long ans = 0;
        for (int a = 1; a < 26; a++) { // 枚举所有组对
            for (int b = 0; b < a; b++) {
                int m = 0; // 交集的大小
                for (auto& s : groups[a]) {
                    m += groups[b].count(s);
                }
                ans += (long long) (groups[a].size() - m) * (groups[b].size() - m);
            }
        }
        return ans * 2; // 乘 2 放到最后
    }
};
```

```go [sol-Go]
func distinctNames(ideas []string) (ans int64) {
    group := [26]map[string]bool{}
    for i := range group {
        group[i] = map[string]bool{}
    }
    for _, s := range ideas {
        group[s[0]-'a'][s[1:]] = true // 按照首字母分组
    }

    for i, a := range group { // 枚举所有组对
        for _, b := range group[:i] {
            m := 0 // 交集的大小
            for s := range a {
                if b[s] {
                    m++
                }
            }
            ans += int64(len(a)-m) * int64(len(b)-m)
        }
    }
    return ans * 2 // 乘 2 放到最后
}
```

```js [sol-JavaScript]
var distinctNames = function(ideas) {
    const groups = Array.from({ length: 26 }, () => new Set());
    for (const s of ideas) {
        groups[s.charCodeAt(0) - 'a'.charCodeAt(0)].add(s.slice(1)); // 按照首字母分组
    }

    let ans = 0;
    for (let a = 1; a < 26; a++) { // 枚举所有组对
        for (let b = 0; b < a; b++) {
            let m = 0; // 交集的大小
            for (const s of groups[a]) {
                if (groups[b].has(s)) {
                    m++;
                }
            }
            ans += (groups[a].size - m) * (groups[b].size - m);
        }
    }
    return ans * 2; // 乘 2 放到最后
};
```

```rust [sol-Rust]
use std::collections::HashSet;

impl Solution {
    pub fn distinct_names(ideas: Vec<String>) -> i64 {
        let mut groups = vec![HashSet::new(); 26];
        for s in ideas {
            groups[(s.as_bytes()[0] - b'a') as usize].insert(s[1..].to_string()); // 按照首字母分组
        }

        let mut ans = 0i64;
        for a in 1..26 { // 枚举所有组对
            for b in 0..a {
                let m = groups[a].iter().filter(|&s| groups[b].contains(s)).count(); // 交集的大小
                ans += (groups[a].len() - m) as i64 * (groups[b].len() - m) as i64;
            }
        }
        ans * 2 // 乘 2 放到最后
    }
}
```

#### 复杂度分析

- 时间复杂度：$O(nm|\Sigma|)$，其中 $n$ 为 $\textit{ideas}$ 的长度，$m\le 10$ 为单个字符串的长度，$|\Sigma|=26$ 是字符集合的大小。注意枚举组对的逻辑看上去是 $O(nm|\Sigma|^2)$ 的，但去掉内层 $\mathcal{O}(|\Sigma|)$ 的循环后，剩余循环相当于把 $\textit{ideas}$ 遍历了一遍，是 $\mathcal{O}(nm)$，所以总的时间复杂度是 $O(nm|\Sigma|)$。
- 空间复杂度：$O(nm+|\Sigma|)$。

## 方法二：按照后缀分组

下文把去掉首字母后的剩余部分称作后缀。

横看成岭侧成峰，换一个角度计算交集大小 $|A\cap B|$。

在遍历 $\textit{ideas}=[\texttt{aa},\texttt{ab},\texttt{ac},\texttt{bc},\texttt{bd},\texttt{be}]$ 的过程中，当我们遍历到 $\texttt{bc}$ 时，发现之前遍历过一个后缀也为 $\texttt{c}$ 的字符串 $\texttt{ac}$，这就对交集大小 $|A\cap B|$ 产生了 $1$ 的贡献，也就是交集大小 $|A\cap B|$ 增加 $1$。

具体来说，在遍历 $\textit{ideas}$ 的过程中，维护如下信息：

1. 集合大小 $\textit{size}[a]$。遍历到 $s=\textit{ideas}[i]$ 时，把 $\textit{size}[s[0]]$ 加一。
2. 交集大小 $\textit{intersection}[a][b]$。遍历到 $s=\textit{ideas}[i]$ 时，设 $b=s[0]$，把 $\textit{intersection}[a][b]$ 和 $\textit{intersection}[b][a]$ 加一，其中 $a$ 是和 $s$ 同后缀的其他字符串的首字母。
3. 为了计算有哪些字符串和 $s$ 有着相同的后缀，用一个哈希表 $\textit{groups}$ 维护，key 是后缀，value 是后缀对应的首字母列表。注意题目保证所有字符串互不相同。

代码实现时，可以用把首字母列表压缩成一个二进制数，原理见 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```py [sol-Python3]
class Solution:
    def distinctNames(self, ideas: List[str]) -> int:
        size = [0] * 26  # 集合大小
        intersection = [[0] * 26 for _ in range(26)]  # 交集大小
        groups = defaultdict(list)  # 后缀 -> 首字母列表
        for s in ideas:
            b = ord(s[0]) - ord('a')
            size[b] += 1  # 增加集合大小
            g = groups[s[1:]]
            for a in g:  # a 是和 s 有着相同后缀的字符串的首字母
                intersection[a][b] += 1  # 增加交集大小
                intersection[b][a] += 1
            g.append(b)

        ans = 0
        for a in range(1, 26):  # 枚举所有组对
            for b in range(a):
                m = intersection[a][b]
                ans += (size[a] - m) * (size[b] - m)
        return ans * 2  # 乘 2 放到最后
```

```java [sol-Java]
class Solution {
    public long distinctNames(String[] ideas) {
        int[] size = new int[26]; // 集合大小
        int[][] intersection = new int[26][26]; // 交集大小
        Map<String, Integer> groups = new HashMap<>(); // 后缀 -> 首字母
        for (String s : ideas) {
            int b = s.charAt(0) - 'a';
            size[b]++; // 增加集合大小
            String suffix = s.substring(1);
            int mask = groups.getOrDefault(suffix, 0);
            groups.put(suffix, mask | 1 << b); // 把 b 加到 mask 中
            for (int a = 0; a < 26; a++) { // a 是和 s 有着相同后缀的字符串的首字母
                if ((mask >> a & 1) > 0) { // a 在 mask 中
                    intersection[b][a]++; // 增加交集大小
                    intersection[a][b]++;
                }
            }
        }

        long ans = 0;
        for (int a = 1; a < 26; a++) { // 枚举所有组对
            for (int b = 0; b < a; b++) {
                int m = intersection[a][b];
                ans += (long) (size[a] - m) * (size[b] - m);
            }
        }
        return ans * 2; // 乘 2 放到最后
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long distinctNames(vector<string>& ideas) {
        int size[26]{}; // 集合大小
        int intersection[26][26]{}; // 交集大小
        unordered_map<string, int> groups; // 后缀 -> 首字母
        for (auto& s : ideas) {
            int b = s[0] - 'a';
            size[b]++; // 增加集合大小
            auto suffix = s.substr(1);
            int mask = groups[suffix];
            groups[suffix] = mask | 1 << b; // 把 b 加到 mask 中
            for (int a = 0; a < 26; a++) { // a 是和 s 有着相同后缀的字符串的首字母
                if (mask >> a & 1) { // a 在 mask 中
                    intersection[b][a]++; // 增加交集大小
                    intersection[a][b]++;
                }
            }
        }

        long long ans = 0;
        for (int a = 1; a < 26; a++) { // 枚举所有组对
            for (int b = 0; b < a; b++) {
                int m = intersection[a][b];
                ans += (long long) (size[a] - m) * (size[b] - m);
            }
        }
        return ans * 2; // 乘 2 放到最后
    }
};
```

```go [sol-Go]
func distinctNames(ideas []string) (ans int64) {
    size := [26]int{} // 集合大小
    intersection := [26][26]int{} // 交集大小
    groups := map[string]int{} // 后缀 -> 首字母
    for _, s := range ideas {
        b := s[0] - 'a'
        size[b]++ // 增加集合大小
        suffix := s[1:]
        mask := groups[suffix]
        groups[suffix] = mask | 1<<b // 把 b 加到 mask 中
        for a := 0; a < 26; a++ { // a 是和 s 有着相同后缀的字符串的首字母
            if mask>>a&1 > 0 { // a 在 mask 中
                intersection[b][a]++ // 增加交集大小
                intersection[a][b]++
            }
        }
    }

    for a := 1; a < 26; a++ { // 枚举所有组对
        for b := 0; b < a; b++ {
            m := intersection[a][b]
            ans += int64(size[a]-m) * int64(size[b]-m)
        }
    }
    return ans * 2 // 乘 2 放到最后
}
```

```js [sol-JavaScript]
var distinctNames = function(ideas) {
    const size = Array(26).fill(0); // 集合大小
    const intersection = Array.from({ length: 26 }, () => Array(26).fill(0)); // 交集大小
    const groups = {}; // 后缀 -> 首字母
    for (let s of ideas) {
        const b = s.charCodeAt(0) - 'a'.charCodeAt(0);
        size[b]++; // 增加集合大小
        const suffix = s.slice(1);
        const mask = groups[suffix] ?? 0;
        groups[suffix] = mask | 1 << b; // 把 b 加到 mask 中
        for (let a = 0; a < 26; a++) { // a 是和 s 有着相同后缀的字符串的首字母
            if (mask >> a & 1) { // a 在 mask 中
                intersection[b][a]++; // 增加交集大小
                intersection[a][b]++;
            }
        }
    }

    let ans = 0;
    for (let a = 1; a < 26; a++) { // 枚举所有组对
        for (let b = 0; b < a; b++) {
            const m = intersection[a][b];
            ans += (size[a] - m) * (size[b] - m);
        }
    }
    return ans * 2; // 乘 2 放到最后
};
```

```rust [sol-Rust]
use std::collections::HashMap;

impl Solution {
    pub fn distinct_names(ideas: Vec<String>) -> i64 {
        let mut size = vec![0; 26]; // 集合大小
        let mut intersection = vec![vec![0; 26]; 26]; // 交集大小
        let mut groups = HashMap::new(); // 后缀 -> 首字母
        for s in ideas {
            let b = (s.as_bytes()[0] - b'a') as usize;
            size[b] += 1; // 增加集合大小
            let suffix = &s[1..];
            let mask = *groups.get(suffix).unwrap_or(&0);
            groups.insert(suffix.to_string(), mask | 1 << b); // 把 b 加到 mask 中
            for a in 0..26 { // a 是和 s 有着相同后缀的首字母
                if (mask >> a & 1) > 0 { // a 在 mask 中
                    intersection[b][a] += 1; // 增加交集大小
                    intersection[a][b] += 1;
                }
            }
        }

        let mut ans = 0i64;
        for a in 1..26 { // 枚举所有组对
            for b in 0..a {
                let m = intersection[a][b];
                ans += (size[a] - m) as i64 * (size[b] - m) as i64;
            }
        }
        ans * 2 // 乘 2 放到最后
    }
}
```

#### 复杂度分析

- 时间复杂度：$O(n(m+|\Sigma|) + |\Sigma|^2)$，其中 $n$ 为 $\textit{ideas}$ 的长度，$m\le 10$ 为单个字符串的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$O(nm+|\Sigma|^2)$。

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
