package main

import "fmt"

type stack struct {
	items []string
}

func (s *stack) add(item string) {
	s.items = append(s.items[:len(s.items)], item)
}

func (s *stack) pop() (string, bool) {
	if len(s.items) == 0 {
		return "", false
	}
	element := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return element, true
}

func isBalanced(s string) bool {

	dict := make(map[string]string)
	dict["{"] = "}"
	dict["("] = ")"
	dict["["] = "]"

	stack := stack{}
	for _, value := range s {
		c := string(value)
		if c == "(" || c == "{" || c == "[" {
			stack.add(c)
			continue
		}

		pop, ok := stack.pop()
		if !ok {
			return false
		}

		closedBracket, ok := dict[pop]
		if !ok || closedBracket != c {
			return false
		}
	}

	return len(stack.items) == 0
}

func main() {
	fmt.Println(isBalanced("{[()]}"))
}
