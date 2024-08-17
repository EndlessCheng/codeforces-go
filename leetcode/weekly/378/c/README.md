由于特殊子串只包含单一字母，我们按照相同字母分组，每组统计相同字母连续出现的长度。例如字符串 aaaabbbabb 分成 aaaa+bbb+a+bb 四组，字母 a 有长度 $4$ 和长度 $1$，字母 b 有长度 $3$ 和长度 $2$。所以字母 a 的长度列表为 $[4,1]$，字母 b 的长度列表为 $[3,2]$。

遍历每个字母对应的长度列表 $a$，把 $a$ 从大到小排序。

有哪些取出三个特殊子串的方法呢？

- 从最长的特殊子串（$a[0]$）中取三个长度均为 $a[0]-2$ 的特殊子串。例如示例 1 的 aaaa 可以取三个 aa。
- 或者，从最长和次长的特殊子串（$a[0],a[1]$）中取三个长度一样的特殊子串：
  - 如果 $a[0]=a[1]$，那么可以取三个长度均为 $a[0]-1$ 的特殊子串。
  - 如果 $a[0]>a[1]$，那么可以取三个长度均为 $a[1]$ 的特殊子串：从最长中取两个，从次长中取一个。
  - 这两种情况可以合并成 $\min(a[0]-1, a[1])$，如果 $a[0]-1 < a[1]$，这只能是第一种情况，因为 $a[0]\ge a[1]$，我们取二者较小值 $a[0]-1$；如果 $a[0]-1\ge a[1]$，即 $a[0] > a[1]$，这是第二种情况，我们也取的是二者较小值 $a[1]$。
- 又或者，从最长、次长、第三长的的特殊子串（$a[0],a[1],a[2]$）中各取一个长为 $a[2]$ 的特殊子串。

这三种情况取最大值，即

$$
\max(a[0]-2, \min(a[0]-1, a[1]), a[2])
$$

对每个长度列表计算上式，取最大值即为答案。

如果答案是 $0$，返回 $-1$。

代码实现时，在数组末尾加两个 $0$，就无需特判 $a$ 长度小于 $3$ 的情况了。

附：[视频讲解](https://www.bilibili.com/video/BV1XG411B7bX?t=1m44s) 第二题。

```py [sol-Python3]
class Solution:
    def maximumLength(self, s: str) -> int:
        groups = defaultdict(list)
        cnt = 0
        for i, ch in enumerate(s):
            cnt += 1
            if i + 1 == len(s) or ch != s[i + 1]:
                groups[ch].append(cnt)  # 统计连续字符长度
                cnt = 0

        ans = 0
        for a in groups.values():
            a.sort(reverse=True)
            a.extend([0, 0])  # 假设还有两个空串
            ans = max(ans, a[0] - 2, min(a[0] - 1, a[1]), a[2])

        return ans if ans else -1
```

```java [sol-Java]
class Solution {
    public int maximumLength(String S) {
        char[] s = S.toCharArray();
        List<Integer>[] groups = new ArrayList[26];
        Arrays.setAll(groups, i -> new ArrayList<>());
        int cnt = 0;
        for (int i = 0; i < s.length; i++) {
            cnt++;
            if (i + 1 == s.length || s[i] != s[i + 1]) {
                groups[s[i] - 'a'].add(cnt); // 统计连续字符长度
                cnt = 0;
            }
        }

        int ans = 0;
        for (List<Integer> a : groups) {
            if (a.isEmpty()) continue;
            a.sort(Collections.reverseOrder());
            a.add(0);
            a.add(0); // 假设还有两个空串
            ans = Math.max(ans, Math.max(a.get(0) - 2, Math.max(Math.min(a.get(0) - 1, a.get(1)), a.get(2))));
        }

        return ans > 0 ? ans : -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumLength(string s) {
        vector<int> groups[26];
        int cnt = 0, n = s.length();
        for (int i = 0; i < n; i++) {
            cnt++;
            if (i + 1 == n || s[i] != s[i + 1]) {
                groups[s[i] - 'a'].push_back(cnt); // 统计连续字符长度
                cnt = 0;
            }
        }

        int ans = 0;
        for (auto& a: groups) {
            if (a.empty()) continue;
            ranges::sort(a, greater());
            a.push_back(0);
            a.push_back(0); // 假设还有两个空串
            ans = max({ans, a[0] - 2, min(a[0] - 1, a[1]), a[2]});
        }

        return ans ? ans : -1;
    }
};
```

```go [sol-Go]
func maximumLength(s string) int {
    groups := [26][]int{}
    cnt := 0
    for i := range s {
        cnt++
        if i+1 == len(s) || s[i] != s[i+1] {
            groups[s[i]-'a'] = append(groups[s[i]-'a'], cnt) // 统计连续字符长度
            cnt = 0
        }
    }

    ans := 0
    for _, a := range groups {
        if len(a) == 0 {
            continue
        }
        slices.SortFunc(a, func(a, b int) int { return b - a })
        a = append(a, 0, 0) // 假设还有两个空串
        ans = max(ans, a[0]-2, min(a[0]-1, a[1]), a[2])
    }

    if ans == 0 {
        return -1
    }
    return ans
}
```

```js [sol-JavaScript]
var maximumLength = function(s) {
    const n = s.length;
    const groups = Array.from({ length: 26 }, () => []);
    let cnt = 0;
    for (let i = 0; i < n; i++) {
        cnt++;
        if (i + 1 === n || s[i] !== s[i + 1]) {
            groups[s[i].charCodeAt(0) - 'a'.charCodeAt(0)].push(cnt); // 统计连续字符长度
            cnt = 0;
        }
    }

    let ans = 0;
    for (let a of groups) {
        if (a.length === 0) {
            continue;
        }
        a.sort((x, y) => y - x);
        a.push(0, 0); // 假设还有两个空串
        ans = Math.max(ans, a[0] - 2, Math.min(a[0] - 1, a[1]), a[2]);
    }

    return ans ? ans : -1;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn maximum_length(s: String) -> i32 {
        let mut groups = vec![vec![]; 26];
        let s = s.as_bytes();
        let mut cnt = 0;
        for (i, &c) in s.iter().enumerate() {
            cnt += 1;
            if i + 1 == s.len() || c != s[i + 1] {
                groups[(c - b'a') as usize].push(cnt); // 统计连续字符长度
                cnt = 0;
            }
        }

        let mut ans = 0;
        for a in groups.iter_mut() {
            if a.is_empty() {
                continue;
            }
            a.sort_unstable_by(|x, y| y.cmp(x));
            a.push(0);
            a.push(0); // 假设还有两个空串
            ans = ans.max(a[0] - 2).max(a[1].min(a[0] - 1)).max(a[2]);
        }

        if ans > 0 { ans } else { -1 }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $s$ 的长度。如果改用堆维护前三大可以做到 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

额外输入一个正整数 $k$，把「至少三次」改成「至少 $k$ 次」，怎么做？

欢迎在评论区分享你的思路/代码。

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
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
