package userstatehandler

type UserStateHandler interface {
	GetUserState(user string, platform string) *UserState
	DelUserState(user string, platform string) error
	SetUserState(user string, platform string, currentCommand string, userData map[string]any) error
}

type UserState struct {
	User           *string        `json:"user"`
	Platform       *string        `json:"platform"`
	CurrentCommand *string        `json:"current_command"`
	UserData       map[string]any `json:"user_data"`
}
