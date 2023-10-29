推荐先完成本题的简单版本：[2262. 字符串的总引力](https://leetcode.cn/problems/total-appeal-of-a-string/)，这题和本题都在 [视频讲解](https://www.bilibili.com/video/BV1Tz4y1N7Wx/) 第四题中讲了。

为方便描述，下文将 $\textit{nums}$ 简记为 $a$。

## 提示 1

把右端点相同的子数组，分为同一组。

右端点为 $i$ 的子数组，可以看成是右端点为 $i-1$ 的子数组，在末尾添加上 $a[i]$。

添加后，右端点为 $i-1$ 的这些子数组的「不同计数的平方」之和**增加**了多少？

## 提示 2

假设一个子数组的「不同计数」为 $x$，那么它的「不同计数的平方」为 $x^2$。

如果这个子数组的「不同计数」增加了 $1$，那么它的「不同计数的平方」的**增加量**为

$$
(x+1)^2 - x^2 = 2x+1
$$

## 提示 3

假设 $i=3$，那么右端点为 $i-1$ 的子数组有

- $a[0..2]$，设其「不同计数」为 $x_0$。
- $a[1..2]$，设其「不同计数」为 $x_1$。
- $a[2..2]$，设其「不同计数」为 $x_2$。

其中 $a[k..i]$ 表示从 $a[k]$ 到 $a[i]$ 的子数组。

考虑从子数组 $a[k..i-1]$ 到子数组 $a[k..i]$，分类讨论：

- 如果 $a[i]$ 之前没有遇到过（例如 $a=[1,2,3,4]$），那么这些子数组的「不同计数」都会增加 $1$。根据提示 2，「不同计数的平方」之和的**增加量**为 $(x_0+x_1+x_2)\cdot 2 + 3$。
- 如果 $a[i]$ 之前遇到过，设其上次出现的下标为 $j$，那么：
    - 对于子数组 $a[0..i-1],\ a[1..i-1],\ a[2..i-1],\cdots,a[j..i-1]$，在其末尾添加 $a[i]$ 后，这些子数组的「不同计数」是不会变化的，因为 $a[i]$ 已经在 $a[j]$ 处出现过了。 
    - 对于子数组 $a[j+1..i-1],\ a[j+2..i-1],\cdots,a[i-1..i-1]$，由于不包含 $a[i]$，这些子数组的「不同计数」都会增加 $1$，「不同计数的平方」之和的**增加量**计算方式同上。
- 别忘了 $a[i..i]$ 也是一个子数组，把它的「不同计数」加一。

所以，我们需要一个这样的数据结构，用来维护子数组的「不同计数」：

1. 定义 $f[k]$ 表示左端点为 $k$ 的子数组的「不同计数」，如果当前遍历到 $\textit{nums}[i]$，那么 $f[k]$ 就对应着子数组 $a[k..i]$。
2. 区间加一：例如 $a[1..3]$，$a[2..3]$ 和 $a[3..3]$ 的「不同计数」都增加了 $1$，那么就把区间 $[1,3]$ 的「不同计数」加一。
3. 询问区间元素和：为了计算出「不同计数的平方」之和的**增加量**，需要知道从 $f[j+1]$ 到 $f[i-1]$ 的「不同计数」，这里 $j$ 为 $a[i]$ 上次出现的下标。

这可以用 lazy 线段树实现，具体请看[【双周赛 98】](https://www.bilibili.com/video/BV15D4y1G7ms/)第四题的讲解（[2569. 更新数组后处理求和查询](https://leetcode.cn/problems/handling-sum-queries-after-update/)）

## 提示 4

用一个变量 $s$ 维护右端点为 $i$ 的子数组的「不同计数的平方」之和。

遍历 $\textit{nums}$，每次循环按照上述规则更新 $s$：

1. 为了方便调用线段树，假设下标从 $1$ 开始。
2. 设 $a[i]$ 上次出现的下标为 $j$（不存在则为 $0$）。询问 $[j+1,i]$ 的元素和，设为 $s_1$。把 $s_1\cdot 2 + i-j$ 加到 $s$ 中。
3. 把区间 $[j+1,i]$ 都加一。
4. 把 $s$ 加到答案中。
5. 更新 $a[i]$ 的上一次出现位置为 $i$。

代码实现时，由于查询的区间和更新的区间是同一个，可以同时完成。

```py [sol-Python3]
class Solution:
    def sumCounts(self, nums: List[int]) -> int:
        n = len(nums)
        sum = [0] * (n * 4)
        todo = [0] * (n * 4)

        def do(o: int, l: int, r: int, add: int) -> None:
            sum[o] += add * (r - l + 1)
            todo[o] += add

        # o=1  [l,r] 1<=l<=r<=n
        # 把 [L,R] 加一，同时返回加一之前的区间和
        def query_and_add1(o: int, l: int, r: int, L: int, R: int) -> int:
            if L <= l and r <= R:
                res = sum[o]
                do(o, l, r, 1)
                return res

            m = (l + r) // 2
            add = todo[o]
            if add:
                do(o * 2, l, m, add)
                do(o * 2 + 1, m + 1, r, add)
                todo[o] = 0

            res = 0
            if L <= m: res += query_and_add1(o * 2, l, m, L, R)
            if m < R:  res += query_and_add1(o * 2 + 1, m + 1, r, L, R)
            sum[o] = sum[o * 2] + sum[o * 2 + 1]
            return res

        ans = s = 0
        last = {}
        for i, x in enumerate(nums, 1):
            j = last.get(x, 0)
            s += query_and_add1(1, 1, n, j + 1, i) * 2 + i - j
            ans += s
            last[x] = i
        return ans % 1_000_000_007
```

```java [sol-Java]
class Solution {
    private long[] sum;
    private int[] todo;

    public int sumCounts(int[] nums) {
        int n = nums.length;
        sum = new long[n * 4];
        todo = new int[n * 4];

        long ans = 0, s = 0;
        var last = new HashMap<Integer, Integer>();
        for (int i = 1; i <= n; i++) {
            int x = nums[i - 1];
            int j = last.getOrDefault(x, 0);
            s += queryAndAdd1(1, 1, n, j + 1, i) * 2 + i - j;
            ans = (ans + s) % 1_000_000_007;
            last.put(x, i);
        }
        return (int) ans;
    }

    private void do_(int o, int l, int r, int add) {
        sum[o] += (long) add * (r - l + 1);
        todo[o] += add;
    }

    // o=1  [l,r] 1<=l<=r<=n
    // 把 [L,R] 加一，同时返回加一之前的区间和
    private long queryAndAdd1(int o, int l, int r, int L, int R) {
        if (L <= l && r <= R) {
            long res = sum[o];
            do_(o, l, r, 1);
            return res;
        }

        int m = (l + r) / 2;
        int add = todo[o];
        if (add != 0) {
            do_(o * 2, l, m, add);
            do_(o * 2 + 1, m + 1, r, add);
            todo[o] = 0;
        }

        long res = 0;
        if (L <= m) res += queryAndAdd1(o * 2, l, m, L, R);
        if (m < R)  res += queryAndAdd1(o * 2 + 1, m + 1, r, L, R);
        sum[o] = sum[o * 2] + sum[o * 2 + 1];
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<long long> sum;
    vector<int> todo;

    void do_(int o, int l, int r, int add) {
        sum[o] += (long long) add * (r - l + 1);
        todo[o] += add;
    }

    // o=1  [l,r] 1<=l<=r<=n
    // 把 [L,R] 加一，同时返回加一之前的区间和
    long long query_and_add1(int o, int l, int r, int L, int R) {
        if (L <= l && r <= R) {
            long long res = sum[o];
            do_(o, l, r, 1);
            return res;
        }

        int m = (l + r) / 2;
        int add = todo[o];
        if (add != 0) {
            do_(o * 2, l, m, add);
            do_(o * 2 + 1, m + 1, r, add);
            todo[o] = 0;
        }

        long long res = 0;
        if (L <= m) res += query_and_add1(o * 2, l, m, L, R);
        if (m < R)  res += query_and_add1(o * 2 + 1, m + 1, r, L, R);
        sum[o] = sum[o * 2] + sum[o * 2 + 1];
        return res;
    }

public:
    int sumCounts(vector<int> &nums) {
        int n = nums.size();
        sum.resize(n * 4);
        todo.resize(n * 4);

        long long ans = 0, s = 0;
        unordered_map<int, int> last;
        for (int i = 1; i <= n; i++) {
            int x = nums[i - 1];
            int j = last.count(x) ? last[x] : 0;
            s += query_and_add1(1, 1, n, j + 1, i) * 2 + i - j;
            ans = (ans + s) % 1'000'000'007;
            last[x] = i;
        }
        return ans;
    }
};
```

```go [sol-Go]
type lazySeg []struct{ sum, todo int }

func (t lazySeg) do(o, l, r, add int) {
	t[o].sum += add * (r - l + 1)
	t[o].todo += add
}

// o=1  [l,r] 1<=l<=r<=n
// 把 [L,R] 加一，同时返回加一之前的区间和
func (t lazySeg) queryAndAdd1(o, l, r, L, R int) (res int) {
	if L <= l && r <= R {
		res = t[o].sum
		t.do(o, l, r, 1)
		return
	}
	m := (l + r) >> 1
	if add := t[o].todo; add != 0 {
		t.do(o<<1, l, m, add)
		t.do(o<<1|1, m+1, r, add)
		t[o].todo = 0
	}
	if L <= m {
		res = t.queryAndAdd1(o<<1, l, m, L, R)
	}
	if m < R {
		res += t.queryAndAdd1(o<<1|1, m+1, r, L, R)
	}
	t[o].sum = t[o<<1].sum + t[o<<1|1].sum
	return
}

func sumCounts(nums []int) (ans int) {
	last := map[int]int{}
	n := len(nums)
	t := make(lazySeg, n*4)
	s := 0
	for i, x := range nums {
		i++
		j := last[x]
		s += t.queryAndAdd1(1, 1, n, j+1, i)*2 + i - j
		ans = (ans + s) % 1_000_000_007
		last[x] = i
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
