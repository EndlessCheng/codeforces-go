// https://loj.ac/submission/56207

#include <cstdio>
#include <cassert>
#include <cmath>
#include <cstring>

#include <iostream>
#include <algorithm>
#include <vector>
#include <map>
#include <set>
#include <functional>
#include <stack>
#include <queue>

#include <tuple>

#define getchar getchar_unlocked
#define putchar putchar_unlocked

#define _rep(_1, _2, _3, _4, name, ...) name
#define rep2(i, n) rep3(i, 0, n)
#define rep3(i, a, b) rep4(i, a, b, 1)
#define rep4(i, a, b, c) for (int i = int(a); i < int(b); i += int(c))
#define rep(...) _rep(__VA_ARGS__, rep4, rep3, rep2, _)(__VA_ARGS__)

using namespace std;

using i64 = long long;
using u8 = unsigned char;
using u32 = unsigned;
using u64 = unsigned long long;
using f80 = long double;

int get_int() {
    int n, c, sign = 0;
    while ((c = getchar()) < '-')
        ;
    if (c == '-')
        sign = 1, n = 0;
    else
        n = c - '0';
    while ((c = getchar()) >= '0') n = n * 10 + c - '0';
    return sign ? -n : n;
}

template <typename CapType, typename TotalCapType, typename CostType, typename TotalCostType>
class CostScaling {
private:
    static const int alpha = 8;  // eps <- max(1, eps / alpha)

    using cap_t = CapType;
    using tcap_t = TotalCapType;
    using cost_t = CostType;  // > max{|C|} * (2 * |V|)
    using tcost_t = TotalCostType;
    static constexpr cost_t Inf = (tcap_t(1) << (sizeof(tcap_t) * 8 - 2)) - 1;

    struct InputEdge {
        int from, to;
        cap_t b, c;
        cost_t cost;
    };
    struct Edge {
        int to, rev;
        cap_t cap;
        cost_t cost;
    };

    class Dinic {
    public:
        Dinic(int N, const vector<int>& ofs, vector<Edge>& edges, vector<tcap_t>& capacity)
            : N(N), ofs(ofs), edges(edges), capacity(capacity), last(N) {}

        bool succeeded() {
            // s -> u: capacity[u]
            // u -> t: capacity[u + N]
            tcap_t f = 0;
            for (int u = 0; u < N; ++u) f += capacity[u];
            vector<int> que(N);
            while (f) {
                dist.assign(N, -1);
                int qh = 0, qt = 0, lv = N;
                for (int u = 0; u < N; ++u)
                    if (capacity[u] > 0)
                        que[qt++] = u, dist[u] = 0;
                for (; qh < qt;) {
                    int u = que[qh++];
                    if (lv == N && capacity[u + N] > 0)
                        lv = dist[u];
                    if (dist[u] > lv)
                        break;
                    for (int ei = ofs[u]; ei < ofs[u + 1]; ++ei) {
                        int v = edges[ei].to;
                        if (edges[ei].cap > 0 && dist[v] == -1) {
                            que[qt++] = v, dist[v] = dist[u] + 1;
                        }
                    }
                }
                if (lv == N)
                    break;
                for (int u = 0; u < N; ++u) last[u] = ofs[u];
                for (int u = 0; u < N; ++u)
                    if (capacity[u] > 0) {
                        auto df = block_flow(u, capacity[u]);
                        f -= df, capacity[u] -= df;
                    }
            }
            return f == 0;
        }

    private:
        tcap_t block_flow(int u, tcap_t f) {
            tcap_t ret = 0;
            if (capacity[u + N] > 0) {
                tcap_t df = min(f, capacity[u + N]);
                capacity[u + N] -= df;
                return df;
            }
            for (auto& ei = last[u]; ei < ofs[u + 1]; ++ei) {
                auto& e = edges[ei];
                int v = e.to;
                if (e.cap == 0 || dist[v] <= dist[u])
                    continue;
                cap_t df = block_flow(v, min<cap_t>(e.cap, f));
                if (df == 0)
                    continue;
                e.cap -= df, edges[e.rev].cap += df;
                f -= df, ret += df;
                if (f == 0)
                    break;
            }
            return ret;
        }

        int N;
        const vector<int>& ofs;
        vector<Edge>& edges;
        vector<tcap_t>& capacity;
        vector<int> last, dist;
    };

public:
    CostScaling(int N, int M = 0) : N(N), capacity(2 * N) {
        if (M > 0)
            in.reserve(M);
    }

    void add_directed_edge(int u, int v, cap_t b, cap_t c, cost_t cost) {
        if (b > 0)
            capacity[v] += b, capacity[u + N] += b;
        else
            capacity[u] += -b, capacity[v + N] += -b;
        in.push_back({ u, v, b, c, cost });
    }

    pair<bool, tcost_t> minimum_cost_circulation() {
        construct();
        if (!has_feasible_circulation())
            return { false, 0 };

        const int cost_multiplier = 2 << __lg(N);  // should be > |V|
        cost_t eps = 0;
        for (auto& e : edges) e.cost *= cost_multiplier, eps = max(eps, e.cost);

        while (eps > 1) refine(eps = max<cost_t>(1, eps / alpha));

        tcost_t ret = initial_cost;
        for (auto& e : edges) ret -= (e.cost / cost_multiplier) * e.cap;
        return { true, ret / 2 };
    }

private:
    void refine(const cost_t eps) {
        auto cost_p = [&](int u, const Edge& e) { return e.cost + potential[u] - potential[e.to]; };
        for (int u = 0; u < N; ++u)
            for (int i = ofs[u]; i < ofs[u + 1]; ++i) {
                auto& e = edges[i];
                if (cost_p(u, e) < 0)
                    edges[e.rev].cap += e.cap, e.cap = 0;
            }
        vector<tcap_t> excess(initial_excess);
        for (auto& e : edges) excess[e.to] -= e.cap;

        vector<int> stack;
        stack.reserve(N);
        for (int u = 0; u < N; ++u)
            if (excess[u] > 0)
                stack.push_back(u);

        auto residue = [&](const Edge& e) -> cap_t { return e.cap; };
        auto push = [&](int u, Edge& e, cap_t df) {
            e.cap -= df;
            edges[e.rev].cap += df;
            excess[e.to] += df;
            excess[u] -= df;
            if (excess[e.to] > 0 && excess[e.to] <= df) {
                stack.push_back(e.to);
            }
        };
        auto relabel = [&](int u, cost_t delta) { potential[u] -= delta + eps; };
        auto relabel_in_advance = [&](int u) {
            if (excess[u] != 0)
                return false;
            auto delta = Inf;
            for (int ei = ofs[u]; ei < ofs[u + 1]; ++ei) {
                auto& e = edges[ei];
                if (residue(e) == 0)
                    continue;
                if (cost_p(u, e) < 0)
                    return false;
                else
                    delta = min<tcost_t>(delta, cost_p(u, e));
            }
            relabel(u, delta);
            return true;
        };
        auto discharge = [&](int u) {
            auto delta = Inf;
            for (int ei = ofs[u]; ei < ofs[u + 1]; ++ei) {
                auto& e = edges[ei];
                if (residue(e) == 0)
                    continue;
                if (cost_p(u, e) < 0) {
                    if (relabel_in_advance(e.to)) {
                        --ei;
                        continue;  // modify ei (!)
                    }
                    cap_t df = min<tcap_t>(excess[u], residue(e));
                    push(u, e, df);
                    if (!excess[u])
                        return;
                } else
                    delta = min<tcost_t>(delta, cost_p(u, e));
            }
            relabel(u, delta);
            stack.push_back(u);
        };
        while (!stack.empty()) {
            auto u = stack.back();
            stack.pop_back();
            discharge(u);
        }
    }

    void construct() {
        ofs.assign(N + 1, 0);
        edges.resize(2 * in.size());
        initial_excess.assign(N, 0);
        initial_cost = 0;
        potential.assign(N, 0);
        for (auto& e : in) ofs[e.from + 1]++, ofs[e.to + 1]++;
        for (int i = 1; i <= N; ++i) ofs[i] += ofs[i - 1];
        for (auto& e : in) {
            initial_excess[e.to] += e.c;
            initial_excess[e.from] += -e.b;
            initial_cost += tcost_t(e.cost) * (e.c + e.b);
            edges[ofs[e.from]++] = { e.to, ofs[e.to], e.c - e.b, e.cost };
            edges[ofs[e.to]++] = { e.from, ofs[e.from] - 1, 0, -e.cost };
        }
        for (int i = N; i > 0; --i) ofs[i] = ofs[i - 1];
        ofs[0] = 0;
    }

    bool has_feasible_circulation() { return Dinic(N, ofs, edges, capacity).succeeded(); }

private:
    int N;
    vector<InputEdge> in;
    vector<tcap_t> capacity;

    vector<int> ofs;
    vector<Edge> edges;

    tcost_t initial_cost;
    vector<tcap_t> initial_excess;
    vector<tcost_t> potential;
};
// cap, total_cap, cost * (2 * |V|), total_cost
using MCC = CostScaling<int64_t, int64_t, int64_t, int64_t>;
// using MCC = CostScaling<int, int, int, int>;

void solve() {
    int N, M;
    while (~scanf("%d %d", &N, &M)) {
        MCC mcc1(N, M), mcc2(N, M);
        rep(i, M) {
            int u = get_int() - 1, v = get_int() - 1;
            int c = get_int(), cost = get_int();
            mcc1.add_directed_edge(u, v, 0, c, 0);
            mcc2.add_directed_edge(u, v, 0, c, cost);
        }
        mcc1.add_directed_edge(N - 1, 0, 0, 1e18, -1);
        auto max_flow = -mcc1.minimum_cost_circulation().second;
        mcc2.add_directed_edge(N - 1, 0, max_flow, max_flow, 0);
        printf("%lld %lld\n", max_flow, mcc2.minimum_cost_circulation().second);
    }
}

int main() {
    clock_t beg = clock();
    solve();
    clock_t end = clock();
    fprintf(stderr, "%.3f sec\n", double(end - beg) / CLOCKS_PER_SEC);
    return 0;
}