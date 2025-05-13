package userstatehandler

type UserStateHandlerType int

const (
	UserStateHandlerTypeInMemory UserStateHandlerType = iota
)

var userStateHandler UserStateHandler

func InitUserStateHandler(t UserStateHandlerType) {
	if t == UserStateHandlerTypeInMemory {
		userStateHandler = initInMemoryUserStateHandler()
	}
}

func GetUserState(user string, platform string) *UserState {
	return userStateHandler.GetUserState(user, platform)
}

func DelUserState(user string, platform string) error {
	return userStateHandler.DelUserState(user, platform)
}

func SetUserState(user string, platform string, currentCommand string, userData map[string]any) error {
	return userStateHandler.SetUserState(user, platform, currentCommand, userData)
}
