package piscine

func ListReverse(l *List) {
	if l.Head == nil || l.Head.Next == nil {
		return
	}

	var prev, current, next *NodeL
	current = l.Head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head, l.Tail = l.Tail, l.Head
}
