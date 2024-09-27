/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	// 1. check if head == nil
	// 2. find the length
	//3. find the mid
	//4. remove the middle element
	// 5 print the element

	if head.Next == nil {
		return nil
	}

	count := 0
	mid := 0
	temp := head
	for temp != nil {
		count++
		temp = temp.Next
	}
	mid = int(count / 2)
	temp = head
	for i := 0; i < mid-1; i++ {
		temp = temp.Next
	}
	temp.Next = temp.Next.Next
	return head
}
