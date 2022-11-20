package hw5

import "context"

const (
	IoCRegister = "IoC.Register"
)

type scopeNameCtxKey struct{}

//go:generate mockgen -source=hw5.go -destination=mock.go -package=${GOPACKAGE}
type ICommand interface {
	Execute() error
}

type IoC struct {
	scopeStorage *scopeStorage
}

func NewIoC() *IoC {
	return &IoC{
		scopeStorage: newScopeStorage(),
	}
}

func (ioc *IoC) Resolve(ctx context.Context, key string, args ...any) any {
	scopeName := extractScopeName(ctx)
	if key == IoCRegister {
		return newRegisterCommand(
			scopeName,
			args[0].(string),
			args[1].(func(...any) any),
			ioc.scopeStorage,
		)
	}
	f := ioc.scopeStorage.get(scopeName, key)
	if f == nil {
		return nil
	}
	return f(args...)
}

func extractScopeName(ctx context.Context) string {
	v := ctx.Value(scopeNameCtxKey{})
	if v == nil {
		panic("scope not set")
	}
	return v.(string)
}

type registerCommand struct {
	scopeName    string
	key          string
	f            func(...any) any
	scopeStorage *scopeStorage
}

func newRegisterCommand(scopeName, key string, f func(...any) any, scopeStorage *scopeStorage) *registerCommand {
	return &registerCommand{
		scopeName:    scopeName,
		key:          key,
		f:            f,
		scopeStorage: scopeStorage,
	}
}

func (rc *registerCommand) Execute() error {
	rc.scopeStorage.set(rc.scopeName, rc.key, rc.f)
	return nil
}
