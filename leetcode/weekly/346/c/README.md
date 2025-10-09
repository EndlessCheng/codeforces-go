把 $i^2$ 转成字符串 $s$，然后写一个递归，枚举 $s$ 分割出的第一个子串、第二个子串、……，把每个子串对应的整数值 $x$ 加到 $\textit{sum}$ 中。

递归到终点时，如果 $\textit{sum}=i$，则说明 $i$ 符合要求。

如果你不清楚怎么写这个递归，请看 [回溯算法套路①子集型回溯【基础算法精讲 14】](https://www.bilibili.com/video/BV1mG4y1A7Gu/)。

代码实现时，可以预处理 $[1,1000]$ 内的所有数的惩罚数。

[视频讲解](https://www.bilibili.com/video/BV1Qm4y1t7cx/) 第三题。

## 写法一：使用字符串

```py [sol-Python3]
PRE_SUM = [0] * 1001
for i in range(1, 1001):
    s = str(i * i)
    n = len(s)
    def dfs(p: int, sum: int) -> bool:
        if p == n:  # 递归终点
            return sum == i  # i 符合要求
        x = 0
        for j in range(p, n):  # 枚举分割出从 s[p] 到 s[j] 的子串
            x = x * 10 + int(s[j])  # 子串对应的整数值
            if dfs(j + 1, sum + x):
                return True
        return False
    PRE_SUM[i] = PRE_SUM[i - 1] + (i * i if dfs(0, 0) else 0)

class Solution:
    def punishmentNumber(self, n: int) -> int:
        return PRE_SUM[n]
```

```java [sol-Java]
class Solution {
    private static final int[] PRE_SUM = new int[1001];

    static {
        for (int i = 1; i <= 1000; i++) {
            char[] s = Integer.toString(i * i).toCharArray();
            PRE_SUM[i] = PRE_SUM[i - 1] + (dfs(s, i, 0, 0) ? i * i : 0);
        }
    }

    private static boolean dfs(char[] s, int i, int p, int sum) {
        if (p == s.length) { // 递归终点
            return sum == i; // i 符合要求
        }
        int x = 0;
        for (int j = p; j < s.length; j++) { // 枚举分割出从 s[p] 到 s[j] 的子串
            x = x * 10 + s[j] - '0'; // 子串对应的整数值
            if (dfs(s, i, j + 1, sum + x)) {
                return true;
            }
        }
        return false;
    }

    public int punishmentNumber(int n) {
        return PRE_SUM[n];
    }
}
```

```cpp [sol-C++]
int PRE_SUM[1001];

int init = []() {
    for (int i = 1; i <= 1000; i++) {
        string s = to_string(i * i);
        int n = s.length();
        auto dfs = [&](this auto&& dfs, int p, int sum) -> bool {
            if (p == n) { // 递归终点
                return sum == i; // i 符合要求
            }
            int x = 0;
            for (int j = p; j < n; j++) { // 枚举分割出从 s[p] 到 s[j] 的子串
                x = x * 10 + s[j] - '0'; // 子串对应的整数值
                if (dfs(j + 1, sum + x)) {
                    return true;
                }
            }
            return false;
        };
        PRE_SUM[i] = PRE_SUM[i - 1] + (dfs(0, 0) ? i * i : 0);
    }
    return 0;
}();

class Solution {
public:
    int punishmentNumber(int n) {
        return PRE_SUM[n];
    }
};
```

```go [sol-Go]
var preSum [1001]int

func init() {
	for i := 1; i <= 1000; i++ {
		s := strconv.Itoa(i * i)
		n := len(s)
		var dfs func(int, int) bool
		dfs = func(p, sum int) bool {
			if p == n { // 递归终点
				return sum == i // i 符合要求
			}
			x := 0
			for j := p; j < n; j++ { // 枚举分割出从 s[p] 到 s[j] 的子串
				x = x*10 + int(s[j]-'0') // 子串对应的整数值
				if dfs(j+1, sum+x) {
					return true
				}
			}
			return false
		}
		preSum[i] = preSum[i-1]
		if dfs(0, 0) { // i 符合要求
			preSum[i] += i * i // 计算前缀和
		}
	}
}

func punishmentNumber(n int) int {
	return preSum[n]
}
```

```js [sol-JavaScript]
function dfs(p, sum, s, i) {
    const n = s.length;
    if (p === n) { // 递归终点
        return sum === i; // i 符合要求
    }
    let x = 0;
    for (let j = p; j < n; j++) { // 枚举分割出从 s[p] 到 s[j] 的子串
        x = x * 10 + parseInt(s[j]); // 子串对应的整数值
        if (dfs(j + 1, sum + x, s, i)) {
            return true;
        }
    }
    return false;
}

const PRE_SUM = new Array(1001).fill(0);
for (let i = 1; i <= 1000; i++) {
    const s = (i * i).toString();
    PRE_SUM[i] = PRE_SUM[i - 1] + (dfs(0, 0, s, i) ? i * i : 0);
}

var punishmentNumber = function (n) {
    return PRE_SUM[n];
};
```

```rust [sol-Rust]
fn dfs(p: usize, sum: i32, i: i32, s: &Vec<u8>) -> bool {
    if p == s.len() { // 递归终点
        return sum == i; // i 符合要求
    }
    let mut x = 0;
    for j in p..s.len() { // 枚举分割出从 s[p] 到 s[j] 的子串
        x = x * 10 + (s[j] & 0xf) as i32; // 子串对应的整数值
        if dfs(j + 1, sum + x, i, s) {
            return true;
        }
    }
    false
}

static mut initialized: bool = false;
static mut pre_sum: [i32; 1001] = [0; 1001];

fn init_once() {
    unsafe {
        if initialized { // 之前初始化过了
            return;
        }
        initialized = true;
        for i in 1..1001 {
            let s = (i * i).to_string().bytes().collect();
            pre_sum[i as usize] = pre_sum[i as usize - 1] + if dfs(0, 0, i, &s) { i * i } else { 0 };
        }
    }
}

impl Solution {
    pub fn punishment_number(n: i32) -> i32 {
        init_once();
        unsafe { pre_sum[n as usize] }
    }
}
```

## 写法二：不使用字符串

```py [sol-Python3]
PRE_SUM = [0] * 1001
for i in range(1, 1001):
    def dfs(val: int, sum: int) -> bool:
        if val == 0:  # 递归终点
            return sum == i  # i 符合要求
        x = 0
        pow10 = 1
        while val:
            val, d = divmod(val, 10)
            x += d * pow10
            if dfs(val, sum + x):
                return True
            pow10 *= 10
        return False
    PRE_SUM[i] = PRE_SUM[i - 1] + (i * i if dfs(i * i, 0) else 0)

class Solution:
    def punishmentNumber(self, n: int) -> int:
        return PRE_SUM[n]
```

```java [sol-Java]
class Solution {
    private static final int[] PRE_SUM = new int[1001];

    static {
        for (int i = 1; i <= 1000; i++) {
            PRE_SUM[i] = PRE_SUM[i - 1] + (dfs(i * i, i) ? i * i : 0);
        }
    }

    private static boolean dfs(int val, int sum) {
        if (val == 0) { // 递归终点
            return sum == 0; // i 符合要求
        }
        int pow10 = 1;
        for (int x = 0; val > 0; val /= 10) {
            x += val % 10 * pow10;
            if (dfs(val / 10, sum - x)) {
                return true;
            }
            pow10 *= 10;
        }
        return false;
    }

    public int punishmentNumber(int n) {
        return PRE_SUM[n];
    }
}
```

```cpp [sol-C++]
int PRE_SUM[1001];

int init = []() {
    for (int i = 1; i <= 1000; i++) {
        auto dfs = [&](this auto&& dfs, int val, int sum) -> bool {
            if (val == 0) { // 递归终点
                return sum == i; // i 符合要求
            }
            int pow10 = 1;
            for (int x = 0; val > 0; val /= 10) {
                x += val % 10 * pow10;
                if (dfs(val / 10, sum + x)) {
                    return true;
                }
                pow10 *= 10;
            }
            return false;
        };
        PRE_SUM[i] = PRE_SUM[i - 1] + (dfs(i * i, 0) ? i * i : 0);
    }
    return 0;
}();

class Solution {
public:
    int punishmentNumber(int n) {
        return PRE_SUM[n];
    }
};
```

```go [sol-Go]
var preSum [1001]int

func init() {
	for i := 1; i <= 1000; i++ {
		var dfs func(int, int) bool
		dfs = func(val, sum int) bool {
			if val == 0 { // 递归终点
				return sum == i // i 符合要求
			}
			for x, pow10 := 0, 1; val > 0; val /= 10 {
				x += val % 10 * pow10
				if dfs(val/10, sum+x) {
					return true
				}
				pow10 *= 10
			}
			return false
		}
		preSum[i] = preSum[i-1]
		if dfs(i*i, 0) { // i 符合要求
			preSum[i] += i * i // 计算前缀和
		}
	}
}

func punishmentNumber(n int) int {
	return preSum[n]
}
```

```js [sol-JavaScript]
function dfs(val, sum) {
    if (val === 0) { // 递归终点
        return sum === 0; // i 符合要求
    }
    let pow10 = 1;
    for (let x = 0; val > 0; val = Math.floor(val / 10)) {
        x += val % 10 * pow10;
        if (dfs(Math.floor(val / 10), sum - x)) {
            return true;
        }
        pow10 *= 10;
    }
    return false;
}

const PRE_SUM = new Array(1001).fill(0);
for (let i = 1; i <= 1000; i++) {
    PRE_SUM[i] = PRE_SUM[i - 1] + (dfs(i * i, i) ? i * i : 0);
}

var punishmentNumber = function (n) {
    return PRE_SUM[n];
};
```

```rust [sol-Rust]
fn dfs(mut val: i32, mut sum: i32) -> bool {
    if val == 0 { // 递归终点
        return sum == 0; // i 符合要求
    }
    let mut pow10 = 1;
    let mut x = 0;
    while val > 0 {
        x += val % 10 * pow10;
        val /= 10;
        if (dfs(val, sum - x)) {
            return true;
        }
        pow10 *= 10;
    }
    false
}

static mut initialized: bool = false;
static mut pre_sum: [i32; 1001] = [0; 1001];

fn init_once() {
    unsafe {
        if initialized { // 之前初始化过了
            return;
        }
        initialized = true;
        for i in 1..1001 {
            pre_sum[i as usize] = pre_sum[i as usize - 1] + if dfs(i * i, i) { i * i } else { 0 };
        }
    }
}

impl Solution {
    pub fn punishment_number(n: i32) -> i32 {
        init_once();
        unsafe { pre_sum[n as usize] }
    }
}
```

#### 复杂度分析

预处理的时间：$\mathcal{O}(U^{1 + 2\log_{10} 2})\approx\mathcal{O}(U^{1.602})$，其中 $U=1000$。对于数字 $i^2$，它的十进制字符串的长度为 $m=\lfloor1+2\log_{10} i\rfloor$。我在【基础算法精讲】中讲过，划分型题目的本质就是枚举子集，所以递归需要 $\mathcal{O}(2^m)=\mathcal{O}(i^{2\log_{10} 2})$ 的时间，对其积分可知，整个预处理需要 $\mathcal{O}(U^{1 + 2\log_{10} 2})$ 的时间。

力扣的计时规则是，预处理的时间不计入，所以两种写法都是 $\mathcal{O}(1)$ 时间。

- 时间复杂度：$\mathcal{O}(1)$
- 空间复杂度：$\mathcal{O}(1)$。

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
