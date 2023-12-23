package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/pluto037/pluto-go-bot/bussiness/dispatch"
	"github.com/pluto037/pluto-go-bot/config/kook"
	"github.com/pluto037/pluto-go-bot/module/basemodule"
	"log"
	"strconv"
	"sync"
	"time"
)

type WebSocket struct{}

var (
	Conn           *websocket.Conn
	receiveChannel = make(chan []byte) // 用于接收消息的通道
	wg             sync.WaitGroup
)

func ConnectServer() {
	if kook.KOOK_WSS_URL == "" {
		basemodule.KookConnInit()
	}
	c, _, err := websocket.DefaultDialer.Dial(kook.KOOK_WSS_URL, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket:", err)
		return
	}
	Conn = c
	wg.Add(2)
	go func() {
		defer wg.Done()
		WsReceive()
	}()
	go func() {
		defer wg.Done()
		StartPingLoop()
	}()
	wg.Wait()
}

func WsReceive() {
	for {
		_, p, err := Conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}
		log.Printf("Received message: %s", p)

		// 将消息发送到通道
		receiveChannel <- p
	}
}

func StartPingLoop() {
	pingTicker := time.NewTicker(30 * time.Second) // 30秒定时器
	defer pingTicker.Stop()

	for {
		select {
		case <-pingTicker.C:
			sn := GetMaxSN()
			pingMessage := CreatePingMessage(sn)
			err := Conn.WriteMessage(websocket.PingMessage, pingMessage)
			if err != nil {
				log.Println("Error sending Ping:", err)
			}
		case <-time.After(6 * time.Second):
			// 在规定时间内等待 Pong

		case p := <-receiveChannel:
			// 接收到消息，继续循环
			dispatch.Process(p)
		}
	}
}

// CreatePingMessage 创建 Ping 消息
func CreatePingMessage(sn int) []byte {
	// 创建 Ping 消息的逻辑，使用传入的 sn
	// 示例中使用 JSON 格式的消息
	// 你需要根据实际情况自定义
	return []byte(`{"s": 2, "sn": ` + strconv.Itoa(sn) + `}`)
}

// GetMaxSN 获取当前的最大 sn 的逻辑，你需要根据实际情况实现
func GetMaxSN() int {
	// 获取最大 sn 的逻辑，示例中直接返回一个随机数
	return kook.SN
}
