package logging

type Logger []string

func (s *Logger) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Logger) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack. In a real scenario this would change for an external service
}

func (s *Logger) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Logger) GetItems() []string {
	return *s
}
