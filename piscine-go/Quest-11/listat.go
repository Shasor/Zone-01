package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	current := l

	for i := 0; current != nil && i < pos; i++ {
		current = current.Next
	}

	return current
}
