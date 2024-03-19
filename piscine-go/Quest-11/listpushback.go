package piscine

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	currentNode := l.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	currentNode.Next = newNode

	l.Tail = newNode
}
