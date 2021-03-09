package intf

type Intf1 interface {
	Method1V2(string, bool, int) error
	Method2V2(num int)
}

type Intf2 interface {
	Method1V1(str string) (int, error)
	Method2V1(char rune) string
}
