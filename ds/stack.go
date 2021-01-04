package ds

type Stack []interface{}

func (s *Stack) Push(i interface{}) {
	*s = append(*s, i)
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return (*s)[s.Len()-1]
}

func (s *Stack) Pop() interface{} {
	defer func() { *s = (*s)[0 : s.Len()-1] }()
	return s.Peek()
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}
