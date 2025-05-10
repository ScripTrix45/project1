package main

import "fmt"

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) addToList(val int) { // дженерики для всех числовых типов

	newNode := &Node{Value: val, Next: l.Head}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	} // записали новую голову 0. next - nil
	l.Head.Prev = newNode // записываем новую голову 1. next - 0, prev у 0 стал 1
	l.Head = newNode      // теперь новая голова 1. next у нее l.Head, как и описано выше

}

func main() {
	var l LinkedList
	sl := []int{0, 1, 2, 3, 4, 5}

	for _, val := range sl {
		l.addToList(val)
	}

	current := l.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}

	current2 := l.Tail
	for current2 != nil {
		fmt.Println(current2.Value)
		current2 = current2.Prev
	}

}
