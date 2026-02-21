package pilhas

type Stack struct {
	Items []any
}

func (s *Stack) push(item any) {
	s.Items = append(s.Items, item)
}

func (s *Stack) pop() any {
	if len(s.Items) != 0 {
		item := s.Items[len(s.Items)-1]
		s.Items = s.Items[:len(s.Items)-1]
		return item
	}
	return nil
}

func (s *Stack) peek() any {
	if len(s.Items) == 0 {
		return nil
	}
	return s.Items[len(s.Items)-1]
}

func (s *Stack) isEmpity() bool {
	if len(s.Items) != 0 {
		return false
	}
	return true
}
