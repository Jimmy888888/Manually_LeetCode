package jumpgame

func canJump(nums []int) bool {
	reachable := 0
	for i, n := range nums {
		if i > reachable {
			return false
		}
		if i+n > reachable {
			reachable = i + n
		}
	}
	return true
}
