package per

// typ 一个interface 目的是为了解决不同的输出。
type typ interface {
	//
	put(s string) error
}

//输出是md的类型
type Md struct {
	s string
}

//输出是文本类型
type Plain struct {
	s string
}

//输出是json类型
type Json struct {
	s string
}

//md 类型实现interface
func (m Md) put(s string) error {

}

// 文本类型实现接口
func (p Plain) put(s string) error {

}

// json类型实现接口
func (j Json) put(s string) error {

}
