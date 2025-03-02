## 分析

根据题意，最大波动值只由 $s$ 中的两种字母决定，至于是哪两种我们还不知道，可以枚举这两种字母。

由于 $s$ 只包含小写字母，我们可以从 $26$ 个小写字母中选出 $2$ 个不同的字母（相同字母无需考虑，波动值为 $0$），并假设这两个字母是答案子串中出现次数最多的和最少的。这一共需要枚举 $A_{26}^2=26\cdot 25=650$ 种不同的字母组合。

例如子串只有 $3$ 个 $a$ 和 $2$ 个 $b$，那么波动值为 $3-2=1$。若把 $a$ 视作 $1$，$b$ 视作 $-1$，也可以算出波动值为 $1+1+1+(-1)+(-1)=1$。

设出现次数最多的字母为 $a$，出现次数最少的字母为 $b$。把 $a$ 视作 $1$，$b$ 视作 $-1$，其余字母视作 $0$，最大波动值就等同于 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)。

## 状态机 DP

和 53 题的不同之处是，$a$ 和 $b$ **必须都在子串中**。

- 只包含 $a$ 的子串，实际波动值为 $0$，不能用 $a$ 的个数作为波动值。
- 只包含 $b$ 的子串，由于我们把 $b$ 当作 $-1$，这样算出的「子数组和」是负数，不会更新答案的最大值。
- 当子串包含 $b$ 且「子数组和」是正数，那么子串一定包含 $a$。所以只要保证子串包含 $b$ 就行（子串全为 $b$ 也没关系，不会影响答案）。

**状态定义**。在 53 题的基础上，用一个额外的参数 $j$ 表示是否需要包含 $b$：

- 定义 $f[i+1][0]$ 表示以 $s[i]$ 结尾的最大子数组和，包不包含 $b$ 都可以。加一是方便定义初始值。
- 定义 $f[i+1][1]$ 表示以 $s[i]$ 结尾的、一定包含 $b$ 的最大子数组和。

**状态转移方程**：

- 对于 $f[i+1][0]$，转移方程同 53 题，即 $f[i+1][0] = \max(f[i][0], 0) + v$，其中 $v$ 等于 $1$、$-1$ 或 $0$，见前文的分析。
- 对于 $f[i+1][1]$：
  - 如果 $s[i]=a$，那么只能在以 $s[i-1]$ 结尾的、一定包含 $b$ 的子数组后面加上 $s[i]$，即 $f[i+1][1] = f[i][1] + 1$。
  - 如果 $s[i]=b$，那么问题等价于以 $s[i]$ 结尾的最大子数组和（必然包含 $b$），即 $f[i+1][1] = f[i+1][0]$。
  - 其他情况和 $s[i]=a$ 是一样的，只能在以 $s[i-1]$ 结尾的、一定包含 $b$ 的子数组后面加上 $s[i]$，即 $f[i+1][1] = f[i][1] + 0 = f[i][1]$。

**初始值**：

- $f[0][0]=0$。一开始什么也没有，子数组和为 $0$。
- $f[0][1]=-\infty$。一开始什么也没有，一定包含 $b$ 的情况不存在，用 $-\infty$ 表示，这样计算 $\max$ 不会取到 $-\infty$。

**答案**：$\max\limits_{i=1}^{n} f[i][1]$。注意答案不是 $f[n][1]$，因为这仅仅表示以 $s[n-1]$ 结尾的子串。

代码实现时，$f$ 的第一个维度可以去掉，用两个变量 $f_0$ 和 $f_1$ 分别表示 $f[i][0]$ 和 $f[i][1]$：

- $f_0 = \max(f_0, 0) + v$，这和 53 题完全一样。
- 如果 $s[i]=a$，$f_1 = f_1 + 1$。
- 如果 $s[i]=b$，$f_1 = f_0$。
- 其余情况 $f_1$ 不变。

初始值 $f_0=0,\ f_1=-\infty$。

循环末尾用 $f_1$ 更新答案的最大值。

```py [sol-Python3]
class Solution:
    def largestVariance(self, s: str) -> int:
        ans = 0
        for a, b in permutations(ascii_lowercase, 2):  # 枚举所有小写字母对
            f0, f1 = 0, -inf
            for ch in s:
                if ch == a:
                    f0 = max(f0, 0) + 1
                    f1 += 1
                elif ch == b:
                    f1 = f0 = max(f0, 0) - 1
                # else: f0 = max(f0, 0) 可以留到 ch 等于 a 或者 b 的时候计算，f1 不变
                ans = max(ans, f1)
        return ans
```

```java [sol-Java]
class Solution {
    public int largestVariance(String S) {
        char[] s = S.toCharArray();
        int ans = 0;
        for (char a = 'a'; a <= 'z'; a++) {
            for (char b = 'a'; b <= 'z'; b++) {
                if (b == a) {
                    continue;
                }
                int f0 = 0;
                int f1 = Integer.MIN_VALUE;
                for (char ch : s) {
                    if (ch == a) {
                        f0 = Math.max(f0, 0) + 1;
                        f1++;
                    } else if (ch == b) {
                        f1 = f0 = Math.max(f0, 0) - 1;
                    } // else f0 = Math.max(f0, 0); 可以留到 ch 等于 a 或者 b 的时候计算，f1 不变
                    ans = Math.max(ans, f1);
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int largestVariance(string s) {
        int ans = 0;
        for (char a = 'a'; a <= 'z'; a++) {
            for (char b = 'a'; b <= 'z'; b++) {
                if (b == a) {
                    continue;
                }
                int f0 = 0, f1 = INT_MIN;
                for (char ch : s) {
                    if (ch == a) {
                        f0 = max(f0, 0) + 1;
                        f1++;
                    } else if (ch == b) {
                        f1 = f0 = max(f0, 0) - 1;
                    } // else f0 = max(f0, 0); 可以留到 ch 等于 a 或者 b 的时候计算，f1 不变
                    ans = max(ans, f1);
                }
            }
        }
        return ans;
    }
};
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

int largestVariance(char* s) {
    int ans = 0;
    for (char a = 'a'; a <= 'z'; a++) {
        for (char b = 'a'; b <= 'z'; b++) {
            if (b == a) {
                continue;
            }
            int f0 = 0, f1 = INT_MIN;
            for (int i = 0; s[i]; i++) {
                if (s[i] == a) {
                    f0 = MAX(f0, 0) + 1;
                    f1++;
                } else if (s[i] == b) {
                    f1 = f0 = MAX(f0, 0) - 1;
                } // else f0 = MAX(f0, 0); 可以留到 ch 等于 a 或者 b 的时候计算，f1 不变
                ans = MAX(ans, f1);
            }
        }
    }
    return ans;
}
```

```go [sol-Go]
func largestVariance(s string) (ans int) {
    for a := 'a'; a <= 'z'; a++ {
        for b := 'a'; b <= 'z'; b++ {
            if b == a {
                continue
            }
            f0, f1 := 0, math.MinInt
            for _, ch := range s {
                if ch == a {
                    f0 = max(f0, 0) + 1
                    f1++
                } else if ch == b {
                    f1, f0 = max(f0, 0)-1, max(f0, 0)-1
                } // else { f0 = max(f0, 0) } 可以留到 ch 等于 a 或者 b 的时候计算，f1 不变
                ans = max(ans, f1)
            }
        }
    }
    return
}
```

```js [sol-JavaScript]
var largestVariance = function(s) {
    let ans = 0;
    for (let a = 97; a <= 122; a++) { // 'a'.charCodeAt(0) === 97
        for (let b = 97; b <= 122; b++) {
            if (b === a) {
                continue;
            }
            let f0 = 0, f1 = -Infinity;
            for (const ch of s) {
                if (ch.charCodeAt(0) === a) {
                    f0 = Math.max(f0, 0) + 1;
                    f1++;
                } else if (ch.charCodeAt(0) === b) {
                    f1 = f0 = Math.max(f0, 0) - 1;
                } // else f0 = Math.max(f0, 0); 可以留到 ch 等于 a 或者 b 的时候计算，f1 不变
                ans = Math.max(ans, f1);
            }
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn largest_variance(s: String) -> i32 {
        let mut ans = 0;
        for a in b'a'..=b'z' {
            for b in b'a'..=b'z' {
                if b == a {
                    continue;
                }
                let mut f0 = 0;
                let mut f1 = i32::MIN;
                for ch in s.bytes() {
                    if ch == a {
                        f0 = f0.max(0) + 1;
                        f1 += 1;
                    } else if ch == b {
                        f1 = f0.max(0) - 1;
                        f0 = f1;
                    } // else { f0 = f0.max(0); } 可以留到 ch 等于 a 或者 b 的时候计算，f1 不变
                    ans = ans.max(f1);
                }
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|^2)$，其中 $|\Sigma|=26$ 为字符集合的大小。
- 空间复杂度：$\mathcal{O}(1)$。

## 优化

在上面的代码中，我们只在 $s[i]=a$ 或者 $s[i]=b$ 时才更新状态，因为其他情况下 $f$ 值是不变的。

这意味着，我们浪费了大量时间，在遍历那些既不等于 $a$ 又不等于 $b$ 的 $s[i]$ 上了。

珍惜时间，改为只遍历 $s$ 一次，同时更新所有会变化的状态。

创建两个 $26\times 26$ 的矩阵 $f_0[a][b]$ 和 $f_1[a][b]$，在遍历 $s$ 的同时，只更新那些会变化的状态，即 $a=s[i]$ 或者 $b=s[i]$ 的状态。

```py [sol-Python3]
class Solution:
    def largestVariance(self, s: str) -> int:
        ans = 0
        f0 = [[0] * 26 for _ in range(26)]
        f1 = [[-inf] * 26 for _ in range(26)]
        for ch in map(ord, s):
            ch -= ord('a')
            # 遍历到 ch 时，只需计算 a=ch 或者 b=ch 的状态，其他状态和 ch 无关，f 值不变
            for i in range(26):
                if i == ch:
                    continue
                # 假设出现次数最多的字母 a=ch，更新所有 b=i 的状态
                f0[ch][i] = max(f0[ch][i], 0) + 1
                f1[ch][i] += 1
                # 假设出现次数最少的字母 b=ch，更新所有 a=i 的状态
                f1[i][ch] = f0[i][ch] = max(f0[i][ch], 0) - 1
                ans = max(ans, f1[ch][i], f1[i][ch])
        return ans
```

```py [sol-Python3 手写 max]
class Solution:
    def largestVariance(self, s: str) -> int:
        ans = 0
        f0 = [[0] * 26 for _ in range(26)]
        f1 = [[-inf] * 26 for _ in range(26)]
        for ch in map(ord, s):
            ch -= ord('a')
            # 遍历到 ch 时，只需计算 a=ch 或者 b=ch 的状态，其他状态和 ch 无关，f 值不变
            for i in range(26):
                if i == ch:
                    continue
                # 假设出现次数最多的字母 a=ch，更新所有 b=i 的状态
                if f0[ch][i] < 0:
                    f0[ch][i] = 1
                else:
                    f0[ch][i] += 1
                f1[ch][i] += 1
                v = f1[ch][i]
                if v > ans:
                    ans = v
                # 假设出现次数最少的字母 b=ch，更新所有 a=i 的状态
                if f0[i][ch] < 0:
                    f0[i][ch] = -1
                else:
                    f0[i][ch] -= 1
                f1[i][ch] = v = f0[i][ch]
                if v > ans:
                    ans = v
        return ans
```

```java [sol-Java]
class Solution {
    public int largestVariance(String s) {
        int ans = 0;
        int[][] f0 = new int[26][26];
        int[][] f1 = new int[26][26];
        for (int[] row : f1) {
            Arrays.fill(row, Integer.MIN_VALUE);
        }

        for (char ch : s.toCharArray()) {
            ch -= 'a';
            // 遍历到 ch 时，只需计算 a=ch 或者 b=ch 的状态，其他状态和 ch 无关，f 值不变
            for (int i = 0; i < 26; i++) {
                if (i == ch) {
                    continue;
                }
                // 假设出现次数最多的字母 a=ch，更新所有 b=i 的状态
                f0[ch][i] = Math.max(f0[ch][i], 0) + 1;
                f1[ch][i]++;
                // 假设出现次数最少的字母 b=ch，更新所有 a=i 的状态
                f1[i][ch] = f0[i][ch] = Math.max(f0[i][ch], 0) - 1;
                ans = Math.max(ans, Math.max(f1[ch][i], f1[i][ch]));
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int largestVariance(string s) {
        int ans = 0;
        int f0[26][26]{}, f1[26][26];
        memset(f1, -0x3f, sizeof(f1)); // 初始化成一个很小的负数

        for (char ch : s) {
            ch -= 'a';
            // 遍历到 ch 时，只需计算 a=ch 或者 b=ch 的状态，其他状态和 ch 无关，f 值不变
            for (int i = 0; i < 26; i++) {
                if (i == ch) {
                    continue;
                }
                // 假设出现次数最多的字母 a=ch，更新所有 b=i 的状态
                f0[ch][i] = max(f0[ch][i], 0) + 1;
                f1[ch][i]++;
                // 假设出现次数最少的字母 b=ch，更新所有 a=i 的状态
                f1[i][ch] = f0[i][ch] = max(f0[i][ch], 0) - 1;
                ans = max(ans, max(f1[ch][i], f1[i][ch])); // 或者 max({ans, f1[ch][i], f1[i][ch]})
            }
        }
        return ans;
    }
};
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

int largestVariance(char* s) {
    int ans = 0;
    int f0[26][26] = {}, f1[26][26];
    memset(f1, -0x3f, sizeof(f1)); // 初始化成一个很小的负数

    for (int k = 0; s[k]; k++) {
        int ch = s[k] - 'a';
        // 遍历到 ch 时，只需计算 a=ch 或者 b=ch 的状态，其他状态和 ch 无关，f 值不变
        for (int i = 0; i < 26; i++) {
            if (i == ch) {
                continue;
            }
            // 假设出现次数最多的字母 a=ch，更新所有 b=i 的状态
            f0[ch][i] = MAX(f0[ch][i], 0) + 1;
            f1[ch][i]++;
            // 假设出现次数最少的字母 b=ch，更新所有 a=i 的状态
            f1[i][ch] = f0[i][ch] = MAX(f0[i][ch], 0) - 1;
            ans = MAX(ans, MAX(f1[ch][i], f1[i][ch]));
        }
    }
    return ans;
}
```

```go [sol-Go]
func largestVariance(s string) (ans int) {
	var f0, f1 [26][26]int
	for i := range f1 {
		for j := range f1[i] {
			f1[i][j] = math.MinInt
		}
	}

	for _, ch := range s {
		ch -= 'a'
		// 遍历到 ch 时，只需计算 a=ch 或者 b=ch 的状态，其他状态和 ch 无关，f 值不变
		for i := range 26 {
			if i == int(ch) {
				continue
			}
			// 假设出现次数最多的字母 a=ch，更新所有 b=i 的状态
			f0[ch][i] = max(f0[ch][i], 0) + 1
			f1[ch][i]++
			// 假设出现次数最少的字母 b=ch，更新所有 a=i 的状态
			f0[i][ch] = max(f0[i][ch], 0) - 1
			f1[i][ch] = f0[i][ch]
			ans = max(ans, f1[ch][i], f1[i][ch])
		}
	}
	return
}
```

```js [sol-JavaScript]
var largestVariance = function(s) {
    let ans = 0;
    const f0 = Array.from({ length: 26 }, () => Array(26).fill(0));
    const f1 = Array.from({ length: 26 }, () => Array(26).fill(-Infinity));

    for (let ch of s) {
        ch = ch.charCodeAt(0) - 97; // 'a'.charCodeAt(0) === 97
        // 遍历到 ch 时，只需计算 a=ch 或者 b=ch 的状态，其他状态和 ch 无关，f 值不变
        for (let i = 0; i < 26; i++) {
            if (i === ch) {
                continue;
            }
            // 假设出现次数最多的字母 a=ch，更新所有 b=i 的状态
            f0[ch][i] = Math.max(f0[ch][i], 0) + 1;
            f1[ch][i]++;
            // 假设出现次数最少的字母 b=ch，更新所有 a=i 的状态
            f1[i][ch] = f0[i][ch] = Math.max(f0[i][ch], 0) - 1;
            ans = Math.max(ans, f1[ch][i], f1[i][ch]);
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn largest_variance(s: String) -> i32 {
        let mut ans = 0;
        let mut f0 = [[0; 26]; 26];
        let mut f1 = [[i32::MIN; 26]; 26];

        for ch in s.bytes() {
            let ch = (ch - b'a') as usize;
            // 遍历到 ch 时，只需计算 a=ch 或者 b=ch 的状态，其他状态和 ch 无关，f 值不变
            for i in 0..26 {
                if i == ch {
                    continue;
                }
                // 假设出现次数最多的字母 a=ch，更新所有 b=i 的状态
                f0[ch][i] = f0[ch][i].max(0) + 1;
                f1[ch][i] += 1;
                // 假设出现次数最少的字母 b=ch，更新所有 a=i 的状态
                f0[i][ch] = f0[i][ch].max(0) - 1;
                f1[i][ch] = f0[i][ch];
                ans = ans.max(f1[ch][i]).max(f1[i][ch]);
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|)$，其中 $|\Sigma|=26$ 为字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|^2)$。

更多相似题目，见下面动态规划题单中的「**§1.3 最大子数组和**」以及「**六、状态机 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
