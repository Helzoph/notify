package pushplus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Bot struct {
	Token string // 用户令牌，可直接加到请求地址后
}

func (b *Bot) Send(m *Message) (err error) {
	url := "https://www.pushplus.plus/send"
	method := "POST"

	// 检查消息内容
	if err = b.checkMessage(m); err != nil {
		return fmt.Errorf("Send:Check message-> %w", err)
	}

	// 序列化消息内容
	jsonData, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("Send:Marshal message to json-> %w", err)
	}

	// 添加用户令牌到json数据中
	jsonData = bytes.Replace(jsonData, []byte("{"), []byte(`{"token":"`+b.Token+`",`), 1)

	// 创建请求
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("Send:Set new request-> %w", err)
	}

	// 设置请求头
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Content-Type", "application/json")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Send:Send request-> %w", err)
	}
	defer resp.Body.Close()

	// 解析响应内容
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return fmt.Errorf("Send:Decode response-> %w", err)
	}
	if response.Code != 200 {
		return fmt.Errorf("Send:Response error-> %s", response.Msg)
	}

	return
}

// 检查消息内容
func (b *Bot) checkMessage(m *Message) (err error) {
	if m.Content == "" {
		return fmt.Errorf("checkMessage-> content is empty")
	}

	return
}
