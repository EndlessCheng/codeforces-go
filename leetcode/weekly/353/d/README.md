[视频讲解](https://www.bilibili.com/video/BV1XW4y1f7Wv/) 第四题。

## 提示 1

想一想，如果 $\textit{nums}[0]>0$，我们必须要执行什么样的操作，才能让 $\textit{nums}[0]=0$？

## 提示 2

对于 $\textit{nums}[0]>0$ 的情况，必须把 $\textit{nums}[0]$ 到 $\textit{nums}[k-1]$ 都减去 $\textit{nums}[0]$。

然后思考 $\textit{nums}[1]$ 要怎么处理，依此类推。

## 提示 3

子数组同时加上/减去一个数，非常适合用 [差分数组](https://leetcode.cn/circle/discuss/FfMCgb/) 来维护，请至少做一道差分数组题目再往下阅读。

设差分数组为 $d$。那么把 $\textit{nums}[i]$ 到 $\textit{nums}[i+k-1]$ 同时减去 $1$，等价于把 $d[i]$ 减 $1$，$d[i+k]$ 加 $1$。

注意子数组长度必须恰好等于 $k$，所以当 $i+k\le n$ 时，才能执行上述操作。

遍历数组的同时，用变量 $\textit{sumD}$ 累加差分值。遍历到 $\textit{nums}[i]$ 时，$\textit{nums}[i]+\textit{sumD}$ 就是 $\textit{nums}[i]$ 的实际值了。

分类讨论：

- 如果 $\textit{nums}[i]<0$，由于无法让元素值增大，返回 `false`。
- 如果 $\textit{nums}[i]=0$，无需操作，遍历下一个数。
- 如果 $\textit{nums}[i]>0$：
  - 如果 $i+k> n$，无法执行操作，所以 $\textit{nums}[i]$ 无法变成 $0$，返回 `false`。
  - 如果 $i+k\le n$，按照上面说的执行操作，修改差分数组，遍历下一个数。

如果遍历中途没有返回 `false`，那么最后返回 `true`。

```py [sol-Python3]
class Solution:
    def checkArray(self, nums: List[int], k: int) -> bool:
        n = len(nums)
        d = [0] * (n + 1)
        sum_d = 0
        for i, x in enumerate(nums):
            sum_d += d[i]
            x += sum_d
            if x == 0: continue  # 无需操作
            if x < 0 or i + k > n: return False  # 无法操作
            sum_d -= x  # 直接加到 sum_d 中
            d[i + k] += x
        return True
```

```java [sol-Java]
class Solution {
    public boolean checkArray(int[] nums, int k) {
        int n = nums.length, sumD = 0;
        var d = new int[n + 1];
        for (int i = 0; i < n; i++) {
            sumD += d[i];
            int x = nums[i];
            x += sumD;
            if (x == 0) continue; // 无需操作
            if (x < 0 || i + k > n) return false; // 无法操作
            sumD -= x; // 直接加到 sumD 中
            d[i + k] += x;
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool checkArray(vector<int> &nums, int k) {
        int n = nums.size(), sum_d = 0;
        vector<int> d(n + 1);
        for (int i = 0; i < n; i++) {
            sum_d += d[i];
            int x = nums[i];
            x += sum_d;
            if (x == 0) continue; // 无需操作
            if (x < 0 || i + k > n) return false; // 无法操作
            sum_d -= x; // 直接加到 sum_d 中
            d[i + k] += x;
        }
        return true;
    }
};
```

```go [sol-Go]
func checkArray(nums []int, k int) bool {
	n := len(nums)
	d := make([]int, n+1)
	sumD := 0
	for i, x := range nums {
		sumD += d[i]
		x += sumD
		if x == 0 { // 无需操作
			continue
		}
		if x < 0 || i+k > n { // 无法操作
			return false
		}
		sumD -= x // 直接加到 sumD 中
		d[i+k] += x
	}
	return true
}
```

```js [sol-JavaScript]
var checkArray = function (nums, k) {
    const n = nums.length;
    let d = new Array(n + 1).fill(0);
    let sumD = 0;
    for (let i = 0; i < n; i++) {
        sumD += d[i];
        let x = nums[i];
        x += sumD;
        if (x == 0) continue; // 无需操作
        if (x < 0 || i + k > n) return false; // 无法操作
        sumD -= x; // 直接加到 sumD 中
        d[i + k] += x;
    }
    return true;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

#### 相似题目

- [2528. 最大化城市的最小供电站数目](https://leetcode.cn/problems/maximize-the-minimum-powered-city/)
