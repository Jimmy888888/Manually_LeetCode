# 141 Linked List Cycle

## hash map

```go
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
```

## fast slow pointer

```js
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} head
 * @return {boolean}
 */
var hasCycle = function(head) {

    let point = head;
    while (point !== null && point.next !== null){
        head = head.next;
        point = point.next.next;

        if (point === head){
            return true;
        }
    }
    return false;
};
```

## marked

```py
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        while head is not None :
            if head.val > 100000 :
                return True
            head.val = 100001 # Node.val <= 10^5
            head = head.next
        
        return False
```