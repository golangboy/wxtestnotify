# 这是什么
这是使用微信测试号用来下发通知用的，用于个人开发者实现微信通知，具体使用方法请参考。

# 如何使用
在[这里](http://mp.weixin.qq.com/debug/cgi-bin/sandboxinfo?action=showinfo&t=sandbox/index)注册好，添加一个如下的模版。
![模版](./imgs/1.png)

```go
package main

import "github.com/golangboy/wxtestnotify"

func main() {
	notify, err := wxtestnotify.WxNotify("wxxxxxxx", "477yyyyyyy", "userOpenID", "templateID", "测试通知")
	if err != nil {
		return
	}
	println(notify)
}

```