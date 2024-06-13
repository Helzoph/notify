package notify

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(filename string) (config *Config, err error) {
	// 读取配置文件
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Load:os.ReadFile-> %w", err)
	}

	// 解析配置文件
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("Load:yaml.Unmarshal-> %w", err)
	}

	return
}
