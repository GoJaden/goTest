package main

func main() {
	sVals := map[int]S{1: {"A"}}

	// 你只能通过值调用 Read
	sVals[1].Read()

	// 这不能编译通过：

	//sVals[1].Write("test")

	sPtrs := map[int]*S{1: {"A"}}

	// 通过指针既可以调用 Read，也可以调用 Write 方法
	sPtrs[1].Read()
	sPtrs[1].Write("test")
}

type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}

func (s *S) Write(str string) {
	s.data = str

}
