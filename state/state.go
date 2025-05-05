package state

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

type State struct {
	User           *string                `json:"user"`
	Platform       *string                `json:"platform"`
	CurrentCommand *string                `json:"current_command"`
	UserData       map[string]interface{} `json:"user_data"`
}

var stateCache *cache.Cache

func init() {
	stateCache = cache.New(5*time.Minute, 10*time.Minute)
}

func buildStateKey(user string, platform string) string {
	return fmt.Sprintf("%s:%s", platform, user)
}

func GetUserState(user string, platform string) *State {
	stateKey := buildStateKey(user, platform)
	userState, found := stateCache.Get(stateKey)
	if !found {
		return nil
	}
	return userState.(*State)
}

func DelUserState(user string, platform string) error {
	stateKey := buildStateKey(user, platform)
	stateCache.Delete(stateKey)
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
	stateCache.Set(stateKey, newState, cache.DefaultExpiration)
	return nil
}
