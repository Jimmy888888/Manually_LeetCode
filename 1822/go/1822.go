func arraySign(nums []int) int {
    res := 1
    for _, num := range nums {
        if num == 0 {
            return 0
        }else if num < 0 {
            res = res * -1
        }
    }

    return res
}