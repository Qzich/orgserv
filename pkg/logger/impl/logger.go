package impl

import (
	"context"
	"fmt"
)

func New() loggerImpl {
	return loggerImpl{}
}

type loggerImpl struct{}

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
