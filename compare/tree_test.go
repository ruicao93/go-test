package compare

import "testing"

// Definition of binary tree node.
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// Find max depth of binary tree.
func FindMaxDepth(head *Node) int {
	if head == nil {
		return 0
	}
	leftMax := FindMaxDepth(head.Left)
	rightMax := FindMaxDepth(head.Right)
	if leftMax >= rightMax {
		return leftMax + 1
	} else {
		return rightMax + 1
	}
}

// Generate binary tree from array. For non-exist node, use number: -1.
func GenerateBinaryTree(nums []int) *Node {
	if len(nums) <= 0 {
		return nil
	}
	var nodeList []*Node
	for index, num := range nums {
		if num == -1 {
			continue
		}
		newNode := Node{Data: num}
		nodeList = append(nodeList, &newNode)
		if index <= 0 {
			continue
		} else if (index-1)%2 == 0 {
			nodeList[(index-1)/2].Left = &newNode
		} else {
			nodeList[(index-1)/2].Right = &newNode
		}
	}
	return nodeList[0]
}

type TestCase struct {
	Tree        []int
	ExpectedRes int
}

// Test Function
func TestFindMaxDepth(t *testing.T) {
	testCaseList := []TestCase{
		{[]int{}, 0},
		{[]int{1, -1, 2}, 2},
		{[]int{3, 9, 20, -1, -1, 15, 7}, 3},
	}
	for _, testCase := range testCaseList {
		res := FindMaxDepth(GenerateBinaryTree(testCase.Tree))
		if testCase.ExpectedRes != res {
			t.Fatalf("Fail. For case: %v. Expected result is: %v, but got: %v", testCase, testCase.ExpectedRes, res)
		}
	}
	// TODO: more test cases.
}
