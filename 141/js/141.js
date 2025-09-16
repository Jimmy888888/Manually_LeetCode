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