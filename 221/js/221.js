/**
 * @param {character[][]} matrix
 * @return {number}
 */
var maximalSquare = function(matrix) {
    // make 2d dp
    let m = matrix.length;
    let n = matrix[0].length;
    let dp = new Array(m).fill(null).map(() => Array(n).fill(0));
    let maxSqu = 0;

    // State transition equation
    for (let i = 0; i < m; i++){
        for (let j = 0; j < n; j++){
            if (matrix[i][j] === "1") {
                if (i === 0 || j === 0){
                    dp[i][j] = 1;
                }else{
                    dp[i][j] = 1 + Math.min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]);
                }

                if (dp[i][j] > maxSqu) {
                    maxSqu = dp[i][j];
                }
            }
        }
    }

    return maxSqu * maxSqu;
};