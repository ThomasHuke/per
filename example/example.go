package main

import (
	p "github.com/thomashuke/per"
)

func main() {
	url := "http://zhihu.com"
	ty := p.Md{"outPut"}
	p.Run(url, ty)
}
