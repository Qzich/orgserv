package pkg

import "context"

type Logger interface {
	Warn(ctx context.Context, messages ...string)
	Error(ctx context.Context, messages ...string)
	Info(ctx context.Context, messages ...string)
	Debug(ctx context.Context, messages ...string)
}
