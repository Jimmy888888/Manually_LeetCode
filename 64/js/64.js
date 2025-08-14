/**
 * @param {number[][]} grid
 * @return {number}
 */
var minPathSum = function(grid) {

    // handle base cases
    for (let i = 1; i < grid.length; i++){
       grid[i][0] = grid[i][0] + grid[i - 1][0];
    }

    for (let j = 1; j < grid[0].length; j++){
       grid[0][j] = grid[0][j] + grid[0][j - 1];
    }

    // recursive case
    for (let i = 1; i < grid.length; i++){
        for (let j = 1; j < grid[i].length; j++){
            grid[i][j] = grid[i][j] + Math.min(grid[i - 1][j], grid[i][j - 1]);
        }
    }

    return grid[grid.length - 1][grid[0].length - 1];
};