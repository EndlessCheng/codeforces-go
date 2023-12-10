[本题视频讲解](https://www.bilibili.com/video/BV1Lj411s7ga/)

用 [差分数组](https://leetcode.cn/circle/discuss/FfMCgb/) 的思想可以做到 $\mathcal{O}(n)$，方法如下：

1. 初始化 $\textit{dec}=0$，表示需要减一的次数。
2. 设 $x=\textit{batteryPercentages}[i]$，如果 $x - \textit{dec} > 0$，即 $x > \textit{dec}$，那么后面的数都要减一，把 $\textit{dec}$ 加一即可。
3. 答案就是 $\textit{dec}$。因为每次遇到 $x > \textit{dec}$ 都把 $\textit{dec}$ 加一，这正是题目要求统计的。

```py [sol-Python3]
class Solution:
    def countTestedDevices(self, batteryPercentages: List[int]) -> int:
        dec = 0
        for x in batteryPercentages:
            if x > dec:
                dec += 1
        return dec
```

```java [sol-Java]
class Solution {
    public int countTestedDevices(int[] batteryPercentages) {
        int dec = 0;
        for (int x : batteryPercentages) {
            if (x > dec) {
                dec++;
            }
        }
        return dec;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countTestedDevices(vector<int> &batteryPercentages) {
        int dec = 0;
        for (int x : batteryPercentages) {
            dec += x > dec;
        }
        return dec;
    }
};
```

```go [sol-Go]
func countTestedDevices(batteryPercentages []int) int {
	dec := 0
	for _, x := range batteryPercentages {
		if x > dec {
			dec++
		}
	}
	return dec
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{batteryPercentages}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。
