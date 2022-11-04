package hw5

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

// Динамическая подмена scope реализована через каналы.
func TestIoC_Resolve_Parallel(t *testing.T) {
	scopeNames := []string{"scope1", "scope2"}
	goroutinesNum := 2
	ctxCh := make(chan context.Context, 1)
	ioc := NewIoC()
	ctrl := gomock.NewController(t)

	ctx := setScopeName(context.Background(), scopeNames[0])
	cmd1 := NewMockICommand(ctrl)
	cmd1.EXPECT().Execute().MinTimes(1)
	cmd := ioc.Resolve(ctx, IoCRegister, "cmd", func(args ...any) any {
		return cmd1
	}).(ICommand)
	_ = cmd.Execute()

	ctx = setScopeName(ctx, scopeNames[1])
	cmd2 := NewMockICommand(ctrl)
	cmd2.EXPECT().Execute().MinTimes(1)
	cmd = ioc.Resolve(ctx, IoCRegister, "cmd", func(args ...any) any {
		return cmd2
	}).(ICommand)
	_ = cmd.Execute()

	for i := 0; i < goroutinesNum; i++ {
		go func(ctx context.Context) {
			for {
				select {
				case newCtx, ok := <-ctxCh:
					if !ok {
						return
					}
					ctx = newCtx
				default:
					cmd, ok := ioc.Resolve(ctx, "cmd").(ICommand)
					assertEqual(ok, true)

					err := cmd.Execute()
					assertNoError(err)
				}
			}
		}(ctx)
	}

	time.Sleep(100 * time.Millisecond)
	ctxCh <- setScopeName(context.Background(), scopeNames[0])
	time.Sleep(100 * time.Millisecond)
	close(ctxCh)
}

func TestIoC_Resolve_Register(t *testing.T) {
	tests := []struct {
		name      string
		scopeName string
		key       string
		f         func(...any) any
		cmdArgs   []any
	}{
		{
			name:      "register test",
			scopeName: "scope1",
			key:       "HelloCommand",
			f: func(a ...any) any {
				return &command{
					f: func() error {
						fmt.Printf("Hello %v\n", a[0])
						return nil
					},
					duration: 1 * time.Millisecond,
				}
			},
			cmdArgs: []any{"some arg"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ioc := NewIoC()
			ctxScoped := setScopeName(ctx, tt.scopeName)

			cmdRaw := ioc.Resolve(ctxScoped, IoCRegister, tt.key, tt.f)

			// Вернулась команда для регистрации.
			cmd, ok := cmdRaw.(ICommand)
			assertEqual(ok, true)

			// Команда регистрации выполнилась без ошибок.
			err := cmd.Execute()
			assertNoError(err)

			cmd, ok = ioc.Resolve(ctxScoped, tt.key, tt.cmdArgs...).(ICommand)

			// Вернулась команда по ключу.
			assertEqual(ok, true)

			// Команда выполнилась без ошибок.
			err = cmd.Execute()
			assertNoError(err)

			// Смена скоупа.
			ctxScoped = setScopeName(ctx, tt.scopeName+strconv.Itoa(rand.Int()))

			// Попытка получения команды не из того скоупа.
			_, ok = ioc.Resolve(ctxScoped, tt.key).(ICommand)

			// Команда не вернулась по ключу.
			assertNotEqual(ok, true)
		})
	}
}

func setScopeName(ctx context.Context, scopeName string) context.Context {
	return context.WithValue(ctx, scopeNameCtxKey{}, scopeName)
}

type command struct {
	f        func() error
	duration time.Duration
}

func (c *command) Execute() error {
	select {
	case <-time.After(c.duration):
		return c.f()
	}
}

func assertNoError(err error) {
	if err != nil {
		panic(fmt.Sprintf("err not nil: %s", err))
	}
}

func assertEqual[T comparable](got T, want T) {
	if got != want {
		panic(fmt.Sprintf("not equal: got = %v, want = %v", got, want))
	}
}

func assertNotEqual[T comparable](got T, want T) {
	if got == want {
		panic(fmt.Sprintf("should be not equal: got = %v, want = %v", got, want))
	}
}

func assertNotNil(got any) {
	if got == nil {
		panic(fmt.Sprintf("should not be equal: got = %v, want = nil", got))
	}
}
