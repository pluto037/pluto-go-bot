package dispatch

import (
	. "github.com/pluto037/pluto-go-bot/bussiness/message"
	"github.com/pluto037/pluto-go-bot/module/botmodule/aichat"
	logger "github.com/pluto037/pluto-go-logger"
	"github.com/pluto037/pluto-go/system"
	"strings"
	"sync"
)

var (
	ChattingUsers = make(map[string]bool)
	ChatUser      []string
	mu            sync.Mutex // 用于保护 chattingUsers 映射的互斥锁
)

func CommandDispatch(user string, message_ string, msgId string, targetId string) []byte {
	trimmed := strings.TrimSpace(message_)
	message := strings.TrimSpace(strings.TrimPrefix(trimmed, "@布鲁托"))
	mu.Lock()
	defer mu.Unlock()
	switch message {
	case "/weather":
		return []byte("")
	case "/status":
		return ResultPackage(targetId, system.Query(), msgId)
	case "/chat":
		ChattingUsers[user] = true
		logger.Info("CommandDispatch,user add in ai chat list", 0, "")
		return ResultPackage(targetId, "与AI聊天模式已启动，发送消息将自动转发给AI。", msgId)
	case "/exit":
		delete(ChattingUsers, user)
		logger.Info("CommandDispatch,user delete in ai chat list", 0, "")
		return ResultPackage(targetId, "与AI聊天模式已关闭，进入机器人模式。", msgId)
	case "/help":
		return ResultPackage(targetId, "/help 查看命令，/status 查看服务器运行状态，/chat 进入AI聊天模式，/exit 退出AI聊天模式", msgId)
	}
	if ChattingUsers[user] {
		return ResultPackage(targetId, aichat.Chat(message), msgId)
	}
	return ResultPackage(targetId, "命令不正确哈QAQ~", msgId)
}
