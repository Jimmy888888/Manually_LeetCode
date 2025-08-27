func jump(nums []int) int {

    // check length
    if len(nums) <= 1 {
        return 0;
    }


    // Greedy
    // check every step

    minSteps := 0
    curFarthest := 0
    curBound := 0

    for i, n := range nums {
        if curFarthest < i + n {
            curFarthest = i + n
        }

        if i == curBound {
            curBound = curFarthest
            minSteps += 1
            if curBound >= len(nums) - 1{
                return minSteps
            }
        }
    }

    return minSteps
}