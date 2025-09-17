// Backtracking

// number of left "(" >= number of right ")" always

func generateParenthesis(n int) []string {
    var result []string

    var Backtracking func(curStr string, leftNum int, rightNum int)

    Backtracking = func(curStr string, leftNum int, rightNum int) {
        if len(curStr) == 2 * n {
            result = append(result, curStr)
            return
        }

        if leftNum < n {
            Backtracking(curStr + "(", leftNum + 1, rightNum)
        }

        if rightNum < leftNum {
            Backtracking(curStr + ")", leftNum, rightNum + 1)
        }
    }

    Backtracking("", 0, 0)

    return result
}