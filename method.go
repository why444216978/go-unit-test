package main

type S struct{}

func (s *S) A() int {
	return s.B() + s.b()
}

func (s *S) AA() int {
	return s.B() + s.b() + s.B() + s.b()
}

func (s *S) B() int {
	return 0
}

func (s *S) b() int {
	return 0
}
