package confighandler

type ApiKey struct {
	Aviationstack string `mapstructure:"aviationstack"`
	Telegram      string `mapstructure:"telegram"`
}

type ConfigStruct struct {
	ApiKey ApiKey `mapstructure:"apiKey"`
}

var (
	Config ConfigStruct
)
