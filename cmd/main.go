package main

import (
	"github.com/pluto037/pluto-go-bot/module/basemodule"
	"github.com/pluto037/pluto-go-bot/module/websocket"
)

func main() {
	//basemodule.WebInit()
	basemodule.KookConnInit()
	websocket.ConnectServer()
}
