type ListNode struct {
	val  int
	Next *ListNode
}

func ReverseLinkedList(head *ListNode) *ListNode {
	var (
		rev  *ListNode
		temp *ListNode
	)
	temp = head
	for temp != nil {
		next := temp.Next
		temp.Next = rev
		rev = temp
		temp = next
	}
	return rev
}

