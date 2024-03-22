package main

import "fmt"

type Bracket uint8

const (
	Parentheses = iota
	Square
	Curly
)

func (b Bracket) String() string {
	switch b {
	case Parentheses:
		return ")"
	case Square:
		return "]"
	case Curly:
		return "}"
	default:
		return ""
	}
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(a T) {
	s.elements = append(s.elements, a)
}

func (s *Stack[T]) Pop() (T, error) {
	res, err := s.Peek()
	if err != nil {
		return res, fmt.Errorf("Pop Error, stack is empty")
	}
	s.elements = s.elements[:len(s.elements)-1]
	return res, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("Stack Error, stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func isValid(s string) bool {
	stack := Stack[Bracket]{}
	for _, r := range s {
		fmt.Printf("%c \n", r)
		switch r {
		case '(':
			stack.Push(Parentheses)
		case '[':
			stack.Push(Square)
		case '{':
			stack.Push(Curly)
		case ')':
			bracket, err := stack.Pop()
			if err != nil || bracket != Parentheses {
				return false
			}
		case ']':
			bracket, err := stack.Pop()
			if err != nil || bracket != Square {
				return false
			}
		case '}':
			bracket, err := stack.Pop()
			if err != nil || bracket != Curly {
				return false
			}
		default:
			return false
		}
	}
	if stack.IsEmpty() {
		return true
	}
	return false
}

func main() {
	s := "([(]())"

	fmt.Println(isValid(s))
}
