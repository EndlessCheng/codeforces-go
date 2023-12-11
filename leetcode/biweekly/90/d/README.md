## 前置知识：单调栈

请看视频讲解：[单调栈【基础算法精讲 26】](https://www.bilibili.com/video/BV1VN411J7S7/)，欢迎点赞关注~

## 思路

本题要找「下下个更大元素」。注意不是下一个更大元素的下一个更大元素（$a<x<y$ 找 $y$），而是右边第二个大于当前元素的数（$a<x$ 且 $a<y$ 找 $y$）。

首先回想一下，怎么找下一个更大元素，即右侧最近的更大元素。

使用视频中讲的第二种做法，从左到右遍历数组，用一个（递减）单调栈 $s$ 维护遍历过的元素，如果当前元素 $x$ 比栈顶大，那么栈顶的下一个更大元素就是 $x$，并弹出栈顶。

在这个做法中，栈顶元素弹出后，就没有用了。但对于本题，我们需要的正是弹出去的元素！

再用一个（递减）单调栈 $t$ 记录从 $s$ 中弹出去的元素。继续向后遍历，如果又找到一个元素 $y$ 比 $t$ 的栈顶大，那么栈顶的**下下个**更大元素就是 $y$，并弹出栈顶。

以示例 1 为例，其中 $s$ 和 $t$ 维护每轮循环**结束**时的栈中元素，$\textit{ans}$ 为答案：

|  $i$ | $\textit{nums}[i]$  | $s$  | $t$ | $\textit{ans}$ |
|---|:---:|---|---|---|
|  $0$ | $2$  | $[2]$  | $[]$  |  $[-1,-1,-1,-1,-1]$|
|  $1$ | $4$  | $[4]$  | $[2]$  |  $[-1,-1,-1,-1,-1]$|
|  $2$ | $0$  | $[4,0]$  | $[2]$  | $[-1,-1,-1,-1,-1]$|
|  $3$ | $9$  | $[9]$  | $[4,0]$  | $[9, -1, -1, -1, -1]$|
|  $4$ | $6$  | $[9,6]$  | $[]$  | $[9, 6, 6, -1, -1]$|

遍历到 $9$ 时，发现它比 $t$ 的栈顶 $2$ 大，说明 $2$ 在找到一个比 $2$ 大的元素 $4$ 后，又找到了一个比 $2$ 大的元素 $9$，即 $2$ 的下下个更大元素是 $9$，所以 $\textit{ans}[0] = 9$。

遍历到 $6$ 时，发现它比 $t$ 的栈顶 $4$ 和 $0$ 大，同理，$\textit{ans}[1] = \textit{ans}[2] = 6$。

## 算法

1. 初始化 $\textit{ans}$ 数组，长度和 $\textit{nums}$ 相同，初始值均为 $-1$。
2. 初始化两个空栈 $s$ 和 $t$。
3. 从左到右遍历 $\textit{nums}$。
4. 设 $x=\textit{nums}[i]$。如果 $x$ 比 $t$ 的栈顶大，则栈顶元素对应的答案为 $x$，并弹出 $t$ 的栈顶。
5. 如果 $x$ 比 $s$ 的栈顶大，则弹出 $s$ 的栈顶，并移入 $t$ 的栈顶。由于要保证 $t$ 也是一个递减的单调栈，我们可以直接把一整段从 $s$ 中弹出的元素插入到 $t$ 的末尾。
6. 把 $x$ 加到 $s$ 的栈顶。代码实现时，为方便记录答案，改为把 $x$ 的下标 $i$ 加到栈顶。

## 答疑

**问**：为什么可以直接把从 $s$ 中弹出的元素加到 $t$，如果加入的数比 $t$ 的栈顶大，不就破坏了 $t$ 的单调性吗？

**答**：$t$ 仍然是单调的。注意在弹出 $s$ 栈顶之前，我们已经把 $t$ 栈顶的小于 $x$ 的数都弹出了，而后面从 $s$ 栈顶弹出的数，都是比 $x$ 小的，而 $x\le$ 此时 $t$ 的栈顶，根据不等式的传递性，从 $s$ 栈顶弹出的数，也同样比 $t$ 的栈顶小，所以这些数入栈后，$t$ 仍然是单调的。

```py [sol-Python3]
class Solution:
    def secondGreaterElement(self, nums: List[int]) -> List[int]:
        ans = [-1] * len(nums)
        s = []
        t = []
        for i, x in enumerate(nums):
            while t and nums[t[-1]] < x:
                ans[t.pop()] = x  # t 栈顶的下下个更大元素是 x
            j = len(s) - 1
            while j >= 0 and nums[s[j]] < x:
                j -= 1  # s 栈顶的下一个更大元素是 x
            t += s[j + 1:]  # 把从 s 弹出的这一整段元素加到 t
            del s[j + 1:]  # 弹出一整段元素
            s.append(i)  # 当前元素（的下标）加到 s 栈顶
        return ans
```

```java [sol-Java]
// 更快的写法见【数组】版本
class Solution {
    public int[] secondGreaterElement(int[] nums) {
        int n = nums.length;
        int[] ans = new int[n];
        Arrays.fill(ans, -1);
        List<Integer> s = new ArrayList<>();
        List<Integer> t = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            while (!t.isEmpty() && nums[t.get(t.size() - 1)] < x) {
                ans[t.get(t.size() - 1)] = x; // t 栈顶的下下个更大元素是 x
                t.remove(t.size() - 1);
            }
            int j = s.size();
            while (j > 0 && nums[s.get(j - 1)] < x) {
                j--; // s 栈顶的下一个更大元素是 x
            }
            List<Integer> popped = s.subList(j, s.size());
            t.addAll(popped); // 把从 s 弹出的这一整段元素加到 t
            popped.clear(); // 弹出一整段元素
            s.add(i); // 当前元素（的下标）加到 s 栈顶
        }
        return ans;
    }
}
```

```java [sol-Java 数组]
class Solution {
    private static final int MX = 100000;
    private static final int[] s = new int[MX];
    private static final int[] t = new int[MX];

    public int[] secondGreaterElement(int[] nums) {
        int n = nums.length, lenS = 0, lenT = 0;
        int[] ans = new int[n];
        Arrays.fill(ans, -1);
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            while (lenT > 0 && nums[t[lenT - 1]] < x) {
                ans[t[--lenT]] = x; // t 栈顶的下下个更大元素是 x
            }
            int tmp = lenS;
            while (lenS > 0 && nums[s[lenS - 1]] < x) {
                lenS--; // s 栈顶的下一个更大元素是 x
            }
            System.arraycopy(s, lenS, t, lenT, tmp - lenS); // 把从 s 弹出的这一整段元素加到 t
            lenT += tmp - lenS;
            s[lenS++] = i; // 当前元素（的下标）加到 s 栈顶
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> secondGreaterElement(vector<int> &nums) {
        int n = nums.size();
        vector<int> ans(n, -1), s, t;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            while (!t.empty() && nums[t.back()] < x) {
                ans[t.back()] = x; // t 栈顶的下下个更大元素是 x
                t.pop_back();
            }
            int j = s.size();
            while (j && nums[s[j - 1]] < x) {
                j--; // s 栈顶的下一个更大元素是 x
            }
            t.insert(t.end(), s.begin() + j, s.end()); // 把从 s 弹出的这一整段元素加到 t
            s.resize(j); // 弹出一整段元素
            s.push_back(i); // 当前元素（的下标）加到 s 栈顶
        }
        return ans;
    }
};
```

```go [sol-Go]
func secondGreaterElement(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}
	s, t := []int{}, []int{}
	for i, x := range nums {
		for len(t) > 0 && nums[t[len(t)-1]] < x {
			ans[t[len(t)-1]] = x // t 栈顶的下下个更大元素是 x
			t = t[:len(t)-1]
		}
		j := len(s) - 1
		for j >= 0 && nums[s[j]] < x {
			j-- // s 栈顶的下一个更大元素是 x
		}
		t = append(t, s[j+1:]...) // 把从 s 弹出的这一整段元素加到 t
		s = append(s[:j+1], i) // 当前元素（的下标）加到 s 栈顶
	}
	return ans
}
```

```js [sol-JavaScript]
var secondGreaterElement = function (nums) {
    const n = nums.length;
    const ans = Array(n).fill(-1);
    const s = [];
    const t = [];
    for (let i = 0; i < n; i++) {
        const x = nums[i];
        while (t.length > 0 && nums[t[t.length - 1]] < x) {
            ans[t.pop()] = x; // t 栈顶的下下个更大元素是 x
        }
        let j = s.length - 1;
        while (j >= 0 && nums[s[j]] < x) {
            j--; // s 栈顶的下一个更大元素是 x
        }
        t.push(...s.splice(j + 1)); // 把从 s 弹出的这一整段元素加到 t
        s.push(i); // 当前元素（的下标）加到 s 栈顶
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn second_greater_element(nums: Vec<i32>) -> Vec<i32> {
        let mut ans = vec![-1; nums.len()];
        let mut s = Vec::new();
        let mut t = Vec::new();
        for (i, &x) in nums.iter().enumerate() {
            while !t.is_empty() && nums[*t.last().unwrap()] < x {
                ans[t.pop().unwrap()] = x; // t 栈顶的下下个更大元素是 x
            }
            let mut j = s.len();
            while j > 0 && nums[s[j - 1]] < x {
                j -= 1; // s 栈顶的下一个更大元素是 x
            }
            t.extend(s.drain(j..)); // 把从 s 弹出的这一整段元素加到 t
            s.push(i); // 当前元素（的下标）加到 s 栈顶
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。每个元素至多入栈出栈各两次。
- 空间复杂度：$\mathcal{O}(n)$。

## 题单

#### 单调栈

- [739. 每日温度](https://leetcode.cn/problems/daily-temperatures/)（单调栈模板题）
- [1475. 商品折扣后的最终价格](https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop/)
- [496. 下一个更大元素 I](https://leetcode.cn/problems/next-greater-element-i/)
- [503. 下一个更大元素 II](https://leetcode.cn/problems/next-greater-element-ii/)
- [1019. 链表中的下一个更大节点](https://leetcode.cn/problems/next-greater-node-in-linked-list/) 1571
- [901. 股票价格跨度](https://leetcode.cn/problems/online-stock-span/) 1709
- [1124. 表现良好的最长时间段](https://leetcode.cn/problems/longest-well-performing-interval/) 1908
- [456. 132 模式](https://leetcode.cn/problems/132-pattern/) ~2000
- [2866. 美丽塔 II](https://leetcode.cn/problems/beautiful-towers-ii/) 2072
- [2454. 下一个更大元素 IV](https://leetcode.cn/problems/next-greater-element-iv/) 2175
- [2289. 使数组按非递减顺序排列](https://leetcode.cn/problems/steps-to-make-array-non-decreasing/) 2482
- [2832. 每个元素为最大值的最大范围](https://leetcode.cn/problems/maximal-range-that-each-element-is-maximum-in-it/)（会员题）

#### 矩形系列

- [84. 柱状图中最大的矩形](https://leetcode.cn/problems/largest-rectangle-in-histogram/)
- [85. 最大矩形](https://leetcode.cn/problems/maximal-rectangle/)
- [1504. 统计全 1 子矩形](https://leetcode.cn/problems/count-submatrices-with-all-ones/) 1845

#### 字典序最小

- [316. 去除重复字母](https://leetcode.cn/problems/remove-duplicate-letters/)
- [316 扩展：重复个数不超过 limit](https://leetcode.cn/contest/tianchi2022/problems/ev2bru/)
- [402. 移掉 K 位数字](https://leetcode.cn/problems/remove-k-digits/) ~1800
- [1673. 找出最具竞争力的子序列](https://leetcode.cn/problems/find-the-most-competitive-subsequence/) 1802
- [321. 拼接最大数](https://leetcode.cn/problems/create-maximum-number/)

#### 贡献法（计算所有子数组的……的和）

- [907. 子数组的最小值之和](https://leetcode.cn/problems/sum-of-subarray-minimums/) 1976
- [2104. 子数组范围和（最大值-最小值）](https://leetcode.cn/problems/sum-of-subarray-ranges/) $\mathcal{O}(n)$ 做法 ~2000
- [1856. 子数组最小乘积的最大值](https://leetcode.cn/problems/maximum-subarray-min-product/) 2051
- [2818. 操作使得分最大](https://leetcode.cn/problems/apply-operations-to-maximize-score/) 2397
- [2281. 巫师的总力量和（最小值*和）](https://leetcode.cn/problems/sum-of-total-strength-of-wizards/) 2621

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

更多精彩题解，请看 [往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
