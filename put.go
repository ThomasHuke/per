package per

import (
	"fmt"
	"os"
)

// typ 一个interface 目的是为了解决不同的输出。
type typ interface {
	// 文件的输出
	put(m string)
}

//输出是md的类型
type Md struct {
	//文件名
	S string
}

//输出是文本类型
type Plain struct {
	S string
}

//输出是json类型
type JSON struct {
	S string
}

//md 类型实现interface m 是信息。 md.S 是文件名
var file *os.File

func (md Md) put(m string) {
	s := md.S + ".md"
	file, err := os.Create(s)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString(m)
	if err != nil {
		fmt.Println("错误出现在写入区域")
		fmt.Println(err)
	}

}

// 文本类型实现接口
func (p Plain) put(m string) {
	s := p.S + ".text"
	file, err := os.Create(s)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString(m)
	if err != nil {
		fmt.Println("错误出现在写入区域")
		fmt.Println(err)
	}

}

// json类型实现接口
func (j JSON) put(m string) {
	s := j.S + ".json"
	file, err := os.Create(s)
	if err != nil {
		fmt.Println(err)
	}

	m = "{" + "\"message\"" + ":" + "\"" + m + "\"" + "}"
	_, err = file.WriteString(m)
	if err != nil {
		fmt.Println("错误出现在写入区域")
		fmt.Println(nil)
	}
}
