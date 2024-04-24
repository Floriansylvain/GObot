package config

var (
	Token string
	BotPrefix string

	config *Config
)

type Config struct {
	TOKEN string `json:"TOKEN"`
	BOT_PREFIX string `json:"BOT_PREFIX"`
}

func LoadConfig() error {
	fmt.printLn("Loading config file")
}
