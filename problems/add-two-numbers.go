package problems

// https://leetcode.com/problems/add-two-numbers/
// resolve from https://leetcode.com/problems/add-two-numbers/discuss/1340/A-summary-about-how-to-solve-Linked-List-problem-C%2B%2B

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var c int = 0
	var newHead *ListNode = &ListNode{}
	var t *ListNode = newHead
	for ok := true; ok; ok = c > 0 || l1 != nil || l2 != nil {
		var a int = 0
		if l1 != nil {
			a = l1.Val
		} else {
			a = 0
		}

		var b int = 0
		if l2 != nil {
			b = l2.Val
		} else {
			b = 0
		}

		c += a + b

		t.Next = &ListNode{
			Val:  c % 10,
			Next: nil,
		}
		t = t.Next
		c /= 10

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	return newHead.Next
}
