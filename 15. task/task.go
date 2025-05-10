package main

import "fmt"

type Stack struct {
	slice []int
}

type Queue struct {
	slice []int
}

func (s *Stack) Push(val int) {
	s.slice = append(s.slice, val)
}

func (s *Stack) Pop() int {
	if len(s.slice) == 0 {
		panic("stack is empty") // Или: fmt.Println("Stack is empty")
		return -1
	}
	lastElem := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	return lastElem
}

func (s *Stack) Peek() int {
	if len(s.slice) == 0 {
		panic("stack is empty")
		return -1
	}
	return s.slice[len(s.slice)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.slice) == 0
}

func (q *Queue) Enqueue(val int) {
	q.slice = append(q.slice, val)
}

func (q *Queue) Dequeue() int {
	if len(q.slice) == 0 {
		panic("queue is empty")
		return -1
	}
	firstVal := q.slice[0]
	q.slice = q.slice[1:]
	return firstVal
}

func (q *Queue) Peek() int {
	if len(q.slice) == 0 {
		panic("queue is empty")
		return -1
	}
	return q.slice[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.slice) == 0
}

func main() {
	var s Stack
	s.Push(0)
	s.Push(1)
	s.Push(2)

	fmt.Println("Stack after Push operations:", s.slice)

	num := s.Pop()
	fmt.Println("Pop:", num)
	fmt.Println("Stack after Pop:", s.slice)

	lastNum := s.Peek()
	fmt.Println("Peek:", lastNum)

	check := s.IsEmpty()
	fmt.Println("Is stack empty?", check)

	// Очередь
	var q Queue
	q.Enqueue(7)
	q.Enqueue(8)
	q.Enqueue(9)

	fmt.Println("Queue after Enqueue operations:", q.slice)

	fVal := q.Dequeue()
	fmt.Println("Dequeue:", fVal)
	fmt.Println("Queue after Dequeue:", q.slice)

	newFval := q.Peek()
	fmt.Println("Peek:", newFval)

	checkQueue := q.IsEmpty()
	fmt.Println("Is queue empty?", checkQueue)
}
