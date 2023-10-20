package chatAi

import "github.com/pluto037/pluto-go/kdxf"

var (
	hostUrl   = "wss://spark-api.xf-yun.com/v2.1/chat"
	hostUrlv1 = "wss://spark-api.xf-yun.com/v1.1/chat"
	appid     = "5a98fc41"
	apiSecret = "MTBhNmJhODlkOTVhZjAxNTQyNGE1ZmZk"
	apiKey    = "22e2b5b6c733c95fd5624638a7ef08e0"
	domianv1  = "general"
	domianv2  = "generalv2"
)

func init() {

}

func Chat(question string) string {
	return kdxf.StartChat(hostUrlv1, appid, apiSecret, apiKey, question, "11122333", domianv1)
}
