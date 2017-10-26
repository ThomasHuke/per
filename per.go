package per

import (
	"fmt"
	"log"
	"net/http"
)

//Run 作为对外暴露的主函数。这里t是一个interface的实例当然也就是一个struct，所以我们只要在输入数据的时候
//输入Type Plain Json Md 类型的 struct即可。
func Run(url, ty string, t typ) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error! Here are some messages: ")
		log.Fatalln(err)
	}
	fmt.Print("The output is :", *res)

	var body string = *res
	message := dealWith(body)
	//所以只需要实现一个
	err = t.put(message)
	if err != nil {
		fmt.Println("Error! Here are some messages: ")
		log.Fatal(err)
	}

}
