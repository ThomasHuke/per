package main

import (
	"fmt"

	"github.com/github.com/thomashuke/per"
)

type name struct {
	text per.Plain
	md   per.Md
	json per.Json
}

func main() {
	fmt.Println("hey i am using the great crawler which is named per and  created by ThomasHuke")
	per.Run()
}
