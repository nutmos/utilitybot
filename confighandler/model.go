package confighandler

type ApiKey struct {
	Aviationstack string `yaml:"aviationstack"`
	Telegram      string `yaml:"telegram"`
}

type ConfigStruct struct {
	ApiKey ApiKey `yaml:"apiKey"`
}

var (
	Config ConfigStruct
)
