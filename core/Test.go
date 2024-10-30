package core

type TestStruct struct {
	Name string
}

func (t TestStruct) Test() string {
	return t.Name
}
