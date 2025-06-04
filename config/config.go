package config

type Config struct {
	APIKey   string
	ApiToken string
	Debug    bool // 是否启用调试模式
	Sandbox  bool // 是否为沙箱环境
	Timeout  int  // HTTP 超时设定（单位：秒）
}

func LoadConfig() *Config {
	return &Config{
		APIKey:   "100000",                           // "99000015",
		ApiToken: "D6140AA383FD8515B09028C586493DDB", // "E33A3973221DB08128F8FF436EFDB8F4",
		Debug:    true,
		Sandbox:  true,
		Timeout:  360,
	}
}
