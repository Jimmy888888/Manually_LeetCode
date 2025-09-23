# 746

## DP

assume dp[i] is the min cost at cost[i]

### state transition equation:

Can jump one or two steps at once

dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])

if (i+1 > len(cost) || i+2 > len(cost)) {end}

### base cases:

dp[0], dp[1] = 0, 0

```c
int getMin(int a, int b) {
    return (a < b)? a:b ;
}

int minCostClimbingStairs(int* cost, int costSize) {
    // make dp
    int dp[costSize+1];
    memset(dp, 0, sizeof(dp));

    // state transition equation
    for (int i = 2; i < costSize+1; i++){
        dp[i] = getMin(dp[i-1] + cost[i-1], dp[i-2] + cost[i-2]);
    }

    return dp[costSize];
}
```
