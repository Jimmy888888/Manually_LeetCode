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