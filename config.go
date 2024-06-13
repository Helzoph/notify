package notify

type Config struct {
	PushPlus PushPlusConfig `yaml:"pushplus"`
}

type PushPlusConfig struct {
	Token string `json:"token" yaml:"token"`
}
