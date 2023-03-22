### 前置知识：二分

见[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

### 本题视频讲解

见[【双周赛 100】](https://www.bilibili.com/video/BV1WM411H7UE/)。

### 思路

如果可以用 $t$ 分钟修完所有车，那么同样可以用大于 $t$ 分钟的时间修完所有车。

如果无法用 $t$ 分钟修完所有车，那么同样无法用小于 $t$ 分钟的时间修完所有车。

有单调性，可以二分答案。

根据题意，需要满足 

$$
rn^2 \le t
$$

解得

$$
n\le \sqrt\dfrac{t}{r}
$$

所以能力值为 $r$ 的机械工最多可以修

$$
\left\lfloor\sqrt\dfrac{t}{r}\right\rfloor
$$

辆车。

如果

$$
\sum_{i=0}^{i=n-1}\left\lfloor\sqrt\dfrac{t}{\textit{ranks}[i]}\right\rfloor \ge \textit{cars}
$$

则说明可以在 $t$ 分钟修理完所有汽车，根据这个公式来二分答案。

二分上界为 $\min(\textit{ranks}) \cdot \textit{cars}^2$，即让能力值最低的机械工修理所有汽车。

### 答疑

**问**：开方直接取整的做法是否会有精度误差？

**答**：代码中对整数开方，只要整数转浮点没有丢失精度（在 $2^{53}-1$ 内），开方出来的整数部分就是正确的。具体可以参考 IEEE 754。

```py [sol1-Python3]
class Solution:
    def repairCars(self, ranks: List[int], cars: int) -> int:
        s = lambda t: sum(isqrt(t // r) for r in ranks)
        return bisect_left(range(min(ranks) * cars * cars), cars, key=s)
```

```java [sol1-Java]
class Solution {
    public long repairCars(int[] ranks, int cars) {
        int minR = ranks[0];
        for (int r : ranks) minR = Math.min(minR, r);
        long left = 0, right = (long) minR * cars * cars;
        while (left + 1 < right) { // 开区间
            long mid = (left + right) / 2, s = 0;
            for (int r : ranks)
                s += Math.sqrt(mid / r);
            if (s >= cars) right = mid;
            else left = mid;
        }
        return right;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    long long repairCars(vector<int> &ranks, int cars) {
        int min_r = *min_element(ranks.begin(), ranks.end());
        long long left = 0, right = 1LL * min_r * cars * cars;
        while (left + 1 < right) { // 开区间
            long long mid = (left + right) / 2, s = 0;
            for (int r : ranks)
                s += sqrt(mid / r);
            (s >= cars ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol1-Go]
func repairCars(ranks []int, cars int) int64 {
	minR := ranks[0]
	for _, r := range ranks {
		if r < minR {
			minR = r
		}
	}
	return int64(sort.Search(minR*cars*cars, func(t int) bool {
		s := 0
		for _, r := range ranks {
			s += int(math.Sqrt(float64(t / r)))
		}
		return s >= cars
	}))
}
```

### 复杂度分析

- 时间复杂度：$O(n\log(mc^2))$，其中 $n$ 为 $\textit{ranks}$ 的长度，$m=\min(\textit{ranks})$，$c=\textit{cars}$。
- 空间复杂度：$O(1)$。仅用到若干额外变量。

### 优化

注意到 $r=\textit{ranks}[i]$ 值域很小，在 $t$ 相同时，相同 $r$ 值的人修的车数是一样的，所以可以直接按 $r$ 分组计算。

```py [sol2-Python3]
class Solution:
    def repairCars(self, ranks: List[int], cars: int) -> int:
        cnt = Counter(ranks)
        s = lambda t: sum(isqrt(t // r) * c for r, c in cnt.items())
        return bisect_left(range(min(cnt) * cars * cars), cars, key=s)
```

```java [sol2-Java]
class Solution {
    public long repairCars(int[] ranks, int cars) {
        int minR = ranks[0];
        var cnt = new int[101];
        for (int r : ranks) {
            minR = Math.min(minR, r);
            ++cnt[r];
        }
        long left = 0, right = (long) minR * cars * cars;
        while (left + 1 < right) { // 开区间
            long mid = (left + right) / 2, s = 0;
            for (int r = minR; r <= 100; ++r)
                s += (long) Math.sqrt(mid / r) * cnt[r];
            if (s >= cars) right = mid;
            else left = mid;
        }
        return right;
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    long long repairCars(vector<int> &ranks, int cars) {
        int min_r = ranks[0], cnt[101]{};
        for (int r : ranks) {
            min_r = min(min_r, r);
            ++cnt[r];
        }
        long long left = 0, right = 1LL * min_r * cars * cars;
        while (left + 1 < right) { // 开区间
            long long mid = (left + right) / 2, s = 0;
            for (int r = min_r; r <= 100; ++r)
                s += (long long) sqrt(mid / r) * cnt[r];
            (s >= cars ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol2-Go]
func repairCars(ranks []int, cars int) int64 {
	minR := ranks[0]
	cnt := [101]int{}
	for _, r := range ranks {
		if r < minR {
			minR = r
		}
		cnt[r]++
	}
	return int64(sort.Search(minR*cars*cars, func(t int) bool {
		s := 0
		for r := minR; r <= 100; r++ {
			s += int(math.Sqrt(float64(t/r))) * cnt[r]
		}
		return s >= cars
	}))
}
```

### 复杂度分析

- 时间复杂度：$O(n + U\log(mc^2))$，其中 $n$ 为 $\textit{ranks}$ 的长度，$U=\max(\textit{ranks})$，$m=\min(\textit{ranks})$，$c=\textit{cars}$。
- 空间复杂度：$O(U)$。

### 相似题目

- [875. 爱吃香蕉的珂珂](https://leetcode.cn/problems/koko-eating-bananas/)
- [2187. 完成旅途的最少时间](https://leetcode.cn/problems/minimum-time-to-complete-trips/)
- [2226. 每个小孩最多能分到多少糖果](https://leetcode.cn/problems/maximum-candies-allocated-to-k-children/)
- [1552. 两球之间的磁力](https://leetcode.cn/problems/magnetic-force-between-two-balls/)
- [2439. 最小化数组中的最大值](https://leetcode.cn/problems/minimize-maximum-of-array/)
- [2513. 最小化两个数组中的最大值](https://leetcode.cn/problems/minimize-the-maximum-of-two-arrays/)
- [2517. 礼盒的最大甜蜜度](https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/)
- [2528. 最大化城市的最小供电站数目](https://leetcode.cn/problems/maximize-the-minimum-powered-city/)

### 思考题

如果 $\textit{ranks}[i]=i+1$，是否存在某种数学做法？
