## 需要交替播种吗？

假设有两枚种子要播种，$1$ 号种子需要 $3$ 天，$2$ 号种子需要 $2$ 天。

- 如果播种顺序为 $11122$，即先播种完 $1$ 号种子，再播种 $2$ 号种子，那么 $1$ 号种子在第 $3$ 天播种完毕，$2$ 号种子在第 $5$ 天播种完毕，完成播种共需 $5$ 天。
- 如果交替播种，比如 $12112$，那么 $2$ 号种子完成播种的时间是不变的，但对于 $1$ 号种子来说，完成播种的时间却延后了一天。

这说明如果要交替播种，至少有一枚种子的完成时间要延后，并且所有种子完成播种的时间是不变的。也就是说，交替播种不仅没有得到任何好处，反而会拖慢其中一些种子的播种进度。

所以不应交替播种，应当种完一枚种子再开始另一枚。

## 谁先播种？

对于两枚种子，设其播种所需天数为 $p_1$ 和 $p_2$，生长所需天数为 $g_1$ 和 $g_2$。

不妨设 $g_1\ge g_2$。我们来比较哪种播种顺序更优：

- 先 $1$ 后 $2$ 时的最晚开花时间：
    $$
    \max(p_1+g_1,p_1+p_2+g_2)
    $$
- 先 $2$ 后 $1$ 时的最晚开花时间：
    $$
    \max(p_1+p_2+g_1,p_2+g_2)
    $$
    由于 $g_1\ge g_2$ 且 $p_1>0$，所以 $p_1+p_2+g_1>p_2+g_2$，上式可化简为 
    $$
    p_1+p_2+g_1
    $$

由于 $p_1+g_1 < p_1+p_2+g_1$ 且 $p_1+p_2+g_2 \le p_1+p_2+g_1$，因此

$$
\max(p_1+g_1,p_1+p_2+g_2) \le p_1+p_2+g_1 = \max(p_1+p_2+g_1,p_2+g_2)
$$

上式表明，按照先 $1$ 后 $2$ 的顺序播种，最晚开花时间不会晚于按照先 $2$ 后 $1$ 播种时的最晚开花时间。

这意味着**按照生长天数从大到小排序**后，交换任意两枚种子的播种顺序，不会让最晚开花时间提前。

假设存在其它更优的种子排列，我们可以交换「生长天数小且排在前面的种子」与「生长天数大且排在后面的种子」，从而得到更早的最晚开花时间，因此假设不成立，按照生长天数从大到小的顺序播种是最优的。

对于两枚生长天数相同的种子，由于无论按照何种顺序播种，这两枚种子的最晚开花时间都是相同的，因此无需考虑生长天数相同的种子的播种顺序，在排序时，仅需按生长天数从大到小排序。

```py [sol-Python3]
class Solution:
    def earliestFullBloom(self, plantTime: List[int], growTime: List[int]) -> int:
        ans = days = 0
        for p, g in sorted(zip(plantTime, growTime), key=lambda z: -z[1]):
            days += p  # 累加播种天数
            ans = max(ans, days + g)  # 再加上生长天数，就是这个种子的开花时间
        return ans
```

```java [sol-Java]
class Solution {
    public int earliestFullBloom(int[] plantTime, int[] growTime) {
        int n = plantTime.length;
        var id = new Integer[n];
        Arrays.setAll(id, i -> i);
        Arrays.sort(id, (i, j) -> growTime[j] - growTime[i]);
        int ans = 0, days = 0;
        for (int i : id) {
            days += plantTime[i]; // 累加播种天数
            ans = Math.max(ans, days + growTime[i]); // 再加上生长天数，就是这个种子的开花时间
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int earliestFullBloom(vector<int> &plantTime, vector<int> &growTime) {
        vector<int> id(plantTime.size());
        iota(id.begin(), id.end(), 0); // id[i] = i
        sort(id.begin(), id.end(), [&](int i, int j) { return growTime[i] > growTime[j]; });
        int ans = 0, days = 0;
        for (int i : id) {
            days += plantTime[i]; // 累加播种天数
            ans = max(ans, days + growTime[i]); // 再加上生长天数，就是这个种子的开花时间
        }
        return ans;
    }
};
```

```go [sol-Go]
func earliestFullBloom(plantTime, growTime []int) (ans int) {
	type pair struct{ p, g int }
	a := make([]pair, len(plantTime))
	for i, p := range plantTime {
		a[i] = pair{p, growTime[i]}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].g > a[j].g })
	days := 0
	for _, p := range a {
		days += p.p // 累加播种天数
		ans = max(ans, days+p.g) // 再加上生长天数，就是这个种子的开花时间
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

```js [sol-JavaScript]
var earliestFullBloom = function (plantTime, growTime) {
    let ans = 0, days = 0;
    for (const [p, g] of _.zip(plantTime, growTime).sort((a, b) => b[1] - a[1])) {
        days += p; // 累加播种天数
        ans = Math.max(ans, days + g); // 再加上生长天数，就是这个种子的开花时间
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn earliest_full_bloom(plant_time: Vec<i32>, grow_time: Vec<i32>) -> i32 {
        let mut id: Vec<usize> = (0..grow_time.len()).collect();
        id.sort_unstable_by(|&i, &j| grow_time[j].cmp(&grow_time[i]));
        let mut ans = 0;
        let mut days = 0;
        for &i in &id {
            days += plant_time[i]; // 累加播种天数
            ans = ans.max(days + grow_time[i]); // 再加上生长天数，就是这个种子的开花时间
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{plantTime}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

- [1665. 完成所有任务的最少初始能量](https://leetcode.cn/problems/minimum-initial-energy-to-finish-tasks/)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

更多精彩题解，请看 [往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
