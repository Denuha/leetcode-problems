package problems

import "reflect"

func AddTwoNumbersTest() bool {
	var L1a *ListNode = &ListNode{
		Val:  3,
		Next: nil,
	}

	var L1b *ListNode = &ListNode{
		Val:  4,
		Next: L1a,
	}

	var L1 *ListNode = &ListNode{
		Val:  2,
		Next: L1b,
	}

	var L2a *ListNode = &ListNode{
		Val:  4,
		Next: nil,
	}

	var L2b *ListNode = &ListNode{
		Val:  6,
		Next: L2a,
	}

	var L2 *ListNode = &ListNode{
		Val:  5,
		Next: L2b,
	}

	var resa *ListNode = &ListNode{
		Val:  8,
		Next: nil,
	}

	var resb *ListNode = &ListNode{
		Val:  0,
		Next: resa,
	}

	var res *ListNode = &ListNode{
		Val:  7,
		Next: resb,
	}

	if !reflect.DeepEqual(addTwoNumbers(L1, L2), res) {
		return false
	}

	return true
}
