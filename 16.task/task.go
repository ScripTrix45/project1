package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type Stack struct {
	Head *Node
}

type Queue struct {
	Head *Node
	Tail *Node
}

// стек
func (s *Stack) Push(val int) {
	newNode := &Node{Value: val, Next: s.Head, Prev: nil}

	if s.Head != nil {
		s.Head.Prev = newNode
	}
	s.Head = newNode

}

func (s *Stack) Pop() int {
	if s.Head == nil {
		fmt.Println("Error. Stack is empty")
		return -1 // error,
	}
	value := s.Head.Value
	s.Head = s.Head.Next
	if s.Head != nil {
		s.Head.Prev = nil
	} else {
		s.Head = nil
	}
	return value

}

func (s *Stack) Peek() int {
	if s.Head == nil {
		fmt.Println("Error. Stack is empty")
		return -1
	}
	return s.Head.Value
}

func (s *Stack) isEmpty() bool {
	return s.Head == nil
}

// очередь
func (q *Queue) Enqueue(val int) {
	newNode := &Node{Value: val, Next: nil, Prev: q.Tail}
	if q.Tail != nil {
		q.Tail.Next = newNode
	} else {
		q.Head = newNode
	}

	q.Tail = newNode
}

func (q *Queue) Dequeue() int {
	if q.Head == nil {
		fmt.Println("Eror. Queue is empty")
		return -1
	}
	res := q.Head.Value
	q.Head = q.Head.Next
	if q.Head != nil {
		q.Head.Prev = nil
	} else {
		q.Head = nil
		q.Tail = nil
	}

	return res
}

func (q *Queue) Peek() int {
	if q.Head == nil {
		fmt.Println("Eror. Queue is empty")
		return -1
	}
	return q.Head.Value
}

func (q *Queue) isEmpty() bool {
	return q.Head == nil
}

func main() {
	fmt.Println("=== Stack Test ===")
	var s Stack
	s.Push(10)
	s.Push(20)
	s.Push(30)

	fmt.Println("Peek:", s.Peek())       // Ожидаем 30
	fmt.Println("Pop:", s.Pop())         // Ожидаем 30
	fmt.Println("Pop:", s.Pop())         // Ожидаем 20
	fmt.Println("IsEmpty:", s.isEmpty()) // Ожидаем false
	fmt.Println("Pop:", s.Pop())         // Ожидаем 10
	fmt.Println("IsEmpty:", s.isEmpty()) // Ожидаем true
	fmt.Println("Pop:", s.Pop())         // Ошибка: стек пуст

	fmt.Println("\n=== Queue Test ===")
	var q Queue
	q.Enqueue(100)
	q.Enqueue(200)
	q.Enqueue(300)

	fmt.Println("Peek:", q.Peek())       // Ожидаем 100
	fmt.Println("Dequeue:", q.Dequeue()) // Ожидаем 100
	fmt.Println("Dequeue:", q.Dequeue()) // Ожидаем 200
	fmt.Println("IsEmpty:", q.isEmpty()) // Ожидаем false
	fmt.Println("Dequeue:", q.Dequeue()) // Ожидаем 300
	fmt.Println("IsEmpty:", q.isEmpty()) // Ожидаем true
	fmt.Println("Dequeue:", q.Dequeue()) // Ошибка: очередь пуста
}
