Go字符串处理小工具

```bash
go get -u github.com/gotail/strx
```

## 背景

近期处理一些字符串相关的工作，增加一个工具用于链式调用，节约开发时间与代码的可读性。
项目中积累的一些字符串工具

具体参见 str_test.go


## 演示

你可以像这样处理字符串
```go
var newStr string = New(" 这*里|||有人民yyyy币         $yy#{money}   ..").
		NoConsecutiveSpaces().
		Trim(" ").
		Clean("*", "|", "yy").
		Replace("$", "￥").
		Replace("#{money}", "250").
		TrimRight(".").
		Append("元").
		Val()
	fmt.Println(newStr)
```

若你对这个小工具感兴趣，欢迎加入。

## LICENSE

anti-996 LICENSE
