package main

import (
	"fmt"
	"github.com/pluto037/pluto-go/system"
)

var (
	hostUrl   = "wss://spark-api.xf-yun.com/v2.1/chat"
	hostUrlv1 = "wss://spark-api.xf-yun.com/v1.1/chat"
	appid     = "5a98fc41"
	apiSecret = "MTBhNmJhODlkOTVhZjAxNTQyNGE1ZmZk"
	apiKey    = "22e2b5b6c733c95fd5624638a7ef08e0"
	domianv1  = "general"
	domianv2  = "generalv2"
)

func main() {
	//chat := kdxf.StartChat(hostUrlv1, appid, apiSecret, apiKey, "西安天气如何", "1", domianv1)
	//fmt.Println(chat)
	fmt.Println(system.Query())
}
