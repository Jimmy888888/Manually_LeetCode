func plusOne(digits []int) []int {
    n := len(digits)
    carry := 0

    if digits[n-1] + 1 < 10 {
        digits[n-1] = digits[n-1] + 1
        return digits
    }else{
        digits[n-1] = (digits[n-1] + 1) % 10
        carry = 1
    }

    for i := n-2; i >= 0; i-- {
        if digits[i] + 1 < 10 {
            digits[i] = digits[i] + 1
            carry = 0
            break
        }else{
            digits[i] = (digits[i] + 1) % 10
            carry = 1
        }
    }

    if carry == 1 {
        digits = append([]int{1}, digits...)
    }

    return digits
}