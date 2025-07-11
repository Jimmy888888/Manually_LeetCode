class TreeNode:
  def __init__(self, val=0, left=None, right=None):
    self.val = val
    self.left = left
    self.right = right

class Solution:
  def isSymmetric(self, root: Optional[TreeNode]) -> bool: # Python automatically passes self
    if not root:
      return True

    return self.isMirror(root.left, root.right)

  def isMirror(self, t1: Optional[TreeNode], t2: Optional[TreeNode]) -> bool:
    if not ti and not t2:
      return True
    if not ti or not t2:
      return False

    return (t1.val == t2.val) and self.isMirror(t1.left, t2.right) and self.isMirror(t1.right, t2.left)
