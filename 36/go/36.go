import (
    "fmt"
)

func isValidSudoku(board [][]byte) bool {
    // use map to record repeatition
    n := len(board)
    rowCheck := make([]map[byte]int, 9)
    colCheck := make([]map[byte]int, 9)
    boxCheck := make([]map[byte]int, 9)

    for i := 0; i< n; i++ {
        rowCheck[i] = make(map[byte]int)
        colCheck[i] = make(map[byte]int)
        boxCheck[i] = make(map[byte]int)
    }

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++{
            str := board[i][j]
            if str == '.'{
                continue
            }

            // map the rows and columns
            r := i
            c := j
            // map the nine boxes
            b := (i/3)*3 + (j/3)

            if rowCheck[r][str] > 0 || colCheck[c][str] > 0 || boxCheck[b][str] > 0{
                return false
            }

            rowCheck[r][str] += 1
            colCheck[c][str] += 1
            boxCheck[b][str] += 1

        }
    }
    return true
}