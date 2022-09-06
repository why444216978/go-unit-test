package main

func A() int {
	return B()
}

func AA() int {
	return B() + B()
}

func B() int {
	return 0
}
