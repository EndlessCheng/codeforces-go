## 方法一：从左到右遍历

从 $i$ 到 $j$：

- 如果不跨过 $\textit{words}$ 的首尾，那么距离为 $|i-j|$。
- 如果跨过 $\textit{words}$ 的首尾呢？在地球上，从 $A$ 到 $B$，倒着绕地球走一圈，移动距离为一圈的长度，减去直接从 $A$ 到 $B$ 的距离，所以跨过 $\textit{words}$ 的首尾的距离为 $n - |i-j|$。

一般地，遍历 $\textit{words}$，如果 $\textit{words}[i] = \textit{target}$，那么：

- 从 $\textit{startIndex}$ 开始移动，不跨过 $\textit{words}$ 的首尾，距离为 $|i-\textit{startIndex}|$。
- 从 $\textit{startIndex}$ 开始移动，跨过 $\textit{words}$ 的首尾，距离为 $n - |i-\textit{startIndex}|$。

两种情况取最小值，去更新答案的最小值。

```py [sol-Python3]
class Solution:
    def closestTarget(self, words: List[str], target: str, startIndex: int) -> int:
        ans = n = len(words)
        for i, word in enumerate(words):
            if word == target:
                d = abs(i - startIndex)
                ans = min(ans, d, n - d)
        return -1 if ans == n else ans
```

```java [sol-Java]
class Solution {
    public int closestTarget(String[] words, String target, int startIndex) {
        int n = words.length;
        int ans = n;
        for (int i = 0; i < n; i++) {
            if (words[i].equals(target)) {
                int d = Math.abs(i - startIndex);
                ans = Math.min(ans, Math.min(d, n - d));
            }
        }
        return ans == n ? -1 : ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int closestTarget(vector<string>& words, string target, int startIndex) {
        int n = words.size();
        int ans = n;
        for (int i = 0; i < n; i++) {
            if (words[i] == target) {
                int d = abs(i - startIndex);
                ans = min(ans, min(d, n - d));
            }
        }
        return ans == n ? -1 : ans;
    }
};
```

```c [sol-C]
int closestTarget(char** words, int wordsSize, char* target, int startIndex) {
    int n = wordsSize;
    int ans = n;
    for (int i = 0; i < n; i++) {
        if (strcmp(words[i], target) == 0) {
            int d = abs(i - startIndex);
            ans = MIN(ans, MIN(d, n - d));
        }
    }
    return ans == n ? -1 : ans;
}
```

```go [sol-Go]
func closestTarget(words []string, target string, startIndex int) int {
	n := len(words)
	ans := n
	for i, word := range words {
		if word == target {
			d := abs(i - startIndex)
			ans = min(ans, d, n-d)
		}
	}
	if ans == n {
		return -1
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

```js [sol-JavaScript]
var closestTarget = function(words, target, startIndex) {
    const n = words.length;
    let ans = n;
    for (let i = 0; i < n; i++) {
        if (words[i] === target) {
            const d = Math.abs(i - startIndex);
            ans = Math.min(ans, d, n - d);
        }
    }
    return ans === n ? -1 : ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn closest_target(words: Vec<String>, target: String, start_index: i32) -> i32 {
        let n = words.len() as i32;
        let mut ans = n;
        for (i, word) in words.into_iter().enumerate() {
            if word == target {
                let d = (i as i32 - start_index).abs();
                ans = ans.min(d).min(n - d);
            }
        }
        if ans == n { -1 } else { ans }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 是 $\textit{words}$ 的长度，$m$ 是 $\textit{target}$ 的长度。比较字符串是否相等需要 $\mathcal{O}(m)$ 的时间。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：从 startIndex 往两侧遍历

枚举答案 $k=0,1,2,\ldots,\left\lfloor\dfrac{n}{2}\right\rfloor$，判断 $\textit{nums}[\textit{startIndex}-k] = \textit{target}$ 或者 $\textit{nums}[\textit{startIndex}+k] = \textit{target}$ 是否成立，如果成立，直接返回 $k$。

由于 $\textit{words}$ 是环形数组，下标 $-1$ 即下标 $n-1$，下标 $-2$ 即下标 $n-2$，依此类推，通过计算 $(i+n)\bmod n$，任意下标 $i$ 可以通过取模调整到 $[0,n-1]$ 中。

```py [sol-Python3]
class Solution:
    def closestTarget(self, words: list[str], target: str, startIndex: int) -> int:
        n = len(words)
        for k in range(n // 2 + 1):
            if words[startIndex - k] == target or words[(startIndex + k) % n] == target:
                return k
        return -1
```

```java [sol-Java]
class Solution {
    public int closestTarget(String[] words, String target, int startIndex) {
        int n = words.length;
        for (int k = 0; k <= n / 2; k++) {
            if (words[(startIndex - k + n) % n].equals(target) || words[(startIndex + k) % n].equals(target)) {
                return k;
            }
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int closestTarget(vector<string>& words, string target, int startIndex) {
        int n = words.size();
        for (int k = 0; k <= n / 2; k++) {
            if (words[(startIndex - k + n) % n] == target || words[(startIndex + k) % n] == target) {
                return k;
            }
        }
        return -1;
    }
};
```

```c [sol-C]
int closestTarget(char** words, int wordsSize, char* target, int startIndex) {
    int n = wordsSize;
    for (int k = 0; k <= n / 2; k++) {
        if (strcmp(words[(startIndex - k + n) % n], target) == 0 ||
            strcmp(words[(startIndex + k) % n], target) == 0) {
            return k;
        }
    }
    return -1;
}
```

```go [sol-Go]
func closestTarget(words []string, target string, startIndex int) int {
	n := len(words)
	for k := range n/2 + 1 {
		if words[(startIndex-k+n)%n] == target || words[(startIndex+k)%n] == target {
			return k
		}
	}
	return -1
}
```

```js [sol-JavaScript]
var closestTarget = function(words, target, startIndex) {
    const n = words.length;
    for (let k = 0; k <= n / 2; k++) {
        if (words[(startIndex - k + n) % n] === target || words[(startIndex + k) % n] === target) {
            return k;
        }
    }
    return -1;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn closest_target(words: Vec<String>, target: String, start_index: i32) -> i32 {
        let n = words.len();
        let start_index = start_index as usize;
        for k in 0..=n / 2 {
            if words[(start_index + n - k) % n] == target || words[(start_index + k) % n] == target {
                return k as _;
            }
        }
        -1
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 是 $\textit{words}$ 的长度，$m$ 是 $\textit{target}$ 的长度。比较字符串是否相等需要 $\mathcal{O}(m)$ 的时间。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

[1848. 到目标元素的最小距离](https://leetcode.cn/problems/minimum-distance-to-the-target-element/)

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
