package module

import (
	"fmt"
	chatAi "github.com/pluto037/pluto-go-bot/module/chatAi"

	"sync"
)

var (
	chattingUsers = make(map[string]bool)
	mu            sync.Mutex // 用于保护 chattingUsers 映射的互斥锁
)

func CommandDispatch(user string, message string) string {
	mu.Lock()
	defer mu.Unlock()
	if chattingUsers[user] {
		return fmt.Sprintln(chatAi.Chat(message))
	}
	switch message {
	case "/chat":
		chattingUsers[user] = true
		return "与AI聊天模式已启动，发送消息将自动转发给AI。"
	case "/exit":
		chattingUsers[user] = false
		return "与AI聊天模式已关闭，进入机器人模式。"
	}
	return ""
}
