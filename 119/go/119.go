func getRow(rowIndex int) []int {
    if rowIndex == 0 {
        return []int{1}
    }

    var result [][]int
    result = append(result, []int{1})

    for i := 1; i <= rowIndex; i++ {
        cur := make([]int, i+1)
        for j := 0; j < i + 1; j ++ {
            if j == 0 || j == i {
                cur[j] = 1
            }else{
                cur[j] = result[i-1][j-1] + result[i-1][j]
            }
        }
        result = append(result, cur)
    }

    return result[rowIndex]
}