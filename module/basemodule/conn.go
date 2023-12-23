package basemodule

import (
	"encoding/json"
	"fmt"
	"github.com/pluto037/pluto-go-bot/config/kook"
	logger "github.com/pluto037/pluto-go-logger"
	"net/http"
)

func KookConnInit() {
	client := &http.Client{}

	request, err := http.NewRequest("GET", kook.SOCKET_URL_QUERT, nil)
	if err != nil {
		logger.Warning("http request build fail", 1, err)
		return
	}
	request.Header.Add("Authorization", kook.TOKEN)
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println(data)
	info := data["data"].(map[string]interface{})
	kook.KOOK_WSS_URL = info["url"].(string)
	logger.Info("kook websocket url="+kook.KOOK_WSS_URL, 0, "")
}
