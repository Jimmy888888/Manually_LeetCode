// DP

// contiguous
// at every step i:
// if nums[i] > nums[i] + curS -> curS = nums[i] , else curS += nums[i]
// then:
// if golS < curS -> golS = curS

func maxSubArray(nums []int) int {
    // check 
    n := len(nums)
    if n == 1 {
        return nums[0]
    }

    curS := nums[0]
    golS := nums[0]

    for i := 1; i < n; i++ {
        num := nums[i]
        if num > num + curS {
            curS = num
        }else{
            curS += num
        }

        if golS < curS {
            golS = curS
        }
    }
    return golS
}