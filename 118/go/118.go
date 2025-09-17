func generate(numRows int) [][]int {
    train := make([][]int, numRows)
    train[0] = []int{1}

    for i := 1; i < numRows; i++ {
        newR := make([]int, i+1)
        for j := 0; j < i+1; j++{
            if j == 0 || j == i {
                newR[j] = 1
            }else{
                newR[j] = train[i-1][j] + train[i-1][j-1]
            }
        }
        train[i] = newR
    }
    return train
}