package notify

import "github.com/helzoph/notify/pushplus"

func NewPushPlusBot(config *Config) *pushplus.Bot {
	return &pushplus.Bot{
		Token: config.PushPlus.Token,
	}
}
