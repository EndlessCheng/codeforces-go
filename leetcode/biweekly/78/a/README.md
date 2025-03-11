## 方法一：字符串

把 $\textit{num}$ 转成十进制字符串 $s$。

枚举 $s$ 的所有长为 $k$ 的子串，设子串对应的数值为 $x$，如果 $x>0$ 且 $\textit{num}\bmod x = 0$，那么子串符合题目要求，答案加一。

```py [sol-Python3]
class Solution:
    def divisorSubstrings(self, num: int, k: int) -> int:
        s = str(num)
        ans = 0
        for i in range(k, len(s) + 1):
            x = int(s[i - k: i])  # 长为 k 的子串
            if x > 0 and num % x == 0:  # 子串能整除 num 
                ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int divisorSubstrings(int num, int k) {
        String s = String.valueOf(num);
        int ans = 0;
        for (int i = k; i <= s.length(); i++) {
            int x = Integer.parseInt(s.substring(i - k, i)); // 长为 k 的子串
            if (x > 0 && num % x == 0) { // 子串能整除 num
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int divisorSubstrings(int num, int k) {
        string s = to_string(num);
        int ans = 0;
        for (int i = k; i <= s.size(); i++) {
            int x = stoi(s.substr(i - k, k)); // 长为 k 的子串
            if (x > 0 && num % x == 0) { // 子串能整除 num
                ans++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func divisorSubstrings(num, k int) (ans int) {
    s := strconv.Itoa(num)
    for i := k; i <= len(s); i++ {
        x, _ := strconv.Atoi(s[i-k : i]) // 长为 k 的子串
        if x > 0 && num%x == 0 { // 子串能整除 num 
            ans++
        }
    }
    return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n-k)k)$，其中 $n=\mathcal{O}(\log \textit{num})$ 是 $\textit{num}$ 的十进制长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：数学

例如 $\textit{num}=1234,\ k=2$。

设 $m=10^k=100$。

- 初始化 $n=\textit{num}$。
- 计算 $n\bmod m = 34$，这是最右边的长为 $k$ 的子串。判断是否合法。然后去掉 $n$ 的个位数，也就是把 $n$ 除以 $10$ 下取整，现在 $n = 123$。
- 计算 $n\bmod m = 23$，这是中间的长为 $k$ 的子串。判断是否合法。然后去掉 $n$ 的个位数，也就是把 $n$ 除以 $10$ 下取整，现在 $n = 12$。
- 计算 $n\bmod m = 12$，这是最左边的长为 $k$ 的子串。判断是否合法。然后去掉 $n$ 的个位数，也就是把 $n$ 除以 $10$ 下取整，现在 $n = 1 < m/10 = 10^{k-1}$，说明 $n$ 的十进制长度不足 $k$，结束。

```py [sol-Python3]
class Solution:
    def divisorSubstrings(self, num: int, k: int) -> int:
        M = 10 ** k
        ans = 0
        n = num
        while n >= M // 10:
            x = n % M
            if x > 0 and num % x == 0:
                ans += 1
            n //= 10
        return ans
```

```java [sol-Java]
class Solution {
    public int divisorSubstrings(int num, int k) {
        final long M = (long) Math.pow(10, k);
        int ans = 0;
        for (int n = num; n >= M / 10; n /= 10) {
            int x = (int) (n % M);
            if (x > 0 && num % x == 0) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int divisorSubstrings(int num, int k) {
        const long long M = pow(10, k);
        int ans = 0;
        for (int n = num; n >= M / 10; n /= 10) {
            int x = n % M;
            if (x > 0 && num % x == 0) {
                ans++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func divisorSubstrings(num, k int) (ans int) {
    m := int(math.Pow10(k))
    for n := num; n >= m/10; n /= 10 {
        x := n % m
        if x > 0 && num%x == 0 {
            ans++
        }
    }
    return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n-k)$，其中 $n=\mathcal{O}(\log \textit{num})$ 是 $\textit{num}$ 的十进制长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
