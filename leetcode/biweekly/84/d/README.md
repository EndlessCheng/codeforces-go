下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

#### 提示 1

每个数字要么不变，要么分成更小的数字。

#### 提示 2

最后一个数字需要操作吗？

不需要，如果操作，前面的数字就需要变得更小，这会让操作次数增多。

#### 提示 3

倒着遍历 $\textit{nums}$。设当前操作出的最小值为 $m$，如果 $\textit{nums}[i]>m$，那么需要拆分 $\textit{nums}[i]$，使得拆分出的数字的最大值不超过 $m$。

设拆分出了 $x$ 个数字，由于这 $x$ 个数字都不超过 $m$，即

$$
\textit{nums}[i] = v_1+v_2+\cdots+v_x \le m+m+\cdots+m = mx
$$ 
 
得 

$$
x\ge\left\lceil\dfrac{\textit{nums}[i]}{m}\right\rceil
$$

为了使操作次数尽量小，应取等号，即

$$
x=\left\lceil\dfrac{\textit{nums}[i]}{m}\right\rceil
$$

操作次数为 $k=x-1$。

为了使拆分出的数字的最小值尽可能地大，拆分出的最小数字应为 $\left\lfloor\dfrac{\textit{nums}[i]}{x}\right\rfloor$，证明如下：

> 若这 $x$ 个数均为 $\left\lfloor\dfrac{\textit{nums}[i]}{x}\right\rfloor$，那么有
> $$
> x\left\lfloor\dfrac{\textit{nums}[i]}{x}\right\rfloor\le \textit{nums}[i]
> $$
> 若这 $x$ 个数均为 $\left\lfloor\dfrac{\textit{nums}[i]}{x}\right\rfloor+1$，那么有
> $$
> x\left(\left\lfloor\dfrac{\textit{nums}[i]}{x}\right\rfloor+1\right)\ge x\left\lceil\dfrac{\textit{nums}[i]}{x}\right\rceil \ge \textit{nums}[i]
> $$
> 联合得到
> $$
> x\left\lfloor\dfrac{\textit{nums}[i]}{x}\right\rfloor\le \textit{nums}[i]\le x\left(\left\lfloor\dfrac{\textit{nums}[i]}{x}\right\rfloor+1\right)
> $$
> 据此，我们可以给出一个拆分方案：将这 $x$ 个数均初始化为 $\left\lfloor\dfrac{\textit{nums}[i]}{x}\right\rfloor$，然后给其中的 $\textit{nums}[i]-x\left\lfloor\dfrac{\textit{nums}[i]}{x}\right\rfloor$ 个数字加一，这样可以使这 $x$ 数的和恰好为 $\textit{nums}[i]$。上面的不等式说明这样的方案是存在的。

代码实现时，无需判断 $\textit{nums}[i]$ 与 $m$ 的大小关系。

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干额外变量。

```py [sol1-Python3]
class Solution:
    def minimumReplacement(self, nums: List[int]) -> int:
        ans, m = 0, nums[-1]
        for i in range(len(nums) - 2, -1, -1):
            k = (nums[i] - 1) // m
            ans += k
            m = nums[i] // (k + 1)
        return ans
```

```java [sol1-Java]
class Solution {
    public long minimumReplacement(int[] nums) {
        var ans = 0L;
        var m = nums[nums.length - 1];
        for (var i = nums.length - 2; i >= 0; --i) {
            var k = (nums[i] - 1) / m;
            ans += k;
            m = nums[i] / (k + 1);
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    long long minimumReplacement(vector<int> &nums) {
        long ans = 0L;
        int m = nums.back();
        for (int i = int(nums.size()) - 2; i >= 0; --i) {
            int k = (nums[i] - 1) / m;
            ans += k;
            m = nums[i] / (k + 1);
        }
        return ans;
    }
};
```

```go [sol1-Go]
func minimumReplacement(nums []int) (ans int64) {
	m := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		k := (nums[i] - 1) / m
		ans += int64(k)
		m = nums[i] / (k + 1)
	}
	return
}
```

#### 思考题

每个数字需要拆分成若干**整数**（目前题目没有说清楚这一点，读者可以通过测试 $[5,5,3]$ 这个输入来确认，预期结果为 $3$）。

如果可以拆分成实数，答案会是多少呢？比如 $[5,5,3]$ 可以拆分成 $[2.5,2.5,2.5,2.5,3]$，答案为 $2$。

如何在计算过程中避免使用浮点数？
