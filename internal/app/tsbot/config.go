package tsbot

type Config struct {
	ServerAddress string `toml:"server_address"`
	ServerPort    string `toml:"server_port"`
	QueryLogin    string `toml:"query_login"`
	QueryPassword string `toml:"query_password"`
}

func NewConfig() *Config {
	return &Config{
		ServerAddress: "localhost",
		ServerPort:    "10011",
		QueryLogin:    "Yol",
		QueryPassword: "UkL9s9vE",
	}
}
