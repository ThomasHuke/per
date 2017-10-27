package per

import (
	"fmt"
	"os"
)

// typ 一个interface 目的是为了解决不同的输出。
type typ interface {
	//
	put(m string) error
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
func (md Md) put(m string) {
	s := "~/Desktop" + md.s + ".md"
	file, err := os.Create(s)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString(m)
	if err != nil {
		fmt.Println(err)
	}

}

// 文本类型实现接口
func (p Plain) put(m string) {
	s := "~/Desktop" + p.s + ".text"
	file, err := os.Create(s)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString(m)
	if err != nil {
		fmt.Println(err)
	}

}

// json类型实现接口
func (j Json) put(m string) {
	s := "~/Desktop" + j.s + ".json"
	file, err := os.Create(s)
	if err != nil {
		fmt.Println(err)
	}
	m = "{" + "message" + ":" + "\"" + m + "\"" + "}"
	_, err = file.WriteString(m)
	if err != nil {
		fmt.Println(nil)
	}
}
