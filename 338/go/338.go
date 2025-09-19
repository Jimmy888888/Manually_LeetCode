import "math/bits"
func countBits(n int) []int {
    res := make([]int, n+1)

    for i := 0; i < n+1; i++ {
        // go native function for counting the amount of 1 in a binary 
        res[i] = bits.OnesCount(uint(i))
    }
    return res
}