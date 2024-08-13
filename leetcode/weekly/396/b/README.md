根据题意，我们只能选择首字母下标为 $0,k,2k,3k,\cdots,n-k$ 的长为 $k$ 的子串来操作（替换）。

并且，$k$ 周期字符串意味着，所有首字母下标为 $0,k,2k,3k,\cdots,n-k$ 的长为 $k$ 的子串均相等。

为使操作次数尽量少，我们可以计算最多保留多少个子串**不变**。也就是统计 $\textit{word}$ 中的这些首字母下标为 $0,k,2k,3k,\cdots,n-k$ 的长为 $k$ 的子串中，出现次数最多的子串的出现次数 $\textit{mx}$。**用出现次数最多的子串，替换其余子串。**

所以用子串的个数 $\dfrac{n}{k}$ 减去 $\textit{mx}$，就是最少操作次数。

请看 [视频讲解](https://www.bilibili.com/video/BV1Nf421U7em/) 第二题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minimumOperationsToMakeKPeriodic(self, word: str, k: int) -> int:
        n = len(word)
        cnt = Counter(word[i - k: i] for i in range(k, n + 1, k))
        mx = max(cnt.values())
        return n // k - mx
```

```java [sol-Java]
class Solution {
    public int minimumOperationsToMakeKPeriodic(String word, int k) {
        int n = word.length();
        int mx = 0;
        HashMap<String, Integer> cnt = new HashMap<>();
        for (int i = k; i <= n; i += k) {
            String sub = word.substring(i - k, i);
            int c = cnt.merge(sub, 1, Integer::sum); // c = ++cnt[sub]
            mx = Math.max(mx, c);
        }
        return n / k - mx;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumOperationsToMakeKPeriodic(string word, int k) {
        int n = word.length(), mx = 0;
        unordered_map<string, int> cnt;
        for (int i = k; i <= n; i += k) {
            mx = max(mx, ++cnt[word.substr(i - k, k)]);
        }
        return n / k - mx;
    }
};
```

```go [sol-Go]
func minimumOperationsToMakeKPeriodic(word string, k int) int {
	n := len(word)
	cnt := map[string]int{}
	for i := k; i <= n; i += k {
		cnt[word[i-k:i]]++
	}
	mx := 0
	for _, c := range cnt {
		mx = max(mx, c)
	}
	return n/k - mx
}
```

```js [sol-JavaScript]
var minimumOperationsToMakeKPeriodic = function(word, k) {
    const n = word.length;
    const cnt = new Map();
    for (let i = k; i <= n; i += k) {
        const sub = word.slice(i - k, i);
        cnt.set(sub, (cnt.get(sub) ?? 0) + 1);
    }
    const mx = Math.max(...Array.from(cnt.values()));
    return Math.floor(n / k) - mx;
};
```

```rust [sol-Rust]
use std::collections::HashMap;

impl Solution {
    pub fn minimum_operations_to_make_k_periodic(word: String, k: i32) -> i32 {
        let n = word.len();
        let k = k as usize;
        let mut cnt = HashMap::new();
        for i in (k..=n).step_by(k) {
            *cnt.entry(&word[i - k..i]).or_insert(0) += 1;
        }
        let mx = *cnt.values().max().unwrap();
        (n / k) as i32 - mx
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{word}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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
