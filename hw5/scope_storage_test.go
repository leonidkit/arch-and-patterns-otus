package hw5

import "testing"

func Test_scopeStorage_set(t *testing.T) {
	tests := []struct {
		name      string
		scopeName string
		key       string
		f         func(...any) any
	}{
		{
			scopeName: "scope1",
			key:       "move",
			f: func(a ...any) any {
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scopeStorage := newScopeStorage()

			scopeStorage.set(tt.scopeName, tt.key, tt.f)
			assertEqual(len(scopeStorage.storage), 1)
			assertEqual(len(scopeStorage.storage[tt.scopeName]), 1)
		})
	}

}

func Test_scopeStorage_get(t *testing.T) {
	tests := []struct {
		name      string
		scopeName string
		key       string
		f         func(...any) any
	}{
		{
			scopeName: "scope1",
			key:       "move",
			f: func(a ...any) any {
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scopeStorage := newScopeStorage()

			scopeStorage.set(tt.scopeName, tt.key, tt.f)

			res := scopeStorage.get(tt.scopeName, tt.key)
			assertNotNil(res)
		})
	}
}
