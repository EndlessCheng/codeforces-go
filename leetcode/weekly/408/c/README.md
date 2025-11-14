## 分析

设 $\textit{cnt}_0$ 为子串中的 $0$ 的个数，$\textit{cnt}_1$ 为子串中的 $1$ 的个数。

根据题意，$1$ 显著子串必须满足

$$
0\le \textit{cnt}_0^2 \le \textit{cnt}_1 \le n
$$

解得

$$
0\le \textit{cnt}_0 \le \sqrt{n}
$$

所以 $1$ 显著子串**至多**有 $\lfloor \sqrt{n}\rfloor$ 个 $0$。本题 $\lfloor \sqrt{n}\rfloor\le 200$。

## 核心思路

枚举子串右端点，分别计算：

- **恰好**有 $0$ 个 $0$ 的子串有多少个？
- **恰好**有 $1$ 个 $0$ 的子串有多少个？
- **恰好**有 $2$ 个 $0$ 的子串有多少个？
- ……
- **恰好**有 $\lfloor \sqrt{n}\rfloor$ 个 $0$ 的子串有多少个？

记录 $s$ 中的 $0$ 的下标，方便计算子串个数。

## 详细思路

设子串右端点为 $r$。

设上一个 $0$ 的下标为 $i$。

那么右端点**固定**为 $r$ 的子串 $[i+1,r],[i+2,r],\ldots,[r,r]$ 都是不包含 $0$ 的，一定是 $1$ 显著子串，这有 $r-i$ 个。

设上上一个 $0$ 的下标为 $j$。

那么子串 $[j+1,r],[j+2,r],\ldots,[i,r]$ 都恰好有 $1$ 个 $0$。

- 如果 $s_r = 1$，那么这些子串都至少有 $1$ 个 $1$，都是 $1$ 显著子串。
- 如果 $s_r = 0$，也就是 $i = r$，那么最右边的 $[i,r]$ 不是 $1$ 显著子串。

这意味着我们还要考虑子串是否**有足够的** $1$。

下面考虑更一般的情况。

![lc3234-3c.png](https://pic.leetcode.cn/1763092355-WgVIYY-lc3234-3c.png){:width=600px}

上图用红线/蓝线标记的子串，右端点相同，左端点不同，都恰好有 $2$ 个 $0$。

子串 $1$ 的个数必须 $\ge$ 子串 $0$ 的个数的平方，也就是至少有 $2^2=4$ 个 $1$。上图中的红色子串不满足要求，蓝色子串满足要求。

在包含的 $0$ 的个数不变的前提下，子串左端点越小，$1$ 的个数越多，越满足要求；子串左端点越大，$1$ 的个数越少，越不满足要求。

算出**刚好满足要求**的那个蓝色子串的左端点（即左端点的最大值），以及更左边的 $0$ 的位置（图中的 $p$，加一得到左端点的最小值），就能算出蓝色子串的个数。

一般地，外层循环枚举子串右端点 $r=0,1,2,\ldots,n-1$，内层循环枚举子串**恰好**有 $\textit{cnt}_0 = 0,1,2,\ldots$ 个 $0$，如果 $\textit{cnt}_0^2$ 超过了 $[0,r]$ 中的 $1$ 的个数，则跳出枚举 $\textit{cnt}_0$ 的循环。

设当前枚举的 $\textit{cnt}_0$ 对应的左右两个 $0$ 的下标分别为 $p$ 和 $q$。

右端点为 $r$ 且恰好有 $\textit{cnt}_0$ 个 $0$ 的最短子串为 $[q,r]$，设其有 $\textit{cnt}_1$ 个 $1$。

分类讨论：

- 如果 $\textit{cnt}_0^2\le \textit{cnt}_1$，那么子串左端点可以是 $p+1,p+2,\ldots, q$，一共 $q-p$ 个合法左端点。
- 如果 $\textit{cnt}_0^2> \textit{cnt}_1$，那么还需要包含至少 $\textit{cnt}_0^2 - \textit{cnt}_1$ 个 $1$，所以子串左端点的最大值为 $q - (\textit{cnt}_0^2 - \textit{cnt}_1)$，最小值为 $p+1$，一共 $q - (\textit{cnt}_0^2 - \textit{cnt}_1) - p$ 个合法左端点。如果个数是负数，则没有合法左端点。

综合一下，合法左端点最小是 $p+1$，最大是 $q - \max(\textit{cnt}_0^2 - \textit{cnt}_1, 0)$，所以一共有

$$
q - \max(\textit{cnt}_0^2 - \textit{cnt}_1, 0) - p
$$

个合法左端点。考虑到上式可能是负数，所以还要再与 $0$ 取最大值，得

$$
\max\left(q - \max(\textit{cnt}_0^2 - \textit{cnt}_1, 0) - p, 0\right)
$$

累加上式，即为答案。

为了知道 $0$ 的下标，我们用一个列表记录遍历到的 $0$ 的下标。

```py [sol-Python3]
# 手写 max 更快
max = lambda a, b: b if b > a else a

class Solution:
    def numberOfSubstrings(self, s: str) -> int:
        pos0 = [-1]  # 哨兵，方便处理 cnt0 达到最大时的计数
        total1 = 0  # [0,r] 中的 1 的个数
        ans = 0

        for r, ch in enumerate(s):
            if ch == '0':
                pos0.append(r)  # 记录 0 的下标
            else:
                total1 += 1
                ans += r - pos0[-1]  # 单独计算不含 0 的子串个数

            # 倒着遍历 pos0，就相当于在从小到大枚举 cnt0
            for i in range(len(pos0) - 1, 0, -1):
                cnt0 = len(pos0) - i
                if cnt0 * cnt0 > total1:
                    break
                p, q = pos0[i - 1], pos0[i]
                cnt1 = r - q + 1 - cnt0  # [q,r] 中的 1 的个数 = [q,r] 的长度 - cnt0
                ans += max(q - max(cnt0 * cnt0 - cnt1, 0) - p, 0)

        return ans
```

```java [sol-Java]
class Solution {
    public int numberOfSubstrings(String s) {
        int n = s.length();
        int ans = 0;
        int total1 = 0; // [0,r] 中的 1 的个数
        int[] pos0 = new int[n + 1]; // 0 的下标
        pos0[0] = -1; // 加个 -1 哨兵，方便处理 cnt0 达到最大时的计数
        int size = 1;

        for (int r = 0; r < n; r++) {
            if (s.charAt(r) == '0') {
                pos0[size++] = r; // 记录 0 的下标
            } else {
                total1++;
                ans += r - pos0[size - 1]; // 单独计算不含 0 的子串个数
            }

            // 倒着遍历 pos0，那么 cnt0 = size - i
            for (int i = size - 1; i > 0 && (size - i) * (size - i) <= total1; i--) {
                int p = pos0[i - 1];
                int q = pos0[i];
                int cnt0 = size - i;
                int cnt1 = r - q + 1 - cnt0; // [q,r] 中的 1 的个数 = [q,r] 的长度 - cnt0
                ans += Math.max(q - Math.max(cnt0 * cnt0 - cnt1, 0) - p, 0);
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfSubstrings(string s) {
        vector<int> pos0 = {-1}; // 加个 -1 哨兵，方便处理 cnt0 达到最大时的计数
        int total1 = 0; // [0,r] 中的 1 的个数
        int ans = 0;
        for (int r = 0; r < s.size(); r++) {
            char ch = s[r];
            if (ch == '0') {
                pos0.push_back(r); // 记录 0 的下标
            } else {
                total1++;
                ans += r - pos0.back(); // 单独计算不含 0 的子串个数
            }

            int m = pos0.size();
            // 倒着遍历 pos0，那么 cnt0 = m - i
            for (int i = m - 1; i > 0 && (m - i) * (m - i) <= total1; i--) {
                int p = pos0[i - 1], q = pos0[i];
                int cnt0 = m - i;
                int cnt1 = r - q + 1 - cnt0; // [q,r] 中的 1 的个数 = [q,r] 的长度 - cnt0
                ans += max(q - max(cnt0 * cnt0 - cnt1, 0) - p, 0);
            }
        }
        return ans;
    }
};
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

int numberOfSubstrings(char* s) {
    int n = strlen(s);
    int* pos0 = malloc((n + 1) * sizeof(int)); // 0 的下标
    pos0[0] = -1; // 加个 -1 哨兵，方便处理 cnt0 达到最大时的计数
    int size = 1;
    int total1 = 0; // [0,r] 中的 1 的个数
    int ans = 0;

    for (int r = 0; r < n; r++) {
        if (s[r] == '0') {
            pos0[size++] = r; // 记录 0 的下标
        } else {
            total1++;
            ans += r - pos0[size - 1]; // 单独计算不含 0 的子串个数
        }

        // 倒着遍历 pos0，那么 cnt0 = size - i
        for (int i = size - 1; i > 0 && (size - i) * (size - i) <= total1; i--) {
            int p = pos0[i - 1], q = pos0[i];
            int cnt0 = size - i;
            int cnt1 = r - q + 1 - cnt0; // [q,r] 中的 1 的个数 = [q,r] 的长度 - cnt0
            ans += MAX(q - MAX(cnt0 * cnt0 - cnt1, 0) - p, 0);
        }
    }

    free(pos0);
    return ans;
}
```

```go [sol-Go]
func numberOfSubstrings(s string) (ans int) {
	pos0 := []int{-1} // 加个 -1 哨兵，方便处理 cnt0 达到最大时的计数
	total1 := 0 // [0,r] 中的 1 的个数
	for r, ch := range s {
		if ch == '0' {
			pos0 = append(pos0, r) // 记录 0 的下标
		} else {
			total1++
			ans += r - pos0[len(pos0)-1] // 单独计算不含 0 的子串个数
		}

		m := len(pos0)
		// 倒着遍历 pos0，那么 cnt0 = m-i
		for i := m - 1; i > 0 && (m-i)*(m-i) <= total1; i-- {
			p, q := pos0[i-1], pos0[i]
			cnt0 := m - i
			cnt1 := r - q + 1 - cnt0 // [q,r] 中的 1 的个数 = [q,r] 的长度 - cnt0
			ans += max(q-max(cnt0*cnt0-cnt1, 0)-p, 0)
		}
	}
	return
}
```

```js [sol-JavaScript]
var numberOfSubstrings = function(s) {
    const pos0 = [-1]; // 加个 -1 哨兵，方便处理 cnt0 达到最大时的计数
    let total1 = 0; // [0,r] 中的 1 的个数
    let ans = 0;
    for (let r = 0; r < s.length; r++) {
        if (s[r] === '0') {
            pos0.push(r); // 记录 0 的下标
        } else {
            total1++;
            ans += r - pos0[pos0.length - 1]; // 单独计算不含 0 的子串个数
        }

        const m = pos0.length;
        // 倒着遍历 pos0，那么 cnt0 = m - i
        for (let i = m - 1; i > 0 && (m - i) * (m - i) <= total1; i--) {
            const p = pos0[i - 1], q = pos0[i];
            const cnt0 = m - i;
            const cnt1 = r - q + 1 - cnt0; // [q,r] 中的 1 的个数 = [q,r] 的长度 - cnt0
            ans += Math.max(q - Math.max(cnt0 * cnt0 - cnt1, 0) - p, 0);
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn number_of_substrings(s: String) -> i32 {
        let mut pos0 = vec![-1]; // 加个 -1 哨兵，方便处理 cnt0 达到最大时的计数
        let mut total1 = 0; // [0,r] 中的 1 的个数
        let mut ans = 0;
        for (r, ch) in s.bytes().enumerate() {
            let r = r as i32;
            if ch == b'0' {
                pos0.push(r); // 记录 0 的下标
            } else {
                total1 += 1;
                ans += r - pos0.last().unwrap(); // 单独计算不含 0 的子串个数
            }

            // 倒着遍历 pos0，就相当于在从小到大枚举 cnt0
            for i in (1..pos0.len()).rev() {
                let cnt0 = (pos0.len() - i) as i32;
                if cnt0 * cnt0 > total1 {
                    break;
                }
                let p = pos0[i - 1];
                let q = pos0[i];
                let cnt1 = r - q + 1 - cnt0; // [q,r] 中的 1 的个数 = [q,r] 的长度 - cnt0
                ans += 0.max(q - 0.max(cnt0 * cnt0 - cnt1) - p);
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt{n})$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

**注**：使用队列，只保存 $r$ 及其左侧的 $\mathcal{O}(\sqrt{n})$ 个 $0$ 的下标，可以把空间复杂度优化到 $\mathcal{O}(\sqrt{n})$。

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
