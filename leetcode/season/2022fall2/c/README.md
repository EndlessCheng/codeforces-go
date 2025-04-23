[视频讲解](https://www.bilibili.com/video/BV1rT411P7NA) 已出炉，**包括本题滑窗的原理和时间复杂度分析**，欢迎点赞三连，在评论区分享你对这场力扣杯的看法~

**注**：本题测试数据比较弱，不取模也能过。正确做法是需要取模的，因为 $10^5$ 个 $1$ 算出的答案会 $\ge 10^9+7$。

```py [sol-Python3]
class Solution:
    def beautifulBouquet(self, flowers: List[int], cnt: int) -> int:
        ans = left = 0
        c = defaultdict(int)
        for right, x in enumerate(flowers):
            c[x] += 1
            while c[x] > cnt:
                c[flowers[left]] -= 1
                left += 1
            ans += right - left + 1
        return ans % 1_000_000_007
```

```java [sol-Java]
class Solution {
    public int beautifulBouquet(int[] flowers, int cnt) {
        long ans = 0;
        Map<Integer, Integer> c = new HashMap<>();
        int left = 0;
        for (int right = 0; right < flowers.length; right++) {
            int x = flowers[right];
            c.merge(x, 1, Integer::sum); // c[x]++
            while (c.get(x) > cnt) {
                c.merge(flowers[left], -1, Integer::sum);
                left++;
            }
            ans += right - left + 1;
        }
        return (int) (ans % 1_000_000_007);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int beautifulBouquet(vector<int>& flowers, int cnt) {
        long long ans = 0;
        unordered_map<int, int> c;
        int left = 0;
        for (int right = 0; right < flowers.size(); right++) {
            int x = flowers[right];
            c[x]++;
            while (c[x] > cnt) {
                c[flowers[left]]--;
                left++;
            }
            ans += right - left + 1;
        }
        return ans % 1'000'000'007;
    }
};
```

```go [sol-Go]
func beautifulBouquet(flowers []int, cnt int) (ans int) {
	c := map[int]int{}
	left := 0
	for right, x := range flowers {
		c[x]++
		for c[x] > cnt {
			c[flowers[left]]--
			left++
		}
		ans += right - left + 1
	}
	return ans % 1_000_000_007
}
```