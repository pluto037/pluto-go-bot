package message

import (
	"fmt"
	"github.com/goccy/go-json"
)

type ResultMessage struct {
	Type         int    `json:"type"`
	TargetID     string `json:"target_id"`
	Content      string `json:"content"`
	Quote        string `json:"quote,omitempty"`
	Nonce        string `json:"nonce,omitempty"`
	TempTargetID string `json:"temp_target_id,omitempty"`
}

func ResultPackage(targetId string, content string, quote string) []byte {
	message := ResultMessage{
		Type:     1,
		TargetID: targetId,
		Content:  content,
		Quote:    quote,
		//Nonce:        nonce,
		//TempTargetID: tempTargetId,
	}
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		fmt.Println("JSON转换出错:", err)
		return nil
	}
	return jsonMessage
}
