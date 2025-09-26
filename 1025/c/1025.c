bool divisorGame(int n) {
    // check
    if (n == 1){
        return false;
    }
    
    // make dp
    bool* dp = (bool*)malloc((n+1) * sizeof(bool));
    memset(dp, 0, (n+1) * sizeof(bool));
    // dp[1] = false

    // go through , and Alice go first
    for (int i = 2; i <= n; i++){
        for (int j = 1; j < i; j++){
            if (i % j == 0){
                if (dp[i-j] == false){
                    dp[i] = true;
                    break;
                }
            }
        }
    }

    bool res = dp[n];
    free(dp);

    return res;
}