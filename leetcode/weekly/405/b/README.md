## 题意

生成所有长为 $n$ 的二进制字符串，要求字符串不包含 $\texttt{"00"}$。

## 方法一：回溯（爆搜）

类似 [17. 电话号码的字母组合](https://leetcode.cn/problems/letter-combinations-of-a-phone-number/)，用递归生成所有长为 $n$ 的符合要求的字符串。如果你没有写过这题，可以看 [视频讲解【基础算法精讲 14】](https://www.bilibili.com/video/BV1mG4y1A7Gu/)。

本题枚举字符串的位置 $i$ 填 $1$ 还是 $0$：

- 填 $1$，没有任何约束，可以直接填。
- 填 $0$，需要满足 $i=0$，或者 $i-1$ 这个位置（前一个位置）填的是 $1$。
- 填完后，继续往下递归，考虑 $i+1$ 怎么填。

从 $i=0$ 开始递归，到 $i=n$ 结束。

⚠**注意**：由于字符串长度固定为 $n$，本题不需要恢复现场，**直接覆盖**之前填的数据就行。

```py [sol-Python3]
class Solution:
    def validStrings(self, n: int) -> List[str]:
        ans = []
        path = [''] * n

        def dfs(i: int) -> None:
            if i == n:
                ans.append(''.join(path))  # 注意 join 需要 O(n) 时间
                return

            # 填 1
            path[i] = '1'
            dfs(i + 1)

            # 填 0
            if i == 0 or path[i - 1] == '1':
                path[i] = '0'  # 直接覆盖
                dfs(i + 1)

        dfs(0)
        return ans
```

```java [sol-Java]
class Solution {
    public List<String> validStrings(int n) {
        List<String> ans = new ArrayList<>();
        char[] path = new char[n];
        dfs(0, n, path, ans);
        return ans;
    }

    private void dfs(int i, int n, char[] path, List<String> ans) {
        if (i == n) {
            ans.add(new String(path));
            return;
        }

        // 填 1
        path[i] = '1';
        dfs(i + 1, n, path, ans);

        // 填 0
        if (i == 0 || path[i - 1] == '1') {
            path[i] = '0'; // 直接覆盖
            dfs(i + 1, n, path, ans);
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<string> validStrings(int n) {
        vector<string> ans;
        string path(n, 0);

        auto dfs = [&](auto&& dfs, int i) -> void {
            if (i == n) {
                ans.push_back(path); // 注意这里复制了一份 path，需要 O(n) 时间
                return;
            }

            // 填 1
            path[i] = '1';
            dfs(dfs, i + 1);

            // 填 0
            if (i == 0 || path[i - 1] == '1') {
                path[i] = '0'; // 直接覆盖
                dfs(dfs, i + 1);
            }
        };

        dfs(dfs, 0);
        return ans;
    }
};
```

```go [sol-Go]
func validStrings(n int) (ans []string) {
    path := make([]byte, n)
    var dfs func(i int)
    dfs = func(i int) {
        if i == n {
            ans = append(ans, string(path)) // 注意 string(path) 需要 O(n) 时间
            return
        }

        // 填 1
        path[i] = '1'
        dfs(i + 1)

        // 填 0
        if i == 0 || path[i-1] == '1' {
            path[i] = '0' // 直接覆盖
            dfs(i + 1)
        }
    }
    dfs(0)
    return
}
```

```js [sol-JavaScript]
var validStrings = function(n) {
    const ans = [];
    const path = Array(n);

    function dfs(i) {
        if (i === n) {
            ans.push(path.join('')); // 注意 join 需要 O(n) 时间
            return;
        }

        // 填 1
        path[i] = '1';
        dfs(i + 1);

        // 填 0
        if (i === 0 || path[i - 1] === '1') {
            path[i] = '0'; // 直接覆盖
            dfs(i + 1);
        }
    }

    dfs(0);
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn valid_strings(n: i32) -> Vec<String> {
        fn dfs(i: usize, path: &mut Vec<char>, ans: &mut Vec<String>) {
            if i == path.len() {
                ans.push(path.iter().collect());
                return;
            }

            // 填 1
            path[i] = '1';
            dfs(i + 1, path, ans);

            // 填 0
            if i == 0 || path[i - 1] == '1' {
                path[i] = '0'; // 直接覆盖
                dfs(i + 1, path, ans);
            }
        }

        let mut ans = vec![];
        let mut path = vec!['\0'; n as usize];
        dfs(0, &mut path, &mut ans);
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度（粗略分析）：$\mathcal{O}(n2^n)$。至多生成 $2^n$ 个字符串，每个字符串的长度为 $\mathcal{O}(n)$，这一共需要 $\mathcal{O}(n2^n)$ 的时间。
- 时间复杂度（精确分析）：$\mathcal{O}(n\varphi^n)$，其中 $\varphi=\dfrac{1+\sqrt 5}{2}\approx 1.618$。有多少个有效字符串？这个问题可以用动态规划解决，考虑字符串的最后一位填 $1$ 还是填 $0$。如果填 $1$，子问题是长为 $n-1$ 的有效字符串有多少个；如果填 $0$，那么倒数第二位必然填 $1$，子问题是长为 $n-2$ 的有效字符串有多少个。所以有递推式 $f(n) = f(n-1) + f(n-2)$，初始值 $f(0)=1,f(1)=2$。由斐波那契数列通项公式可得，有效字符串的个数为 $\mathcal{O}(\varphi^n)$，再算上生成每个字符串需要 $\mathcal{O}(n)$ 的时间，所以总的时间复杂度为 $\mathcal{O}(n\varphi^n)$。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 方法二：位运算

枚举 $[0,2^n-1]$ 中的 $i$，如果 $i$ 的长为 $n$ 的二进制中，没有相邻的 $0$，那么将其二进制字符串加入答案。

怎么判断二进制中是否有相邻的 $0$？

我们可以把 $i$ 取反（保留低 $n$ 位），记作 $x$。问题变成：判断 $x$ 中是否有相邻的 $1$。

需要一个一个地遍历二进制数 $x$ 的每一位吗？

不需要，我们可以用 `x & (x >> 1)` 来判断，如果这个值不为零，则说明 $x$ 中有相邻的 $1$，反之没有。例如 $x=110$，右移一位得 $011$，可以发现这两个二进制数的次低位都是 $1$，所以计算 AND 的结果必然不为 $0$。

代码实现时，可以直接枚举取反后的值 $x$，如果 `x & (x >> 1)` 等于 $0$，就把 $x$ 取反后的值（也就是 $i$）加入答案。

如何取反？

1. 创建一个低 $n$ 位全为 $1$ 的二进制数 `mask = (1 << n) - 1`。
2. 计算 `x ^ mask`，由于 $0$ 和 $1$ 异或后变成了 $1$，$1$ 和 $1$ 异或后变成了 $0$，所以 $x$ 的低 $n$ 位都取反了。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Ry411q71f/) 第二题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def validStrings(self, n: int) -> List[str]:
        ans = []
        mask = (1 << n) - 1
        for x in range(1 << n):
            if (x >> 1) & x == 0:
                # 0{n}b 表示长为 n 的有前导零的二进制
                ans.append(f"{x ^ mask:0{n}b}")
        return ans
```

```py [sol-Python3 写法二]
class Solution:
    def validStrings(self, n: int) -> List[str]:
        mask = (1 << n) - 1
        return [f"{x ^ mask:0{n}b}" for x in range(1 << n) if (x >> 1) & x == 0]
```

```java [sol-Java]
class Solution {
    public List<String> validStrings(int n) {
        List<String> ans = new ArrayList<>();
        int mask = (1 << n) - 1;
        for (int x = 0; x < (1 << n); x++) {
            if (((x >> 1) & x) == 0) {
                int i = x ^ mask;
                // 一种生成前导零的写法：在 i 前面插入 1<<n，转成字符串后再去掉插入的 1<<n
                ans.add(Integer.toBinaryString((1 << n) | i).substring(1));
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<string> validStrings(int n) {
        vector<string> ans;
        int mask = (1 << n) - 1;
        for (int x = 0; x < (1 << n); x++) {
            if (((x >> 1) & x) == 0) {
                ans.push_back(bitset<18>(x ^ mask).to_string().substr(18 - n));
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func validStrings(n int) (ans []string) {
    mask := 1<<n - 1
    for x := range 1 << n {
        if x>>1&x == 0 {
            ans = append(ans, fmt.Sprintf("%0*b", n, x^mask))
        }
    }
    return
}
```

```js [sol-JavaScript]
var validStrings = function(n) {
    const ans = [];
    const mask = (1 << n) - 1;
    for (let x = 0; x < (1 << n); x++) {
        if (((x >> 1) & x) === 0) {
            ans.push((x ^ mask).toString(2).padStart(n, '0'));
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    fn valid_strings(n: i32) -> Vec<String> {
        let mask = (1 << n) - 1;
        (0..1 << n)
            .filter(|&x| (x >> 1) & x == 0)
            .map(|x| format!("{:0w$b}", x ^ mask, w = n as usize))
            .collect()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(2^n)$。由方法一可知，生成所有字符串的时间复杂度，比枚举 $i$ 的时间复杂度更低。所以瓶颈在枚举上，时间复杂度为 $\mathcal{O}(2^n)$，
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

你能想出一个更快的枚举方式吗？如何直接跳到下一个合法的二进制数？

## 思考题

1. 额外输入一个长度不超过 $n$ 的二进制字符串 $t$，你生成的字符串**不能包含** $t$。本题相当于 $t=\texttt{"00"}$。
2. 把 $n$ 的数据范围改成 $1\le n\le 1000$，有多少个长为 $n$ 且不包含 $t$ 的二进制字符串？模 $10^9+7$。
3. 再额外输出一个二进制字符串 $\textit{limit}$，有多少个长为 $n$，二进制不超过 $\textit{limit}$ 且不包含 $t$ 的二进制字符串？模 $10^9+7$。做法同 [1397. 找到所有好字符串](https://leetcode.cn/problems/find-all-good-strings/)。

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
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
