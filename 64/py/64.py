class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        # base cases
        for i in range(1, len(grid), 1):
            grid[i][0] = grid[i][0] + grid[i-1][0]
        for j in range(1, len(grid[0]), 1):
            grid[0][j] = grid[0][j] + grid[0][j-1]

        # recursive case
        for i in range(1, len(grid), 1):
            for j in range(1, len(grid[i]), 1):
                grid[i][j] = grid[i][j] + min(grid[i-1][j], grid[i][j-1])

        return grid[len(grid) - 1][len(grid[0]) - 1]