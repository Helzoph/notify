package pushplus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Bot struct {
	Token string
}

func (b *Bot) Send(m *Message) (response *Response, err error) {
	url := "https://www.pushplus.plus/send"
	method := "POST"

	// check message content
	if err = b.checkMessage(m); err != nil {
		return nil, fmt.Errorf("Send:Check message-> %w", err)
	}

	// marshal message to json
	jsonData, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("Send:Marshal message to json-> %w", err)
	}

	// add token to json
	jsonData = bytes.Replace(jsonData, []byte("{"), []byte(`{"token":"`+b.Token+`",`), 1)

	// send request
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("Send:Set new request-> %w", err)
	}

	// set request headers
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Content-Type", "application/json")

	// send request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Send:Send request-> %w", err)
	}
	defer resp.Body.Close()

	// decode response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, fmt.Errorf("Send:Decode response-> %w", err)
	}
	if response.Code != 200 {
		return response, fmt.Errorf("Send:Response error-> %s", response.Msg)
	}

	return response, nil
}

// checkMessage checks the content of the given message.
// It returns an error if the content is empty.
func (b *Bot) checkMessage(m *Message) (err error) {
	if m.Content == "" {
		return fmt.Errorf("checkMessage-> content is empty")
	}

	return
}
