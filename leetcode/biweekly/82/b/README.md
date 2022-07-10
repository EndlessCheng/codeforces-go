下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

排序后，模拟乘客上车的过程。

模拟结束后：

- 如果最后一班公交还有空位，我们可以在发车时到达公交站，如果此刻有人，我们可以顺着他往前找到没人到达的时刻；
- 如果最后一班公交没有空位，我们可以找到上一个上车的乘客，顺着他往前找到一个没人到达的时刻。

```py [sol1-Python3]
class Solution:
    def latestTimeCatchTheBus(self, buses: List[int], passengers: List[int], capacity: int) -> int:
        buses.sort()
        passengers.sort()
        j = 0
        for t in buses:
            c = capacity
            while c and j < len(passengers) and passengers[j] <= t:
                c -= 1
                j += 1
        j -= 1
        ans = buses[-1] if c else passengers[j]
        while j >= 0 and passengers[j] == ans:  # 往前找没人到达的时刻
            ans -= 1
            j -= 1
        return ans
```

```java [sol1-Java]
class Solution {
    public int latestTimeCatchTheBus(int[] buses, int[] passengers, int capacity) {
        Arrays.sort(buses);
        Arrays.sort(passengers);
        int j = 0, c = 0;
        for (int t : buses)
            for (c = capacity; c > 0 && j < passengers.length && passengers[j] <= t; ++j)
                --c;
        --j;
        int ans = c > 0 ? buses[buses.length - 1] : passengers[j];
        while (j >= 0 && passengers[j--] == ans) --ans;
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int latestTimeCatchTheBus(vector<int> &buses, vector<int> &passengers, int capacity) {
        sort(buses.begin(), buses.end());
        sort(passengers.begin(), passengers.end());
        int j = 0, c;
        for (int t: buses)
            for (c = capacity; c && j < passengers.size() && passengers[j] <= t; ++j)
                --c;
        --j;
        int ans = c ? buses.back() : passengers[j];
        while (j >= 0 && passengers[j--] == ans) --ans;
        return ans;
    }
};
```

```go [sol1-Go]
func latestTimeCatchTheBus(buses, passengers []int, capacity int) (ans int) {
	sort.Ints(buses)
	sort.Ints(passengers)
	j, c := 0, 0
	for _, t := range buses {
		for c = capacity; c > 0 && j < len(passengers) && passengers[j] <= t; j++ {
			c--
		}
	}
	if c > 0 {
		ans = buses[len(buses)-1] // 最后一班公交还有空位，在它发车时到达
	} else {
		ans = passengers[j-1] // 上一个上车的乘客
	}
	for j--; j >= 0 && passengers[j] == ans; j-- { // 往前找没人到达的时刻
		ans--
	}
	return
}
```
