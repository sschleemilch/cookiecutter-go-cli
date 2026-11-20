package config

type Logging struct {
	Level  string `json:"level"`
	Caller bool   `json:"caller"`
}

type Config struct {
	Log Logging `json:"log"`
}

func Defaults() Config {
	return Config{
		Log: Logging{
			Level:  "info",
			Caller: false,
		},
	}
}
