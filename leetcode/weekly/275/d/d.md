贪心及其证明

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
    由于 $g_1\ge g_2$ 且 $p_1>0$，所以 $p_1+p_2+g_1>p_2+g_2$，因此上式即为 $p_1+p_2+g_1$。
    
由于 $p_1+g_1 < p_1+p_2+g_1$ 且 $p_1+p_2+g_2 \le p_1+p_2+g_1$，因此我们有

$$
\max(p_1+g_1,p_1+p_2+g_2) \le p_1+p_2+g_1 = \max(p_1+p_2+g_1,p_2+g_2)
$$

上式表明，按照先 $1$ 后 $2$ 的顺序播种，最晚开花时间不会晚于按照先 $2$ 后 $1$ 播种时的最晚开花时间。

因此，我们可以按照生长天数从大到小的顺序播种。对于两枚生长天数相同的种子，由于无论按照何种顺序播种，这两枚种子的最晚开花时间都是相同的，因此无需考虑生长天数相同的种子的播种顺序，所以在排序时，仅需对生长天数从大到小排序。

```go [sol1-Go]
func earliestFullBloom(plantTime, growTime []int) (ans int) {
	type pair struct{ p, g int }
	a := make([]pair, len(plantTime))
	for i, p := range plantTime {
		a[i] = pair{p, growTime[i]}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].g > a[j].g })
	day := 0
	for _, p := range a {
		day += p.p
		if day+p.g > ans {
			ans = day + p.g
		}
	}
	return
}
```

```C++ [sol1-C++]
class Solution {
public:
    int earliestFullBloom(vector<int> &plantTime, vector<int> &growTime) {
        vector<int> id(plantTime.size());
        iota(id.begin(), id.end(), 0);
        sort(id.begin(), id.end(), [&](int i, int j) { return growTime[i] > growTime[j]; });
        int ans = 0, day = 0;
        for (int i : id) {
            day += plantTime[i];
            ans = max(ans, day + growTime[i]);
        }
        return ans;
    }
};
```

```Python [sol1-Python3]
class Solution:
    def earliestFullBloom(self, plantTime: List[int], growTime: List[int]) -> int:
        a = list(zip(plantTime, growTime))
        a.sort(key=lambda x: -x[1])
        ans, day = 0, 0
        for p in a:
            day += p[0]
            ans = max(ans, day + p[1])
        return ans
```


