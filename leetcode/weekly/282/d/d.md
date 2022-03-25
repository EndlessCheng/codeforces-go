#### 提示 1

连续使用同一个轮胎时，耗费的时间是指数增长的。

#### 提示 2

连续使用同一个轮胎 $i$ 跑 $x$ 圈，第 $x$ 圈的耗时不应超过 $\textit{changeTime} + f_i$，否则直接更换轮胎会更优。结合提示 1，连续使用同一个轮胎的次数不会很大。

#### 提示 3

考虑线性 DP。

---

#### 上界分析

连续使用同一个轮胎 $i$ 跑 $x$ 圈，第 $x$ 圈的耗时不应超过 $\textit{changeTime} + f_i$，即

$$
f_i\cdot r_i^{x-1} \le \textit{changeTime} + f_i
$$

考虑 $x$ 至多能是多少。由于 $f_i$ 越小 $x$ 的上界越大，以及 $r_i$ 越小 $x$ 的上界越大，那么取 $f_i=1,r_i=2$，则有

$$
2^{x-1}\le \textit{changeTime} + 1
$$

解得

$$
x\le \log_2(\textit{changeTime}+1)+1
$$

由于 $x$ 是个整数，因此 $x$ 的上界为 $\lfloor \log_2(\textit{changeTime}+1)+1 \rfloor$。

根据题目的数据范围，代码实现时可将上界视为 $17$。

#### 算法

首先预处理出连续使用同一个轮胎跑 $x$ 圈的最小耗时，记作 $\textit{minSec}[x]$，这可以通过遍历每个轮胎计算出来。

然后定义 $f[i]$ 表示跑 $i$ 圈的最小耗时。为方便计算，初始值 $f[0]=-\textit{changeTime}$。

考虑最后一个轮胎连续跑了 $j$ 圈，我们可以从 $f[i-j]$ 转移过来，因此有转移方程

$$
f[i] = \textit{changeTime} + \min\limits_{j=1}^{\min(17,i)} f[i-j] + \textit{minSec}[j]
$$

最后答案为 $f[\textit{numLaps}]$。

```go [sol1-Go]
func minimumFinishTime(tires [][]int, changeTime, numLaps int) int {
	minSec := [18]int{}
	for i := range minSec {
		minSec[i] = math.MaxInt32
	}
	for _, tire := range tires {
		f, r := tire[0], tire[1]
		for x, time, sum := 1, f, 0; time <= changeTime+f; x++ {
			sum += time
			minSec[x] = min(minSec[x], sum)
			time *= r
		}
	}

	f := make([]int, numLaps+1)
	f[0] = -changeTime
	for i := 1; i <= numLaps; i++ {
		f[i] = math.MaxInt32
		for j := 1; j <= 17 && j <= i; j++ {
			f[i] = min(f[i], f[i-j]+minSec[j])
		}
		f[i] += changeTime
	}
	return f[numLaps]
}

func min(a, b int) int { if a > b { return b }; return a }
```

```C++ [sol1-C++]
class Solution {
public:
    int minimumFinishTime(vector<vector<int>> &tires, int changeTime, int numLaps) {
        vector<int> minSec(18, INT_MAX / 2); // 除二是防止下面计算状态转移时溢出
        for (auto &tire : tires) {
            long time = tire[0];
            for (int x = 1, sum = 0; time <= changeTime + tire[0]; ++x) {
                sum += time;
                minSec[x] = min(minSec[x], sum);
                time *= tire[1];
            }
        }

        vector<int> f(numLaps + 1, INT_MAX);
        f[0] = -changeTime;
        for (int i = 1; i <= numLaps; ++i) {
            for (int j = 1; j <= min(17, i); ++j)
                f[i] = min(f[i], f[i - j] + minSec[j]);
            f[i] += changeTime;
        }
        return f[numLaps];
    }
};
```

```Python [sol1-Python3]
class Solution:
    def minimumFinishTime(self, tires: List[List[int]], changeTime: int, numLaps: int) -> int:
        min_sec = [inf] * 18
        for f, r in tires:
            x, time, sum = 1, f, 0
            while time <= changeTime + f:
                sum += time
                min_sec[x] = min(min_sec[x], sum)
                time *= r
                x += 1

        f = [0] * (numLaps + 1)
        f[0] = -changeTime
        for i in range(1, numLaps + 1):
            f[i] = changeTime + min(f[i - j] + min_sec[j] for j in range(1, min(18, i + 1)))
        return f[numLaps]
```

```java [sol1-Java]
class Solution {
    public int minimumFinishTime(int[][] tires, int changeTime, int numLaps) {
        var minSec = new int[18];
        Arrays.fill(minSec, Integer.MAX_VALUE / 2); // 除二是防止下面计算状态转移时溢出
        for (var tire : tires) {
            long time = tire[0];
            for (int x = 1, sum = 0; time <= changeTime + tire[0]; ++x) {
                sum += time;
                minSec[x] = Math.min(minSec[x], sum);
                time *= tire[1];
            }
        }

        var f = new int[numLaps + 1];
        Arrays.fill(f, Integer.MAX_VALUE);
        f[0] = -changeTime;
        for (var i = 1; i <= numLaps; ++i) {
            for (var j = 1; j <= Math.min(17, i); ++j)
                f[i] = Math.min(f[i], f[i - j] + minSec[j]);
            f[i] += changeTime;
        }
        return f[numLaps];
    }
}
```
