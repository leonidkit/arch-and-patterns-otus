package hw5

import "sync"

type scopeStorage struct {
	mx      *sync.RWMutex
	storage map[string]map[string]func(...any) any
}

func newScopeStorage() *scopeStorage {
	return &scopeStorage{
		mx:      &sync.RWMutex{},
		storage: make(map[string]map[string]func(...any) any),
	}
}

func (s *scopeStorage) set(scopeName, key string, f func(...any) any) {
	s.mx.Lock()
	defer s.mx.Unlock()

	if _, ok := s.storage[scopeName]; !ok {
		s.storage[scopeName] = make(map[string]func(...any) any)
	}
	s.storage[scopeName][key] = f
}

func (s *scopeStorage) get(scopeName, key string) func(...any) any {
	s.mx.RLock()
	defer s.mx.RUnlock()

	if s.storage[scopeName] == nil {
		return nil
	}
	return s.storage[scopeName][key]
}
