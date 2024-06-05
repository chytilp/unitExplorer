package formatter

import "fmt"

type SimpleFormatter struct {
	ElementType string
}

func (s *SimpleFormatter) Print(elements []fmt.Stringer) {
	var element fmt.Stringer
	for _, el := range elements {
		element = el
		fmt.Printf("%s -> %s\n", s.ElementType, element)
	}
}
