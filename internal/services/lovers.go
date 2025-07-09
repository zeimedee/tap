package services

import (
	"errors"
	"strings"
	"sync"
)

type LoversService struct {
	ids   map[string]string
	mutex sync.RWMutex
}

func NewLoversService() *LoversService {
	return &LoversService{
		ids:   make(map[string]string),
		mutex: sync.RWMutex{},
	}
}

func (ws *LoversService) StoreWord(id, token string) {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()

	lowerCaseId := strings.ToLower(id)

	ws.ids[lowerCaseId] = token
}

func (ws *LoversService) GetToken(id string) (string, error) {
	token, ok := ws.ids[id]
	if !ok {
		return "", errors.New("token not found")
	}
	return token, nil
}

func (ws *LoversService) GetAll() map[string]string {
	return ws.ids
}
