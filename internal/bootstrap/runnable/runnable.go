package runnable

import "context"

type Runnable interface {
	Start() error
	ShutDown(ctx context.Context) error
}
