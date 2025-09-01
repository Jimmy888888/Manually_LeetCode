/**
 * @param {number[][]} obstacleGrid
 * @return {number}
 */


var uniquePathsWithObstacles = function(obstacleGrid) {
    let m = obstacleGrid.length;
    let n = obstacleGrid[0].length;
    // make a 2D array
    let DP = Array(m).fill(0).map(() => Array(n).fill(0));

    // base cases
    for (let i = 0; i < m; i++){
        if (obstacleGrid[i][0] === 1){
            break;
        }
        DP[i][0] = 1;
    }

    for (let j = 0; j < n; j++){
        if (obstacleGrid[0][j] === 1){
            break;
        }
        DP[0][j] = 1;
    }

    // transition equation
    for (let i = 1; i < m; i++){
        for (let j = 1; j < n; j++){
            if(obstacleGrid[i][j] === 1){
                continue;
            }

            DP[i][j] = DP[i - 1][j] + DP[i][j - 1];
        }
    }

    return DP[m-1][n-1];
};