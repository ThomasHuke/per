# API

- [type Md struct](#md)
- [type Json struct](#json)
- [type Plain struct](#plain)
- [func Run(url, ty string, t typ) error](#run)
#### md

``` go
 // 输出类型是Md
 type MD{
    s string
 }
```
#### json

```go
  // 输出类型是json
    type Json {
        s string
    }
```

#### plain

 ```go
  // 输出类型是文本
    type Plain {
        s string
    }
```

**typ是一个interface，内部已经通过几个不同的struct类型实现了这个interface，所以在这个位置上你只需要传入一个struct类型即可**

#### run

```go
//输入你要爬取的URL
func Run (url string, ty p.typ, t int)

```
这个最后的t是判断输出是否要删除英文，是的话t != 0 如何想删除英语请让 t == 0
