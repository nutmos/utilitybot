package userstatehandler

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

type InMemoryStateHandler struct {
	cache *cache.Cache
}

func initInMemoryUserStateHandler() *InMemoryStateHandler {
	return &InMemoryStateHandler{
		cache: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (m *InMemoryStateHandler) buildStateKey(user string, platform string) string {
	return fmt.Sprintf("%s:%s", platform, user)
}

func (m *InMemoryStateHandler) GetUserState(user string, platform string) *UserState {
	stateKey := m.buildStateKey(user, platform)
	userState, found := m.cache.Get(stateKey)
	if !found {
		return nil
	}
	return userState.(*UserState)
}

func (m *InMemoryStateHandler) DelUserState(user string, platform string) error {
	stateKey := m.buildStateKey(user, platform)
	m.cache.Delete(stateKey)
	return nil
}

func (m *InMemoryStateHandler) SetUserState(user string, platform string, currentCommand string, userData map[string]any) error {
	stateKey := m.buildStateKey(user, platform)
	newState := &UserState{
		User:           &user,
		Platform:       &platform,
		CurrentCommand: &currentCommand,
		UserData:       userData,
	}
	m.cache.Set(stateKey, newState, cache.DefaultExpiration)
	return nil
}
