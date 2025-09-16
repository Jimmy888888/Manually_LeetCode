/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func hasCycle(head *ListNode) bool {
    // check
    visited := make(map[*ListNode]bool)
    curN := head

    for curN != nil {
        if _, find := visited[curN]; find {
            return true
        }

        visited[curN] = true
        curN = curN.Next
    }

    return false
}