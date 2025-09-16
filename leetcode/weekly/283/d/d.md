如果有三个相邻且可以合并的数 $x,y,z$，那么先合并 $x$ 和 $y$ 再合并 $z$，还是先合并 $y$ 和 $z$ 再合并 $x$，结果是一样的吗？

是一样的，由于 $\text{LCM}$ 本质是质因数分解中质数的指数取最大值（见初等数论教材），对于 $\max(a,b,c)$ 来说，我们有

$$
\max(a,b,c) = \max(\max(a,b),c) = \max(a,\max(b,c))
$$

所以 $x,y,z$ 先合并 $x,y$ 还是先合并 $y,z$ 都可以。得到的结果均为 $\text{LCM}(x,y,z)$。

该结论可以推广到更多元素的情况。

据此，对于任意一种合并顺序，我们总是可以将该顺序**重排**成：

- 优先选择最左边的能合并的相邻元素，将其合并。

这可以用栈模拟：

1. 创建一个空栈。
2. 从左到右遍历 $\textit{nums}$。
3. 设 $x = \textit{nums}[i]$。如果栈不为空且 $x$ 与栈顶不互质，那么弹出栈顶 $y$，更新 $x$ 为 $\text{LCM}(x,y)$。循环直到栈为空或者 $x$ 与栈顶互质。
4. 把 $x$ 入栈。
5. 遍历结束，栈即为答案。

```py [sol-Python3]
class Solution:
    def replaceNonCoprimes(self, nums: List[int]) -> List[int]:
        st = []
        for x in nums:
            while st and gcd(x, st[-1]) > 1:
                x = lcm(x, st.pop())
            st.append(x)
        return st
```

```java [sol-Java]
class Solution {
    public List<Integer> replaceNonCoprimes(int[] nums) {
        List<Integer> st = new ArrayList<>();
        for (int x : nums) {
            while (!st.isEmpty() && gcd(x, st.getLast()) > 1) {
                x = lcm(x, st.removeLast());
            }
            st.add(x);
        }
        return st;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }

    // 先除后乘，避免溢出
    private int lcm(int a, int b) {
        return a / gcd(a, b) * b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> replaceNonCoprimes(vector<int>& nums) {
        vector<int> st;
        for (int x : nums) {
            while (!st.empty() && gcd(x, st.back()) > 1) {
                x = lcm(x, st.back());
                st.pop_back();
            }
            st.push_back(x);
        }
        return st;
    }
};
```

```c [sol-C]
int gcd(int a, int b) {
    while (a != 0) {
        int tmp = a;
        a = b % a;
        b = tmp;
    }
    return b;
}

// 先除后乘，避免溢出
int lcm(int a, int b) {
    return a / gcd(a, b) * b;
}

int* replaceNonCoprimes(int* nums, int numsSize, int* returnSize) {
    int* stack = nums; // 把 nums 当作栈用
    int top = -1; // 栈顶下标
    for (int i = 0; i < numsSize; i++) {
        int x = nums[i];
        while (top >= 0 && gcd(x, stack[top]) > 1) {
            x = lcm(x, stack[top--]); // 出栈
        }
        stack[++top] = x; // 入栈
    }
    *returnSize = top + 1;
    return stack;
}
```

```go [sol-Go]
func replaceNonCoprimes(nums []int) []int {
	st := nums[:0] // 把 nums 当作栈用
	for _, x := range nums {
		for len(st) > 0 && gcd(x, st[len(st)-1]) > 1 {
			x = lcm(x, st[len(st)-1])
			st = st[:len(st)-1]
		}
		st = append(st, x)
	}
	return st
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
```

```js [sol-JavaScript]
var replaceNonCoprimes = function(nums) {
    const st = [];
    for (let x of nums) {
        while (st.length > 0 && gcd(x, st[st.length - 1]) > 1) {
            x = lcm(x, st.pop());
        }
        st.push(x);
    }
    return st;
};

function gcd(a, b) {
    while (a !== 0) {
        const tmp = a;
        a = b % a;
        b = tmp;
    }
    return b;
};

function lcm(a, b) {
    return a / gcd(a, b) * b;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn replace_non_coprimes(nums: Vec<i32>) -> Vec<i32> {
        fn gcd(mut a: i32, mut b: i32) -> i32 {
            while a != 0 {
                let tmp = a;
                a = b % a;
                b = tmp;
            }
            b
        }
        fn lcm(a: i32, b: i32) -> i32 {
            a / gcd(a, b) * b
        }

        let mut st = vec![];
        for mut x in nums {
            while !st.is_empty() && gcd(x, *st.last().unwrap()) > 1 {
                x = lcm(x, st.pop().unwrap());
            }
            st.push(x);
        }
        st
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$。其中 $n$ 是 $\textit{nums}$ 的长度，题目保证 $U\le 10^8$。由于每个元素至多入栈出栈各一次，所以二重循环的总循环次数是 $\mathcal{O}(n)$ 的。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。如果把 $\textit{nums}$ 当作栈用，可以做到 $\mathcal{O}(1)$ 空间。

## 专题训练

见下面数据结构题单的「**§3.3 邻项消除**」。

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
