package symmetric_test

import (
	"testing"

	symmetric "example.com/symmetric"
)

func TestIsSymmetric(t *testing.T) {
	// Test case 1: A symmetric tree
	//       1
	//      / \
	//     2   2
	//    / \ / \
	//   3  4 4  3
	tree1 := &symmetric.TreeNode{
		Val: 1,
		Left: &symmetric.TreeNode{
			Val: 2,
			Left: &symmetric.TreeNode{Val: 3},
			Right: &symmetric.TreeNode{Val: 4},
		},
		Right: &symmetric.TreeNode{
			Val: 2,
			Left: &symmetric.TreeNode{Val: 4},
			Right: &symmetric.TreeNode{Val: 3},
		},
	}
	if !symmetric.IsSymmetric(tree1) {
		t.Errorf("IsSymmetric(tree1) = false; want true")
	}

	// Test case 2: An asymmetric tree
	//       1
	//      / \
	//     2   2
	//      \   \
	//      3    3
	tree2 := &symmetric.TreeNode{
		Val: 1,
		Left: &symmetric.TreeNode{
			Val: 2,
			Right: &symmetric.TreeNode{Val: 3},
		},
		Right: &symmetric.TreeNode{
			Val: 2,
			Right: &symmetric.TreeNode{Val: 3},
		},
	}
	if symmetric.IsSymmetric(tree2) {
		t.Errorf("IsSymmetric(tree2) = true; want false")
	}

	// Test case 3: Root is nil
	if !symmetric.IsSymmetric(nil) {
		t.Errorf("IsSymmetric(nil) = false; want true")
	}

	// Test case 4: A single node tree
	tree4 := &symmetric.TreeNode{Val: 1}
	if !symmetric.IsSymmetric(tree4) {
		t.Errorf("IsSymmetric(tree4) = false; want true")
	}
}