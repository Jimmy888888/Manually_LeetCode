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