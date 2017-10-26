# API

 ``` go
 // 输出类型是Md
 type MD{
    s string
 }
 ```

 ```go
  // 输出类型是json
    type Json {
        s string
    }
```


 ```go
  // 输出类型是文本
    type Plain {
        s string
    }
```

```go
//输入你要爬取的URL
typ是一个interface，内部已经通过几个不同的struct类型实现了这个interface，所以在这个位置上你只需要传入一个struct类型即可。
func Run (url, ty string, t typ) error

```
