package jumpgame

import "fmt"

func canJump(nums []int) bool {

    i := 0
    step := nums[0]
    newStep := step
    position := 0
    for i < len(nums){
        // check success
        if i + step >= len(nums) - 1 {
            return true
        }else if i + step <= i {
            return false
        }
        
        // find the next position
        newStep = 0
        position = i+1
        for j := i+1; j <= i + step; j++ {
            if j + nums[j] >= i + step {
                newStep = nums[j]
                position = j
                break
                // fmt.Print(position)
                // fmt.Print(newStep)
            }
        }

        step = newStep

        i = position
        // fmt.Print(i)
        // fmt.Print(step)
    }

    return true
}

// greedy : in every step, choose the biggest number
// success condition : the position reach the last index
// values :
// // step : 
// // position : 
