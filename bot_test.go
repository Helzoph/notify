package notify

import (
	"testing"

	"github.com/helzoph/notify/pushplus"
)

func TestSend(t *testing.T) {
	config, err := LoadConfig("config.yaml")
	if err != nil {
		t.Errorf("LoadConfig() error = %v", err)
		return
	}

	bot := NewPushPlusBot(config)
	if bot.Token == "" {
		t.Errorf("NewPushPlusBot() error = %v", bot.Token)
		return
	}

	_, err = bot.Send(&pushplus.Message{
		Content: "Hello, World!",
	})
	if err != nil {
		t.Errorf("Send() error = %v", err)
		return
	}
}
