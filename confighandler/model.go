package confighandler

type ApiKey struct {
	Aviationstack string `mapstructure:"aviationstack"`
	Telegram      string `mapstructure:"telegram"`
}

type ConfigStruct struct {
	ApiKey         ApiKey `mapstructure:"apiKey"`
	DeploymentMode string `mapstructure:"deploymentMode"`
}

var (
	Config ConfigStruct
)

const (
	DeploymentModePull   string = "Pull"
	DeploymentModeLambda string = "Lambda"
)
