package stack

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(el string) {
	*s = append(*s, el)
}

func (s *Stack) Pop() (string, bool) {
	if (*s).IsEmpty() {
		return "", false
	}

	idx := len(*s) - 1
	el := (*s)[idx]
	*s = (*s)[:idx]
	return el, true
}
