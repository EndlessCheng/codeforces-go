## 方法一：有序集合

用一个哈希表 $\textit{foodMap}$ 记录每个食物名称对应的食物评分和烹饪方式，同时一个哈希表套有序集合 $\textit{cuisineMap}$ 记录每个烹饪方式对应的食物评分和食物名字。

- $\texttt{changeRating}$：先从 $\textit{cuisineMap}$ 中删掉旧数据，然后将 $\textit{newRating}$ 和 $\textit{food}$ 记录到 $\textit{cuisineMap}$ 和 $\textit{foodMap}$ 中。
- $\texttt{highestRated}$：从有序集合 $\textit{cuisineMap}[\textit{cuisine}]$ 中找到最优数据。

```py [sol-Python3]
class FoodRatings:
    def __init__(self, foods: List[str], cuisines: List[str], ratings: List[int]):
        self.food_map = {}
        self.cuisine_map = defaultdict(SortedList)  # sortedcontainers
        for food, cuisine, rating in zip(foods, cuisines, ratings):
            self.food_map[food] = [rating, cuisine]
            # 取负号，保证 rating 相同时，字典序更小的 food 排在前面
            self.cuisine_map[cuisine].add((-rating, food))

    def changeRating(self, food: str, newRating: int) -> None:
        rating, cuisine = self.food_map[food]
        sl = self.cuisine_map[cuisine]
        sl.discard((-rating, food))  # 移除旧数据
        sl.add((-newRating, food))  # 添加新数据
        self.food_map[food][0] = newRating  # 更新 food 的 rating

    def highestRated(self, cuisine: str) -> str:
        return self.cuisine_map[cuisine][0][1]
```

```java [sol-Java]
class FoodRatings {
    private final Map<String, Pair<Integer, String>> foodMap = new HashMap<>();
    private final Map<String, TreeSet<Pair<Integer, String>>> cuisineMap = new HashMap<>();

    public FoodRatings(String[] foods, String[] cuisines, int[] ratings) {
        for (int i = 0; i < foods.length; i++) {
            String food = foods[i];
            String cuisine = cuisines[i];
            int rating = ratings[i];
            foodMap.put(food, new Pair<>(rating, cuisine));
            cuisineMap.computeIfAbsent(cuisine, k -> new TreeSet<>((a, b) ->
                    !Objects.equals(a.getKey(), b.getKey()) ? b.getKey() - a.getKey() : a.getValue().compareTo(b.getValue())
            )).add(new Pair<>(rating, food));
        }
    }

    public void changeRating(String food, int newRating) {
        Pair<Integer, String> p = foodMap.get(food);
        TreeSet<Pair<Integer, String>> st = cuisineMap.get(p.getValue());
        st.remove(new Pair<>(p.getKey(), food)); // 移除旧数据
        st.add(new Pair<>(newRating, food)); // 添加新数据
        foodMap.put(food, new Pair<>(newRating, p.getValue()));
    }

    public String highestRated(String cuisine) {
        return cuisineMap.get(cuisine).first().getValue();
    }
}
```

```cpp [sol-C++]
class FoodRatings {
    unordered_map<string, pair<int, string>> food_map;
    unordered_map<string, set<pair<int, string>>> cuisine_map;

public:
    FoodRatings(vector<string>& foods, vector<string>& cuisines, vector<int>& ratings) {
        for (int i = 0; i < foods.size(); i++) {
            auto& food = foods[i];
            auto& cuisine = cuisines[i];
            int rating = ratings[i];
            food_map[food] = {rating, cuisine};
            // 取负号，保证 rating 相同时，字典序更小的 food 排在前面
            cuisine_map[cuisine].emplace(-rating, food);
        }
    }

    void changeRating(string food, int newRating) {
        auto& [rating, cuisine] = food_map[food];
        auto& st = cuisine_map[cuisine];
        st.erase({-rating, food}); // 移除旧数据
        st.emplace(-newRating, food); // 添加新数据
        rating = newRating;
    }

    string highestRated(string cuisine) {
        return cuisine_map[cuisine].begin()->second;
    }
};
```

```go [sol-Go]
import "github.com/emirpasic/gods/v2/trees/redblacktree"

type pair struct {
	rating int
	food   string
}

type FoodRatings struct {
	foodMap    map[string]pair
	cuisineMap map[string]*redblacktree.Tree[pair, struct{}]
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	foodMap := map[string]pair{}
	cuisineMap := map[string]*redblacktree.Tree[pair, struct{}]{}
	for i, food := range foods {
		rating, cuisine := ratings[i], cuisines[i]
		foodMap[food] = pair{rating, cuisine}
		if cuisineMap[cuisine] == nil {
			cuisineMap[cuisine] = redblacktree.NewWith[pair, struct{}](func(a, b pair) int {
				return cmp.Or(b.rating-a.rating, strings.Compare(a.food, b.food))
			})
		}
		cuisineMap[cuisine].Put(pair{rating, food}, struct{}{})
	}
	return FoodRatings{foodMap, cuisineMap}
}

func (r FoodRatings) ChangeRating(food string, newRating int) {
	p := r.foodMap[food]
	t := r.cuisineMap[p.food]
	t.Remove(pair{p.rating, food})           // 移除旧数据
	t.Put(pair{newRating, food}, struct{}{}) // 添加新数据
	p.rating = newRating
	r.foodMap[food] = p
}

func (r FoodRatings) HighestRated(cuisine string) string {
	return r.cuisineMap[cuisine].Left().Key.food
}
```

#### 复杂度分析

- 时间复杂度：初始化 $\mathcal{O}(n\log n)$，其余 $\mathcal{O}(\log n)$。其中 $n$ 为 $\textit{foods}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：懒删除堆

另一种做法是用懒删除堆：

- $\texttt{changeRating}$：直接往 $\textit{cuisineMap}$ 中记录，不做任何删除操作。
- $\texttt{highestRated}$：查看堆顶的食物评分是否等于其实际值，若不相同则意味着对应的元素已被替换成了其他值，堆顶存的是旧数据，直接弹出堆顶；否则堆顶就是答案。

```py [sol-Python3]
class FoodRatings:
    def __init__(self, foods: List[str], cuisines: List[str], ratings: List[int]):
        self.food_map = {}
        self.cuisine_map = defaultdict(list)
        for food, cuisine, rating in zip(foods, cuisines, ratings):
            self.food_map[food] = [rating, cuisine]
            self.cuisine_map[cuisine].append((-rating, food))
        # 这样可以保证初始化是线性复杂度
        for h in self.cuisine_map.values():
            heapify(h)

    def changeRating(self, food: str, newRating: int) -> None:
        p = self.food_map[food]
        # 直接添加新数据，后面 highestRated 再删除旧的
        heappush(self.cuisine_map[p[1]], (-newRating, food))
        p[0] = newRating

    def highestRated(self, cuisine: str) -> str:
        h = self.cuisine_map[cuisine]
        # 堆顶的食物评分不等于其实际值
        while -h[0][0] != self.food_map[h[0][1]][0]:
            heappop(h)
        return h[0][1]
```

```java [sol-Java]
class FoodRatings {
    private final Map<String, Pair<Integer, String>> foodMap = new HashMap<>();
    private final Map<String, PriorityQueue<Pair<Integer, String>>> cuisineMap = new HashMap<>();

    public FoodRatings(String[] foods, String[] cuisines, int[] ratings) {
        for (int i = 0; i < foods.length; i++) {
            String food = foods[i];
            String cuisine = cuisines[i];
            int rating = ratings[i];
            foodMap.put(food, new Pair<>(rating, cuisine));
            cuisineMap.computeIfAbsent(cuisine, k -> new PriorityQueue<>((a, b) ->
                    !Objects.equals(a.getKey(), b.getKey()) ? b.getKey() - a.getKey() : a.getValue().compareTo(b.getValue())
            )).offer(new Pair<>(rating, food));
        }
    }

    public void changeRating(String food, int newRating) {
        String cuisine = foodMap.get(food).getValue();
        // 直接添加新数据，后面 highestRated 再删除旧的
        cuisineMap.get(cuisine).offer(new Pair<>(newRating, food));
        foodMap.put(food, new Pair<>(newRating, cuisine));
    }

    public String highestRated(String cuisine) {
        PriorityQueue<Pair<Integer, String>> pq = cuisineMap.get(cuisine);
        // 堆顶的食物评分不等于其实际值
        while (!Objects.equals(pq.peek().getKey(), foodMap.get(pq.peek().getValue()).getKey())) {
            pq.poll();
        }
        return pq.peek().getValue();
    }
}
```

```cpp [sol-C++]
class FoodRatings {
    using pis = pair<int, string>;

    unordered_map<string, pair<int, string>> food_map;
    unordered_map<string, priority_queue<pis, vector<pis>, greater<>>> cuisine_map;

public:
    FoodRatings(vector<string>& foods, vector<string>& cuisines, vector<int>& ratings) {
        for (int i = 0; i < foods.size(); i++) {
            auto& food = foods[i];
            auto& cuisine = cuisines[i];
            int rating = ratings[i];
            food_map[food] = {rating, cuisine};
            cuisine_map[cuisine].emplace(-rating, food);
        }
    }

    void changeRating(string food, int newRating) {
        auto& [rating, cuisine] = food_map[food];
        // 直接添加新数据，后面 highestRated 再删除旧的
        cuisine_map[cuisine].emplace(-newRating, food);
        rating = newRating;
    }

    string highestRated(string cuisine) {
        auto& pq = cuisine_map[cuisine];
        // 堆顶的食物评分不等于其实际值
        while (-pq.top().first != food_map[pq.top().second].first) {
            pq.pop();
        }
        return pq.top().second;
    }
};
```

```go [sol-Go]
type pair struct {
	rating int
	s      string
}

type FoodRatings struct {
	foodMap    map[string]pair
	cuisineMap map[string]*hp
}

func Constructor(foods, cuisines []string, ratings []int) FoodRatings {
	foodMap := map[string]pair{}
	cuisineMap := map[string]*hp{}
	for i, food := range foods {
		rating, cuisine := ratings[i], cuisines[i]
		foodMap[food] = pair{rating, cuisine}
		if cuisineMap[cuisine] == nil {
			cuisineMap[cuisine] = &hp{}
		}
		heap.Push(cuisineMap[cuisine], pair{rating, food})
	}
	return FoodRatings{foodMap, cuisineMap}
}

func (r FoodRatings) ChangeRating(food string, newRating int) {
	p := r.foodMap[food]
	// 直接添加新数据，后面 highestRated 再删除旧的
	heap.Push(r.cuisineMap[p.s], pair{newRating, food})
	p.rating = newRating
	r.foodMap[food] = p
}

func (r FoodRatings) HighestRated(cuisine string) string {
	h := r.cuisineMap[cuisine]
	// 堆顶的食物评分不等于其实际值
	for h.Len() > 0 && (*h)[0].rating != r.foodMap[(*h)[0].s].rating {
		heap.Pop(h)
	}
	return (*h)[0].s
}

type hp []pair
func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.rating > b.rating || a.rating == b.rating && a.s < b.s
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any     { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```

#### 复杂度分析

- 时间复杂度：初始化 $\mathcal{O}(n)$ 或 $\mathcal{O}(n\log n)$，$\texttt{changeRating}$ 为 $\mathcal{O}(\log (n+q))$，$\texttt{highestRated}$ 为**均摊** $\mathcal{O}(\log (n+q))$。其中 $n$ 为 $\textit{foods}$ 的长度，$q$ 为 $\texttt{changeRating}$ 的调用次数。
- 空间复杂度：$\mathcal{O}(n+q)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
