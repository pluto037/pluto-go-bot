package dispatch

import (
	"encoding/json"
	"fmt"
	"github.com/pluto037/pluto-go-bot/config/kook"
	"github.com/pluto037/pluto-go-bot/tools"
)

type Message struct {
	S int `json:"s"`
	D struct {
		ChannelType string `json:"channel_type"`
		Type        int    `json:"type"`
		TargetID    string `json:"target_id"`
		AuthorID    string `json:"author_id"`
		Content     string `json:"content"`
		Extra       struct {
			Type        int    `json:"type"`
			Code        string `json:"code"`
			GuildID     string `json:"guild_id"`
			GuildType   int    `json:"guild_type"`
			ChannelName string `json:"channel_name"`
			Author      struct {
				ID               string      `json:"id"`
				Username         string      `json:"username"`
				IdentifyNum      string      `json:"identify_num"`
				Online           bool        `json:"online"`
				OS               string      `json:"os"`
				Status           int         `json:"status"`
				Avatar           string      `json:"avatar"`
				VipAvatar        string      `json:"vip_avatar"`
				Banner           string      `json:"banner"`
				Nickname         string      `json:"nickname"`
				Roles            []int       `json:"roles"`
				IsVip            bool        `json:"is_vip"`
				VipAmp           bool        `json:"vip_amp"`
				IsAiReduce       bool        `json:"is_ai_reduce_noise"`
				IsPersonalBg     bool        `json:"is_personal_card_bg"`
				Bot              bool        `json:"bot"`
				DecorationsIDMap interface{} `json:"decorations_id_map"`
				IsSys            bool        `json:"is_sys"`
			} `json:"author"`
			VisibleOnly  interface{}   `json:"visible_only"`
			Mention      []interface{} `json:"mention"`
			MentionAll   bool          `json:"mention_all"`
			MentionRoles []int         `json:"mention_roles"`
			MentionHere  bool          `json:"mention_here"`
			NavChannels  []interface{} `json:"nav_channels"`
			KMarkdown    struct {
				RawContent      string        `json:"raw_content"`
				MentionPart     []interface{} `json:"mention_part"`
				MentionRolePart []struct {
					RoleID    int           `json:"role_id"`
					Name      string        `json:"name"`
					Color     int           `json:"color"`
					ColorType int           `json:"color_type"`
					ColorMap  []interface{} `json:"color_map"`
				} `json:"mention_role_part"`
				ChannelPart []interface{} `json:"channel_part"`
			} `json:"kmarkdown"`
			Emoji          []interface{} `json:"emoji"`
			LastMsgContent string        `json:"last_msg_content"`
			SendMsgDevice  int           `json:"send_msg_device"`
		} `json:"extra"`
		MsgID        string `json:"msg_id"`
		MsgTimestamp int64  `json:"msg_timestamp"`
		Nonce        string `json:"nonce"`
		FromType     int    `json:"from_type"`
	} `json:"d"`
	Extra struct {
		VerifyToken string `json:"verifyToken"`
		EncryptKey  string `json:"encryptKey"`
		CallbackUrl string `json:"callbackUrl"`
	} `json:"extra"`
	Sn int `json:"sn"`
}

func Process(text []byte) {
	var message Message
	kook.SN = message.Sn
	err := json.Unmarshal(text, &message)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	mentionRoles := message.D.Extra.Mention
	for _, item := range mentionRoles {
		if item == "1099668604" {
			dispatch := CommandDispatch(message.D.Extra.Author.ID, message.D.Extra.KMarkdown.RawContent, message.D.MsgID, message.D.TargetID)
			client := tools.NewHTTPClient()
			result, err := client.SendMessage(dispatch)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(result)
			return
		}
	}
	if ChattingUsers[message.D.Extra.Author.ID] {
		dispatch := CommandDispatch(message.D.Extra.Author.ID, message.D.Content, message.D.MsgID, message.D.TargetID)
		client := tools.NewHTTPClient()
		result, err := client.SendMessage(dispatch)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	}

}
