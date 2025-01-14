package impl

import (
	"context"
	"fmt"
)

type loggerImpl struct{}

func New() loggerImpl {
	return loggerImpl{}
}

func (l loggerImpl) Warn(ctx context.Context, messages ...string) {
	for _, v := range messages {
		if len(v) == 0 {
			continue
		}
		fmt.Printf("%s\n", v)
	}
}

func (l loggerImpl) Error(ctx context.Context, messages ...string) {
	for _, v := range messages {
		if len(v) == 0 {
			continue
		}
		fmt.Printf("%s\n", v)
	}
}

func (l loggerImpl) Info(ctx context.Context, messages ...string) {
	for _, v := range messages {
		if len(v) == 0 {
			continue
		}
		fmt.Printf("%s\n", v)
	}
}

func (l loggerImpl) Debug(ctx context.Context, messages ...string) {
	for _, v := range messages {
		if len(v) == 0 {
			continue
		}
		fmt.Printf("%s\n", v)
	}
}
