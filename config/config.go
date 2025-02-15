package config

type Config struct {
	APIKey   string
	Apitoken string
	Debug    bool // 是否启用调试模式
	Sandbox  bool // 是否为沙箱环境
	Timeout  int  // HTTP 超时设定（单位：秒）
}

func LoadConfig() *Config {
	return &Config{
		APIKey:   "99000015",
		Apitoken: "E33A3973221DB08128F8FF436EFDB8F4",
		Debug:    true,
		Sandbox:  true,
		Timeout:  360,
	}
}
