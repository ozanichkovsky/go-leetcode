package p2AddTwoNumbers

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	cur := l3
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if carry == 1 {
			sum += 1
		}

		if sum >= 10 {
			sum = sum % 10
			carry = 1
		} else {
			carry = 0
		}

		cur.Val = sum

		if !(l1 == nil && l2 == nil) {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}

	if carry == 1 {
		cur.Next = &ListNode{}
		cur.Next.Val = 1
	}

	return l3
}
