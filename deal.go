package per

import (
	"regexp"
	"strings"
)

//这个文件的目的就是为了处理get获得的各种资源，使用正则即可。
func dealWith(body string, t int) string {
	body = simpleDealwith(body)
	body = plusDealwith(t, body)
	return body
}

// 简单的处理
func simpleDealwith(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n\n")
	return strings.TrimSpace(src)
}

// // 更加复杂的处理
func plusDealwith(s int, src string) string {
	if s == 0 {
		re, _ := regexp.Compile("[[:alpha:]]*")
		src = re.ReplaceAllString(src, "")
		// 去除连续两个或者大于两个的换行符
		re, _ = regexp.Compile("\\s{2,}")
		src = re.ReplaceAllString(src, "")
		return strings.TrimSpace(src)
	} else {
		return src
	}

}
