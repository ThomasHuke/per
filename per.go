package per

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//RunALL批量处理数据
func RunALL(url []string, t []Typ) {
	for _, i := range url {
		for _, z := range t {
			Run(i, z)
		}

	}

}

func Run(url string, t Typ) {
	ace := tt(url)
	// dealWith处理接收到的信息。
	message := dealWith(ace)
	//然后将信息传递给put方法，这样put就可以生成一个拥有信息的一个文件，并且这个文件默认是放置到桌面的
	//后者是说目前我强制它放置到桌面了。
	t.put(message)
}

//Run 作为对外暴露的主函数。这里t是一个interface的实例当然也就是一个struct，所以我们只要在输入数据的时候
//输入Type Plain Json Md 类型的 struct即可。
func tt(url string) string {

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("错误出现在主函数获取资源的时候 ")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("错误出现在资源转换的时候")
	}

	return string(body)
}
