package notify

import "testing"

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("config.yaml")
	if err != nil {
		t.Errorf("LoadConfig() error = %v", err)
		return
	}

	if config.PushPlus.Token == "" {
		t.Errorf("config.PushPlus.Token is empty")
		return
	}
}
