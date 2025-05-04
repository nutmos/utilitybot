package state

import (
	"fmt"
)

type State struct {
	User           *string                `json:"user"`
	Platform       *string                `json:"platform"`
	CurrentCommand *string                `json:"current_command"`
	UserData       map[string]interface{} `json:"user_data"`
}

var stateDatabase map[string]*State

func init() {
	stateDatabase = map[string]*State{}
}

func buildStateKey(user string, platform string) string {
	return fmt.Sprintf("%s:%s", platform, user)
}

func GetUserState(user string, platform string) *State {
	stateKey := buildStateKey(user, platform)
	userState, ok := stateDatabase[stateKey]
	if !ok {
		return nil
	}
	return userState
}

func DelUserState(user string, platform string) error {
	stateKey := buildStateKey(user, platform)
	delete(stateDatabase, stateKey)
	return nil
}

func SetUserState(user string, platform string, currentCommand string, userData map[string]interface{}) error {
	stateKey := buildStateKey(user, platform)
	newState := &State{
		User:           &user,
		Platform:       &platform,
		CurrentCommand: &currentCommand,
		UserData:       userData,
	}
	stateDatabase[stateKey] = newState
	return nil
}
