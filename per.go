package per

import (
	"fmt"
	"log"
	"net/http"
)

func Run(url []string, ty string, t typ, d chan *http.Response) {
	c := make(chan string)
	go tt(url, ty, t, d)
	for i := range c {
		// dealWith处理接收到的信息。
		message := dealWith(i)
		//然后将信息传递给put方法，这样put就可以生成一个拥有信息的一个文件，并且这个文件默认是放置到桌面的
		//后者是说目前我强制它放置到桌面了。
		t.put(message)
	}
}

//Run 作为对外暴露的主函数。这里t是一个interface的实例当然也就是一个struct，所以我们只要在输入数据的时候
//输入Type Plain Json Md 类型的 struct即可。
func tt(url []string, ty string, t typ, d chan *http.Response) {
	for _, value := range url {
		res, err := http.Get(value)
		defer res.Body.Close()
		if err != nil {
			fmt.Println("Error! Here are some messages: ")
			log.Fatalln(err)
		}
		fmt.Print("The output is :", *res)
		d <- res
	}
	close(d)
}
