func minPathSum(grid [][]int) int {

    // base cases
    for i := 1; i < len(grid); i++ {
        grid[i][0] = grid[i][0] + grid[i-1][0]
    }

    for j := 1; j < len(grid[0]); j++ {
        grid[0][j] = grid[0][j] + grid[0][j-1]
    }

    // recursive case
    for i := 1; i < len(grid); i++ {
        for j := 1; j < len(grid[i]); j++{
            grid[i][j] = grid[i][j] + minNumber(grid[i-1][j], grid[i][j-1])
        }
    }

    return grid[len(grid) - 1][len(grid[0]) - 1]
}

func minNumber( a int, b int) int {
    if a > b {
        return b
    }
    return a
}