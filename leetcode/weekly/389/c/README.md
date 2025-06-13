**定理**：一定存在一种字母是不需要删除的。比如示例 1 没有删除字母 $\texttt{b}$。

**证明**：反证法，假设每种字母都至少删除一个。我们可以把每种字母的剩余出现次数都增加一，这不会影响字母数量之差，且总删除次数更少，矛盾。故原命题成立。

统计 $\textit{word}$ 每种字母的出现次数，记到一个数组 $\textit{cnt}$ 中。

枚举 $i$ 作为出现次数最小的字母，为了保留尽量多的字母，字母 $i$ 肯定不需要删除。此外，出现次数最多的字母，出现次数不能超过 $\textit{cnt}[i]+k$。

分类讨论：

- 出现次数小于 $\textit{cnt}[i]$ 的字母，全部删除。
- 出现次数大于等于 $\textit{cnt}[i]$ 的字母 $j$，保留 $\min(\textit{cnt}[j], \textit{cnt}[i] + k)$ 个。累加保留的字母个数，取最大值，得到最多保留的字母个数 $\textit{maxSave}$。

最后，用 $\textit{word}$ 的长度，减去 $\textit{maxSave}$，即为答案。

代码实现时，为方便计算那些字母的出现次数小于 $\textit{cnt}[i]$，哪些大于等于 $\textit{cnt}[i]$，把 $\textit{cnt}$ 排序。

[本题视频讲解](https://www.bilibili.com/video/BV1RH4y1W7DP/?t=5m37s)

```py [sol-Python3]
class Solution:
    def minimumDeletions(self, word: str, k: int) -> int:
        cnt = sorted(Counter(word).values())
        max_save = 0
        for i, base in enumerate(cnt):
            s = sum(min(c, base + k) for c in cnt[i:])  # 至多保留 base+k 个
            max_save = max(max_save, s)
        return len(word) - max_save
```

```py [sol-Python3 写法二]
class Solution:
    def minimumDeletions(self, word: str, k: int) -> int:
        cnt = sorted(Counter(word).values())
        max_save = max(sum(min(c, base + k) for c in cnt[i:])
                       for i, base in enumerate(cnt))
        return len(word) - max_save
```

```java [sol-Java]
class Solution {
    public int minimumDeletions(String word, int k) {
        int[] cnt = new int[26];
        for (char c : word.toCharArray()) {
            cnt[c - 'a']++;
        }
        Arrays.sort(cnt);

        int maxSave = 0;
        for (int i = 0; i < 26; i++) {
            int sum = 0;
            for (int j = i; j < 26; j++) {
                sum += Math.min(cnt[j], cnt[i] + k); // 至多保留 cnt[i]+k 个
            }
            maxSave = Math.max(maxSave, sum);
        }

        return word.length() - maxSave;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumDeletions(string word, int k) {
        int cnt[26]{};
        for (char c: word) {
            cnt[c - 'a']++;
        }
        ranges::sort(cnt);

        int max_save = 0;
        for (int i = 0; i < 26; i++) {
            int sum = 0;
            for (int j = i; j < 26; j++) {
                sum += min(cnt[j], cnt[i] + k); // 至多保留 cnt[i]+k 个
            }
            max_save = max(max_save, sum);
        }

        return word.length() - max_save;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))
#define MAX(a, b) ((b) > (a) ? (b) : (a))

int cmp(const void* a, const void* b) {
    return *(int*)a - *(int*)b;
}

int minimumDeletions(char* word, int k) {
    int cnt[26] = {};
    int n = 0;
    for (; word[n]; n++) {
        cnt[word[n] - 'a']++;
    }
    qsort(cnt, 26, sizeof(int), cmp);

    int max_save = 0;
    for (int i = 0; i < 26; i++) {
        int sum = 0;
        for (int j = i; j < 26; j++) {
            sum += MIN(cnt[j], cnt[i] + k); // 至多保留 cnt[i]+k 个
        }
        max_save = MAX(max_save, sum);
    }

    return n - max_save;
}
```

```go [sol-Go]
func minimumDeletions(word string, k int) int {
	cnt := make([]int, 26)
	for _, b := range word {
		cnt[b-'a']++
	}
	slices.Sort(cnt)

	maxSave := 0
	for i, base := range cnt {
		sum := 0
		for _, c := range cnt[i:] {
			sum += min(c, base+k) // 至多保留 base+k 个
		}
		maxSave = max(maxSave, sum)
	}

	return len(word) - maxSave
}
```

```js [sol-JavaScript]
var minimumDeletions = function(word, k) {
    const cnt = Array(26).fill(0);
    for (const c of word) {
        cnt[c.charCodeAt(0) - 97]++;
    }
    cnt.sort((a, b) => a - b);

    let maxSave = 0;
    for (let i = 0; i < 26; i++) {
        let sum = 0;
        for (let j = i; j < 26; j++) {
            sum += Math.min(cnt[j], cnt[i] + k); // 至多保留 cnt[i]+k 个
        }
        maxSave = Math.max(maxSave, sum);
    }

    return word.length - maxSave;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn minimum_deletions(word: String, k: i32) -> i32 {
        let mut cnt = [0; 26];
        for c in word.bytes() {
            cnt[(c - b'a') as usize] += 1;
        }
        cnt.sort_unstable();

        let mut max_save = 0;
        for (i, &base) in cnt.iter().enumerate() {
            let mut sum = 0;
            for &c in &cnt[i..] {
                sum += c.min(base + k); // 至多保留 base+k 个
            }
            max_save = max_save.max(sum);
        }

        word.len() as i32 - max_save
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|^2)$，其中 $n$ 为 $\textit{word}$ 的长度，$|\Sigma|=26$ 为字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

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
