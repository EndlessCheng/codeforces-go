## 一条线段

从特殊到一般，先想想只有一条线段要怎么做。

如果线段的右端点没有奖品，我们可以把线段左移，使其右端点恰好有奖品，这不会让线段覆盖的奖品个数变少。所以只需枚举 $\textit{prizePositions}[\textit{right}]$ 为线段的右端点，然后需要算出最远（最小）覆盖的奖品的位置 $\textit{prizePositions}[\textit{left}]$，此时覆盖的奖品的个数为

$$
\textit{right} - \textit{left} + 1
$$

由于 $\textit{right}$ 变大时，$\textit{left}$ 也会变大，有单调性，可以用**滑动窗口**快速算出 $\textit{left}$。原理见 [滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

⚠**注意**：$\textit{prizePositions}[\textit{left}]$ 不一定是线段的左端点。$\textit{prizePositions}[\textit{left}]$ 只是最左边的被线段覆盖的那个奖品的位置，线段左端点可能比 $\textit{prizePositions}[\textit{left}]$ 更小。

## 两条线段

两条线段一左一右。

贪心地想，两条线段**不相交**肯定比相交更好，覆盖的奖品可能更多。

设**第二条线段**右端点在 $\textit{prizePositions}[\textit{right}]$ 时，最远（最小）覆盖的奖品的位置为 $\textit{prizePositions}[\textit{left}]$。

我们需要计算在 $\textit{prizePositions}[\textit{left}]$ 左侧的**第一条线段**最多可以覆盖多少个奖品。这可以保证两条线段不相交。

定义 $\textit{mx}[i+1]$ 表示第一条线段右端点 $\le \textit{prizePositions}[i]$ 时，最多可以覆盖多少个奖品。特别地，定义 $\textit{mx}[0]=0$。

如何计算 $\textit{mx}$？

根据 $\textit{mx}$ 的定义，我们相当于在计算 $i - \textit{left}_i + 1$ 的**前缀最大值**，其中 $\textit{left}_i$ 表示右端点覆盖奖品 $\textit{prizePositions}[i]$ 时，最左边的被线段覆盖的奖品。

所以有

$$
\textit{mx}[i + 1] = \max(\textit{mx}[i], i - \textit{left}_i + 1)
$$

如何计算两条线段可以覆盖的奖品个数？

- 第二条线段覆盖的奖品个数为 $\textit{right} - \textit{left} + 1$。
- 第一条线段覆盖的奖品个数为线段右端点 $\le \textit{prizePositions}[\textit{left}-1]$ 时，最多覆盖的奖品个数，即 $\textit{mx}[\textit{left}]$。

综上，两条线段可以覆盖的奖品个数为

$$
\textit{mx}[\textit{left}] + \textit{right}-\textit{left}+1
$$

枚举 $\textit{right}$ 的过程中，取上式的最大值，即为答案。

我们遍历了所有的奖品作为第二条线段的右端点，通过 $\textit{mx}[\textit{left}]$ 保证第一条线段与第二条线段不相交，且第一条线段覆盖了第二条线段左侧的最多奖品。那么这样遍历后，算出的答案就一定是所有情况中的最大值。

如果脑中没有一幅直观的图像，可以看看 [视频讲解【双周赛 97】](https://www.bilibili.com/video/BV1rM4y1X7z9/)的第三题。

**小优化**：如果 $2k+1\ge \textit{prizePositions}[n-1] - \textit{prizePositions}[0]$，说明所有奖品都可以被覆盖，直接返回 $n$。

```py [sol-Python3]
class Solution:
    def maximizeWin(self, prizePositions: List[int], k: int) -> int:
        n = len(prizePositions)
        if k * 2 + 1 >= prizePositions[-1] - prizePositions[0]:
            return n
        ans = left = 0
        mx = [0] * (n + 1)
        for right, p in enumerate(prizePositions):
            while p - prizePositions[left] > k:
                left += 1
            ans = max(ans, mx[left] + right - left + 1)
            mx[right + 1] = max(mx[right], right - left + 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int maximizeWin(int[] prizePositions, int k) {
        int n = prizePositions.length;
        if (k * 2 + 1 >= prizePositions[n - 1] - prizePositions[0]) {
            return n;
        }
        int ans = 0;
        int left = 0;
        int[] mx = new int[n + 1];
        for (int right = 0; right < n; right++) {
            while (prizePositions[right] - prizePositions[left] > k) {
                left++;
            }
            ans = Math.max(ans, mx[left] + right - left + 1);
            mx[right + 1] = Math.max(mx[right], right - left + 1);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximizeWin(vector<int>& prizePositions, int k) {
        int n = prizePositions.size();
        if (k * 2 + 1 >= prizePositions[n - 1] - prizePositions[0]) {
            return n;
        }
        int ans = 0, left = 0;
        vector<int> mx(n + 1);
        for (int right = 0; right < n; right++) {
            while (prizePositions[right] - prizePositions[left] > k) {
                left++;
            }
            ans = max(ans, mx[left] + right - left + 1);
            mx[right + 1] = max(mx[right], right - left + 1);
        }
        return ans;
    }
};
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

int maximizeWin(int* prizePositions, int n, int k) {
    if (k * 2 + 1 >= prizePositions[n - 1] - prizePositions[0]) {
        return n;
    }
    int ans = 0, left = 0;
    int* mx = calloc(n + 1, sizeof(int));
    for (int right = 0; right < n; right++) {
        int p = prizePositions[right];
        while (p - prizePositions[left] > k) {
            left++;
        }
        ans = MAX(ans, mx[left] + right - left + 1);
        mx[right + 1] = MAX(mx[right], right - left + 1);
    }
    free(mx);
    return ans;
}
```

```go [sol-Go]
func maximizeWin(prizePositions []int, k int) (ans int) {
    n := len(prizePositions)
    if k*2+1 >= prizePositions[n-1]-prizePositions[0] {
        return n
    }
    mx := make([]int, n+1)
    left := 0
    for right, p := range prizePositions {
        for p-prizePositions[left] > k {
            left++
        }
        ans = max(ans, mx[left]+right-left+1)
        mx[right+1] = max(mx[right], right-left+1)
    }
    return
}
```

```js [sol-JavaScript]
var maximizeWin = function(prizePositions, k) {
    const n = prizePositions.length;
    if (k * 2 + 1 >= prizePositions[n - 1] - prizePositions[0]) {
        return n;
    }
    const mx = Array(n + 1).fill(0);
    let ans = 0, left = 0;
    for (let right = 0; right < n; right++) {
        const p = prizePositions[right];
        while (p - prizePositions[left] > k) {
            left++;
        }
        ans = Math.max(ans, mx[left] + right - left + 1);
        mx[right + 1] = Math.max(mx[right], right - left + 1);
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn maximize_win(prize_positions: Vec<i32>, k: i32) -> i32 {
        let n = prize_positions.len();
        if k * 2 + 1 >= prize_positions[n - 1] - prize_positions[0] {
            return n as _;
        }
        let mut ans = 0;
        let mut left = 0;
        let mut mx = vec![0; n + 1];
        for (right, &p) in prize_positions.iter().enumerate() {
            while p - prize_positions[left] > k {
                left += 1;
            }
            ans = ans.max(mx[left] + right - left + 1);
            mx[right + 1] = mx[right].max(right - left + 1);
        }
        ans as _
    }
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{prizePositions}$ 的长度。虽然写了个二重循环，但是内层循环中对 $\textit{left}$ 加一的**总**执行次数不会超过 $n$ 次，所以总的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

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
